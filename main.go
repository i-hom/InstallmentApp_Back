package main

import "installment_back/src"

func main() {
	var server src.Server

	server.Init()
	server.Run()
}
