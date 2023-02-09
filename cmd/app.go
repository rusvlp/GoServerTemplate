package main

import (
	config2 "CustomServerTemplate/internal/config"
	server2 "CustomServerTemplate/internal/server"
	"fmt"
)

func main() {
	reflectionTest()
	/*err := runApp()
	if err != nil {
		logrus.Error(err)
	} */
}

func runApp() error {

	config, err := config2.Initialize()
	if err != nil {
		return err
	}
	fmt.Println("Server is going to start on port " + config.Port)
	server := server2.New(config)
	err = server.Start()
	if err != nil {
		return err
	}
	return nil
	// Логгер, который журналирует в файл
}
