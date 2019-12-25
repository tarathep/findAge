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
		/*
		if len(input) != 4 {
			return errors.New("error input year format (yyyy)")
		}
		*/

		iYear, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("error input convert year format")
		}
		if !(iYear > 0) {
			return errors.New("error input year format")
		}
	case 3:
		if !(input == "EN" || input == "TH") {
			return errors.New("error input option language format")
		}
	}
	return nil
}

func Find(inputs []string) string {
	if !(len(inputs) == 4 || len(inputs) == 3) {
		fmt.Println("error input params")
		return ""
	}
	day := strings.Trim(inputs[0], "\n")
	month := strings.Trim(inputs[1], "\n")
	year := strings.Trim(inputs[2], "\n")
	var language string = "EN"

	// CHECK ERROR AND PRINT
	var errTmp error = nil
	if err := validation(day, 0); err != nil {
		fmt.Println(err.Error())
		errTmp = err
	}
	if err := validation(month, 1); err != nil {
		fmt.Println(err.Error())
		errTmp = err
	}
	if err := validation(year, 2); err != nil {
		fmt.Println(err.Error())
		errTmp = err
	}
	if len(inputs) == 4 {
		//input EN/TH
		language = strings.ToUpper(strings.Trim(inputs[3], "\n"))

		if err := validation(language, 3); err != nil {
			fmt.Println(err.Error())
			errTmp = err
		}
	}
	//break flow
	if errTmp != nil {
		return ""
	}

	/*
		if language == "TH" {
			y, _ := strconv.Atoi(year)
			year = fmt.Sprint((y - 543))
		}
	*/

	start, _ := time.Parse("02-01-2006", day+"-"+month+"-"+year)
	diffYear, diffMonth, diffDay, _, _, _ := diff(start, time.Now())
	if diffYear < 0 {
		return "error input year"
	}
	if language == "TH" {
		return fmt.Sprint(diffYear) + " ปี  " + fmt.Sprint(diffMonth) + " เดือน  " + fmt.Sprint(diffDay) + " วัน"
	} else {
		return fmt.Sprint(diffYear) + " year(s)  " + fmt.Sprint(diffMonth) + " month(s)  " + fmt.Sprint(diffDay) + " day(s)"
	}
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {

	//location zone time
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	/*
		if a.After(b) {
			a, b = b, a
		}
	*/
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
