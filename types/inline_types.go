// This package define all inline type used in Telegram Bot API.
//
// More information : https://core.telegram.org/bots/api#inline-mode
package types

import (
	_ "encoding/json"
)

type InlineQuery struct {
	Id     string `json:"id"`
	From   *User  `json:"from"`
	Query  string `json:"query"`
	offset string `json:"offset"`
}

type ChosenInlineResult struct {
	ResultId string `json:"result_id"`
	From     *User  `json:"from"`
	Query    string `json:"query"`
}

type InlineQueryResult struct {
	Type                  string  `json:"type"`
	Id                    string  `json:"id"`
	ParseMode             *string `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool   `json:"description,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	InlineQueryResult
	Mpeg4FileId string  `json:"mpeg4_file_id"`
	Title       *string `json:"title,omitempty"`
	Caption     *string `josn:"caption,omitempty"`
}

func NewInlineQueryResultCachedMpeg4Gif() *InlineQueryResultCachedMpeg4Gif {
	ret := new(InlineQueryResultCachedMpeg4Gif)
	ret.Type = "mpeg4_gif"
	return ret
}

type InlineQueryResultArticle struct {
	InlineQueryResult
	Title       string  `json:"titl`
	MessageText string  `json:"message_text"`
	Url         string  `json:"url,omitempty"`
	HideUrl     bool    `json:"hide_url,omitempty"`
	Description string  `json:"description,omitempty"`
	ThumbUrl    string  `json:"thumb_url,omitempty"`
	ThumbWidth  float64 `json:"thumb_width,omitempty"`
	ThumbHeight float64 `json:"thumb_height,omitempty"`
}

func NewInlineQueryResultArticl() *InlineQueryResultArticle {
	ret := new(InlineQueryResultArticle)
	ret.Type = "article"
	return ret
}

type InlineQueryResultPhoto struct {
	InlineQueryResult
	PhotoUrl    string  `json:"photo_url"`
	PhotoWidth  float64 `json:"photo_width,omitempty"`
	PhotoHeight float64 `json:"photo_height,omitempty"`
	ThumbUrl    string  `json:"thumb_url"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Caption     string  `json:"caption,omitempty"`
	MessageText string  `json:"message_text,omitempty"`
}

func NewInlineQueryResultPhoto() *InlineQueryResultPhoto {
	ret := new(InlineQueryResultPhoto)
	ret.Type = "photo"
	return ret
}

type InlineQueryResultGif struct {
	InlineQueryResult
	GifUrl      string  `json:"gif_url"`
	GifWidth    float64 `json:"gif_width,omitempty"`
	GifHeight   float64 `json:"gif_height,omitempty"`
	ThumbUrl    string  `json:"thumb_url"`
	Title       string  `json:"title,omitempty"`
	Caption     string  `json:"caption,omitempty"`
	MessageText string  `json:"message_text,omitempty"`
}

func NewInlineQueryResultGif() *InlineQueryResultGif {
	ret := new(InlineQueryResultGif)
	ret.Type = "gif"
	return ret
}

type InlineQueryResultMpeg4Gif struct {
	InlineQueryResult
	Mpeg4Url    string   `json:"mpeg4_url"`
	Mpeg4Width  *float64 `json:"mpeg4_width,omitempty"`
	Mpeg4Height *float64 `json:"mpeg4_height,omitempty"`
	ThumbUrl    string   `json:"thumb_url"`
	Title       *string  `json:"title,omitempty"`
	Caption     *string  `json:"caption,omitempty"`
	MessageText *string  `json:"message_text,omitempty"`
}

func NewInlineQueryResultMpeg4Gif() *InlineQueryResultMpeg4Gif {
	ret := new(InlineQueryResultMpeg4Gif)
	ret.Type = "mpeg4_gif"
	return ret
}

type InlineQueryResultVideo struct {
	InlineQueryResult
	VideoUrl      string  `json:"video_url"`
	MimeType      string  `json:"mime_type"`
	MessageText   string  `json:"message_text"`
	VideoWidth    float64 `json:"video_width,omitempty"`
	VideoHeight   float64 `json:"video_height,omitempty"`
	VideoDuration float64 `json:"video_duration,omitempty"`
	Title         string  `json:"title"`
	ThumbUrl      string  `json:"thumb_url"`
	Description   string  `json:"description,omitempty"`
}

func NewInlineQueryResultVideo() *InlineQueryResultVideo {
	ret := new(InlineQueryResultVideo)
	ret.Type = "video"
	return ret
}
