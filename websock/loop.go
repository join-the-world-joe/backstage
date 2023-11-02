package main

import (
	"backstage/common/payload"
	"backstage/global/crypto"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

func getClient(userId, port string, encryption bool) *websocket.Conn {
	_encryption = encryption
	url := fmt.Sprintf("ws://127.0.0.1:%s/ws", port)
	header := make(http.Header)

	_userId = userId

	m := map[string]string{
		"user_id": userId,
	}

	for k, v := range m {
		header.Add(k, v)
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		panic(err)
	}
	return conn
}

func loop(userId, port string, encryption bool) error {
	conn := getClient(userId, port, encryption)
	out := make(chan string)
	notify := make(chan bool)

	go console(out, notify)

	go recvRoutine(conn)

	for {
		select {
		case str := <-out:
			msg := strings.TrimSuffix(str, "\r\n")
			if strings.Compare(msg, "exit") == 0 {
				notify <- true
				conn.Close()
				return nil
			} else {
				fmt.Println("console msg: ", string(msg))
				switch msg {
				case "Auth":
					auth(conn)
				case "Ping":
					ping(conn)
				case "Repeat":
					repeat(conn)
				case "Register":
					register(conn)
				case "SendSMSRegister":
					sendVerificationCodeRegister(conn)
				case "SendSMSLogin":
					sendVerificationCodeLogin(conn)
				case "Login":
					login(conn)
				case "Token":
					token(conn)
				default:
					conn.WriteMessage(websocket.BinaryMessage, []byte(msg))
				}
			}
		}
	}
}

func recvRoutine(conn *websocket.Conn) {

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}

		fmt.Println("msg: ", string(msg))

		plainText := msg
		if _encryption {
			plainText, err = crypto.AESDecrypt(msg)
			if err != nil {
				panic(err)
			}
		}

		packet := new(payload.PacketClient)
		err = json.Unmarshal(plainText, packet)
		if err != nil {
			panic(err)
		}

		fmt.Println("Receive from Major: ", packet.Header.Major)
		fmt.Println("Receive from Minor: ", packet.Header.Minor)
		fmt.Println("Message: ", string(packet.Body))
	}
}

func send(conn *websocket.Conn, major, minor, message string) {
	packet := &payload.PacketClient{
		Header: &payload.Header{
			Major: major,
			Minor: minor,
		},
		Body: []byte(message),
	}

	content, err := packet.ToBytes()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}

	if _encryption {
		content, err = encrypt(content)
		if err != nil {
			panic(err)
		}
	}

	err = conn.WriteMessage(websocket.BinaryMessage, content)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println("Send message successfully.")
}
