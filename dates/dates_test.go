package dates

import (
	"testing"
	"time"
)

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

func TestConvertDateStringToTime(t *testing.T) {
	dateStr := "20/12/2019"
	_, err := ConvertToTime(dateStr)
	if err != nil {
		t.Errorf("Error parsing date string: %q", err)
	}
}

func TestExtractShortDate(t *testing.T) {
	dateStr := "20/12/2019"
	datetime, _ := ConvertToTime(dateStr)
	shortDateStr := ExtractShortDate(datetime)
	if dateStr != shortDateStr {
		t.Errorf("Extract short date failed, got %q, want %q", shortDateStr, dateStr)
	}
}

func TestAddingDays(t *testing.T) {
	dateStr, expected := "20/12/2019", "23/12/2019"
	datetime, _ := ConvertToTime(dateStr)
	datetime = AddDays(datetime, 3)
	shortDateStr := ExtractShortDate(datetime)
	if expected != shortDateStr {
		t.Errorf("AddDays failed, got %q, want %q", shortDateStr, expected)
	}

}

func TestElapsedTime(t *testing.T) {
	now := time.Now()
	delta := now.Add(-time.Second * 10)
	expected := "10 secs ago"
	converted := ConvertToTimeElapsed(delta)
	if expected != converted {
		t.Errorf("Elapsed time failed, got %q, want %q", converted, expected)
	}

	now = time.Now()
	delta = now.Add(-time.Minute * 30)
	expected = "30 mins ago"
	converted = ConvertToTimeElapsed(delta)
	if expected != converted {
		t.Errorf("Elapsed time failed, got %q, want %q", converted, expected)
	}

	now = time.Now()
	delta = now.Add(-time.Hour * 5)
	expected = "5 hours ago"
	converted = ConvertToTimeElapsed(delta)
	if expected != converted {
		t.Errorf("Elapsed time failed, got %q, want %q", converted, expected)
	}

	now = time.Now()
	delta = now.Add(-time.Hour * 36)
	expected = "1 days ago"
	converted = ConvertToTimeElapsed(delta)
	if expected != converted {
		t.Errorf("Elapsed time failed, got %q, want %q", converted, expected)
	}

}
