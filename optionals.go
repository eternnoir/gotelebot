package gotelebot

import (
	"net/url"
	"strconv"
)

type Optional interface {
	AppendPayload(payload *map[string][]string)
}

type SendMessageOptional struct {
	ParseMode             *string
	DisableWebPagePreview *bool
	ReplyToMessageId      *int
	ReplyMarkup           *interface{}
}

func (opt *SendMessageOptional) AppendPayload(payload *url.Values) {
	if opt.ParseMode != nil {
		payload.Add("parse_mode", *opt.ParseMode)
	}
	if opt.DisableWebPagePreview != nil {
		payload.Add("disable_web_page_preview", strconv.FormatBool(*opt.DisableWebPagePreview))
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		payload.Add("parse_mode", *opt.ParseMode)
	}
}

type SendPhotoOptional struct {
	Caption          *string
	ReplyToMessageId *int
	ReplyMarkup      *interface{}
}

func (opt *SendPhotoOptional) AppendPayload(payload *url.Values) {
	if opt.Caption != nil {
		payload.Add("caption", *opt.Caption)
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
}

type SendAudioOptional struct {
	Duration         *int
	Performer        *string
	Title            *string
	ReplyToMessageId *int
	ReplyMarkup      *interface{}
}

func (opt *SendAudioOptional) AppendPayload(payload *url.Values) {
	if opt.Duration != nil {
		payload.Add("duration", strconv.Itoa(*opt.Duration))
	}
	if opt.Performer != nil {
		payload.Add("performer", *opt.Performer)
	}
	if opt.Title != nil {
		payload.Add("title", *opt.Title)
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
}
