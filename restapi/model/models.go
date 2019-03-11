package model

type Volunteer struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (v *Volunteer) IsValid() bool {
	return v.Name != "" && v.Email != "" && v.Password != ""
}

type Team struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Motto string `json:"motto"`
}

type TeamCount struct {
	Name      string `json:"team_name"`
	Occupants uint32 `json:"occupants"`
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
