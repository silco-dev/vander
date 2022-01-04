package structs

type User struct {
	Token   string    `bson:"token" json:"token,omitempty"`
	Admin   bool      `bson:"admin" json:"admin,omitempty"`
	Enabled bool      `bson:"enabled" json:"enabled,omitempty"`
	Info    *UserInfo `bson:"info" json:"info,omitempty"`
}

type UserInfo struct {
	Name    string `bson:"name,omitempty" json:"name,omitempty"`
	Contact string `bson:"contact,omitempty" json:"contact,omitempty"`
	Iat     int32  `bson:"iat,omitempty" json:"iat,omitempty"`
}
