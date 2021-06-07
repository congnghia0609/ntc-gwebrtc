package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"

	"github.com/pion/signaler"
)

type SignalServer struct {
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *SignalServer) AuthenticateRequest(params url.Values) (apiKey, room, sessionKey string, ok bool) {
	return "ABC", "ABC", randSeq(16), true
}

func (s *SignalServer) OnClientMessage(ApiKey, Room, SessionKey string, raw []byte) {
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "443"
	}
	fmt.Println(signaler.Start(&SignalServer{}, port))
}