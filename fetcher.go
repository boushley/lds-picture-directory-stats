package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/boushley/lds-picture-stats/models"
)

func main() {
	loginTarget := "https://signin.lds.org/login.html"
	directoryUrl := "https://www.lds.org/directory/services/ludrs/mem/member-list/1975943"

	ignoreMap := loadIgnoreList()

	username := readValue("lds.org username: ")
	password := readValue("lds.org password: ")

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}
	resp, err := client.PostForm(loginTarget, url.Values{"username": {username}, "password": {password}})

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

	found := 0
	checked := 0
	skipped := 0

	for _, h := range households {
		if _, ok := ignoreMap[h.CoupleName]; ok {
			skipped++
			continue
		}

		checked++
		specifics, err := h.FetchSpecifics(client)
		if err != nil {
			log.Fatal(err)
		}
		if !specifics.HasHouseholdPicture() {
			fmt.Printf("%s, %s\n", h.CoupleName, specifics.GetEmails())
		} else {
			found++
		}
	}

	fmt.Printf("Missing: %d Found: %d Checked: %d Skipped: %d\n", (checked - found), found, checked, skipped)
}

func readValue(prompt string) (result string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result = strings.Trim(result, " \n")
	return
}

func loadIgnoreList() (result map[string]bool) {
	result = make(map[string]bool)

	file, err := os.Open("ignore-names")
	if err != nil {
		log.Println("Unable to open ignore-names file, will not ignore any names.")
		log.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result[strings.Trim(scanner.Text(), " \n")] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
