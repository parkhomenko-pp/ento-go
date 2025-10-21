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
  - [x] Captured stones count
  - [ ] Captured areas count
  - [ ] Final score count
- Bot
  - ...
  - [x] return only string from menu (GetReplyMessage)
  - [x] fix menus sorting
  - [x] REGISTRATION: check unique username
  - [x] "new game menu"
    - [x] game creation action
    - [x] goban size selection
  - [x] add "my games" menu
    - [x] add "join game" menu
    - [x] add "invited" menu
    - [x] add "declined" menu
    - [x] add "finished" menu
  - [x] add "game" menu
    - [x] fix empty return message in src/menus/game.go
    - [x] add the ability to make a move on the board
    - [ ] add current game info
  - [x] add settings menu
    - [x] add theme change aviability
  - ...
  - [ ] Full game support

## Credits
- [JetBrains Mono typeface](https://www.jetbrains.com/lp/mono/)
