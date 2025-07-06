// unit test

package address_test

import (
	"testing"
	. "tests/address"
)

func TestAddressType(t *testing.T) {
	t.Parallel()
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

func TestExample(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("This test should not pass")
	}
}
