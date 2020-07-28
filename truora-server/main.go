package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"truora-server/middleware"
	"truora-server/resources"
	"truora-server/services"
)


func main() {
	db := resources.Connect()
	domain := services.ServiceDomain{Database:db}

	router := fasthttprouter.New()
	router.GET("/info/:domain", domain.GetInfo)
	router.GET("/history", domain.GetHistory)
	defer db.Conn.Close()
	log.Fatal(fasthttp.ListenAndServe(":5500", middleware.CORS(router.Handler)))
}