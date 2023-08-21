package websocket

import (
	"cryptocurrency/currency"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return ws, err

	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Printf("Updating price: %+v\n", t)

			items, err := currency.GetCurrency()
			if err != nil {
				fmt.Println(err)
			}

			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}
