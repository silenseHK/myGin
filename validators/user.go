package validators

type UserRegister struct{
	Phone string `form:"phone" json:"phone" binding:"required,min=8,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
	Verify string `form:"verify" json:"verify" binding:"required,len=6"`
}

type UserVerify struct{
	Phone string `form:"phone" json:"phone" binding:"required,min=8,max=15"`
}

type UserLogin struct{
	Phone string `form:"phone" json:"phone" binding:"required,min=8,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
}
