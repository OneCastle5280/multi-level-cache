package mq

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"mlc/cache/log"
)

// SendInvalidLocalCacheEvent
//
//	@Description: send invalid local cache event
//	@param event
func SendInvalidLocalCacheEvent(ctx context.Context, invalidLocalCacheEvent InvalidLocalCacheEvent) error {
	msg, err := json.Marshal(invalidLocalCacheEvent)
	if err != nil {
		log.Error("[SendInvalidLocalCacheEvent] marshal event: %v err", invalidLocalCacheEvent)
		return err
	}
	messageID, err := getProvider(InvalidLocalCacheTopicName).Send(ctx, &pulsar.ProducerMessage{
		Payload: msg,
	})
	if err != nil {
		return err
	}
	log.Info("[SendInvalidLocalCacheEvent] success, msgID: %v", messageID)
	return nil
}
