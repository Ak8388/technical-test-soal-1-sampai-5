package rekrusif

import (
	"math/rand"
	"time"
)

func Acak(index int, nilai ...int) {
	rand.Seed(time.Now().UnixNano())

	if index == 0 {
		return
	}
	in := rand.Intn(index + 1)
	nilai[index-1], nilai[in] = nilai[in], nilai[index-1]
	Acak(index-1, nilai...)
}
