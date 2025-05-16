package main

import (
	"fmt"

	shab_dto "github.com/amirhosseinf79/renthub_service/internal/dto/shab"
	"github.com/amirhosseinf79/renthub_service/pkg"
)

func testGen(dates []string) shab_dto.CalendarBody {
	return shab_dto.CalendarBody{
		Action: "aaaa",
		Dates:  pkg.DatesToJalali(dates, true),
	}
}

func main() {
	dates := []string{"2025-06-18", "2025-06-19", "2025-06-20", "2025-06-21"}
	test := make([]int, 100)
	for range test {
		fmt.Println(testGen(dates))
	}
}
