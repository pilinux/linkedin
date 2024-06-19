package linkedin

// Post struct for LinkedIn posts
type Post struct {
	Paging   Paging        `json:"paging"`
	Elements []ElementPost `json:"elements"`
}

// ElementPost struct for LinkedIn post elements
type ElementPost struct {
	IsReshareDisabledByAuthor bool                   `json:"isReshareDisabledByAuthor"`
	CreatedAt                 int64                  `json:"createdAt"`
	LifecycleState            string                 `json:"lifecycleState"` // e.g. PUBLISHED
	LastModifiedAt            int64                  `json:"lastModifiedAt"`
	Visibility                string                 `json:"visibility"` // e.g. PUBLIC
	PublishedAt               int64                  `json:"publishedAt"`
	Author                    string                 `json:"author"` // e.g. urn:li:organization:123456
	ID                        string                 `json:"id"`     // e.g. urn:li:ugcPost:123456
	ReshareContext            ReshareContextPost     `json:"reshareContext"`
	Distribution              DistributionPost       `json:"distribution"`
	Commentary                string                 `json:"commentary"` // e.g. Hello, world!
	LifecycleStateInfo        LifecycleStateInfoPost `json:"lifecycleStateInfo"`
}

// ReshareContextPost struct for LinkedIn post reshare context
type ReshareContextPost struct {
	Parent string `json:"parent"` // e.g. urn:li:ugcPost:123456
	Root   string `json:"root"`   // e.g. urn:li:ugcPost:123456
}

// DistributionPost struct for LinkedIn post distribution
type DistributionPost struct {
	FeedDistribution string `json:"feedDistribution"` // e.g. MAIN_FEED
}

// LifecycleStateInfoPost struct for LinkedIn post lifecycle state information
type LifecycleStateInfoPost struct {
	IsEditedByAuthor bool `json:"isEditedByAuthor"`
}
