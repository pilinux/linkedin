package linkedin

import (
	"strings"
)

// Paging struct for pagination
type Paging struct {
	Start int    `json:"start"`
	Count int    `json:"count"`
	Links []Link `json:"links"`
	Total int    `json:"total"`
}

// Link struct for pagination
type Link struct {
	Type string `json:"type"` // application/json
	Rel  string `json:"rel"`  // prev, next
	Href string `json:"href"`
}

// GetNext returns the next page URL
func (p *Paging) GetNext() (bool, string) {
	for _, link := range p.Links {
		if link.Rel == "next" {
			// remove "/rest" from the URL
			return true, strings.Replace(link.Href, "/rest", "", 1)
		}
	}
	return false, ""
}

// GetPrev returns the previous page URL
func (p *Paging) GetPrev() (bool, string) {
	for _, link := range p.Links {
		if link.Rel == "prev" {
			// remove "/rest" from the URL
			return true, strings.Replace(link.Href, "/rest", "", 1)
		}
	}
	return false, ""
}
