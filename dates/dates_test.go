package dates

import "testing"

func TestFindDate(t *testing.T) {
	content := "Today's date is 18/12/2019"
	want := "18/12/2019"
	got, err := FindDate(content)
	if err != nil {
		t.Errorf("Failed to find date, error: %q", err)
	}
	if got != want {
		t.Errorf("Failed to find correct date, got %q", got)
	}
}

func TestFindDateReturnsErrorWhenDateNotFound(t *testing.T) {
	content := "Star Wars: The Rise of Skywalker is out tomorrow"
	_, err := FindDate(content)
	if err == nil {
		t.Errorf("Should have non-nil error")
	}
}
