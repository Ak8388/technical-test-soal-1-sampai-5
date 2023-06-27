package filecount

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Count() {
	dir, err := os.Open("C:/Users/akbar/Documents/PelatihanGoLanjutan/tecnical-test/hallo.txt")

	if err != nil {
		fmt.Println("Failed Accses Directory : ", err)
		return
	}
	defer dir.Close()

	scanner := bufio.NewScanner(dir)
	scanner.Split(bufio.ScanLines)

	var sumNilai, jumlahAngka int
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			sumNilai += i
			jumlahAngka++
		}
	}
	fmt.Printf("Total angka pada file: %d\nJumlah semua angka: %d\n", jumlahAngka, sumNilai)
}
