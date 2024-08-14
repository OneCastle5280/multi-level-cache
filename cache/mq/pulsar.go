package mq

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"mlc/cache/log"
	"sync"
)

var (
	pulsarClient pulsar.Client
	providerMap  sync.Map
	consumerMap  sync.Map
	once         sync.Once
)

// InitPulsarClient
//
//	@Description: 创建客户端
func InitPulsarClient(options pulsar.ClientOptions) {
	once.Do(func() {
		client, err := pulsar.NewClient(options)
		if err != nil {
			log.Error(">>>> InitPulsarClient err %v", err)
			panic(err)
		}
		// init client
		pulsarClient = client
	})
}

// GetPulsarClient
//
//	@Description:  get pulsar client
//	@return pulsar.Client
func GetPulsarClient() pulsar.Client {
	return pulsarClient
}

// InitProvider
//
//	@Description: init provider
//	@param topicName
func InitProvider(topicName string) {
	producer, err := GetPulsarClient().CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
	})
	if err != nil {
		log.Error(">>>> Init Pulsar provider err %v", err)
		panic(err)
	}
	providerMap.Store(topicName, producer)
}

// InitConsumer
//
//	@Description: init consumer
//	@param topicName
func InitConsumer(topicName string, subscriptionName string) {
	consumer, err := GetPulsarClient().Subscribe(pulsar.ConsumerOptions{
		SubscriptionName: subscriptionName,
		Topic:            topicName,
	})
	if err != nil {
		log.Error(">>>> Init Pulsar consumer err %v", err)
		panic(err)
	}
	consumerMap.Store(topicName, consumer)
}

// GetProvider
//
//	@Description:
//	@param topicName
//	@return pulsar.Producer
func GetProvider(topicName string) pulsar.Producer {
	if producer, ok := providerMap.Load(topicName); ok {
		if p, ok := producer.(pulsar.Producer); ok {
			return p
		}
	}
	return nil
}
