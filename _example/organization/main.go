// Get list of organizations managed by the authenticated user.
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

	// list of all organizations
	var elements []linkedin.ElementOrganization

	// get the first page of organizations
	count := "100"
	response, data, err := session.Get("/organizationAcls?q=roleAssignee&count=" + count)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
		return
	}

	// parse the response body
	var organization linkedin.Organization
	err = json.Unmarshal(data, &organization)
	if err != nil {
		fmt.Println(err)
		return
	}

	// append the elements
	elements = append(elements, organization.Elements...)

	// loop through the pages
	for {
		// get the next page URL
		hasNext, next := organization.Paging.GetNext()
		if !hasNext {
			fmt.Println("No more pages, breaking the loop...")
			break
		}

		// get the next page
		fmt.Println("Next page:", next)
		response, data, err := session.Get(next)
		if err != nil {
			fmt.Println("LinkedIn: error fetching the next page")
			fmt.Println(err)
			fmt.Println("Breaking the loop...")
			break
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
			fmt.Println("Breaking the loop...")
			break
		}

		// reset the organization
		organization = linkedin.Organization{}

		// parse the response body
		err = json.Unmarshal(data, &organization)
		if err != nil {
			fmt.Println("LinkedIn: error parsing response")
			fmt.Println(err)
			fmt.Println("Breaking the loop...")
			break
		}

		// append the elements
		elements = append(elements, organization.Elements...)
	}

	fmt.Println("Total organizations:", len(elements))

	// print the organizations
	for i, element := range elements {
		fmt.Println("Organization #", i+1)
		fmt.Println("================================")
		fmt.Println("RoleAssignee:", element.RoleAssignee)
		fmt.Println("State:", element.State)
		fmt.Println("Organization:", element.Organization)
		fmt.Println("Organization ID:", element.GetOrganizationID())
		fmt.Println("================================")
	}

	// organization information
	for _, element := range elements {
		organizationID := element.GetOrganizationID()

		// get the organization information
		response, data, err := session.Get("/organizations/" + organizationID)
		if err != nil {
			fmt.Println("LinkedIn: error fetching organization info")
			fmt.Println(err)
			return
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println("LinkedIn: unexpected status code", response.StatusCode)
			fmt.Println("exiting...")
			return
		}

		// parse the response body
		var organizationInfo linkedin.OrganizationInfo
		err = json.Unmarshal(data, &organizationInfo)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Organization Info")
		fmt.Println("================================")
		fmt.Println("VanityName:", organizationInfo.VanityName)
		fmt.Println("LocalizedName:", organizationInfo.LocalizedName)
		fmt.Println("VersionTag:", organizationInfo.VersionTag)
		fmt.Println("OrganizationType:", organizationInfo.OrganizationType)
		fmt.Println("ID:", organizationInfo.GetOrganizationID())
		fmt.Println("AutoCreated:", organizationInfo.AutoCreated)
		fmt.Println("================================")
	}
}
