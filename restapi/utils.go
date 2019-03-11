package restapi

import (
	"github.com/kataras/iris"
	"github.com/lib/pq"
	"log"
	"net"
	"volunteers-api/session"
)

// todo avoid duplicate user extraction
// attempts to retrieve access token from iris.Context and returns User session if found
func getUserForBearer(ctx iris.Context) (*session.User, string) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		return nil, ""
	} else {
		token := bearer[7:]
		return apiSession.GetUser(token), token
	}
}

// utility function to simplify input json parsing and propagating error if any occurs
func parseJsonAndSetCtxFromError(ctx iris.Context, i interface{}, errCode uint32) error {
	if err := ctx.ReadJSON(i); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(getErrorJson(errCode))
		return err
	}
	return nil
}

// utility function to simplify processing of errors related to database operations
func setCtxFromDbError(ctx iris.Context, err error, alreadyExistError uint32) error {
	// fixme handle errors correctly!
	if netErr, ok := err.(net.Error); ok {
		// handle network error
		log.Println("caught network error: ", netErr.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(getErrorJson(networkError))
		return netErr
	} else {
		if pqErr, _ := err.(*pq.Error); pqErr != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			if pqErr.Code == "23505" {
				_, _ = ctx.JSON(getErrorJson(alreadyExistError))
			} else {
				_, _ = ctx.JSON(getErrorJson(dbError))
			}
			return pqErr
		}
	}
	return nil
}

// utility function to set iris.Context error with error response
func setCtxError(ctx iris.Context, errCode uint32) {
	ctx.StatusCode(iris.StatusBadRequest)
	_, _ = ctx.JSON(getErrorJson(errCode))
}
