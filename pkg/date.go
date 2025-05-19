package pkg

import (
	"fmt"
	"sort"
	"time"

	"github.com/mshafiee/jalali"
)

func SortDates(dates []string) []string {
	cDates := make([]string, len(dates))
	copy(cDates, dates)
	if len(cDates) > 1 {
		sort.Strings(cDates)
	}
	return cDates
}

func DatesToIso(dates []string) []string {
	var formattedDates []string
	for _, item := range dates {
		formattedDates = append(formattedDates, fmt.Sprintf("%vT00:00:00.000Z", item))
	}
	sort.Strings(formattedDates)
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
	sort.Strings(jdates)
	fmt.Println("Jalali Dates:", jdates)
	return jdates
}
