// unit test

package address

import (
	"testing"
)

func TestAddressType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Street 123", "Street"},
		{"Avenue 456", "Avenue"},
		{"Road 789", "Road"},
		{"Boulevard 101", "Unknown"},
	}

	for _, test := range tests {
		result := AddressType(test.input)
		if result != test.expected {
			t.Errorf("AddressType(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
