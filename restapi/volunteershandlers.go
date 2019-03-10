package restapi

import (
	"github.com/kataras/iris"
	"volunteers-api/db"
	"volunteers-api/restapi/model"
)

func createVolunteer(ctx iris.Context) {

	v := model.Volunteer{}
	if err := parseJsonAndSetCtxFromError(ctx, &v, invalidInputFormat); err != nil {
		return
	}

	if !v.IsValid() {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(getErrorJson(missingField))
		return
	}

	//lastInsertId := 0
	_, insertErr := db.PostgresDb.Query("insert into volunteers(name,email,password) values($1,$2,$3)", v.Name, v.Email, v.Password)
	setCtxFromDbError(ctx, insertErr)
}

func getVolunteers(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	output := model.Volunteer{}
	err := db.PostgresDb.QueryRow("select name,email from volunteers where id=$1", user.DbId).Scan(&output.Name, &output.Email)
	if err := setCtxFromDbError(ctx, err); err != nil {
		return
	}

	ctx.JSON(output)
}

func updateVolunteers(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	v := model.Volunteer{}
	if err := parseJsonAndSetCtxFromError(ctx, &v, invalidInputFormat); err != nil {
		return
	}

	if v.Email == "" || v.Name == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(getErrorJson(missingField))
		return
	}

	_, insertErr := db.PostgresDb.Query("update volunteers set name=$2, email=$3 where id=$1", user.DbId, v.Name, v.Email)
	setCtxFromDbError(ctx, insertErr)
}
