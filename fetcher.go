package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/boushley/lds-picture-stats/models"
)

func main() {
	loginTarget := "https://signin.lds.org/login.html"
	directoryUrl := "https://www.lds.org/directory/services/ludrs/mem/member-list/1975943"

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	resp, err := client.PostForm(loginTarget, url.Values{"username": {"boushley"}, "password": {"notapassword"}})

	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Get(directoryUrl)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	var households []models.Household
	err = json.Unmarshal(body, &households)
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range households {
		specifics, err := h.FetchSpecifics(client)
		if err != nil {
			log.Fatal(err)
		}
		if !specifics.HasHouseholdPicture() {
			log.Println("Missing picture for: ", h.HouseholdName)
		}
	}
}
