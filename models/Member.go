package models

type Member struct {
	Id            int    `json:"individualId"`
	PreferredName string `json:"preferredName"`
	Gender        string `json:"gender"`
	Surname       string `json:"surname"`
	directoryName string `json:"directoryName"`
}
