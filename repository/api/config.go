package api

import (
	"silence/models"
)

type ConfRep struct{

}

type RechargeConf struct{
	Id int64 `json:"id"`
	ChannelId int64 `json:"channel_id"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Buttons string `json:"buttons"`
	Status int `json:"status"`
	Title string `json:"title"`
}

func (this ConfRep)GetRechargeConf()(list []RechargeConf, err error){
	list = make([]RechargeConf,0)
	var a = models.PREFIX + "recharge_config"
	var b = models.PREFIX + "channels"
	var config RechargeConf
	rows,err := models.DB.Table(a + " AS a").Select("a.id,channel_id,min,max,buttons,b.title").Joins("left join "+b+" AS b on b.id = a.channel_id").Where("status = ?",1).Rows()
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		err := models.DB.ScanRows(rows,&config)
		if err != nil{
			return nil, err
		}
		list = append(list, config)
	}
	return list,err
}
