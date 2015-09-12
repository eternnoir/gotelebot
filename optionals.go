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
}
