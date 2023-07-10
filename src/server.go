package src

import (
	"fmt"
	"installment_back/controllers"
	"log"
	"net/http"
)

func WebServer() {
	var ctrl controllers.Controller
	ctrl.DataBaseModel.Connect("Installment_Front")
	url := "0.0.0.0:7777"
	fmt.Printf("Server started! %s/EndPoint", url)
	http.HandleFunc("/EndPoint", ctrl.Handler)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
	defer ctrl.DataBaseModel.Disconnect()
}
