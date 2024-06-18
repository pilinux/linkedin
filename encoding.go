package linkedin

import "net/url"

// EncodeURL encodes text into URL-encoded format
func EncodeURL(text string) string {
	return url.QueryEscape(text)
}

// DecodeURL decodes URL-encoded text into plain text
func DecodeURL(text string) (string, error) {
	return url.QueryUnescape(text)
}
