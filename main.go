package main 

import (
	"flag"
	"log"

	"github.com/Moksh-10/redis/config"
	"github.com/Moksh-10/redis/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.Port, "port", 7379, "port for the server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("rolling")
	server.RunSyncTCPServer()
}