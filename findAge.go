package findAge

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func validation(input string, offset int) error {
	switch offset {
	case 0:
		iDate, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("error input convert date format")
		}
		if !(iDate > 0 && iDate < 32) {
			return errors.New("error input date range format")
		}
	case 1:
		iMonth, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("error input convert month format")
		}
		if !(iMonth > 0 && iMonth < 13) {
			return errors.New("error input month range format")
		}
	case 2:

		iYear, err := strconv.Atoi(input)

		if err != nil {
			return errors.New("error input convert year format")
		}
		if iYear <= 0 {
			return errors.New("error input year format")

		}
		if iYear > time.Now().Year() {
			return errors.New("error input year is over")
		}
	}
	return nil
}

//Find is หาอายุว่าเท่าไหร่แล้ว??
func Find(inputs []string) ([]error, string) {
	//collecting err
	var errs []error = nil

	//Check Input Params date,moth,year,..
	if !(len(inputs) >= 3) {
		errs = append(errs, errors.New("error input params"))
		return errs, ""
	}

	//Convert DATE parse to int
	daystring := strings.Trim(inputs[0], "\n")
	dayint, _ := strconv.Atoi(daystring)
	day := strconv.Itoa(dayint)
	if len(day) < 2 {
		day = "0" + day
	}

	//Convert MONTH parse to int
	monthstring := strings.Trim(inputs[1], "\n")
	monthint, _ := strconv.Atoi(monthstring)
	month := strconv.Itoa(monthint)
	if len(month) < 2 {
		month = "0" + month
	}

	//Convert DATE parse to int
	yearstring := strings.Trim(inputs[2], "\n")
	yearstring = strings.Trim(yearstring, "\r")
	yearint, err := strconv.Atoi(yearstring)
	if err != nil {
		errs = append(errs, errors.New("error input parsing year"))
		return errs, ""
	}
	year := strconv.Itoa(yearint)

	if len(year) < 2 {
		year = "000" + year
	}

	if len(year) < 3 {
		year = "00" + year
	}

	if len(year) < 4 {
		year = "0" + year
	}

	// CHECK ERROR
	var errTmp error = nil
	if err := validation(day, 0); err != nil {
		errs = append(errs, err)
		errTmp = err
	}
	if err := validation(month, 1); err != nil {
		errs = append(errs, err)
		errTmp = err

	}
	if err := validation(year, 2); err != nil {
		errs = append(errs, err)
		errTmp = err
	}
	//Default Option Language
	language := "EN"

	//check si
	if len(inputs) > 3 {
		//Convert option lang to uppercase
		language = strings.ToUpper(strings.Trim(inputs[3], "\n"))
	}

	//break flow when has error not pass in validate
	if errTmp != nil {
		return errs, ""
	}
	y, _ := strconv.Atoi(year)
	m, _ := strconv.Atoi(month)
	d, _ := strconv.Atoi(day)

	//
	if daysIn(m, y) < d && d > 0 {
		errs = append(errs, errors.New("error Date doesn't have in calendar"))
		return errs, ""
	}

	start, _ := time.Parse("02-01-2006", day+"-"+month+"-"+year)
	diffYear, diffMonth, diffDay, _, _, _ := Diff(start, time.Now())

	//CHECK nagative value year <1
	if diffYear < 0 {
		errs = append(errs, errors.New("error input is over range"))
		return errs, ""
	}

	//CHECK OPTION OUTPUT LANGUAGE
	if language == "TH" {
		return nil, fmt.Sprint(diffYear) + " ปี  " + fmt.Sprint(diffMonth) + " เดือน  " + fmt.Sprint(diffDay) + " วัน"
	}

	formatY, formatM, formatD := " year  ", " month  ", " day"
	if diffYear > 1 {
		formatY = " years  "
	}
	if diffMonth > 1 {
		formatM = " months  "
	}
	if diffDay > 1 {
		formatD = " days"
	}
	return nil, fmt.Sprint(diffYear) + formatY + fmt.Sprint(diffMonth) + formatM + fmt.Sprint(diffDay) + formatD

}

//Diff is Defference Time Between A to B time
func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {

	//location zone time
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}

	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func daysIn(month int, year int) int {
	m := time.Month(month)
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
