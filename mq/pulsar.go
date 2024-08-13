package mq

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"mlc/cache/log"
	"sync"
)

var (
	pulsarClient pulsar.Client
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
