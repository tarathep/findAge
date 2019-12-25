package findAge

import (
	"testing"
	"time"
)

func TestFindDiffTH(t *testing.T) {
	input := []string{"15", "06", "2539", "TH"}
	if (Find(input)) != "23 ปี  6 เดือน  9 วัน" {
		t.Error()
	}

}
func TestFindDiffEN(t *testing.T) {
	input := []string{"15", "06", "1996", "EN"}
	if (Find(input)) != "23 year(s)  6 month(s)  9 day(s)" {
		t.Error()
	}
}

func TestValidationDay(t *testing.T) {
	if validation("15", 0) != nil {
		t.Error()
	}
}
func TestValidationMonth(t *testing.T) {
	if validation("06", 1) != nil {
		t.Error()
	}
}
func TestValidationYear(t *testing.T) {
	if validation("1999", 2) != nil {
		t.Error()
	}
}
func TestValidationLanguage(t *testing.T) {
	if validation("EN", 3) != nil {
		t.Error()
	}
}

func TestValidationDayERROR(t *testing.T) {
	if validation("99", 0) == nil {
		t.Error()
	}
}
func TestValidationMonthERROR(t *testing.T) {
	if validation("-88", 1) == nil {
		t.Error()
	}
}
func TestValidationYearERROR(t *testing.T) {
	if validation("xx", 2) == nil {
		t.Error()
	}
}
func TestValidationLanguageERROR(t *testing.T) {
	if validation("xyz", 3) == nil {
		t.Error()
	}
}

func TestDiff(t *testing.T) {
	start, _ := time.Parse("02-01-2006", "15-06-2539")
	now, _ := time.Parse("02-01-2006", "18-08-2540")
	year, month, day, _, _, _ := diff(start, now)
	if !(year == 1 && month == 2 && day == 3) {
		t.Error()
	}
}
