package main

import (
	"net/mail"
	"os"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validOrderID(orderID string) bool {
	if len(orderID) != 14 {
		return false
	}
	return true
}


func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}