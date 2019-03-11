package restapi

import (
	"github.com/kataras/iris"
)

// middleware for handling authorization (based on tokens)
func handleAuth(ctx iris.Context) {
	if usr, _ := getUserForBearer(ctx); usr != nil {
		ctx.Next()
		return
	}

	ctx.StatusCode(iris.StatusBadRequest)
	_, _ = ctx.JSON(getErrorJson(notAuthorized))
}
