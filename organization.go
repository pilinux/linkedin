package linkedin

import "strings"

// Organization struct for LinkedIn organizations
type Organization struct {
	Paging   Paging                `json:"paging"`
	Elements []ElementOrganization `json:"elements"`
}

// ElementOrganization struct for LinkedIn organization elements
type ElementOrganization struct {
	RoleAssignee string       `json:"roleAssignee"` // e.g. urn:li:person:a1b2c3
	State        string       `json:"state"`        // e.g. APPROVED
	LastModified LastModified `json:"lastModified"`
	Role         string       `json:"role"` // e.g. ADMINISTRATOR
	Created      Created      `json:"created"`
	Organization string       `json:"organization"` // e.g. urn:li:organization:123456789
}

// LastModified struct for last modified timestamp
type LastModified struct {
	Actor        string `json:"actor"`        // e.g. urn:li:person:a1b2c3
	Impersonator string `json:"impersonator"` // e.g. urn:li:servicePrincipal:organization
	Time         int64  `json:"time"`         // e.g. 1612345678901
}

// Created struct for created timestamp
type Created struct {
	Actor        string `json:"actor"`        // e.g. urn:li:person:a1b2c3
	Impersonator string `json:"impersonator"` // e.g. urn:li:servicePrincipal:organization
	Time         int64  `json:"time"`         // e.g. 1612345678901
}

// GetOrganizationID returns the organization ID from the organization URN
func (e *ElementOrganization) GetOrganizationID() string {
	// extract the organization ID from e.g. urn:li:organization:123456789
	return e.Organization[strings.LastIndex(e.Organization, ":")+1:]
}

// OrganizationInfo struct for LinkedIn organization information
type OrganizationInfo struct {
	VanityName              string        `json:"vanityName"`
	LocalizedName           string        `json:"localizedName"`
	Created                 Created       `json:"created"`
	VersionTag              string        `json:"versionTag"`
	CoverPhotoV2            CoverPhotoV2  `json:"coverPhotoV2"`
	OrganizationType        string        `json:"organizationType"` // e.g. SELF_OWNED
	DefaultLocale           DefaultLocale `json:"defaultLocale"`
	LocalizedSpecialties    []string      `json:"localizedSpecialties"`
	Name                    NameLocalized `json:"name"`
	PrimaryOrganizationType string        `json:"primaryOrganizationType"` // e.g. NONE
	LastModified            LastModified  `json:"lastModified"`
	ID                      int64         `json:"id"` // ID of the organization
	LocalizedDescription    string        `json:"localizedDescription"`
	AutoCreated             bool          `json:"autoCreated"`
	LocalizedWebsite        string        `json:"localizedWebsite"`
	LogoV2                  LogoV2        `json:"logoV2"`
}

// CoverPhotoV2 struct for cover photo
type CoverPhotoV2 struct {
	Original string `json:"original"`
}

// LogoV2 struct for logo
type LogoV2 struct {
	Original string `json:"original"`
}

// DefaultLocale struct for default locale
type DefaultLocale struct {
	Country  string `json:"country"`  // e.g. US
	Language string `json:"language"` // e.g. en
}

// NameLocalized struct for localized name
type NameLocalized struct {
	Localized       map[string]string `json:"localized"`
	PreferredLocale DefaultLocale     `json:"preferredLocale"`
}

// GetOrganizationID returns the organization ID from the details of the organization
func (oi *OrganizationInfo) GetOrganizationID() int64 {
	return oi.ID
}
