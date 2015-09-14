package gotelebot

import (
	"errors"
	"fmt"
	"github.com/eternnoir/gotelebot/types"
	"strconv"
)

type Bot interface {
	GetMe() (*types.User, error)
	GetUpdate() ([]*types.Message, error)
	StartPolling() error
}

type TeleBot struct {
	token    string
	Messages chan (*types.Message)
	Offset   float64
}

func InitTeleBot(botToken string) *TeleBot {
	bot := new(TeleBot)
	bot.token = botToken
	ch := make(chan *types.Message)
	bot.Messages = ch
	return bot
}

func (bot *TeleBot) GetMe() (*types.User, error) {
	return getMe(bot.token)
}

func (bot *TeleBot) GetUpdates(offset, limit, timeout string) ([]*types.Update, error) {
	return getUpdates(bot.token, offset, limit, timeout)
}

func (bot *TeleBot) SendMessage(chatid int, text string, opt *SendMessageOptional) (*types.Message, error) {
	return sendMessage(bot.token, strconv.Itoa(chatid), text, opt)
}

func (bot *TeleBot) ForwardMessage(chatid, from_chat_id, message_id int) (*types.Message, error) {
	return forwardMessage(bot.token, strconv.Itoa(chatid), strconv.Itoa(from_chat_id), strconv.Itoa(message_id))
}

func (bot *TeleBot) SendPhoto(chatid int, photo string, opt *SendPhotoOptional) (*types.Message, error) {
	return sendPhoto(bot.token, strconv.Itoa(chatid), photo, opt)
}

func (bot *TeleBot) SendAudio(chatid int, audio string, opt *SendAudioOptional) (*types.Message, error) {
	return sendAudio(bot.token, strconv.Itoa(chatid), audio, opt)
}

func (bot *TeleBot) StartPolling(nonStop bool) error {
	for {
		newUpdates, err := bot.GetUpdates(strconv.Itoa(int(bot.Offset)), "", "")
		if err != nil {
			if !nonStop {
				return err
			} else {
				fmt.Println(err)
			}
		}
		newMessages, errs := bot.processNewUpdate(newUpdates)
		if errs != nil {
			if !nonStop {
				return errs
			} else {
				fmt.Println(errs)
			}
		}
		for _, m := range newMessages {
			bot.Messages <- m
		}
	}
}

func (bot *TeleBot) processNewUpdate(updates []*types.Update) ([]*types.Message, error) {
	retMessages := []*types.Message{}
	for _, update := range updates {
		if update.UpdateId >= bot.Offset {
			bot.Offset = update.UpdateId + 1
		}
		if update.Message == nil {
			return nil, errors.New("[telebot][ERROR] Message is null.")
		}
		retMessages = append(retMessages, update.Message)
	}
	return retMessages, nil
}
