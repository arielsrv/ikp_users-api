package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"time"
)

type Topic struct {
	ARN string
}

type UserProducer interface {
	Publish(ctx context.Context, message, topicARN string) (string, error)
}

type Subscription struct {
	ARN      string
	TopicARN string
	Endpoint string
	Protocol string
}
type SNS struct {
	timeout time.Duration
	client  *sns.SNS
}

type Config struct {
	Address string
	Region  string
	Profile string
	ID      string
	Secret  string
}

func New(config Config) (*session.Session, error) {
	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(config.ID, config.Secret, ""),
				Region:           aws.String(config.Region),
				Endpoint:         aws.String(config.Address),
				S3ForcePathStyle: aws.Bool(true),
			},
			Profile: config.Profile,
		},
	)
}

func NewSNS(session *session.Session, timeout time.Duration) SNS {
	return SNS{
		timeout: timeout,
		client:  sns.New(session),
	}
}

func (s SNS) Publish(ctx context.Context, message, topicARN string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	publishOutput, err := s.client.PublishWithContext(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: aws.String(topicARN),
	})

	if err != nil {
		return "", fmt.Errorf("publish: %w", err)
	}

	return *publishOutput.MessageId, nil
}
