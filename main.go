package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("smsapi")

	Config_Init()
	Modem_Init()
	Routes_Init()

	log.Printf("Connected to %s", ModemAPI_GetCarrier())

	log.Printf("Listening on port %d", config.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil)
}