package gotelebot_test

import (
	"fmt"
)

func Example() {
	bot := InitTeleBot("TOKEN") // Create gotelebot instance
	me, err := bot.GetMe()
	if err != nil {
		fmt.Println("Bot getMe error")
	}
	fmt.Println(me.FirstName)
}
