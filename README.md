# :cop: janitor [![GitHub License](https://img.shields.io/github/license/nightnoryu/janitor)](https://opensource.org/license/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/nightnoryu/janitor)](https://goreportcard.com/report/github.com/nightnoryu/janitor)

Telegram bot for automated channel moderation.

## Building for local development

Prerequisites:

1. Git
2. Docker
3. [bkit](https://github.com/nightnoryu/bkit)

Firstly, clone the repository into your `$GOPATH`:

```shell
mkdir -p $GOPATH/src/github.com/nightnoryu
cd $GOPATH/src/github.com/nightnoryu

git clone git@github.com:nightnoryu/janitor.git
cd janitor
```

Then build the project:

```shell
bkit build
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
