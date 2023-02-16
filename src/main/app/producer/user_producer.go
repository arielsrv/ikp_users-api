package producer

import (
	"context"
	"fmt"
	"github.com/src/main/app/aws"
	"github.com/src/main/app/config"
	"log"
	"time"
)

type UserProducer struct {
	client aws.SNS
	name   string
}

func NewUserProducer() *UserProducer {
	session, err := aws.New(aws.Config{
		Address: config.String("topics.users.address"),
		Region:  config.String("topics.users.region"),
		Profile: config.String("topics.users.profile"),
		ID:      config.String("topics.users.id"),
		Secret:  config.String("topics.users.secret"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	client := aws.NewSNS(session, time.Second*5)

	return &UserProducer{
		client: client,
		name:   fmt.Sprintf("arn:aws:sns:us-east-1:000000000000:%s", config.String("topics.users.name")),
	}
}

func (userProducer *UserProducer) Send(userID int64) {
	userProducer.client.
		Publish(context.Background(), string(userID), userProducer.name)
}
