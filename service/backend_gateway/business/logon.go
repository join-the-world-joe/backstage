package business

import (
	"backstage/common/code"
	"backstage/common/db/mgo/backend/track"
	"backstage/common/macro/permission"
	"backstage/common/major"
	"backstage/common/payload"
	"backstage/common/protocol/admin"
	"backstage/common/protocol/backend_gateway"
	"backstage/common/protocol/gateway"
	"backstage/common/protocol/sms"
	"backstage/global/config"
	"backstage/global/crypto"
	"backstage/global/log"
	"backstage/global/mgo"
	"backstage/service/backend_gateway/runtime"
	"backstage/utils/convert"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

		log.Debug("Logon.ReadMessage.message: ", convert.Bytes2StringArray(message))

		plainText := message
		if runtime.EncryptionEnable() {
			plainText, err = crypto.AESDecrypt(message)
			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
		}

		log.Debug("Logon.ReadMessage.message.PlainText(Bytes): ", convert.Bytes2StringArray(plainText))
		log.Debug("Logon.ReadMessage.message.PlainText(String): ", string(plainText))

		req := &payload.PacketClient{}
		err = json.Unmarshal(plainText, req)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		ma := req.GetHeader().GetMajor()
		mi := req.GetHeader().GetMinor()

		if ma == major.Admin && mi == admin.SignInReq_ {
			session, err := _SignIn(req, conn)
			if err == nil && session != nil {
				return session, nil
			} else {
				log.Error("Logon._SignIn failure, err: ", err)
			}
		} else if ma == major.SMS && mi == sms.SendVerificationCodeReq_ {
			sendSMS(req, conn)
		} else if ma == major.BackendGateway && mi == backend_gateway.FetchRateLimitingConfigReq_ {
			fetchRateLimitingConfig(req, conn)
		} else if ma == major.BackendGateway && mi == backend_gateway.PingReq_ {
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

func _SignIn(req *payload.PacketClient, conn *websocket.Conn) (*payload.Session, error) {
	businessReq := &admin.SignInReq{}
	err := json.Unmarshal(req.GetBody(), businessReq)
	if err != nil {
		return nil, err
	}
	businessRsp := &admin.SignInRsp{}
	err = admin.SignIn(context.Background(), businessReq, businessRsp)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(businessRsp)
	if err != nil {
		return nil, err
	}
	rsp := &payload.PacketClient{
		Header: &payload.Header{
			Major: req.GetHeader().GetMajor(),
			Minor: admin.SignInRsp_,
		},
		Body: bytes,
	}
	bytes, err = json.Marshal(rsp)
	if err != nil {
		return nil, err
	}

	log.Debug(string(bytes))

	if err = _Downstream(
		&payload.PacketClient{
			Header: &payload.Header{
				Major: req.GetHeader().GetMajor(),
				Minor: admin.SignInRsp_,
			},
			Body: bytes,
		},
		conn,
	); err != nil {
		log.ErrorF("_SignIn._Downstream failure, err: %v",
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
			Operator:   businessRsp.Name,
			Major:      req.GetHeader().GetMajor(),
			Minor:      admin.SignInReq_,
			Request:    convert.Bytes2StringArray(req.GetBody()),
			Permission: permission.SignIn,
			Response:   convert.Bytes2StringArray(bytes),
			Timestamp:  time.Now().Unix(),
		},
	); err != nil {
		log.ErrorF("_SignIn failure, err :", err.Error())
	}

	if businessRsp.Code != code.Success {
		return nil, nil
	}

	session := &payload.Session{}
	session.SetUserId(businessRsp.UserId)
	session.SetMemberId(businessRsp.MemberId)
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
				Minor: backend_gateway.FetchRateLimitingConfigRsp_,
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
				Major: major.BackendGateway,
				Minor: backend_gateway.PongRsp_,
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
