package main

import (
	"math/rand"
	"os"
)

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func isWritable(filepath string) bool {
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	file.Close()

	return true
}
