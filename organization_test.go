package linkedin

import "testing"

// TestGetOrganizationIDFromElement tests the ElementOrganization.GetOrganizationID function
func TestGetOrganizationIDFromElement(t *testing.T) {
	// create a new organization element
	e := ElementOrganization{
		Organization: "urn:li:organization:123456789",
	}

	// get the organization ID
	organizationID := e.GetOrganizationID()
	expected := "123456789"

	if organizationID != expected {
		t.Errorf("GetOrganizationID() = %s; want %s", organizationID, expected)
	}
}
