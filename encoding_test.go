package linkedin

import (
	"testing"
)

// TestEncodeURL tests the EncodeURL function
func TestEncodeURL(t *testing.T) {
	text := "urn:li:organization:123456"
	encoded := EncodeURL(text)
	expected := "urn%3Ali%3Aorganization%3A123456"

	if encoded != expected {
		t.Errorf("EncodeURL(%s) = %s; want %s", text, encoded, expected)
	}
}

// TestDecodeURL tests the DecodeURL function
func TestDecodeURL(t *testing.T) {
	text := "urn%3Ali%3Aorganization%3A123456"
	decoded, err := DecodeURL(text)
	expected := "urn:li:organization:123456"

	if decoded != expected || err != nil {
		t.Errorf("DecodeURL(%s) = %s, %v; want %s, nil", text, decoded, err, expected)
	}
}
