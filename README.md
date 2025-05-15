# janitor

Telegram bot for automated channel moderation 👮

---

[![Go Report Card](https://goreportcard.com/badge/github.com/nightnoryu/janitor)](https://goreportcard.com/report/github.com/nightnoryu/janitor)
[![GitHub License](https://img.shields.io/github/license/nightnoryu/janitor)](https://opensource.org/license/MIT)

## Building for local development

Prerequisites:

1. Linux
2. Git
3. Docker
4. (optional) [BrewKit](https://github.com/ispringtech/brewkit)

Firstly, clone the repository into your `$GOPATH`:

```shell
mkdir -p $GOPATH/src/github.com/nightnoryu
cd $GOPATH/src/github.com/nightnoryu
git clone git@github.com:nightnoryu/janitor.git
cd janitor
```

Then build the project:

```shell
brewkit build

# Alternatively, if you don't want to use BrewKit, you can do it the old-fashioned way:
# go build -o ./bin/janitor ./cmd/janitor
```

After that, copy the `docker-compose.override.example.yml` to `docker-compose.override.yml` and set the environment variables:

```yaml
services:
  janitor:
    environment:
      JANITOR_TELEGRAM_BOT_TOKEN: 123:ABC # The token for your bot, obtained from t.me/BotFather
```

And you're set! Use `docker compose` to manage the application:

```shell
# Start
docker compose up -d

# Restart to apply changes
docker restart janitor

# Stop
docker compose down
```
