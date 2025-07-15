package pkg

import (
	"fmt"
	"time"

	"github.com/mshafiee/jalali"
)

func SeperateDates(dates []string) (result [][]string) {
	layout := "2006-01-02"
	var parsedDates []time.Time
	for _, d := range dates {
		t, _ := time.Parse(layout, d)
		parsedDates = append(parsedDates, t)
	}

	var group []string
	for i := 0; i < len(parsedDates); i++ {
		if i == 0 {
			group = append(group, parsedDates[i].Format(layout))
			continue
		}
		diff := parsedDates[i].Sub(parsedDates[i-1]).Hours() / 24
		if diff == 1 {
			group = append(group, parsedDates[i].Format(layout))
		} else {
			result = append(result, group)
			group = []string{parsedDates[i].Format(layout)}
		}
	}

	if len(group) > 0 {
		result = append(result, group)
	}
	return
}

func DatesToIso(dates []string) []string {
	var formattedDates []string
	for _, item := range dates {
		formattedDates = append(formattedDates, fmt.Sprintf("%vT00:00:00.000Z", item))
	}
	return formattedDates
}

func DatesToJalali(dates []string, dash bool) []string {
	fmt.Println("Input Dates:", dates)
	var jdates []string
	for _, date := range dates {
		parsedTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil
		}

		// Now convert to Jalali
		ptobj := jalali.ToJalali(parsedTime)

		// Format it as jYY-jMM-jDD
		var jalaliDate string
		if dash {
			jalaliDate = ptobj.Format("%Y-%m-%d")
		} else {
			jalaliDate = ptobj.Format("%Y/%m/%d")
		}
		jdates = append(jdates, jalaliDate)
	}
	fmt.Println("Jalali Dates:", jdates)
	return jdates
}
