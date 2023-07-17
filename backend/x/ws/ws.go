package ws

import (
	"backend/modules/api/endpoints/messages"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

type WS struct {
	WS *websocket.Conn
}

func NewWS(c echo.Context) (*WS, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return nil, err
	}
	return &WS{
		WS: ws,
	}, nil
}

func (ws *WS) Close() error {
	return ws.WS.Close()
}

func (ws *WS) ReadWrappedMessage() (*messages.Wrapper, error) {
	_, wsMessage, err := ws.WS.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("read_message: %w", err)
	}
	wrapper := &messages.Wrapper{}
	if err := proto.Unmarshal(wsMessage, wrapper); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return wrapper, nil
}

func (ws *WS) WriteWrappedMessage(message *messages.Wrapper) error {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Error in marshal: %w", err)
	}
	if err := ws.WS.WriteMessage(websocket.BinaryMessage, messageBytes); err != nil {
		return fmt.Errorf("Error writing message to WS: %w", err)
	}
	return nil
}
