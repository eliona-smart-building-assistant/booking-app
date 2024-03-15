package apiservices

import (
	"booking-app/apiserver"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eliona-smart-building-assistant/go-utils/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// NewRouter creates a new router for any number of api routers
// This method is copied from generated apiserver/routers.go to add websocket endpoints
func NewRouter(routers ...apiserver.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for name, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = webSocketHandler(handler)
			handler = apiserver.Logger(handler, name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(name).
				Handler(handler)
		}
	}
	return router
}

func webSocketHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.RequestURI, "bookings-subscription") || strings.Contains(r.RequestURI, "bookings-subscription?") {
			listen(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type subscriber struct {
	conn *websocket.Conn
}

var assetSubscriptions = make(map[int32][]subscriber)

func listen(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("listener", "Error upgrading to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	// Create a cancelable context to abort the operation in the services method
	cancelCtx, cancelFunc := context.WithCancel(r.Context())

	// Read message and wait for close method
	go func() {
		for {
			mType, m, err := conn.ReadMessage()
			if mType == websocket.CloseMessage || isCloseError(err) {
				log.Debug("listener", "Close listener because of close message from WebSocket for %s", r.RequestURI)
				cancelFunc() // tells the services method to stop listening for data changes
				return
			}
			if err != nil {
				log.Error("listener", "Error reading message from WebSocket %s: %v", r.RequestURI, err)
				cancelFunc() // tells the services method to stop listening for data changes
				return
			}

			if mType == websocket.TextMessage {
				var request apiserver.SubscribeBookingsRequest
				if err := json.Unmarshal(m, &request); err != nil {
					log.Error("listener", "Error parsing SubscribeBookingsRequest: %v", err)
					continue // Skip processing this message
				}

				for _, assetID := range request.AssetIDs {
					assetSubscriptions[assetID] = append(assetSubscriptions[assetID], subscriber{conn})
				}
			}
		}
	}()

	// Listening for data changes and send message
	msgChan := make(chan []byte)

	// Wait for data changes or until the cancelable context is canceled
	log.Debug("listener", "Start listener for %s", r.RequestURI)

	// Wait for messages and write the message to WebSocket
	for {
		select {
		case msg := <-msgChan:
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Error("listener", "Error writing message to WebSocket %s: %v", r.RequestURI, err)
				cancelFunc()
				return
			}
		case _ = <-cancelCtx.Done():
			return
		}
	}
}

func isCloseError(err error) bool {
	if _, ok := err.(*websocket.CloseError); ok {
		return true
	}
	return false
}

var messageChannelContextKey = &contextKey{
	Name: "message-channel",
}

type contextKey struct {
	Name string
}

// todo: rm?
func getMessageChannelFromContext(ctx context.Context) (chan []byte, error) {
	value := ctx.Value(messageChannelContextKey)
	if value == nil {
		return nil, fmt.Errorf("no value found in context")
	}
	msgChan, ok := value.(chan []byte)
	if !ok {
		return nil, fmt.Errorf("unexpected type in context for key '%v': got %T with value %v, want chan []byte", messageChannelContextKey, value, value)
	}
	return msgChan, nil
}

type ResponseWriterWrapper struct {
	Data       []byte
	StatusCode int
}

func (l *ResponseWriterWrapper) Write(data []byte) (int, error) {
	l.Data = make([]byte, len(data))
	copy(l.Data, data)
	return len(data), nil
}

func (l *ResponseWriterWrapper) Header() http.Header {
	return http.Header{}
}

func (l *ResponseWriterWrapper) WriteHeader(statusCode int) {
	l.StatusCode = statusCode
}
