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

func userInput(){

	//lucky nummbers input 
	var userNumbers[5] int 
	var userNumber = 1

	for i := 0; i < len(userNumbers); i++ {
		
		fmt.Println("Please enter your", i+1 , ". Lotto nummbers : >> ")
		fmt.Scanf("%d", &userNumber)
		
		// Forcing to user give numbers between 1 and 50 
		if userNumber <= 50 && userNumber > 0 {		
			userNumbers[i] = userNumber
		} else {
		fmt.Println("INVALID NUMMBER ! Please give a nummber between 1 and 50 ")
		i--
		}
	}

	//lucky STAR nummbers input 
	var userStars[2] int 
	var userStar int
	
	for i := 0; i < len(userStars); i++ {
		
		fmt.Println("Please enter your", i+1 , ". LOTTO STAR nummbers : >> ")
		fmt.Scanf("%d", &userStar)
		
		// Forcing to user give numbers between 1 and 12 
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
fmt.Printf("your Lucky nummbers : \t %2d \t", userNumbers)
fmt.Printf("your Lucky Star nummbers : \t%2d\n", userStars)
}

func main() {

	userInput()
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
	fmt.Printf("Winning Nummbers : \t %2d \t" , lottaryUniqList)
	fmt.Printf("Winning Star Numbers : \t\t%2d\n", starUniqList)

}
