package gotelebot_test

import (
	"github.com/eternnoir/gotelebot"
)

func Example() {
	// Echo Bot example.
	bot := gotelebot.InitTeleBot("TOKEN") // Create gotelebot instance
	go bot.StartPolling(true)             // Start get new message whit goroutine.
	newMsgChan := bot.Messages
	for {
		m := <-newMsgChan // Get new messaage, when new message arrive.
		if m.Text != "" { // Check message is text message.
			bot.SendMessage(int(m.Chat.Id), m.Text, nil)
		}
	}
}
