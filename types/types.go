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
	Message *Message `json:"message,omitempty"`
	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// Optional. The result of a inline query that was chosen by a user and sent to their chat partner
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
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
	Type      string  `json:"type"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	Title     string  `json:"title"`
}

// Message object.
// Optional means this variable could be nil or empty.
type Message struct {
	// Unique message identifier
	Message_Id int64 `josn:"message_id"`
	// Sender
	From *User `json:"from"`
	// Date the message was sent.
	Date float64 `json:"date"`
	// Conversation the message belongs to.
	Chat *Chat `json:"chat"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For forwarded messages, date the original message was sent.
	ForwardDate float64 `json:"forward_date,omitempty"`
	// Optional. For text messages
	Text string `json:"text,omitempty"`
	// Optional. Message is an audio file
	Audio *Audio `json:"audio,omitempty"`
	// Optional. Message is a general file
	Document *Document `json:"document,omitempty"`
	// Optional. Message is a photo
	Photo *[]*PhotoSize `json:"photo,omitempty"`
	// Optional. Message is a sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Message is a video
	Video *Video `json:"video,omitempty"`
	// Optional. Message is a voice message
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Caption for the photo or video
	Caption string `json:"caption,omitempty"`
	// Optional. Message is a shared contact
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Message is a shared location
	Location *Location `json:"location,omitempty"`
	// Optional. A new member was added to the group
	NewChatParticipant *User `json:"new_chat_participant,omitempty"`
	// Optional. A member was removed from the group
	LeftChatParticipant *User `json:"left_chat_participant,omitempty"`
	//	Optional. A group title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`
	//	Optional. A group photo was change to this value
	NewChatPhoto *[]*PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Informs that the group photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Informs that the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
// Optional means this variable could be nil or empty.
type PhotoSize struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Photo width
	Width float64 `json:"width"`
	// Photo height
	Height float64 `json:"height"`
	// Optional. File size
	FileSize float64 `json:"file_size"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
// Optional means this variable could be nil or empty.
type Audio struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender
	Duration float64 `json:"duration"`
	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer"`
	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`
	// Optional. File size
	FileSize float64 `json:"file_size"`
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
// Optional means this variable could be nil or empty.
type Document struct {
	// Unique file identifier
	FileId string `json:"file_id"`
	// Optional. Document thumbnail as defined by sender
	Thumb PhotoSize `json:"thumb"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type"`
	// Optional. File size
	FileSize float64 `josn:"file_size"`
}

// This object represents a sticker.
// Optional means this variable could be nil or empty.
type Sticker struct {
	// Unique identifier for this file
	FileId string `json:"file_id"`
	// Sticker width
	width float64 `json:"width"`
	// Sticker height
	height float64 `json:"height"`
	// Optional. Sticker thumbnail in .webp or .jpg format
	Thumb PhotoSize `json:"thumb"`
	// Optional. File size
	FileSize float64 `josn:"file_size"`
}

// This object represents a video file.
type Video struct {
	FileId   string    `json:"file_id"`
	width    float64   `json:"width"`
	height   float64   `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Duration float64   `json:"duration"`
	MimeType string    `json:"mime_type"`
	FileSize float64   `josn:"file_size"`
}

//This object represents a voice note.
type Voice struct {
	FileId   string  `json:"file_id"`
	Duration float64 `json:"duration"`
	MimeType string  `json:"mime_type"`
	FileSize float64 `json:"file_size"`
}

// This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
}

// This object represents a point on the map.
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
	OneTimeKeyboard bool       `json:"one_time_keyboard,omitempty"`
	Selective       bool       `json:"selective,omitempty"`
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

type UserProfilePhotos struct {
	TotalCount float64          `json:"total_count"`
	Photos     *[]*[]*PhotoSize `json:"photos"`
}
