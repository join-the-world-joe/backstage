package middleware

import (
	"backstage/common/code"
	"backstage/common/payload"
	"backstage/global/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo: add to dynamic config
var unchecked = map[string]struct{}{
	"/ws":  {},
	"/cmd": {},
}

func Permission() gin.HandlerFunc {
	// every legal user has his token, the token also as a key(hash) in cache and the fields contain user_id, role_id ...
	// since the admin panel change role of user occasionally, the cache also need to be change correspondingly
	// a rule to match token from user_id is needed, every role modification operation has to remove the token in cache
	return func(c *gin.Context) {
		if _, exist := unchecked[c.Request.URL.Path]; exist {
			c.Next()
			return
		}
		// todo: check if permission denied
		rsp := (&payload.Response{}).SetCode(code.AccessDenied).Bytes()
		c.Data(http.StatusOK, "application/json", rsp)
		log.Error(fmt.Sprintf("code: %v, path: %v", code.AccessDenied, c.Request.URL.Path))
		c.Abort()

		c.Next()
	}
}
