# gotelebot

Golang (Go)  implementation for the Telegram Bot API.

This project provide a wrapper around the Telegram Bot API with golang. You can easy to ues telegram bot api
in golang way by use this project. And gotelebot provided polling method let developer easy to get new messages.

Almost all method in Telegram bot api have been implement. 
While official telegram bot api update, this project will update as soon as possible.

## Document

Full godoc document [http://godoc.org/github.com/eternnoir/gotelebot](http://godoc.org/github.com/eternnoir/gotelebot)


## Telegram Bot API Support

## Methods

| Telegram Bot API Method | gotelebot Method     | Status      |
|-------------------------|----------------------|-------------|
| getMe                   | GetMe                | Supported   |
| sendMessage             | SendMessage          | Supported   |
| forwardMessage          | ForwardMessage       | Supported   |
| sendPhoto               | SendPhoto            | Supported   |
| sendAudio               | SendAudio            | Supported   |
| sendDocument            | SendDocument         | Supported   |
| sendSticker             | SendSticker          | Supported   |
| sendVideo               | SendVideo            | Supported   |
| sendVoice               | SendVoice            | Supported   |
| sendLocation            | SendLocation         | Supported   |
| sendChatAction          | SendChatAction       | Supported   |
| getUserProfilePhotos    | GetUserProfilePhotos | Supported   |
| getUpdates              | GetUpdates           | Coming soon |
| setWebhook              | SetWebhook           | Coming soon |
| getFile                 | GetFile              | Supported   |

