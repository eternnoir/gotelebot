package gotelebot_test

import (
	"fmt"
	"github.com/eternnoir/gotelebot"
	"io/ioutil"
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

func ExampleTeleBot_GetMe() {
	bot := gotelebot.InitTeleBot("TOKEN") // Create gotelebot instance
	me, err := bot.GetMe()                //Get user object.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(me.FirstName)
}

func ExampleTeleBot_SendMessage() {
	bot := gotelebot.InitTeleBot("TOKEN") // Create gotelebot instance
	testMsg := "Test Msg"
	chatid := 11111111
	_, err := bot.SendMessage(chatid, testMsg, nil)
	if err != nil {
		fmt.Println("Bot send message error")
	}
}

func ExampleTeleBot_SendPhoto() {
	bot := gotelebot.InitTeleBot("TOKEN") // Create gotelebot instance
	chatid := 11111111
	filePath := "./test_data/go.png"
	_, err := bot.SendPhoto(chatid, filePath, nil)
	if err != nil {
		fmt.Println("Bot send Photo error")
	}
}

func ExampleTeleBot_DownloadFile() {
	token := "TOKEN"
	bot := gotelebot.InitTeleBot(token)
	fi := "BQADBQADnAMAAsYifgZph-iT9_z_rgI"
	file, err := bot.DownloadFile(fi)
	if err != nil {
		fmt.Println("Bot get File error")
		return
	}
	ferr := ioutil.WriteFile("/tmp/data", *file, 0644)
	if ferr != nil {
		fmt.Println("Write to File error")
	}
}
