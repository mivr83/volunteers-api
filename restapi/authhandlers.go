package restapi

import (
	"github.com/google/uuid"
	"github.com/kataras/iris"
	"log"
	"volunteers-api/db"
	"volunteers-api/restapi/model"
	"volunteers-api/session"
)

// handler for route POST:apiV1/login
func loginHandler(ctx iris.Context) {

	v := model.Login{}
	if err := parseJsonAndSetCtxFromError(ctx, &v, invalidInputFormat); err != nil {
		return
	}

	var userId uint32 = 0
	err := db.PostgresDb.QueryRow("select id from volunteers where email=trim($1) and password=trim($2)", v.Email, v.Password).Scan(&userId)
	if err := setCtxFromDbError(ctx, err, volunteerAlreadyExists); err != nil {
		return
	}

	if userId == 0 {
		setCtxError(ctx, invalidCredentials)
		return
	}

	// generate fake token and add create session
	uuidString := uuid.New().String()
	apiSession.AddUser(uuidString, &session.User{DbId: userId, Role: "user"})

	ctx.JSON(model.TokenResponse{Token: uuidString})
}

// handler for route GET:apiV1/logout
func logoutHandler(ctx iris.Context) {
	if usr, token := getUserForBearer(ctx); usr != nil {
		log.Println("logging out user: ", token)
		apiSession.RemoveUser(token)
	}

	ctx.StatusCode(iris.StatusOK)
}
