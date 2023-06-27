package parsing

import (
	"fmt"
	"strings"
	"time"
)

func Parsing(jamS string) {
	fmt.Println("Input Format HH:MM:SS PM/AM :", jamS)
	format := strings.Split(jamS, " ")
	var jam time.Time
	var err error

	if format[1] == "PM" {
		jam, err = time.Parse("03:04:05 PM", jamS)
	} else if format[1] == "AM" {
		jam, err = time.Parse("03:04:05 AM", jamS)
	}

	if err != nil {
		fmt.Println("Invalid Input Format")
		return
	}

	jamS2 := jam.Format("15:04:05")
	fmt.Println("Waktu Dalam Format 24 Jam : ", jamS2)
}
