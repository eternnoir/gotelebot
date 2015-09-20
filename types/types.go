// This package define all type used in Telegram Bot API.
//
// More information : https://core.telegram.org/bots/api#available-types
package types

import (
	"encoding/json"
)

// Update object.
// Optional means this variable could be nil or empty.
type Update struct {
	// The update‘s unique identifier.
	UpdateId float64 `json:"update_id"`
	// Optional. New incoming message of any kind
	Message *Message `json:"message"`
}

// User object.
// Optional means this variable could be nil or empty.
type User struct {
	// Unique identifier for this user or bot
	Id float64 `json:"id"`
	// User‘s or bot’s first name
	FirstName string `json:"first_name"`
	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name"`
	// Optional. User‘s or bot’s username
	Username string `json:"username"`
}

// GroupChat object.
// Optional means this variable could be nil or empty.
type GroupChat struct {
	//	Unique identifier for this group chat
	Id float64 `json:"id"`
	// Group name
	Title string `json:"title"`
}

// Chat object. Representatives GroupChat and User Chat.
// Optional means this variable could be nil or empty.
type Chat struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	Title     string  `json:"title"`
}

// Message object.
// Optional means this variable could be nil or empty.
type Message struct {
	// Unique message identifier
	Message_Id  int64         `josn:"message_id"`
	From        *User         `json:"from"`
	Date        float64       `json:"date"`
	Chat        *Chat         `json:"chat"`
	ForwardFrom *User         `json:"forward_from,omitempty"`
	ForwardDate float64       `json:"forward_date,omitempty"`
	Text        string        `json:"text,omitempty"`
	Audio       *Audio        `json:"audio,omitempty"`
	Document    *Document     `json:"document,omitempty"`
	Photo       *[]*PhotoSize `json:"photo,omitempty"`
	Sticker     *Sticker      `json:"sticker,omitempty"`
	Video       *Video        `json:"video,omitempty"`
	Voice       *Voice        `json:"voice,omitempty"`
	Caption     string        `json:"caption,omitempty"`
	Contact     *Contact      `json:"contact,omitempty"`
	Location    *Location     `json:"location,omitempty"`
}

type PhotoSize struct {
	FileId   string  `json:"file_id"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	FileSize float64 `json:"file_size"`
}

type Audio struct {
	FileId    string  `json:"file_id"`
	Duration  float64 `json:"duration"`
	Performer string  `json:"performer"`
	Title     string  `json:"title"`
	MimeType  string  `json:"mime_type"`
	FileSize  float64 `json:"file_size"`
}

type Document struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize float64   `josn:"file_size"`
}

type Sticker struct {
	FileId   string    `json:"file_id"`
	width    float64   `json:"width"`
	height   float64   `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize float64   `josn:"file_size"`
}

type Video struct {
	FileId   string    `json:"file_id"`
	width    float64   `json:"width"`
	height   float64   `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Duration float64   `json:"duration"`
	MimeType string    `json:"mime_type"`
	FileSize float64   `josn:"file_size"`
}

type Voice struct {
	FileId   string  `json:"file_id"`
	Duration float64 `json:"duration"`
	MimeType string  `json:"mime_type"`
	FileSize float64 `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type ReplyMarkup interface {
	ToJson() (string, error)
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool       `json:one_time_keyboard,omitempty"`
	Selective       bool       `json:selective,omitempty"`
}

func (rkm *ReplyKeyboardMarkup) ToJson() (string, error) {
	b, err := json.Marshal(rkm)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective"`
}

func (rkm *ReplyKeyboardHide) ToJson() (string, error) {
	rkm.HideKeyboard = true
	b, err := json.Marshal(rkm)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

func (rkm *ForceReply) ToJson() (string, error) {
	rkm.ForceReply = true
	b, err := json.Marshal(rkm)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type File struct {
	FileId   string  `json:"file_id"`
	FileSize float64 `json:"file_size"`
	FilePath string  `json:"file_path"`
}
