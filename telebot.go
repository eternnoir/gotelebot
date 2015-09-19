package gotelebot

import (
	"errors"
	"fmt"
	"github.com/eternnoir/gotelebot/types"
	"strconv"
)

type TeleBot struct {
	token           string
	Messages        chan (*types.Message)
	Offset          float64
	stopPollingFlag bool
}

// InitTeleBot is the function to create gotelebot instance.
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

func (bot *TeleBot) SendDocument(chatid int, document string, opt *SendDocumentOptional) (*types.Message, error) {
	return sendDocument(bot.token, strconv.Itoa(chatid), document, opt)
}

func (bot *TeleBot) SendSticker(chatid int, sticker string, opt *SendStickerOptional) (*types.Message, error) {
	return sendSticker(bot.token, strconv.Itoa(chatid), sticker, opt)
}

func (bot *TeleBot) SendVideo(chatid int, video string, opt *SendVideoOptional) (*types.Message, error) {
	return sendVideo(bot.token, strconv.Itoa(chatid), video, opt)
}

func (bot *TeleBot) SendLocation(chatid int, latitude, longitude float64, opt *SendLocationOptional) (*types.Message, error) {
	return sendLocation(bot.token, strconv.Itoa(chatid), strconv.FormatFloat(latitude, 'f', 6, 64), strconv.FormatFloat(longitude, 'f', 6, 64), opt)
}
func (bot *TeleBot) SendChatAction(chatid int, action string) (string, error) {
	return sendChatAction(bot.token, strconv.Itoa(chatid), action)
}

func (bot *TeleBot) GetFile(fileId string) (*types.File, error) {
	return getFile(bot.token, fileId)
}

func (bot *TeleBot) DownloadFile(fileId string) (*[]byte, error) {
	file, err := bot.GetFile(fileId)
	if err != nil {
		return nil, err
	}
	if file.FilePath == "" {
		return nil, errors.New("File path not found")
	}
	filePath := file.FilePath
	return downloadFile(bot.token, filePath)
}

func (bot *TeleBot) StopPolling() {
	bot.stopPollingFlag = true
}

func (bot *TeleBot) StartPolling(nonStop bool) error {
	bot.stopPollingFlag = false
	for {
		if bot.stopPollingFlag == true {
			return nil
		}
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
