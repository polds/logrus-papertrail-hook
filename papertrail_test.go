package logrus_papertrail

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/polds/logrus-papertrail-hook"
	"github.com/stvp/go-udp-testing"
)

func TestWritingToUDP(t *testing.T) {
	port := 16661
	udp.SetAddr(fmt.Sprintf(":%d", port))

	hook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
		Host:     "localhost",
		Port:     port,
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

func TestReal(t *testing.T) {
	port, _ := strconv.Atoi(os.Getenv("PAPERTRAIL_PORT"))

	hook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
		Host:     os.Getenv("PAPERTRAIL_HOST"),
		Port:     port,
		Hostname: "appserver",
		Appname:  "myapp",
	})
	if err != nil {
		t.Errorf("Unable to connect to Papertrail")
	}

	log := logrus.New()
	log.Hooks.Add(hook)

	log.Infoln("testing")
}
