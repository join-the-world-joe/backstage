package business

import (
	"backstage/common/code"
	"backstage/common/db/mgo/server/track"
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/account"
	"backstage/common/protocol/frontend_gateway"
	"backstage/common/protocol/gateway"
	"backstage/common/protocol/sms"
	"backstage/global/config"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/global/mgo"
	"backstage/service/frontend_gateway/runtime"
	"backstage/utils/convert"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"time"
)

// sofar, Agent has not initialized yet;
// since the receive routine and session not readly yet, the upstream(Forward) and downstream(P2P) doesn't work here
func Logon(c *gin.Context, conn *websocket.Conn) (*payload.Session, error) {
	for {

		//if runtime.FeedbackEnable() {
		//	err := _Feedback(conn)
		//	if err != nil {
		//		log.Error(err.Error())
		//		return nil, err
		//	}
		//	time.Sleep(time.Second * time.Duration(runtime.WaitForCloseInterval()))
		//	return nil, errors.New("closed by feedback")
		//}

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
			} else {
				log.Error("Logon._Login failure, err: ", err)
			}
		} else if ma == major.SMS && mi == sms.SendVerificationCodeReq_ {
			sendSMS(req, conn)
		} else if ma == major.FrontendGateway && mi == frontend_gateway.FetchRateLimitingConfigReq_ {
			fetchRateLimitingConfig(req, conn)
		} else if ma == major.FrontendGateway && mi == frontend_gateway.PingReq_ {
			echo(req, conn)
		}
	}
}

func _Downstream(packet *payload.PacketClient, conn *websocket.Conn) error {
	bytes, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	log.Debug(fmt.Sprintf("Downstream.PlainText: [%s]", string(bytes)))

	if runtime.EncryptionEnable() {
		bytes, err = crypto.AESEncrypt(bytes)
		if err != nil {
			return err
		}
	}

	log.Debug(fmt.Sprintf("Downstream.Bytes: [%v]", convert.Bytes2StringArray(bytes)))

	err = conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return err
	}

	return nil
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

	log.Debug(string(bytes))

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

func sendSMS(req *payload.PacketClient, conn *websocket.Conn) {
	businessReq := &sms.SendVerificationCodeReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		log.ErrorF("sendSMS.json.Unmarshal failure, body: %s err: %v",
			string(req.GetBody()),
			err.Error(),
		)
		return
	}
	businessRsp := &sms.SendVerificationCodeRsp{}
	err = sms.SendVerificationCode(context.Background(), businessReq, businessRsp)
	if err != nil {
		log.ErrorF("sendSMS.sms.SendVerificationCode failure(Behavior-%v, CountryCode-%v, PhoneNumber-%v), err: %v",
			businessReq.Behavior,
			businessReq.CountryCode,
			businessReq.PhoneNumber,
			err.Error(),
		)
		return
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		log.ErrorF("sendSMS.json.Marshal failure(Behavior-%v, CountryCode-%v, PhoneNumber-%v), err: %v",
			businessReq.Behavior,
			businessReq.CountryCode,
			businessReq.PhoneNumber,
			err.Error(),
		)
		return
	}

	if err = _Downstream(
		&payload.PacketClient{
			Header: &payload.Header{
				Major: major.SMS,
				Minor: sms.SendVerificationCodeRsp_,
			},
			Body: bytes,
		},
		conn,
	); err != nil {
		log.ErrorF("sendSMS._Downstream failure(Behavior-%v, CountryCode-%v, PhoneNumber-%v), err: %v",
			businessReq.Behavior,
			businessReq.CountryCode,
			businessReq.PhoneNumber,
			err.Error(),
		)
		return
	}
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

	if err = _Downstream(
		&payload.PacketClient{
			Header: &payload.Header{
				Major: major.Account,
				Minor: account.LoginRsp_,
			},
			Body: bytes,
		},
		conn,
	); err != nil {
		log.ErrorF("_Login._Downstream failure, err: %v",
			err.Error(),
		)
		return nil, err
	}

	if _, err = mgo.InsertDoc(
		config.MongoConf(),
		context.Background(),
		track.GetWhich(),
		track.GetDBName(),
		track.GetTableName(),
		&track.Model{
			UserId:    cast.ToString(businessRsp.UserId),
			Major:     req.GetHeader().GetMajor(),
			Minor:     account.LoginReq_,
			Request:   convert.Bytes2StringArray(req.GetBody()),
			Response:  convert.Bytes2StringArray(bytes),
			Timestamp: time.Now().Unix(),
		},
	); err != nil {
		log.ErrorF("_Login failure, err :", err.Error())
	}

	if businessRsp.Code != code.Success {
		return nil, nil
	}

	session := &payload.Session{}
	session.SetMemberId(businessRsp.MemberId)
	session.SetUserId(businessRsp.UserId)

	return session, nil
}

func fetchRateLimitingConfig(req *payload.PacketClient, conn *websocket.Conn) {
	businessReq := &gateway.FetchRateLimitingConfigReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		log.ErrorF("fetchRateLimitingConfig.json.Unmarshal failure, err: %v",
			err.Error(),
		)
		return
	}
	businessRsp := &gateway.FetchRateLimitingConfigRsp{}
	err = FetchRateLimitingConfig(context.Background(), businessReq, businessRsp)
	if err != nil {
		log.ErrorF("fetchRateLimitingConfig.FetchRateLimitingConfig failure, err: %v",
			err.Error(),
		)
		return
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		log.ErrorF("fetchRateLimitingConfig.json.Marshal failure, err: %v",
			err.Error(),
		)
		return
	}

	if err = _Downstream(
		&payload.PacketClient{
			Header: &payload.Header{
				Major: req.GetHeader().GetMajor(),
				Minor: frontend_gateway.FetchRateLimitingConfigRsp_,
			},
			Body: bytes,
		},
		conn,
	); err != nil {
		log.ErrorF("fetchRateLimitingConfig._Downstream failure, err: %v",
			err.Error(),
		)
		return
	}
}

func echo(req *payload.PacketClient, conn *websocket.Conn) {
	businessReq := &gateway.PingReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		log.ErrorF("echo.json.Unmarshal failure, err: %v",
			err.Error(),
		)
		return
	}
	businessRsp := &gateway.PongRsp{}
	err = Echo(context.Background(), businessReq, businessRsp)
	if err != nil {
		log.ErrorF("echo.Echo failure, err: %v",
			err.Error(),
		)
		return
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		if err != nil {
			log.ErrorF("echo.json.Marshal failure, err: %v",
				err.Error(),
			)
			return
		}
	}

	if err = _Downstream(
		&payload.PacketClient{
			Header: &payload.Header{
				Major: major.FrontendGateway,
				Minor: frontend_gateway.PongRsp_,
			},
			Body: bytes,
		},
		conn,
	); err != nil {
		log.ErrorF("echo._Downstream failure, err: %v",
			err.Error(),
		)
		return
	}
}
