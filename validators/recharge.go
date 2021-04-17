package validators

type RechargeRecharge struct{
	ChannelId int64 `form:"channel_id" binding:"required,gte=1"`
	Money float64 `form:"money" binding:"required,gte=1"`
	Method int8 `form:"method" binding:"required"`
}
