package main

import (
	"ftpLeaf/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt)

	engin := gin.Default()
	l, err := net.Listen("tcp", ":9595")
	if err != nil {
		log.Fatal(err)
	}

	controller.Register(engin)

	go func() {
		err := engin.RunListener(l)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-shutdown
}
