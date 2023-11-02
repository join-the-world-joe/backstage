package runtime

import (
	"backstage/common/websock"
)

func WebsocketEndpoint() string {
	return serviceConf.Servant.WebsocketEndpoint
}

func WebsocketAuthReadDeadline() int {
	if serviceConf.Servant.WebsocketAuthReadDeadline == 0 {
		return websock.DefaultAuthReadDeadline
	}
	return serviceConf.Servant.WebsocketAuthReadDeadline
}

func WebsocketReadDeadline() int {
	if serviceConf.Servant.WebsocketReadDeadline == 0 {
		return websock.DefaultReadDeadline
	}
	return serviceConf.Servant.WebsocketReadDeadline
}

func WebsocketReadLimit() int64 {
	if serviceConf.Servant.WebsocketReadLimit == 0 {
		return websock.DefaultReadLimit
	}
	return serviceConf.Servant.WebsocketReadLimit
}

func WebsocketReadBufferSize() int {
	if serviceConf.Servant.WebsocketReadBufferSize == 0 {
		return websock.DefaultReadBufferSize
	}
	return serviceConf.Servant.WebsocketReadBufferSize
}

func WebsocketWriteBufferSize() int {
	if serviceConf.Servant.WebsocketWriteBufferSize == 0 {
		return websock.DefaultWriteBufferSize
	}
	return serviceConf.Servant.WebsocketWriteBufferSize
}
