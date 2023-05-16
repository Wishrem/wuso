package transaction

import (
	"fmt"
	"io"
	"log"

	"github.com/Wishrem/wuso/client/screen"
	"github.com/gorilla/websocket"
)

var ws *_websocket

type _websocket struct {
	dialer    *websocket.Dialer
	conn      *websocket.Conn
	url       string
	recvBytes chan []byte
}

func Close() {
	log.Println("received interruption signal, closing the connection")
	err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("failed to closing connection", err)
	}
}

func handleRecvErr(err error) {
	if err == io.EOF || websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		log.Println("Close Error:", err)
		if !reconnect(3) {
			screen.Notice("错误", "暂时无法连接至服务器", screen.NoticeWrong)
		}
	} else if err == websocket.ErrReadLimit {
		log.Println("Received message is too big")
		// TODO
	} else {
		log.Println("Error reading message:", err)
		// TODO
	}
}

func handleSendErr(err error) bool {
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		log.Println("Close Error:", err)
		if !reconnect(3) {
			screen.Notice("错误", "暂时无法连接至服务器", screen.NoticeWrong)
			return false
		}
		return true
	} else {
		log.Println("Error sending message:", err)
		return false
	}
}

func recvMsg() {
	for {
		t, msg, err := ws.conn.ReadMessage()
		if err != nil {
			handleRecvErr(err)
			continue
		}

		if t != websocket.BinaryMessage {
			// TODO
			continue
		}
		ws.recvBytes <- msg
	}
}

func connect() bool {
	log.Println("connecting the server...")
	screen.Notice("提示", "正在尝试连接服务器...", screen.NoticeAttention)
	conn, _, err := ws.dialer.Dial(ws.url, nil)
	if err != nil {
		log.Println("connect error:", err)
		return reconnect(3)
	}
	ws.conn = conn
	return true
}

func reconnect(times int) bool {
	if times <= 0 {
		times = 1
	}
	if times > 5 {
		times = 5
	}

	success := false
	for i := 0; i < times; i++ {
		log.Println("try to reconnect ", i+1)
		conn, _, err := ws.dialer.Dial(ws.url, nil)
		if err != nil {
			log.Println("reconnect error:", err)
			screen.Notice("提示", fmt.Sprintf("正在尝试第%v次连接", i+1), screen.NoticeAttention)
			continue
		}
		success = true
		ws.conn = conn
		break
	}
	return success
}
