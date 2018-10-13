package model

type User struct {
	UserName string `bson:"username" json:"username"`
}

type WechatModel struct {
	Code string `json:"code"`
	CryptedData string `json:"crypted_data"`
	Iv string `json:"iv"`
}