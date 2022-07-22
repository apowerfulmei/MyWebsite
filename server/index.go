package main

import (
	"MY-WEB/server/admin"
	"MY-WEB/server/db"
	"MY-WEB/server/models"
	"MY-WEB/server/user"
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

var options = []Option{}

func getEngine(opts ...Option) {
	options = append(options, opts...)
}

func arrange() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}

func main() {
	db.Get_init()
	localIP := models.GetOutBoundIP()
	getEngine(admin.Admin, user.User, user.RegUser, user.LoginUser)
	r := arrange()
	r.Run(localIP + ":3000")
}
