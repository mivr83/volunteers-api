package restapi

import (
	"github.com/kataras/iris"
	"volunteers-api/session"
	"volunteers-api/settings"
)

// probably should be retrieved from som config
const apiV1Route = "/api/v1"

var app *iris.Application
var apiSession session.Session

// initializes api, add all routes and starts server on specified configuration
func InitAndServe(connSettings *settings.ConnSettings, session session.Session) {
	app = iris.Default()

	apiSession = session

	apiV1 := app.Party(apiV1Route)
	addRoutes(&apiV1)

	// ignore error, this is last executed statement which automatically prints out errors
	_ = app.Run(iris.Addr(connSettings.GetApiHostWithPort()))
}

func addRoutes(party *iris.Party) {
	lParty := *party

	// router which checks if user is authenticated aka have token
	authenticatedParty := lParty.Party("", handleAuth)

	// ==== auth routes

	// login user
	lParty.Post("/login", loginHandler)

	// logout user
	lParty.Get("/logout", logoutHandler)

	// ==== volunteer routes

	// returns all volunteers profiles (based on access rights, volunteers has rights to retrieve only his own profile)
	authenticatedParty.Get("/volunteers", getVolunteers)

	// returns specific volunteer profile with 'id' (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Get("/volunteers/{id:int}", nil)

	// deletes specific volunteer (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Delete("/volunteers/{id:int}", nil)

	// creates new volunteer
	lParty.Post("/volunteers", createVolunteer)

	// updates specific volunteer (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Put("/volunteers/{id:int}", nil)

	// updates currently logged in volunteer
	authenticatedParty.Put("/volunteers", updateVolunteers)

	// ==== teams routes

	// returns list of teams
	authenticatedParty.Get("/teams", getAllTeams)

	// returns list of teams
	authenticatedParty.Delete("/teams", getAllTeams)

	// returns list of joined teams
	// todo
	authenticatedParty.Get("/teams/joined", getJoinedTeams)

	// creates new team
	authenticatedParty.Post("/teams", createTeam)

	// joins specific team
	// fixme should be done with id but its simpler to join team by name so i used name instead
	authenticatedParty.Post("/teams/{name:string}/join", joinTeam)

	// leaves specific team
	// fixme should be done with id but its simpler to join team by name so i used name instead
	authenticatedParty.Delete("/teams/{name:string}/leave", leaveTeam)

	// deletes team (volunteer can delete team only when he was creator, also team can be deleted
	// while it have active members, this will be removed from team)
	authenticatedParty.Delete("/teams/{name:string}", deleteTeam)

	// returns list of teams with count of occupants
	lParty.Get("/teams/occupants", getOccupants)

}
