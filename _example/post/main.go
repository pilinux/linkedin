// Fetch published posts from a LinkedIn page
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pilinux/linkedin"
)

// GlobalApp - LinkedIn application
var GlobalApp *linkedin.App

// Init - Initialize the application
func Init() error {
	clientID := strings.TrimSpace(os.Getenv("LINKEDIN_CLIENT_ID"))
	clientSecret := strings.TrimSpace(os.Getenv("LINKEDIN_CLIENT_SECRET"))

	GlobalApp = linkedin.New(clientID, clientSecret)

	return nil
}

func main() {
	// initialize the application
	err := Init()
	if err != nil {
		return
	}

	accessToken := strings.TrimSpace(os.Getenv("LINKEDIN_ACCESS_TOKEN"))
	if accessToken == "" {
		fmt.Println("LinkedIn: access token is empty")
		return
	}

	// create a new LinkedIn session
	session := GlobalApp.Session(accessToken)

	// override the default LinkedIn API version
	session.SetLinkedInVersion("202405")

	// set Authorization header
	session.UseAuthorizationHeader()

	// list of published posts
	var elements []linkedin.ElementPost

	// get the first page of posts
	author := "urn:li:organization:123456789"
	author = linkedin.EncodeURL(author)
	count := "100"
	response, data, err := session.Get("/posts?q=author&author=" + author + "&count=" + count + "&sortBy=LAST_MODIFIED")
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
		return
	}

	// parse the response body
	var post linkedin.Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		fmt.Println(err)
		return
	}

	// append the elements
	elements = append(elements, post.Elements...)

	// loop through the pages
	for {
		// get the next page URL
		ok, next := post.Paging.GetNext()
		if !ok {
			fmt.Println("No next page, breaking the loop...")
			break
		}

		// get the next page
		fmt.Println("Fetching next page...")
		fmt.Println("URL:", next)
		response, data, err = session.Get(next)
		if err != nil {
			fmt.Println("LinkedIn: error fetching next page")
			fmt.Println(err)
			fmt.Println("breaking the loop...")
			break
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
			fmt.Println("breaking the loop...")
			break
		}

		// reset the post struct
		post = linkedin.Post{}

		// parse the response body
		err = json.Unmarshal(data, &post)
		if err != nil {
			fmt.Println("LinkedIn: error parsing response")
			fmt.Println(err)
			fmt.Println("breaking the loop...")
			break
		}

		// append the elements
		elements = append(elements, post.Elements...)
	}

	// print the messages of the posts
	for i, element := range elements {
		fmt.Println("Post #", i+1)
		fmt.Println("====================================")
		fmt.Println(element.Commentary)
		fmt.Println("====================================")
	}
}
