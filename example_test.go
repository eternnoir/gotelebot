package gotelebot_test

import (
	"fmt"
	"github.com/eternnoir/gotelebot"
)

func Example() {
	bot := gotelebot.InitTeleBot("TOKEN") // Create gotelebot instance
	me, err := bot.GetMe()
	if err != nil {
		fmt.Println("Bot getMe error")
	}
	fmt.Println(me.FirstName)
}
