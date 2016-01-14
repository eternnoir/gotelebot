// Golang (Go) implementation for the Telegram Bot API.
package gotelebot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/eternnoir/gotelebot/types"
)

type TeleBot struct {
	token               string
	Messages            chan (*types.Message)
	InlineQuerys        chan (*types.InlineQuery)
	ChosenInlineResults chan (*types.ChosenInlineResult)
	Offset              float64
	stopPollingFlag     bool
}

// InitTeleBot is the function to create gotelebot instance.
func InitTeleBot(botToken string) *TeleBot {
	bot := new(TeleBot)
	bot.token = botToken
	bot.Messages = make(chan *types.Message)
	bot.InlineQuerys = make(chan *types.InlineQuery)
	bot.ChosenInlineResults = make(chan *types.ChosenInlineResult)
	return bot
}

// Returns basic information about the bot in form of a User object.
func (bot *TeleBot) GetMe() (*types.User, error) {
	return getMe(bot.token)
}

// Use this method to receive incoming updates using long polling. An Array of Update objects is returned.
func (bot *TeleBot) GetUpdates(offset, limit string, timeout int) ([]*types.Update, error) {
	return getUpdates(bot.token, offset, limit, strconv.Itoa(timeout))
}

// Use this method to send text messages. On success, the sent Message type is returned.
//
// Use SendMessageOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendMessage(chatid int, text string, opt *SendMessageOptional) (*types.Message, error) {
	return sendMessage(bot.token, strconv.Itoa(chatid), text, opt)
}

// Use this method to forward messages of any kind. On success, the sent Message type is returned.
func (bot *TeleBot) ForwardMessage(chatid, from_chat_id, message_id int) (*types.Message, error) {
	return forwardMessage(bot.token, strconv.Itoa(chatid), strconv.Itoa(from_chat_id), strconv.Itoa(message_id))
}

//Use this method to send photos. On success, the sent Message is returned.
//
// Use SendPhotoOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendPhoto(chatid int, photo string, opt *SendPhotoOptional) (*types.Message, error) {
	return sendPhoto(bot.token, strconv.Itoa(chatid), photo, opt)
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format.
// On success, the sent Message is returned.
//
// Use SendAudioOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendAudio(chatid int, audio string, opt *SendAudioOptional) (*types.Message, error) {
	return sendAudio(bot.token, strconv.Itoa(chatid), audio, opt)
}

// Use this method to send general files. On success, the sent Message is returned.
//
// Use SendDocumentOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendDocument(chatid int, document string, opt *SendDocumentOptional) (*types.Message, error) {
	return sendDocument(bot.token, strconv.Itoa(chatid), document, opt)
}

// Use this method to send .webp stickers. On success, the sent Message is returned.
//
// Use SendStickerOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendSticker(chatid int, sticker string, opt *SendStickerOptional) (*types.Message, error) {
	return sendSticker(bot.token, strconv.Itoa(chatid), sticker, opt)
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
//
// Use SendVideoOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendVideo(chatid int, video string, opt *SendVideoOptional) (*types.Message, error) {
	return sendVideo(bot.token, strconv.Itoa(chatid), video, opt)
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
//
// Use SendVideoOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendVoice(chatid int, voice string, opt *SendVoiceOptional) (*types.Message, error) {
	return sendVoice(bot.token, strconv.Itoa(chatid), voice, opt)
}

// Use this method to send point on the map. On success, the sent Message is returned.
//
// Use SendLocationOptional to setup optional parameters. If you don't want use any optional parameters, just asign nil to opt.
func (bot *TeleBot) SendLocation(chatid int, latitude, longitude float64, opt *SendLocationOptional) (*types.Message, error) {
	return sendLocation(bot.token, strconv.Itoa(chatid), strconv.FormatFloat(latitude, 'f', 6, 64), strconv.FormatFloat(longitude, 'f', 6, 64), opt)
}

func (bot *TeleBot) AnswerInlineQuery(inlineQueryId string, results []interface{}, opt *AnswerInlineQueryOptional) (bool, error) {
	return answerInlineQuery(bot.token, inlineQueryId, results, opt)
}

// Use this method when you need to tell the user that something is happening on the bot's side.
//
// action can be :
// "typing" for text messages,
// "upload_photo" for photos,
// "record_video" or upload_video for videos,
// "record_audio" or upload_audio for audio files,
// "upload_document" for general files,
// "find_location" for location data.
func (bot *TeleBot) SendChatAction(chatid int, action string) (string, error) {
	return sendChatAction(bot.token, strconv.Itoa(chatid), action)
}

// Use this method to get a list of profile pictures for a user.
func (bot *TeleBot) GetUserProfilePhotos(userid int, opt *GetUserProfilePhotosOptional) (*types.UserProfilePhotos, error) {
	return getUserProfilePhotos(bot.token, strconv.Itoa(userid), opt)
}

// Use this method to get basic info about a file and prepare it for downloading.
func (bot *TeleBot) GetFile(fileId string) (*types.File, error) {
	return getFile(bot.token, fileId)
}

// Download file by fileId.
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

// Let gotelebot stop polling new messages.
func (bot *TeleBot) StopPolling() {
	bot.stopPollingFlag = true
}

// Let gotelebot always try to get new messages. This function will put new message to gotelebot's Message channel.
func (bot *TeleBot) StartPolling(nonStop bool, timeout int) error {
	bot.stopPollingFlag = false
	for {
		if bot.stopPollingFlag == true {
			return nil
		}
		newUpdates, err := bot.GetUpdates(strconv.Itoa(int(bot.Offset)), "", timeout)
		if err != nil {
			if !nonStop {
				return err
			} else {
				fmt.Println(err)
			}
		}
		bot.processNewUpdate(newUpdates)
	}
}

func (bot *TeleBot) processNewUpdate(updates []*types.Update) {
	for _, update := range updates {
		if update.UpdateId >= bot.Offset {
			bot.Offset = update.UpdateId + 1
		}
		switch {
		case update.Message != nil:
			bot.Messages <- update.Message
		case update.InlineQuery != nil:
			bot.InlineQuerys <- update.InlineQuery
		case update.ChosenInlineResult != nil:
			bot.ChosenInlineResults <- update.ChosenInlineResult
		}
	}
}
