package gotelebot

import (
	"github.com/eternnoir/gotelebot/types"
)

type Bot interface {
	GetMe() *types.User
	GetUpdate() []*types.Message
}

type TeleBot struct {
	token    string
	Messages *chan (string)
}

func InitTeleBot(botToken string) *TeleBot {
	bot := new(TeleBot)
	bot.token = botToken
	ch := make(chan string)
	bot.Messages = &ch
	return bot
}

func (bot *TeleBot) GetMe() (*types.User, error) {
	return getMe(bot.token)
}

func (bot *TeleBot) GetUpdates(offset, limit, timeout string) ([]*types.Update, error) {
	return getUpdates(bot.token, offset, limit, timeout)
}
