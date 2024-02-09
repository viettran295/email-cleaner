package main

import (
	"email-cleaner/service"
	"fmt"
)

func main() {
	mess := service.ListMessId()
	for _, msg := range mess {
		fmt.Println(msg.Id)
		service.GetHeader(msg.Id)
	}
}
