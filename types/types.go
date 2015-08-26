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
	Id float64 `json:"id"`
}

type Message struct {
	MessageId   float64 `josn:"message_id"`
	From        *User   `json:"from"`
	Date        float64 `json:"date"`
	Chat        *Chat   `json:"chat"`
	ForwardFrom *User   `json:"forward_from"`
	ForwardDate float64 `json:"forward_date"`
	//ReplyToMessage Message     `json:"reply_to_message"`
	Text     string        `json:"text"`
	Audio    *Audio        `json:"audio"`
	Document *Document     `json:"document"`
	Photo    *[]*PhotoSize `json:"photo"`
	Sticker  *Sticker      `json:"sticker"`
	Video    *Video        `json:"video"`
	Voice    *Voice        `json:"voice"`
	Caption  string        `json:"caption"`
	Contact  *Contact      `json:"contact"`
	Location *Location     `json:"location"`
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
	FileSize  string  `json:"file_size"`
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
	FileSize string  `json:"file_size"`
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
