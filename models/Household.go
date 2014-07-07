package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Household struct {
	Children          []Member `json:"children"`
	HeadOfHouse       Member   `json:"headOfHouse"`
	Spouse            Member   `json:"spouse"`
	CoupleName        string   `json:"coupleName"`
	HouseholdName     string   `json:"householdName"`
	HeadOfHouseholdId int      `json:"headOfHouseholdIndividualId"`
}

func (h Household) FetchSpecifics(client *http.Client) (specifics HouseholdSpecifics, err error) {
	householdUrl := "https://www.lds.org/directory/services/ludrs/mem/householdProfile/" + string(h.HeadOfHouseholdId)
	log.Println("Requesting household from: ", householdUrl)
	resp, err := client.Get(householdUrl)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &specifics)
	return
}
