package gotelebot

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestGetMe(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	bot := InitTeleBot(token)
	me, err := bot.GetMe()
	if err != nil {
		assert.Fail("Bot getMe error")
	}
	assert.NotEmpty(me.Id)
}

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	testMsg := "Test Msg"
	msg, err := bot.SendMessage(chatid, testMsg, nil)
	if err != nil {
		assert.Fail("Bot send message error")
	}
	assert.EqualValues(testMsg, msg.Text)
}

func TestSendMessageWithOpt(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	testMsg := "Test Msg"
	dis := true
	replayId := 267
	opt := &SendMessageOptional{DisableWebPagePreview: &dis, ReplyToMessageId: &replayId}
	msg, err := bot.SendMessage(chatid, testMsg, opt)
	if err != nil {
		assert.Fail("Bot send message error")
	}
	assert.EqualValues(testMsg, msg.Text)
}

func TestForwardMessage(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	msg, err := bot.ForwardMessage(chatid, chatid, 267)
	if err != nil {
		assert.Fail("Bot forwardMessage error")
	}
	assert.NotEmpty(msg.ForwardFrom)

}

func TestSendPhoto(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/go.png"
	msg, err := bot.SendPhoto(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendPhoto error ")
	}
	assert.NotEmpty(msg.Photo)
}
