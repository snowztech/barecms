package handlers

import "testing"

func TestPaginationParameterBounds(t *testing.T) {
	if value, err := paginationParameter("limit", "", 20, 1, 100); err != nil || value != 20 {
		t.Fatalf("unexpected default: %d %v", value, err)
	}
	for _, raw := range []string{"0", "101", "nope"} {
		if _, err := paginationParameter("limit", raw, 20, 1, 100); err == nil {
			t.Fatalf("expected %q to be rejected", raw)
		}
	}
}
