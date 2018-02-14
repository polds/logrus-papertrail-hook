package logrus_papertrail

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	format = "Jan 2 15:04:05"
)

// private interface to make tests with TCP conn
type conn_int interface {
	Write([]byte) (int, error)
}

// PapertrailHook to send logs to a logging service compatible with the Papertrail API.
type Hook struct {
	// Connection Details
	Host string
	Port int

	// App Details
	Appname  string
	Hostname string

	conn_type string

	levels []logrus.Level

	conn conn_int
}

// NewPapertrailHook creates a UDP hook to be added to an instance of logger.
func NewPapertrailHook(hook *Hook) (*Hook, error) {
	var err error
	hook.conn_type = "udp"
	hook.conn, err = net.Dial(hook.conn_type, fmt.Sprintf("%s:%d", hook.Host, hook.Port))
	return hook, err
}

// NewPapertrailTCPHook creates a TCP/TLS hook to be added to an instance of logger.
func NewPapertrailTCPHook(hook *Hook) (*Hook, error) {
	var err error
	hook.conn_type = "tcp"
	hook.conn, err = tls.Dial(hook.conn_type, fmt.Sprintf("%s:%d", hook.Host, hook.Port), nil)
	return hook, err
}

// Fire is called when a log event is fired.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	date := time.Now().Format(format)
	msg, _ := entry.String()
	payload := fmt.Sprintf("<22> %s %s %s: %s", date, hook.Hostname, hook.Appname, msg)

	bytesWritten, err := hook.conn.Write([]byte(payload))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to send log line to Papertrail via %s. Wrote %d bytes before error: %v", hook.conn_type, bytesWritten, err)
		return err
	}

	return nil
}

// SetLevels specify nessesary levels for this hook.
func (hook *Hook) SetLevels(lvs []logrus.Level) {
	hook.levels = lvs
}

// Levels returns the available logging levels.
func (hook *Hook) Levels() []logrus.Level {

	if hook.levels == nil {
		return []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
		}
	}

	return hook.levels
}
