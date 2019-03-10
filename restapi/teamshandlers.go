package restapi

import (
	"github.com/kataras/iris"
	"volunteers-api/db"
	"volunteers-api/restapi/model"
)

// todo add specific error message when volunteer/team/joined team exists
func createTeam(ctx iris.Context) {

	v := model.Team{}
	if err := parseJsonAndSetCtxFromError(ctx, &v, invalidInputFormat); err != nil {
		return
	}

	if !v.IsValid() {
		setCtxError(ctx, missingField)
		return
	}

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	_, insertErr := db.PostgresDb.Query("insert into teams(name,team_motto,createdById) values($1,$2,$3)", v.Name, v.Motto, user.DbId)
	setCtxFromDbError(ctx, insertErr)
}

func getAllTeams(ctx iris.Context) {

	rows, err := db.PostgresDb.Query("select name, team_motto from teams")
	if err := setCtxFromDbError(ctx, err); err != nil {
		return
	}

	var teamArray = make([]model.Team, 0)

	for rows.Next() {
		tTeam := model.Team{}
		err := rows.Scan(&tTeam.Name, &tTeam.Motto)
		if err := setCtxFromDbError(ctx, err); err != nil {
			return
		} else {
			teamArray = append(teamArray, tTeam)
		}
	}

	ctx.JSON(teamArray)
}

func joinTeam(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	teamName := ctx.Params().Get("name")

	_, insertErr := db.PostgresDb.Query("insert into team_members(team_id, volunteer_id) select id, $1 from teams where name = $2", user.DbId, teamName)
	if err := setCtxFromDbError(ctx, insertErr); err != nil {
		return
	}
}

func getJoinedTeams(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	rows, err := db.PostgresDb.Query("select teams.name, teams.team_motto from teams inner join team_members on teams.id = team_members.team_id where team_members.volunteer_id=$1", user.DbId)
	if err := setCtxFromDbError(ctx, err); err != nil {
		return
	}

	var teamArray = make([]model.Team, 0)

	for rows.Next() {
		tTeam := model.Team{}
		err := rows.Scan(&tTeam.Name, &tTeam.Motto)
		if err := setCtxFromDbError(ctx, err); err != nil {
			return
		} else {
			teamArray = append(teamArray, tTeam)
		}
	}

	ctx.JSON(teamArray)
}

func leaveTeam(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	teamName := ctx.Params().Get("name")

	_, insertErr := db.PostgresDb.Query("delete from team_members where team_id in (select id from teams where name=$2) and volunteer_id=$1;", user.DbId, teamName)
	if err := setCtxFromDbError(ctx, insertErr); err != nil {
		return
	}
}

// todo report back when user is trying to delete room but is not creator, now its returning always 200/ok
func deleteTeam(ctx iris.Context) {

	user, _ := getUserForBearer(ctx)
	if user == nil {
		setCtxError(ctx, unknownError)
		return
	}

	teamName := ctx.Params().Get("name")

	_, insertErr := db.PostgresDb.Query("delete from teams where id in (select id from teams where name=$2) and createdById=$1;", user.DbId, teamName)
	if err := setCtxFromDbError(ctx, insertErr); err != nil {
		return
	}
}

func getOccupants(ctx iris.Context) {

	rows, err := db.PostgresDb.Query("select a.name, b.count from teams a left join (select team_id, count(*) as count from team_members group by team_id) b on (a.id = b.team_id)")
	if err := setCtxFromDbError(ctx, err); err != nil {
		return
	}

	var teamArray = make([]model.TeamCount, 0)

	for rows.Next() {
		tTeam := model.TeamCount{}
		err := rows.Scan(&tTeam.Name, &tTeam.Members)
		if err := setCtxFromDbError(ctx, err); err != nil {
			return
		} else {
			teamArray = append(teamArray, tTeam)
		}
	}

	ctx.JSON(teamArray)
}
