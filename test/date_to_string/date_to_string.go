package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func main() {
	pbody := make(map[string]string)
	dates := []string{"1404-04-26", "1404-04-27", "1404-04-28", "1404-04-29", "1404-04-30", "1404-04-31"}
	// jdates := pkg.DatesToJalali(dates, false)
	for i, date := range dates {
		pbody[fmt.Sprintf("Date_%v", i)] = date
	}
	mbody, _ := json.Marshal(pbody)
	re := regexp.MustCompile(`Date_\d+`)
	ss := re.ReplaceAllString(string(mbody), "Date")
	fmt.Println(ss)
}
