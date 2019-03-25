package string

import "testing"

func Test_String(t *testing.T) {
	var tests = []struct {
		w, want string
	}{
		{"Backward", "drawkcaB"},
		{"", ""},
		{"sonu123", "321unos"},
	}

	for _, c := range tests {
		got := Reverse(c.w)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.w, got, c.want)
		}
	}
}
