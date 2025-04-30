package pkg

import (
	"fmt"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func DatesToIso(dates []string) []string {
	var formattedDates []string
	for _, item := range dates {
		formattedDates = append(formattedDates, fmt.Sprintf("%vT00:00:00.000Z", item))
	}
	return formattedDates
}

func DatesToJalali(dates []string, dash bool) []string {
	var jdates []string
	for _, date := range dates {
		parsedTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil
		}

		// Now convert to Jalali
		ptobj := ptime.New(parsedTime)

		// Format it as jYY-jMM-jDD
		var jalaliDate string
		if dash {
			jalaliDate = ptobj.Format("yyyy-MM-dd")
		} else {
			jalaliDate = ptobj.Format("yyyy/MM/dd")
		}
		jdates = append(jdates, jalaliDate)
		fmt.Println(jdates)
	}
	return jdates
}
