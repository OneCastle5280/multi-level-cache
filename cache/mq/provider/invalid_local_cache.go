package provider

import (
	"context"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"mlc/cache/log"
	"mlc/cache/mq"
	"mlc/cache/mq/event"
)

// SendInvalidLocalCacheEvent
//
//	@Description:
//	@param event
func SendInvalidLocalCacheEvent(ctx context.Context, invalidLocalCacheEvent event.InvalidLocalCacheEvent) error {
	msg, err := json.Marshal(invalidLocalCacheEvent)
	if err != nil {
		log.Error("[SendInvalidLocalCacheEvent] marshal event: %v err", invalidLocalCacheEvent)
		return err
	}
	messageID, err := mq.GetProvider(mq.InvalidLocalCacheTopicName).Send(ctx, &pulsar.ProducerMessage{
		Payload: msg,
	})
	if err != nil {
		return err
	}
	log.Info("[SendInvalidLocalCacheEvent] success, msgID: %v", messageID)
	return nil
}
