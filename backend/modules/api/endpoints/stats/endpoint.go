package stats

import (
	"backend/modules/api/endpoints/messages"
	"backend/x/ws"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func EndpointLatestStatsPush(logger *zap.Logger, dbConn *gorm.DB, ws *ws.WS, pushInterval time.Duration) error {
	wrappedMessage, err := ws.ReadWrappedMessage()
	if err != nil {
		return fmt.Errorf("read_wrapped_message: %w", err)
	}
	pushRequestWrapped, ok := wrappedMessage.Message.(*messages.Wrapper_LatestStatsPushRequest)
	if !ok {
		return fmt.Errorf("Bad message type received")
	}
	request := pushRequestWrapped.LatestStatsPushRequest

	for {
		response := handleLatestStatsPushRequest(logger, dbConn, request)
		wrappedResponse := &messages.Wrapper{
			Message: &messages.Wrapper_LatestStatsPushResponse{
				LatestStatsPushResponse: response,
			},
		}
		if err := ws.WriteWrappedMessage(wrappedResponse); err != nil {
			return fmt.Errorf("write_wrapped_message: %w", err)
		}
		time.Sleep(pushInterval)
	}
}
