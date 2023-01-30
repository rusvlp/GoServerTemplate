package main

import (
	config2 "CustomServerTemplate/internal/config"
	server2 "CustomServerTemplate/internal/server"
	"fmt"
	"log"
)

func main() {
	runApp()
}

func runApp() {

	fmt.Println("Server is going to start...")
	config, err := config2.Initialize()
	server := server2.New(config)
	server.Start()
	if err != nil {
		log.Fatalln(err)
	}

	// Логгер, которые журналирует в файл
}
