package linkedin

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
