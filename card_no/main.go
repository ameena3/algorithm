package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	prefixVisa := []int{4, 1, 0, 0}
	cards := make([]int64, 0)
	timeNow := time.Now()
	for len(cards) <= 10000000 {
		card_no := append(prefixVisa, generateRand()...)
		var sum int
		for i, ci := range card_no {
			if i%2 == 0 {
				nci := ci * 2
				sum += ((int(nci) % 10) + (int(nci/10) % 10))
			} else {
				sum += ci
			}
		}
		for i := 0; i <= 1000; i++ {
			suffix := make([]int, 4)
			suffix[3] = int(i) % 10
			suffix[2] = int(i/10) % 10
			suffix[1] = int(i/100) % 10
			suffix[0] = int(i/1000) % 10
			sumFinal := sum
			for i, ci := range suffix {
				if i%2 == 0 {
					nci := ci * 2
					sumFinal += ((int(nci) % 10) + (int(nci/10) % 10))
				} else {
					sumFinal += ci
				}
			}
			if sumFinal%10 == 0 {
				card_no_final := append(card_no, suffix...)
				fmt.Println("Found card no:")
				var c_no_int int64
				for i, value := range card_no_final {
					c_no_int += int64(int64(math.Pow10(15-i)) * int64(value))
				}
				cards = append(cards, c_no_int)
				fmt.Println(c_no_int)
			}
		}
	}
	timeSince := time.Since(timeNow)
	fmt.Printf("time taken was %#v mins\n", timeSince.Minutes())
}

func generateRand() []int {
	r := rand.Intn(10000000)
	ans := make([]int, 8)
	ans[7] = int(r) % 10
	ans[6] = int(r/10) % 10
	ans[5] = int(r/100) % 10
	ans[4] = int(r/1000) % 10
	ans[3] = int(r/10000) % 10
	ans[2] = int(r/100000) % 10
	ans[1] = int(r/1000000) % 10
	ans[0] = int(r/10000000) % 10
	return ans
}
