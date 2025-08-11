package main

import (
	"math/rand"
	"time"
)


const shortKeyLength = 6
var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateShortKey() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, shortKeyLength)
	for i:= range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}