package business

import (
	"backstage/common/code"
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/service/account"
	"backstage/common/service/backend"
	"backstage/common/service/sms"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/service/gateway/runtime"
	"backstage/utils/convert"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

// sofar, Agent has not initialized yet;
// since the receive routine and session not readly yet, the upstream(Forward) and downstream(P2P) doesn't work here
func Logon(c *gin.Context, conn *websocket.Conn) (*payload.Session, error) {
	for {

		if runtime.FeedbackEnable() { // TODO: 后端用户除外
			err := _Feedback(conn)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
			time.Sleep(time.Second * time.Duration(runtime.FeedbackWaitForClose()))
			return nil, errors.New("closed by feedback")
		}

		conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(runtime.WebsocketAuthReadDeadline())))

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		log.Debug("Message: ", convert.Bytes2StringArray(message))

		plainText := message
		if runtime.EncryptionEnable() {
			plainText, err = crypto.AESDecrypt(message)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
		}

		log.Debug("PlainText(Bytes): ", convert.Bytes2StringArray(plainText))
		log.Debug("PlainText(String): ", string(plainText))

		req := &payload.PacketClient{}
		err = json.Unmarshal(plainText, req)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		ma := req.GetHeader().GetMajor()
		mi := req.GetHeader().GetMinor()

		if ma == major.Account && mi == account.LoginReq_ {
			session, err := _Login(req, conn)
			if err == nil && session != nil {
				return session, nil
			}
		} else if ma == major.Account && mi == account.RegisterReq_ {
			err = _Register(req, conn)
			if err != nil {
				log.Error("Logon._Register failure, err:", err.Error())
				continue
			}
		} else if ma == major.Backend && mi == backend.SignInReq_ {
			session, err := _SignIn(req, conn)
			if err == nil && session != nil {
				return session, nil
			}
		} else if ma == major.SMS && mi == sms.SendVerificationCodeReq_ {
			err = _SendVerificationCode(req, conn)
			if err != nil {
				log.Error("Logon._SendVerificationCode failure, err:", err.Error())
				continue
			}
		}
	}
}

func _Feedback(conn *websocket.Conn) error {
	feedback := &payload.PacketClient{
		Header: &payload.Header{
			Major: runtime.FeedbackMajor(),
			Minor: runtime.FeedbackMinor(),
		},
		Body: []byte(runtime.FeedbackMessage()),
	}
	bytes, err := json.Marshal(feedback)
	if err != nil {
		return err
	}
	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			return err
		}
	}
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return err
	}
	return nil
}

func _SendVerificationCode(req *payload.PacketClient, conn *websocket.Conn) error {
	businessReq := &sms.SendVerificationCodeReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		return err
	}
	businessRsp := &sms.SendVerificationCodeRsp{}
	err = sms.SendVerificationCode(context.Background(), businessReq, businessRsp)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		return err
	}
	rsp := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.SMS,
			Minor: sms.SendVerificationCodeRsp_,
		},
		Body: bytes,
	}
	bytes, err = json.Marshal(rsp)
	if err != nil {
		return err
	}
	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			return err
		}
	}
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return err
	}
	return nil
}

func _Register(req *payload.PacketClient, conn *websocket.Conn) error {
	businessReq := &account.RegisterReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		return err
	}
	businessRsp := &account.RegisterRsp{}
	err = account.Register(context.Background(), businessReq, businessRsp)
	if err != nil {
		log.Error("_Register.account.Register.failure err:", err.Error())
		return err
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		log.Error("_Register.json.Marshal.failure err:", err.Error())
		return err
	}
	rsp := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Account,
			Minor: account.RegisterRsp_,
		},
		Body: bytes,
	}
	bytes, err = json.Marshal(rsp)
	if err != nil {
		log.Error("_Register.json.Marshal.failure err:", err.Error())
		return err
	}
	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			log.Error("_Register.crypto.AESEncrypt.failure err:", err.Error())
			return err
		}
	}
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Error("_Register.conn.WriteMessage.failure err:", err.Error())
		return err
	}
	return nil
}

func _Login(req *payload.PacketClient, conn *websocket.Conn) (*payload.Session, error) {
	businessReq := &account.LoginReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		return nil, err
	}
	businessRsp := &account.LoginRsp{}
	err = account.Login(context.Background(), businessReq, businessRsp)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		return nil, err
	}
	rsp := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Account,
			Minor: account.LoginRsp_,
		},
		Body: bytes,
	}
	bytes, err = json.Marshal(rsp)
	if err != nil {
		return nil, err
	}
	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			return nil, err
		}
	}
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return nil, err
	}

	if businessRsp.Code != code.Success {
		return nil, nil
	}

	session := &payload.Session{}
	session.SetToken(businessRsp.Token)
	session.SetUserId(businessRsp.UserId)

	return session, nil
}

func _SignIn(req *payload.PacketClient, conn *websocket.Conn) (*payload.Session, error) {
	businessReq := &backend.SignInReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		return nil, err
	}
	businessRsp := &backend.SignInRsp{}
	err = backend.SignIn(context.Background(), businessReq, businessRsp)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		return nil, err
	}
	rsp := &payload.PacketClient{
		Header: &payload.Header{
			Major: major.Backend,
			Minor: backend.SignInRsp_,
		},
		Body: bytes,
	}
	bytes, err = json.Marshal(rsp)
	if err != nil {
		return nil, err
	}
	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			return nil, err
		}
	}
	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return nil, err
	}

	if businessRsp.Code != code.Success {
		return nil, nil
	}

	session := &payload.Session{}
	session.SetRole(businessRsp.Role)
	//session.SetToken(businessRsp.Token)
	session.SetUserId(businessRsp.UserId)

	return session, nil
}
