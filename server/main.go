package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"truora-server/config"
	"truora-server/middleware"
	"truora-server/resources"
	"truora-server/services"
)


func main() {
	config.ConfigEnvVariables()
	db := resources.Connect()
	domain := services.ServiceDomain{Database:db}

	router := fasthttprouter.New()
	router.GET("/info/:domain", domain.GetInfo)
	router.GET("/history", domain.GetHistory)
	defer db.Conn.Close()

	port, ok := viper.Get("PORT").(string)
	if !ok {
		port = "5500"
	}
	log.Fatal(fasthttp.ListenAndServe(":" + port, middleware.CORS(router.Handler)))
}