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

type void struct{}

// missing compares two slices and returns slice of differences
func compare(a, b []int) []int {
	// create map with length of the 'a' slice
	ma := make(map[int]void, len(a))
	matchs := []int{}
	// Convert first slice to map with empty struct (0 bytes)
	for _, ka := range a {
		ma[ka] = void{}
	}
	// find matching values in a
	for _, kb := range b {
		if _, ok := ma[kb]; ok {
			matchs = append(matchs, kb)
		}
	}
	return matchs
}

func userInput() {

}

func main() {

	//lucky nummbers input
	var userNumbers [5]int
	var userNumber int

	for i := 0; i < len(userNumbers); i++ {

		fmt.Printf("Please enter your %d. Lotto nummbers : >> ", i+1)
		fmt.Scanf("%d", &userNumber)

		// Forcing to user give numbers between 1 and 50
		if userNumber <= 50 && userNumber > 0 {
			userNumbers[i] = userNumber

			// Here we need AVOID DOUBLE ENTRIES FUNCTION
		} else {
			fmt.Printf("\nINVALID NUMMBER ! Please give a nummber between 1 and 50 \n\n")
			i--
		}

	}

	//lucky STAR nummbers input
	var userStars [2]int
	var userStar int

	for i := 0; i < len(userStars); i++ {

		fmt.Print("Please enter your \t", i+1, ". LOTTO STAR nummbers : >> ")
		fmt.Scanf("%d", &userStar)

		// Forcing user to give only numbers between 1 and 12
		if userStar <= 12 && userStar > 0 {
			userStars[i] = userStar
		} else {
			fmt.Println("INVALID NUMMBER ! Please give a nummber between 1 and 12 ")
			i--
		}

	}

	// Sorting all Values
	sort.Ints(userNumbers[:])
	sort.Ints(userStars[:])

	// Display Luky and Star nummbers
	fmt.Printf("\n***********************************************************************\n\n")
	fmt.Printf("YOUR NUMBERS :\t\t%2d\t", userNumbers)
	fmt.Printf("YOUR STARS:\t%2d\n", userStars)

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

	// Sorting winnings numbers
	sort.Ints(lottaryUniqList)
	sort.Ints(starUniqList)

	// Display Winning Numbers
	fmt.Printf("WINNING NUMBRES :\t%2d\t", lottaryUniqList)
	fmt.Printf("WINNING STARS :\t%2d\n", starUniqList)
	fmt.Printf("\n***********************************************************************\n\n")

	comparedNumbers := compare(lottaryUniqList, userNumbers[:])
	comparedStars:= compare(starUniqList, userStars[:])

	fmt.Printf("MATCHED LUCKY NUMBERS :\t%2d\n",comparedNumbers)

	fmt.Printf("MATCHED LUCKY STARS :\t%2d\n\n", comparedStars)
	fmt.Printf("***********************************************************************\n")
	

}
