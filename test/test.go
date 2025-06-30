package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	method := "POST"
	url := "/auth/otp/send"
	body := `{"mobile":"09334429096","iso2":"IR"}`
	xRequestT := fmt.Sprintf("%d", 1751233777)
	// xRequestT := fmt.Sprintf("%d", time.Now().Unix())

	md5Sum := md5.Sum([]byte(body))
	xRequestB := hex.EncodeToString(md5Sum[:])

	plusStr := "22ab046bd59212160e654bd5a610eb8da87723239db6d059f3074ad451c667b2"

	raw := strings.ToUpper(method) + "/api" + url + xRequestB + xRequestT + plusStr
	fmt.Printf(">>>%s<<<\n", raw)

	md5Raw := md5.Sum([]byte(raw))
	xRequestH := hex.EncodeToString(md5Raw[:])

	fmt.Println("body:", body)
	fmt.Println("row:", raw)

	fmt.Println("x-request-t:", xRequestT)
	fmt.Println("x-request-b:", xRequestB) // c2b3c1136813178ed291ee77ae92417d
	fmt.Println("x-request-h:", xRequestH) // 4452a4ef340b835a57cbc5950d0f5fae
}
