package restapi

import (
	"github.com/kataras/iris"
)

func handleAuth(ctx iris.Context) {
	if usr, _ := getUserForBearer(ctx); usr != nil {
		ctx.Next()
		return
	}

	ctx.StatusCode(iris.StatusBadRequest)
	_, _ = ctx.JSON(getErrorJson(notAuthorized))
}
