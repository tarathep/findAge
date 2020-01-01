package findAge

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindDiffwith15061996TH(t *testing.T) {
	input := []string{"15", "06", "1996", "TH"}
	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")
	err, output := Age{timeMock}.Find(input)
	assert.Equal(t, "23 ปี  6 เดือน  16 วัน", output)
	assert.NotEqual(t, nil, err)
}

func TestFindDiffwith15061996EN(t *testing.T) {
	input := []string{"15", "06", "1996", "EN"}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, output := Age{timeMock}.Find(input)
	assert.Equal(t, "23 years  6 months  16 days", output)
	assert.NotEqual(t, nil, err)
}

func TestFindDiffwith15061996Default(t *testing.T) {
	input := []string{"15", "06", "1996"}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, output := Age{timeMock}.Find(input)
	assert.Equal(t, "23 years  6 months  16 days", output)
	assert.NotEqual(t, nil, err)
}

func TestFindDiffErrorDaywithx(t *testing.T) {
	input := []string{"x", "06", "1996", "EN"}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, _ := Age{timeMock}.Find(input)
	if err == nil {
		t.Error()
	}
	assert.Equal(t, "error input date range format", err[0].Error())
}

func TestFindDiffErrorMonthwithy(t *testing.T) {
	input := []string{"01", "y", "1996", "EN"}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, _ := Age{timeMock}.Find(input)
	if err == nil {
		t.Error()
	}
	assert.Equal(t, "error input month range format", err[0].Error())
}

func TestFindDiffErrorYearwithz(t *testing.T) {
	input := []string{"01", "06", "z", "EN"}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, _ := Age{timeMock}.Find(input)
	if err == nil {
		t.Error()
	}
	assert.Equal(t, "error input parsing year", err[0].Error())
}

func TestFindDiffErrorPatternwithAllParser(t *testing.T) {
	input := []string{"", "", ""}

	timeMock, _ := time.Parse("02-01-2006", "01-01-2020")

	err, _ := Age{timeMock}.Find(input)
	if err == nil {
		t.Error()
	}
	assert.Equal(t, "error input parsing year", err[0].Error())
	assert.Equal(t, "error input date range format", err[1].Error())
	assert.Equal(t, "error input month range format", err[2].Error())

}
