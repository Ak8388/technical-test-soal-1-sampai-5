package main

import (
	"fmt"
	soal1 "test/filecount"
	soal5 "test/parsing"
	soal2 "test/polindrom"
	soal4 "test/rekrusif"
)

func main() {

	//Soal 1
	soal1.Count()

	fmt.Println("")

	//Soal 2
	hasil := soal2.Polindrom("andri")

	if hasil == true {
		fmt.Println("Merupakan Polindrom")
	} else {
		fmt.Println("Bukan Merupakan Polindrom")
	}

	fmt.Println("")

	//Soal 4
	var nilai = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soal4.Acak(len(nilai), nilai...)
	fmt.Println(nilai)

	fmt.Println("")

	//Soal 5
	soal5.Parsing("05:30:40 PM")
}
