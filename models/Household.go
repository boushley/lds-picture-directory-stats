package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Household struct {
	Children          []Member `json:"children"`
	HeadOfHouse       Member   `json:"headOfHouse"`
	Spouse            Member   `json:"spouse"`
	CoupleName        string   `json:"coupleName"`
	HouseholdName     string   `json:"householdName"`
	HeadOfHouseholdId int64    `json:"headOfHouseIndividualId"`
}

func (h Household) FetchSpecifics(client *http.Client) (specifics HouseholdSpecifics, err error) {
	householdUrl := "https://www.lds.org/directory/services/ludrs/mem/householdProfile/" + strconv.FormatInt(h.HeadOfHouseholdId, 10)
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
