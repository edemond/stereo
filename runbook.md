## Browse the database
`stereo` stores its index and now-playing info in a SQLite3 database called `stereo.sqlite3`.

You can open it and browse around with `sqlite3`:
```
sqlite3 stereo.sqlite3
```

## SSH into the Pi
Use account `pi`, password in KeePass.


## Find the IP address of the Pi
You can get the Pi's IP by its Samba share name with `nmblookup`:

```
nmblookup raspberrypi
```

## Put music on the stereo
The Raspberry Pi has a Samba share at `\\raspberrypi\music`.
Use that to drop MP3, FLAC, OGG, and other stuff there and the stereo will index it.

## Connect to Samba share on Raspberry Pi
Nautilus can connect to Samba shares. Click "+ Other Locations" in the left sidebar.
It'll browse for locations automatically and find the Pi. Username and password are in KeePass.
