package logrus_papertrail

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stvp/go-udp-testing"
)

const (
	HOST = "localhost"
	PORT = 16661
)

type test_connect struct {
	buffer []byte
}

func (t *test_connect) Write(b []byte) (int, error) {
	t.buffer = append(t.buffer, b...)
	return len(b), nil
}

func (t *test_connect) Show() []byte {
	return t.buffer
}

func TestWritingToUDP(t *testing.T) {

	udp.SetAddr(fmt.Sprintf(":%d", PORT))

	hook, err := NewPapertrailHook(&Hook{
		Host:     HOST,
		Port:     PORT,
		Hostname: "test.local",
		Appname:  "test",
	})
	if err != nil {
		t.Errorf("Unable to connect to local UDP server.")
	}

	log := logrus.New()
	log.Hooks.Add(hook)

	udp.ShouldReceive(t, "foo", func() {
		log.Info("foo")
	})
}

func TestWritingToTCP_FAKE(t *testing.T) {

	tconn := &test_connect{}

	hook, err := NewPapertrailTCPHook(&Hook{
		Host:     HOST,
		Port:     PORT,
		Hostname: "appserver",
		Appname:  "myapp",
	})

	if err == nil {
		t.Errorf("Fake connection! Must give error!")
	}

	// here i replace conn interface to check if everything is ok with hook
	hook.conn = tconn

	log := logrus.New()
	log.Hooks.Add(hook)

	log.Infoln("testing TCP")

	received := tconn.Show()

	if len(received) > 0 {
		t.Logf("%s", received)
	} else {
		t.Error("Nothing was received!")
	}

}
