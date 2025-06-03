package main

import (
	"fmt"
	"os"
	"os/signal"
	"training/backend/services/database"
	webserver "training/backend/webserver"
	"training/package/log"
	"syscall"

	"training/backend/config"
)

func init() {
	path := config.LoggerPath()
	fmt.Println(path)
	log.SetOptions(
		log.Development(),
		log.WithCaller(true),
		log.WithLogDirs(path),
	)
}
func main() {

	database.Connect()
	defer database.Close()
	go webserver.StartWebserver()

	defer os.Exit(0)

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-stop
	fmt.Println("training frontend webserver is shutting down .... 👋 !")
	log.Infoln("training frontend webserver is shutting down .... 👋 !")

	go func() {
		<-stop
	}()
}
