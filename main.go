package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Function that avoid double entries in WINNING Numbers Slices
func unique(intSlice []int) []int {

	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	// random Lottery Nummbers MIN, MAX and UNIT value declaration
	lotteryMin := 1
	lotteryMax := 50
	lotteryUnits := 5

	// random Lottery STAR Nummbers MIN, MAX and UNIT value declaration
	starMin := 1
	starMax := 12
	starUnits := 2

	// generate random Lottary numbers up to current time
	rand.Seed(time.Now().UnixNano())

	// variable declarations
	var lottaryList []int
	var lottaryUniqList []int

	var starList []int
	var starUniqList []int

	// set a winning numbers randomly
	for i := 0; len(lottaryUniqList) < lotteryUnits; i++ {

		lottaryNumbers := rand.Intn(lotteryMax-lotteryMin+1) + lotteryMin
		lottaryList = append(lottaryList, lottaryNumbers)
		lottaryUniqList = unique(lottaryList)
	}

	// set a winning Star numbers randomly
	for i := 0; len(starUniqList) < starUnits; i++ {

		starNumbers := rand.Intn(starMax-starMin+1) + starMin
		starList = append(starList, starNumbers)
		starUniqList = unique(starList)
	}

	// Sorting winnings random numbers
	sort.Ints(lottaryUniqList)
	sort.Ints(starUniqList)

	// Display Winning Numbers
	fmt.Println("Winning Nummbers : \t", lottaryUniqList)
	fmt.Println("Winning Star Numbers : \t", starUniqList)

}
