package types

type Update struct {
	UpdateId float64  `json:"update_id"`
	Message  *Message `json:"message"`
}

type User struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
}

type GroupChat struct {
	Id    float64 `json:"id"`
	Title string  `json:"title"`
}

type Chat struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	Title     string  `json:"title"`
}

type Message struct {
	Message_Id  int64   `josn:"message_id"`
	From        *User   `json:"from"`
	Date        float64 `json:"date"`
	Chat        *Chat   `json:"chat"`
	ForwardFrom *User   `json:"forward_from,omitempty"`
	ForwardDate float64 `json:"forward_date,omitempty"`
	//ReplyToMessage Message     `json:"reply_to_message"`
	Text     string        `json:"text,omitempty"`
	Audio    *Audio        `json:"audio,omitempty"`
	Document *Document     `json:"document,omitempty"`
	Photo    *[]*PhotoSize `json:"photo,omitempty"`
	Sticker  *Sticker      `json:"sticker,omitempty"`
	Video    *Video        `json:"video,omitempty"`
	Voice    *Voice        `json:"voice,omitempty"`
	Caption  string        `json:"caption,omitempty"`
	Contact  *Contact      `json:"contact,omitempty"`
	Location *Location     `json:"location,omitempty"`
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
