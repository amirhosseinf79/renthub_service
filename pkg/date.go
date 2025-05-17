package pkg

import (
	"fmt"
	"time"

	"github.com/mshafiee/jalali"
)

func DatesToIso(dates []string) []string {
	var formattedDates []string
	for _, item := range dates {
		formattedDates = append(formattedDates, fmt.Sprintf("%vT00:00:00.000Z", item))
	}
	return formattedDates
}

func DatesToJalali(dates []string, dash bool) []string {
	copiedDates := make([]string, len(dates))
	copy(copiedDates, dates)
	fmt.Println("Input Dates:", copiedDates)
	var jdates []string
	for _, date := range copiedDates {
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
