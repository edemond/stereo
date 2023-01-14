package db

import (
	"database/sql"
	"fmt"
	"github.com/edemond/stereo/log"
	"github.com/edemond/stereo/media"
	_ "github.com/mattn/go-sqlite3"
)

// okay what operations does this thing need to support?
// - indexing: create index, update index
// - listing all [songs|albums]
// - retrieving a [song|album] by ID
type Database interface {
}

func Initialize() error {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := `
    CREATE TABLE IF NOT EXISTS artists (
      id integer primary key autoincrement,
      name text
    );

    CREATE TABLE IF NOT EXISTS songs (
      id integer primary key autoincrement,
      artist_id integer, 
      path text,
      title text,
      FOREIGN KEY(artist_id) REFERENCES artists(id)
    );

    CREATE TABLE IF NOT EXISTS albums (
      id integer primary key autoincrement,
      artist_id integer,
      title text,
      FOREIGN KEY(artist_id) REFERENCES artists(id)
    );
    
    CREATE TABLE IF NOT EXISTS nowplaying (
      id integer primary key,
      FOREIGN KEY(id) REFERENCES songs(id)
    );

    CREATE TRIGGER IF NOT EXISTS nowplaying_only_one_row
    BEFORE INSERT ON nowplaying
    WHEN (SELECT COUNT(*) FROM nowplaying) >= 1
    BEGIN
      SELECT RAISE(FAIL, "limit to only one row");
    END;
  `
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func openDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "./stereo.sqlite3")
}

// execInTransaction executes some work in the context of a transaction.
// If the function returns a non-nil error, the transaction will be
// rolled back; otherwise it will be committed.
func execInTransaction(fn func(tx *sql.Tx) error) error {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Couldn't begin transaction: %v", err)
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Couldn't commit transaction: %v", err)
	}

	return nil
}

func deleteAllSongs(tx *sql.Tx) error {
	deleteStatement, err := tx.Prepare(`
      DELETE FROM songs
    `)
	if err != nil {
		return fmt.Errorf("Couldn't prepare songs DELETE statement: %v", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec()
	return err
}

func deleteAllArtists(tx *sql.Tx) error {
	deleteStatement, err := tx.Prepare(`
      DELETE FROM artists
    `)
	if err != nil {
		return fmt.Errorf("Couldn't prepare artists DELETE statement: %v", err)
	}
	defer deleteStatement.Close()

	_, err = deleteStatement.Exec()
	return err
}

func Reindex(contents *media.DirectoryContents) error {
	return execInTransaction(func(tx *sql.Tx) error {
		err := deleteAllSongs(tx)
		if err != nil {
			return fmt.Errorf("Couldn't delete songs from index: %v", err)
		}

		err = deleteAllArtists(tx)
		if err != nil {
			return fmt.Errorf("Couldn't delete artists from index: %v", err)
		}

		artistStmt, err := tx.Prepare(`
      INSERT INTO artists (name) 
      VALUES (?)
    `)
		if err != nil {
			return fmt.Errorf("Couldn't prepare artist INSERT statement: %v", err)
		}
		defer artistStmt.Close()

		songStmt, err := tx.Prepare(`
      INSERT INTO songs (artist_id, path, title) 
      VALUES (?, ?, ?)
    `)
		if err != nil {
			return fmt.Errorf("Couldn't prepare song INSERT statement: %v", err)
		}
		defer songStmt.Close()

		for _, artistSongs := range contents.ArtistSongs {
			result, err := artistStmt.Exec(artistSongs.Artist.Name)
			if err != nil {
				return fmt.Errorf("Couldn't insert artist: %v", err)
			}

			artistId, err := result.LastInsertId()
			if err != nil {
				return fmt.Errorf(
					"Couldn't get last-inserted artist ID after inserting %v.",
					artistSongs.Artist.Name,
				)
			}

			for _, song := range artistSongs.Songs {
				_, err = songStmt.Exec(artistId, song.Path, song.Title)
				if err != nil {
					return fmt.Errorf("Couldn't insert song: %v", err)
				}
			}
		}

		return nil
	})
}

func Search() ([]*media.SearchResult, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	panic("TODO")
}

func GetAllSongs() ([]*media.Song, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, err := db.Prepare(`
    SELECT s.id, a.id, a.name, s.title
    FROM songs s 
    JOIN artists a ON s.artist_id = a.id
  `)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := []*media.Song{}

	for rows.Next() {
		var id string
		var artistId string
		var artist string
		var title string
		err = rows.Scan(&id, &artistId, &artist, &title)
		if err != nil {
			log.Warningf("Error scanning row in query: %v", err)
			continue
		}
		songs = append(songs, &media.Song{
			Id:       id,
			ArtistId: artistId,
			Artist:   artist,
			Title:    title,
		})
	}

	return songs, rows.Err()
}

func GetSongByID(requestedID string) (*media.SongFile, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, err := db.Prepare(`
    SELECT s.id, a.id, a.name, s.path, s.title
    FROM songs s
    JOIN artists a ON a.id = s.artist_id
    WHERE s.id = ?
  `)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(requestedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasRow := rows.Next()
	if !hasRow {
		return nil, rows.Err()
	}

	var id string
	var artistID string
	var artist string
	var path string
	var title string
	err = rows.Scan(&id, &artistID, &artist, &path, &title)
	if err != nil {
		log.Warningf("Error scanning row in query: %v", err)
		return nil, err
	}

	return &media.SongFile{
		Path: path,
		Song: media.Song{
			Id:       id,
			ArtistId: id,
			Artist:   artist,
			Title:    title,
		},
	}, nil
}

func GetArtists() ([]*media.Artist, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, err := db.Prepare(`
    SELECT id, name
    FROM artists
    ORDER BY name COLLATE NOCASE
  `)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasRow := rows.Next()
	if !hasRow {
		return nil, rows.Err()
	}

	artists := []*media.Artist{}
	for rows.Next() {
		var id string
		var name string
		rows.Scan(&id, &name)
		artists = append(artists, &media.Artist{
			Id:   id,
			Name: name,
		})
	}

	return artists, nil
}

func GetNowPlayingSong() (*media.SongFile, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, err := db.Prepare(`
    SELECT s.id, a.name, s.path, s.title
    FROM songs s
    JOIN nowplaying np ON np.id = s.id
    JOIN artists a ON a.id = s.artist_id
  `)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasRow := rows.Next()
	if !hasRow {
		return nil, rows.Err()
	}

	var id string
	var artist string
	var path string
	var title string
	err = rows.Scan(&id, &artist, &path, &title)
	if err != nil {
		log.Warningf("Error scanning row in query: %v", err)
		return nil, err
	}

	return &media.SongFile{
		Path: path,
		Song: media.Song{
			Id:     id,
			Artist: artist,
			Title:  title,
		},
	}, nil
}

func SetNowPlayingSong(song *media.SongFile) error {
	return execInTransaction(func(tx *sql.Tx) error {
		deleteStmt, err := tx.Prepare(`
      DELETE FROM nowplaying
    `)
		if err != nil {
			return fmt.Errorf("Couldn't prepare DELETE statement: %v", err)
		}
		defer deleteStmt.Close()

		_, err = deleteStmt.Exec()
		if err != nil {
			return fmt.Errorf("Couldn't delete data from nowplaying: %v", err)
		}

		stmt, err := tx.Prepare(`
      INSERT INTO nowplaying (id)
      VALUES (?)
    `)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(song.Id)
		if err != nil {
			return fmt.Errorf("Couldn't insert into nowplaying: %v", err)
		}

		return nil
	})
}

func ClearNowPlayingSong() error {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
    DELETE FROM nowplaying
  `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("Couldn't delete from nowplaying: %v", err)
	}

	return nil
}
