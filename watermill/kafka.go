package watermill

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"gitlab.com/shoesmart2.1/backend/packages/functioncaller"
	"gitlab.com/shoesmart2.1/backend/packages/logruslogger"
)

type Kafka struct {
	Brokers   []string
	Publisher message.Publisher
}

var (
	Logger = watermill.NewStdLogger(
		true,  // debug
		false, // trace
	)
)

func NewKafka(brokers []string) Kafka {
	return Kafka{Brokers: brokers}
}

//function to create kafka publisher
func (pub Kafka) CreatePublisher() (res message.Publisher, err error) {
	res, err = kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   pub.Brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		Logger,
	)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "kafka-createPublisher")
		return res, err
	}

	return res, nil
}

//function to create kafka subscriber
func (pub Kafka) CreateSubscriber(handlerName string) (res message.Subscriber, err error) {
	res, err = kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       pub.Brokers,
			Unmarshaler:   kafka.DefaultMarshaler{},
			ConsumerGroup: handlerName,
		},
		Logger,
	)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "kafka-createPublisher")
		return res, err
	}

	return res, nil
}

func(pub Kafka) PublishMessage(publishTopic string,payload map[string]interface{}) (err error){
	payloadByte,_ := json.Marshal(payload)
	message := message.NewMessage(watermill.NewUUID(),payloadByte)
	err = pub.Publisher.Publish(publishTopic,message)
	if err!= nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"kafka-publishQueue")
		return err
	}

	return nil
}
