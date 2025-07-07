package auth

import (
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	header := make(map[string][]string)

	s, e := GetAPIKey(header)
	if !reflect.DeepEqual(s, "") {
		// TODO: Complete
		t.Fatalf("expected: got:")
	}
	if e != ErrNoAuthHeaderIncluded {
		// TODO: Complete
		t.Fatalf("expected: got:")
	}
}
