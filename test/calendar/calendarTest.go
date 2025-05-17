package main

import (
	"fmt"

	"github.com/amirhosseinf79/renthub_service/pkg"
)

func main() {
	dates := []string{"2025-06-20", "2025-06-21", "2025-06-19", "2025-06-18"}
	dates2 := []string{"2025-06-18", "2025-06-19", "2025-06-20", "2025-06-21"}
	test := make([]int, 100)
	for range test {
		go fmt.Println(pkg.DatesToJalali(dates, true))
		go fmt.Println(pkg.DatesToJalali(dates2, false))
	}
}
