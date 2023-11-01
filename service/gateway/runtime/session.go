package runtime

import (
	"backstage/common/payload"
	"backstage/global/log"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

var _session sync.Map

func StoreSession(userId int64, sess *payload.Session) {
	_session.Store(userId, sess)
}

func LoadSession(userId int64) (*payload.Session, error) {
	value, ok := _session.Load(userId)
	if ok {
		return value.(*payload.Session), nil
	}
	return nil, errors.New(fmt.Sprintf("%d doesn't exist", userId))
}

func RemoveSession(userId int64) {
	_session.Delete(userId)
}

func LoadOnlineSequenceList() []uint64 {
	list := []uint64{}
	f := func(key, value any) bool {
		if userId, ok := key.(int64); ok {
			session, err := LoadSession(userId)
			if err != nil {
				return true
			}
			list = append(list, session.Sequence)
		}
		return true
	}
	_session.Range(f)
	return list
}

func DumpSession(userId int64) {
	if sess, exist := _session.Load(userId); exist {
		bytes, err := json.Marshal(sess)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Debug(string(bytes))
		return
	}
	return
}
