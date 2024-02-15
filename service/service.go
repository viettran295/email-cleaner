package service

import (
	"email-cleaner/token"
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
)

var srv = tkn.NewService()
var user = "me"

func ListMessId() []*gmail.Message {
	listCall := srv.Users.Messages.List(user)
	listCall.Q("in:inbox label:INBOX")
	messages, err := listCall.Do()
	if err != nil {
		log.Fatalf("Error to get message ID: ", err)
	}
	return messages.Messages
}

func GetHeader(id string) error {
	mess, err := srv.Users.Messages.Get(user, id).Do()
	if err != nil {
		log.Fatalf("Error to get header: ", err)
		return err
	}
	headers := mess.Payload.Headers
	for _, header := range headers {
		fmt.Println(header.Name, header.Value)
	}
	return nil
}

func DeleteMess(id string) error {
	err := srv.Users.Messages.Delete(user, id).Do()
	if err != nil{
		log.Println("Error while deleting email: ", err)
		return err
	}
	return nil
}
