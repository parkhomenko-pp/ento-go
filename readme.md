# EntoGo

[![Go](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/parkhomenko-pp/go-telegram-bot/actions/workflows/go.yml?query=branch:master)
[![codecov](https://codecov.io/github/parkhomenko-pp/ento-go/graph/badge.svg?token=XRDZ7Q1XRC)](https://codecov.io/github/parkhomenko-pp/ento-go)
[![Telegram Bot](.github/preview/tg-badge.svg)](https://t.me/ento_go_bot)

<img src=".github/preview/icon.png" align="right" width=150 height=150/>

This is a Go game project that allows two players to play Go via a Telegram bot. It also includes support for playing the game from the console. Below you will find instructions on how to run the project and the current roadmap for future development.

<br><br><br>

## Run project

```sh
go run run/bot/main.go      # start telegram bot
```

```sh
go run run/console/main.go  # start console game
```

## Run tests
```sh
go test -v -coverprofile=coverage.out ./...         # with coverage
go tool cover -html=coverage.out -o coverage.html   # generate coverage report
```

## Roadmap
- Goban
  - [x] Draw image
  - [x] Themes support
  - [x] Place stones on board
  - [x] Stones without dame (liberties) determine
  - [x] Stones without dame remove from goban
  - [ ] Captured areas count
  - [ ] Captured stones count
- Refactor
  - [ ] Move Printer from goban
  - [ ] Move Image generator from goban
- Console
  - [ ] ✏️TODO
  - [ ] Full game support
- Bot
  - ...
  - [x] return only string from menu (GetReplyMessage)
  - [x] fix menus sorting
  - [x] REGISTRATION: check unique username
  - [x] add "my games" menu
  - [ ] add "game" menu
    - [x] fix empty return message in src/menus/game.go
    - [ ] add the ability to make a move on the board
    - [ ] add current game info
  - ...
  - [ ] Full game support

## Credits
- [JetBrains Mono typeface](https://www.jetbrains.com/lp/mono/)