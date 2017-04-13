# Papertrail Hook for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" /> [![Build Status](https://travis-ci.org/polds/logrus-papertrail-hook.svg)](https://travis-ci.org/polds/logrus-papertrail-hook)&nbsp;[![godoc reference](https://godoc.org/github.com/polds/logrus-papertrail-hook?status.png)](https://godoc.org/gopkg.in/polds/logrus-papertrail-hook.v2)

[Papertrail](https://papertrailapp.com) provides hosted log management. Once stored in Papertrail, you can [group](http://help.papertrailapp.com/kb/how-it-works/groups/) your logs on various dimensions, [search](http://help.papertrailapp.com/kb/how-it-works/search-syntax) them, and trigger [alerts](http://help.papertrailapp.com/kb/how-it-works/alerts).

In most deployments, you'll want to send logs to Papertrail via their [remote_syslog](http://help.papertrailapp.com/kb/configuration/configuring-centralized-logging-from-text-log-files-in-unix/) daemon, which requires no application-specific configuration. This hook is intended for relatively low-volume logging, likely in managed cloud hosting deployments where installing `remote_syslog` is not possible.

## Usage

You can find your Papertrail port(Accepting TCP/TLS, UDP) on your [Papertrail account page](https://papertrailapp.com/account/destinations). Substitute it below for `YOUR_PAPERTRAIL_PORT`.

For `YOUR_APP_NAME` and `YOUR_HOST_NAME`, substitute a short strings that will readily identify your application and server in the logs.  
If you leave `YOUR_HOST_NAME` empty, papertrail will replace it with your ip.

```go
import (
  log "github.com/Sirupsen/logrus"
  "github.com/polds/logrus-papertrail-hook"
)

func main() {

  hook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
    Host: "logs.papertrailapp.com",
    Port: YOUR_PAPERTRAIL_PORT,
    Hostname: YOUR_HOST_NAME,
    Appname: YOUR_APP_NAME
  })

  hook.SetLevels([]log.Level{log.ErrorLevel, log.WarnLevel})

  if err == nil {
    log.AddHook(hook)
  }

  log.Warning("Here is you message")

}
```

## Changelog
- [gopkg.in/polds/logrus-papertrail-hook.v1](https://godoc.org/gopkg.in/polds/logrus-papertrail-hook.v1)
    - Unchanged from split from [logrus](https://github.com/Sirupsen/logrus)
- [gopkg.in/polds/logrus-papertrail-hook.v2](https://godoc.org/gopkg.in/polds/logrus-papertrail-hook.v2)
    - Adds support for custom hostnames. Major API change.
- [gopkg.in/polds/logrus-papertrail-hook.v3](https://godoc.org/gopkg.in/polds/logrus-papertrail-hook.v3)
    - Low case path. Read more [here](https://github.com/polds/logrus-papertrail-hook/issues/8) and [here](https://github.com/sirupsen/logrus/issues/451)