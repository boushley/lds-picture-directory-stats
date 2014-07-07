package models

type Specifics struct {
	PhotoUrl   string `json:"photoUrl"`
	ImageId    string `json:"imageId"`
	ImageLevel string `json:"imageLevel"`
	email      string `json:"email"`
	emailLevel string `json:"emailLevel"`
	phone      string `json:"phone"`
	phoneLevel string `json:"phoneLevel"`
}
