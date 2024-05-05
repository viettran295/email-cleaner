package service

import (
	"email-cleaner/token"
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

func GetSender(id string) string {
	mess, err := srv.Users.Messages.Get(user, id).Do()
	if err != nil {
		log.Fatalf("Error to get header: ", err)
	}
	headers := mess.Payload.Headers
	for _, header := range headers {
		if header.Name == "From"{
			return header.Value
		}
	}
	return ""
}

func DeleteMess(id string) (*gmail.Message, error)  {
	resp, err := srv.Users.Messages.Trash(user, id).Do()
	if err != nil{
		log.Println("Error while deleting email: ", err)
		return resp, err
	}
	return resp, nil
}
