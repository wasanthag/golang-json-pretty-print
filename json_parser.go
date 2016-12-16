package main
//update the variable link to point the desired json url

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
 )

func main() {
    link := "https://api.github.com/repos/wasanthag/ucs_nexus_cable_check/commits"
    //from https://golang.org/pkg/net/http/
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    response, err := client.Get(link)
    if err != nil {
        fmt.Println("Getting http response error",err)
    }
    defer response.Body.Close()

    content, _ := ioutil.ReadAll(response.Body)
     //s := string(content)
    //fmt.Println(s)

    //generated with http://json2struct.mervine.net/
    // use println of s above to use as input for the tool as using direct input from github api does not seem to capture all json data
    type githubJson struct {
	Sha string `json:"sha"`
	Commit  struct {
		Author struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		Committer    struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL string `json:"url"`
		CommentCount int `json:"comment_count"`
	} `json:"commit"`
	URL string `json:"url"`
	HTMLURL string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"author"`
	Committer struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"committer"`
	Parents []struct {
		HTMLURL string `json:"html_url"`
		Sha     string `json:"sha"`
		URL     string `json:"url"`
	} `json:"parents"`
}

        //from https://golang.org/pkg/encoding/json/#example_Unmarshal
        var commits []githubJson
	err2 := json.Unmarshal(content, &commits)
	if err2 != nil {
		fmt.Println("Unpacking json to slice error:", err2)
	}

	for l := range commits {
		fmt.Println("-------- Commit details in json --------")
		fmt.Printf("%+v\n", commits[l])
		fmt.Println()
		fmt.Println("-------- Commit details in human readable format --------")
		b, err3 := json.MarshalIndent(commits[l], "", "  ")
		if err3 != nil {
        		fmt.Println("json Indenting error",err3)
    		}
                println(string(b))

	}


}
