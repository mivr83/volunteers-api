package model

type Volunteer struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (v *Volunteer) IsValid() bool {
	return v.Name != "" && v.Email != "" && v.Password != ""
}

type Team struct {
	Name  string `json:"name"`
	Motto string `json:"motto"`
}

type TeamCount struct {
	Name    string `json:"team_name"`
	Members uint32 `json:"members"`
}

func (v *Team) IsValid() bool {
	return v.Name != "" && v.Motto != ""
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
