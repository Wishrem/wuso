package websocket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/Wishrem/wuso/client/message"
	"github.com/Wishrem/wuso/client/screen"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/yitter/idgenerator-go/idgen"
)

var ws *_websocket

type _websocket struct {
	dialer    *websocket.Dialer
	conn      *websocket.Conn
	url       string
	recvBytes chan []byte
}

const mask = 0b1111

var endian binary.ByteOrder

func init() {
	var x uint16 = 0x0102
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, x)
	if b[0] == 0x02 && b[1] == 0x01 {
		endian = binary.LittleEndian
	} else {
		endian = binary.BigEndian
	}

	sender = new(sendReq)
}

func Send(data []byte) bool {
	return sender.Send(data)
}

func Connect() {
	if ws == nil {
		u := &url.URL{
			Scheme: "wss",
			Host:   "localhost:8080",
			Path:   "/wuso/ws",
		}
		ws = &_websocket{
			dialer: &websocket.Dialer{},
			url:    u.String(),
		}
	}

	if ws.conn != nil {
		reconnect(3)
		return
	}

	if !connect() {
		screen.Notice("错误", "暂时无法连接至服务器", screen.NoticeWrong)
		return
	}
	defer ws.conn.Close()
	go recvMsg()

	intr := make(chan os.Signal, 1)
	signal.Notify(intr, os.Interrupt)

	for {
		select {
		case <-intr:
			log.Println("收到中断消息，关闭连接")
			err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("关闭连接失败：", err)
			}
			return
		case msg := <-ws.recvBytes:
			if len(msg) < 1 {
				ws.SendErr(ErrRecvWrongBytes)
				continue
			}

			header := uint8(msg[0])
			dt := header & (mask << 4)
			code := header & mask
			if dt^ACKType == 0 {
				sender.ACKhandler(msg[1:])
			} else if dt^CmdType == 0 {
				message.RecvCmdCode(code)
			} else if dt^MsgType == 0 {
				message.RecvMsg(msg[1:])
			} else {
				ws.SendErr(ErrRecvWrongBytes)
			}
		}
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
		ws.SendErr(ErrRecvWrongBytes)
	} else {
		log.Println("Error reading message:", err)
		ws.SendErr(ErrRecvWrongBytes)
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
			ws.SendErr(ErrRecvWrongBytes)
			continue
		}
		ws.recvBytes <- msg
	}
}

func connect() bool {
	screen.Notice("提示", "正在尝试连接服务器...", screen.NoticeAttention)
	conn, _, err := ws.dialer.Dial(ws.url, nil)
	if err != nil {
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
		conn, _, err := ws.dialer.Dial(ws.url, nil)
		if err != nil {
			screen.Notice("提示", fmt.Sprintf("正在尝试第%v次连接", i+1), screen.NoticeAttention)
			continue
		}
		ws.conn = conn
		break
	}
	return success
}

func (ws *_websocket) SendErr(errBytes []byte) {
	ws.conn.WriteMessage(websocket.BinaryMessage, errBytes)
}

var sender *sendReq

type sendReq struct {
	respId chan int64
}

var json = jsoniter.ConfigFastest

type tx struct {
	Id   int64  `json:"id"`
	Data []byte `json:"data"`
}

func (s *sendReq) ACKhandler(data []byte) {
	if len(data) != 8 {
		ws.SendErr(ErrRecvWrongBytes)
		return
	}
	var value int64
	if err := binary.Read(bytes.NewReader(data), endian, &value); err != nil {
		ws.SendErr(ErrRecvWrongBytes)
		return
	}
	s.respId <- value
}

func (s *sendReq) Send(data []byte) bool {
	tx := &tx{
		Id:   idgen.NextId(),
		Data: data,
	}
	bytes, err := json.Marshal(tx)
	if err != nil {
		log.Println("Marshal data failed:", err)
		return false
	}

	for {
		if err := ws.conn.WriteMessage(websocket.BinaryMessage, bytes); err != nil {
			if !handleSendErr(err) {
				return false
			}
		} else {
			break
		}
	}

	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-timer.C:
			return false
		case id := <-s.respId:
			if id == tx.Id {
				return true
			}
		}
	}
}
