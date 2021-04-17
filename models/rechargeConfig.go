package models

type RechargeConfig struct{
	Id int64 `json:"id"`
	ChannelId int64 `json:"channel_id"`
	Channel Channels `json:"channel" gorm:"foreignkey:ChannelId"`
	MerchantId string `json:"merchant_id"`
	Key string `json:"key"`
	PrivateKey string `json:"private_key"`
	PublicKey string `json:"public_key"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Buttons string `json:"buttons"`
	Status int `json:"status"`
}
