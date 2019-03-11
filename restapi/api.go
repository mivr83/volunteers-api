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

// initializes api, add all routes and starts server using specified configuration
// takes as parameter session interface which is used to retrieve/store volunteers sessions
func InitAndServe(connSettings *settings.ConnSettings, session session.Session) {
	app = iris.Default()

	apiSession = session

	// base API route with version
	apiV1 := app.Party(apiV1Route)
	addRoutes(&apiV1)

	// ignore error, this is last executed statement which automatically prints out errors
	_ = app.Run(iris.Addr(connSettings.GetApiHostWithPort()))
}

func addRoutes(party *iris.Party) {
	noAuthParty := *party

	// router which checks if user is authenticated aka have token
	authParty := noAuthParty.Party("", handleAuth)

	// ==== auth routes

	// login user
	// POST apiV1Route/login
	noAuthParty.Post("/login", loginHandler)

	// logout user
	// GET apiV1Route/logout
	noAuthParty.Get("/logout", logoutHandler)

	// ==== volunteer routes

	// returns all volunteers profiles (based on access rights, volunteers has rights to retrieve only his own profile)
	// GET apiV1Route/volunteers
	authParty.Get("/volunteers", getVolunteers)

	// returns specific volunteer profile with 'id' (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Get("/volunteers/{id:int}", nil)

	// deletes specific volunteer (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Delete("/volunteers/{id:int}", nil)

	// creates new volunteer
	// POST apiV1Route/volunteers
	noAuthParty.Post("/volunteers", createVolunteer)

	// updates specific volunteer (based on access rights)
	// todo implementation is omitted because it's not needed
	//lParty.Put("/volunteers/{id:int}", nil)

	// updates currently logged in volunteer
	// PUT apiV1Route/volunteers
	authParty.Put("/volunteers", updateVolunteers)

	// ==== teams routes

	// returns list of teams
	// GET apiV1Route/teams
	authParty.Get("/teams", getAllTeams)

	// returns list of joined teams
	// GET apiV1Route/teams/joined
	authParty.Get("/teams/joined", getJoinedTeams)

	// creates new team
	// POST apiV1Route/teams
	authParty.Post("/teams", createTeam)

	// joins specific team
	// fixme should be done with id but its simpler to join team by name so i used name instead
	// POST apiV1Route/{id:int}/join
	authParty.Post("/teams/{name:string}/join", joinTeam)

	// leaves specific team
	// fixme should be done with id but its simpler to join team by name so i used name instead
	// DELETE apiV1Route/{id:int}/leave
	authParty.Delete("/teams/{name:string}/leave", leaveTeam)

	// deletes team (volunteer can delete team only when he was creator, also team can be deleted
	// while it have active members, this will be removed from team)
	// fixme should be done with id but its simpler to join team by name so i used name instead
	// DELETE apiV1Route/teams/{id:int}
	authParty.Delete("/teams/{name:string}", deleteTeam)

	// returns list of teams with count of occupants
	// GET apiV1Route/teams/occupants
	noAuthParty.Get("/teams/occupants", getOccupants)

}
