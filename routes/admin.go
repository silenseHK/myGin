package routes

import "silence/controllers/admin"

type AdminRouter struct{

}

var (
	AdminLoginCon admin.Login
	AdminDashboard admin.Dashboard
	AdminAdminCon admin.Admin
)

func (this AdminRouter)Register(){
	g := r.Group("/admin")

	gLogin := g.Group("/login")
	gLogin.GET("/",AdminLoginCon.Index)
	gLogin.POST("/",AdminLoginCon.Login)

	gDashboard := g.Group("/dashboard")
	gDashboard.GET("/",AdminDashboard.Index)

	gAdmin := g.Group("/admin")
	gAdminAccount := gAdmin.Group("/account")
	gAdminAccount.GET("/",AdminAdminCon.AccountList)

}
