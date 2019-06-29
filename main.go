package main

import (
	"inventory-toko-ijah/config"
	"inventory-toko-ijah/routes"
	"inventory-toko-ijah/utils/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	confReader := new(config.ConfigReader)
	confReader.Read()
	db.InitSqlite()
	defaultRoute := new(routes.DefaultRoute)
	if viper.GetString("mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	defaultRoute.Init(router)
	config := cors.DefaultConfig()
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	if viper.GetBool("secure") == false {
		router.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	} else {
		router.RunTLS(viper.GetString("server.host")+":"+viper.GetString("server.port"), viper.GetString("crt_file"), viper.GetString("key_file"))
	}
}
