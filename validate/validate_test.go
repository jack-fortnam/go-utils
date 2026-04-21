package validate

import "testing"

func TestEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid simple", "test@example.com", true},
		{"valid dot", "john.doe@gmail.com", true},
		{"valid plus", "user+tag@domain.co.uk", true},

		{"missing at", "testexample.com", false},
		{"missing domain", "test@", false},
		{"double dot", "test..name@example.com", false},
		{"empty", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Email(tt.email)
			got := err == nil

			if got != tt.want {
				t.Errorf("Email(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestPhoneNumber(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{"valid uk style", "+447123456789", true},
		{"valid us", "+15551234567", true},

		{"missing plus", "447123456789", false},
		{"starts with zero", "+0123456789", false},
		{"too short", "+12345", false},
		{"letters", "+44abc123456", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PhoneNumber(tt.phone)
			got := err == nil

			if got != tt.want {
				t.Errorf("PhoneNumber(%q) = %v, want %v", tt.phone, got, tt.want)
			}
		})
	}
}

func TestPostCode(t *testing.T) {
	tests := []struct {
		name     string
		postcode string
		want     bool
	}{
		{"valid SW1A 1AA", "SW1A 1AA", true},
		{"valid no space SW1A1AA", "SW1A1AA", true},
		{"valid EC1A 1BB", "EC1A 1BB", true},
		{"valid M1 1AE", "M1 1AE", true},
		{"valid B33 8TH", "B33 8TH", true},
		{"valid lower case", "sw1a 1aa", true},

		{"empty", "", false},
		{"missing inward code", "SW1A", false},
		{"missing outward code", "1AA", false},
		{"too short", "SW1", false},
		{"invalid characters", "SW1A @1AA", false},
		{"double space", "SW1A  1AA", true},
		{"invalid format", "123 ABC", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCode(tt.postcode)
			got := err == nil

			if got != tt.want {
				t.Errorf("PostCode(%q) = %v, want %v", tt.postcode, got, tt.want)
			}
		})
	}
}

func TestCard(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid cases
		{"Standard Luhn Example", "79927398713", true},
		{"With Hyphens", "7992-7398-713", true},
		{"With Spaces", "7992 7398 713", true},

		// Invalid cases
		{"Incorrect Check Digit", "79927398710", false},
		{"Off by One", "49927398717", false},
		{"Too Short", "1", false},
		{"Empty String", "", false},

		// Error cases
		{"Alphanumeric", "7992739871a", false},
		{"Symbols", "7992739871!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Card(tt.input)
			if result != tt.expected {
				t.Errorf("Card(%s) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUrl(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid cases
		{"Valid URL", "https://google.com", true},
		{"localhost", "http://localhost:8080", true},
		{"complex", "https://sub.domain.example.co.uk/path/to/resource?query=1", true},
		{"different protocol", "ftp://fileserver.local", true},

		//Invalid
		{"missing scheme", "google.com", false},
		{"missing host", "https://", false},
		{"malformed schema", "://example.com", false},
		{"Security!", "javascript:alert('xss')", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Card(tt.input)
			if result != tt.expected {
				t.Errorf("URL(%s) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
