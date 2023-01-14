package log

import(
  "fmt"
  "log"
  "log/syslog"
)

const LOG_TAG = "stereo"

var _writer *syslog.Writer

func getLogger() *syslog.Writer {
  if _writer != nil {
    return _writer
  }

  writer, err := syslog.New(syslog.LOG_INFO | syslog.LOG_DAEMON, LOG_TAG)
  if err != nil {
    log.Printf("Can't connect to syslog!")
  }

  _writer = writer
  return writer
}

func Infof(format string, args ...interface{}) {
  logger := getLogger()
  _ = logger.Info(fmt.Sprintf(format, args...))
  log.Printf(format, args...)
}

func Warningf(format string, args ...interface{}) {
  logger := getLogger()
  _ = logger.Warning(fmt.Sprintf(format, args...))
  log.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
  logger := getLogger()
  _ = logger.Err(fmt.Sprintf(format, args...))
  log.Printf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
  logger := getLogger()
  _ = logger.Err(fmt.Sprintf(format, args...))
  log.Fatalf(format, args...)
}

func Fatal(msg interface{}) {
  Fatalf(fmt.Sprint(msg))
}
