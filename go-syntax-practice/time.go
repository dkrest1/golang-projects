package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	
	// 1704454116
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))

	// Friday, 05-Jan-24 12:29:41 WAT
	fmt.Println(time.Now().Format(time.RFC850))

	fmt.Println(time.Now().Format(time.RFC1123Z))

	fmt.Println(time.Now().Format("20060102150405"))

	fmt.Println(time.Unix(1401403874, 0).Format("02.01.2006 15:04:05"))

	fmt.Println(time.Unix(123456789, 0).Format(time.RFC3339))

	fmt.Println(time.Unix(1234567890, 0).Format(time.RFC822))
}