# gotelebot

[![GoDoc](https://godoc.org/github.com/eternnoir/gotelebot?status.svg)](https://godoc.org/github.com/eternnoir/gotelebot)

Golang (Go)  implementation for the Telegram Bot API.

This project provide a wrapper around the Telegram Bot API with golang. You can easy to ues telegram bot api
in golang way by use this project. And gotelebot provided polling method let developer easy to get new messages.

Almost all method in Telegram bot api have been implement. 
While official telegram bot api update, this project will update as soon as possible.

## Installation

```
go get github.com/eternnoir/gotelebot
```

## Document

Full godoc document [http://godoc.org/github.com/eternnoir/gotelebot](http://godoc.org/github.com/eternnoir/gotelebot)

### Echo Bot Example

```go
package main

import (
	"fmt"
	"github.com/eternnoir/gotelebot"
)

func main() {
	// Echo Bot example.

	// Create gotelebot instance
	bot := gotelebot.InitTeleBot("TOKEN")
	// Start get new message whit goroutine and 60s timeout.
	go bot.StartPolling(true, 60)
	go processNewMessage(bot)
	processNewInlineQuery(bot)
}

func processNewMessage(bot *gotelebot.TeleBot) {
	newMsgChan := bot.Messages
	for {
		m := <-newMsgChan // Get new messaage, when new message arrive.
		fmt.Printf("Get Message:%#v \n", m)
		if m.Text != "" { // Check message is text message.
			bot.SendMessage(int(m.Chat.Id), m.Text, nil)
		}
	}
}

func processNewInlineQuery(bot *gotelebot.TeleBot) {
	newQuery := bot.InlineQuerys
	for {
		q := <-newQuery
		fmt.Printf("Get NewInlineQuery:%#v \n", q)
		if q.Query != "" {	// Only return result when query string not empty.
			result1 := types.NewInlineQueryResultArticl()
			result1.Id = "1"
			result1.Title = "Example"
			result1.MessageText = "Hi" + q.Query
			_, err := bot.AnswerInlineQuery(q.Id, []interface{}{result1}, nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}


```

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
| getUpdates              | GetUpdates           | Supported   |
| getFile                 | GetFile              | Supported   |
| inline mode             | inline mode          | Supported   |



# Change Log

## 2015-10-12
* New type ```Chat``` support. More information : https://core.telegram.org/bots/api#recent-changes
