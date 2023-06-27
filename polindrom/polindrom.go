package polindrom

import (
	"fmt"
	"strings"
)

func Polindrom(kalimat string) bool {
	fmt.Println("Kata/Kalimat Yang di Input : ", kalimat)
	kalimat = strings.ToLower(kalimat)
	kalimat = strings.ReplaceAll(kalimat, " ", "")

	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] != kalimat[len(kalimat)-(i+1)] {
			return false
		}
	}

	return true
}
