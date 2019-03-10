package restapi

import (
	"github.com/kataras/iris"
	"github.com/lib/pq"
	"volunteers-api/session"
)

// todo avoid duplicate user extraction
func getUserForBearer(ctx iris.Context) (*session.User, string) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		return nil, ""
	} else {
		token := bearer[7:]
		return apiSession.GetUser(token), token
	}
}

func parseJsonAndSetCtxFromError(ctx iris.Context, i interface{}, errCode uint32) error {
	if err := ctx.ReadJSON(i); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(getErrorJson(errCode))
		return err
	}
	return nil
}

func setCtxFromDbError(ctx iris.Context, err error) error {
	if err, _ := err.(*pq.Error); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		if err.Code == "23505" {
			_, _ = ctx.JSON(getErrorJson(alreadyExists))
		} else {
			_, _ = ctx.JSON(getErrorJson(dbError))
		}
	}
	return nil
}

func setCtxError(ctx iris.Context, errCode uint32) {
	ctx.StatusCode(iris.StatusBadRequest)
	_, _ = ctx.JSON(getErrorJson(errCode))
}
