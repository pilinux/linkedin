package linkedin

// LinkedIn rest api base urls
//
// See: https://learn.microsoft.com/en-us/linkedin/marketing/versioning?view=li-lms-2024-04
const (
	OauthBaseURL     = "https://www.linkedin.com/oauth/v2/"
	VersionedBaseURL = "https://api.linkedin.com/rest/"
)

// Method - HTTP method for an API call.
type Method string

// HTTP API methods.
const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

// Header - HTTP header data type
type Header string

// HTTP headers used in LinkedIn API requests
//
// https://linkedin.github.io/rest.li/spec/protocol
//
// https://learn.microsoft.com/en-us/linkedin/marketing/versioning?view=li-lms-2024-04#how-your-api-client-should-use-versioning
const (
	ContentType           Header = "Content-Type"
	Connection            Header = "Connection"
	RestLiProtocolVersion Header = "X-RestLi-Protocol-Version"
	RestLiMethodHeader    Header = "X-RestLi-Method"
	LinkedInVersion       Header = "LinkedIn-Version"
	Authorization         Header = "Authorization"
	UserAgent             Header = "user-agent"
	CreatedEntityID       Header = "X-RestLi-Id"
)

// ContentDataType - HTTP content data type
type ContentDataType string

// HTTP content data type
const (
	URLEncoded ContentDataType = "application/x-www-form-urlencoded"
	JSON       ContentDataType = "application/json"
)

// RestLiMethod - LinkedIn Rest.Li method type
type RestLiMethod string

// LinkedIn Rest.Li methods
const (
	Get                RestLiMethod = "GET"
	BatchGet           RestLiMethod = "BATCH_GET"
	GetAll             RestLiMethod = "GET_ALL"
	Finder             RestLiMethod = "FINDER"
	BatchFinder        RestLiMethod = "BATCH_FINDER"
	Create             RestLiMethod = "CREATE"
	BatchCreate        RestLiMethod = "BATCH_CREATE"
	Update             RestLiMethod = "UPDATE"
	BatchUpdate        RestLiMethod = "BATCH_UPDATE"
	PartialUpdate      RestLiMethod = "PARTIAL_UPDATE"
	BatchPartialUpdate RestLiMethod = "BATCH_PARTIAL_UPDATE"
	Delete             RestLiMethod = "DELETE"
	BatchDelete        RestLiMethod = "BATCH_DELETE"
	Action             RestLiMethod = "ACTION"
)

// RestLiMethodToHTTPMethodMap - map of LinkedIn Rest.Li methods to HTTP methods
var RestLiMethodToHTTPMethodMap = map[RestLiMethod]Method{
	Get:                GET,
	BatchGet:           GET,
	GetAll:             GET,
	Finder:             GET,
	BatchFinder:        GET,
	Update:             PUT,
	BatchUpdate:        PUT,
	Create:             POST,
	BatchCreate:        POST,
	PartialUpdate:      POST,
	BatchPartialUpdate: POST,
	Action:             POST,
	Delete:             DELETE,
	BatchDelete:        DELETE,
}

// Rest.Li special characters
const (
	ListPrefix       = "List("
	ListSuffix       = ")"
	ListItemSep      = ","
	ObjPrefix        = "("
	ObjSuffix        = ")"
	ObjKeyValSep     = ":"
	ObjKeyValPairSep = ","
	LeftBracket      = "("
	RightBracket     = ")"
)
