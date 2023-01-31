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

	config, err := config2.Initialize()
	fmt.Println("Server is going to start on port " + config.Port)
	server := server2.New(config)
	server.Start()
	if err != nil {
		log.Fatalln(err)
	}

	// Логгер, который журналирует в файл
}
