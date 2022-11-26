package main

import (
	"github.com/heqingbao/ginchat/router"
	"github.com/heqingbao/ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8080")
}
