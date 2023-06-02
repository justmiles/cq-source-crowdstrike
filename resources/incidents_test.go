package resources

import "testing"

func TestRun(t *testing.T) {

	got := Incidents()

	if got != nil {
		t.Errorf("got %q, wanted %T", got, nil)
	}
}
