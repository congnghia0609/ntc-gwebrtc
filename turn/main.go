package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/pion/logging"
	"github.com/pion/turn"
)

func createAuthHandler() turn.AuthHandler {
	return func(username string, srcAddr net.Addr) (string, bool) {
		return "password", true
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	realm := os.Getenv("REALM")
	if realm == "" {
		// log.Panic("REALM is a required environment variable")
		realm = "localhost"
	}

	udpPortStr := os.Getenv("UDP_PORT")
	if udpPortStr == "" {
		udpPortStr = "3478"
	}
	udpPort, err := strconv.Atoi(udpPortStr)
	if err != nil {
		log.Panic(err)
	}

	var channelBindTimeout time.Duration
	channelBindTimeoutStr := os.Getenv("CHANNEL_BIND_TIMEOUT")
	if channelBindTimeoutStr != "" {
		channelBindTimeout, err = time.ParseDuration(channelBindTimeoutStr)
		if err != nil {
			log.Panicf("CHANNEL_BIND_TIMEOUT=%s is an invalid time Duration", channelBindTimeoutStr)
		}
	}

	s := turn.NewServer(&turn.ServerConfig{
		Realm:              realm,
		AuthHandler:        createAuthHandler(),
		ChannelBindTimeout: channelBindTimeout,
		ListeningPort:      udpPort,
		LoggerFactory:      logging.NewDefaultLoggerFactory(),
		Software:           os.Getenv("SOFTWARE"),
	})

	err = s.Start()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("TURN server is running on port: %s", udpPortStr)

	<-c
	err = s.Close()
	if err != nil {
		log.Panic(err)
	}
}
