package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var githubConnections map[string]string

type githubResponse struct {
	Data struct {
		Viewer struct {
			ID string `json:"id"`
		} `json:"viewer"`
	} `json:"data"`
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "670b383c12b8bcc3fb01",
	ClientSecret: "36a64ec1ac39211bb33cf8bb9949dc321c520218",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./oauth/main/sample.html")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fprint(w, string(file)); err != nil {
		log.Fatal(err)
	}
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "state is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "could not login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)

	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)

	client := oauth2.NewClient(r.Context(), ts)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "could not get user", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var githubResponse githubResponse
	if err := json.NewDecoder(resp.Body).Decode(&githubResponse); err != nil {
		http.Error(w, "github invalid response", http.StatusInternalServerError)
		return
	}

	githubID := githubResponse.Data.Viewer.ID

	userID, ok := githubConnections[githubID]
	if !ok {
		// new user, create account
	}

	// login to account userID using JWT
	fmt.Println(userID)
}
