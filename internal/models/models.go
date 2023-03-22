package models

type Models interface {
	User | Auth
}

type User struct {
	Uuid        string  `json:"uuid"`
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Country     string  `json:"country"`
	Bio         *string `json:"bio"`
	BfpBlob     *string `json:"pfpBlob"`
	PfpMimeType *string `json:"pfpMimeType"`
	DateCreated string  `json:"dateCreated"`
}

type Auth struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Group struct {
	Uuid        string  `json:"uuid"`
	UserUuid    string  `json:"userUuid"`
	PrevItem    *string `json:"prevItem"`
	BelongsTo   string  `json:"belongsTo"`
	GroupName   string  `json:"groupName"`
	DateCreated string  `json:"dateCreated"`
}

type Note struct {
}
