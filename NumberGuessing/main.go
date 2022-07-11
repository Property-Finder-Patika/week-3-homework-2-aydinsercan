package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/* Desc:
Implement a number-guessing game in which the computer computes a four digit number as a secret number and
a player tries to guess that number correctly.
Player would enter her guess and the computer would produce a feedback on the positions of the digits.
Four-digit number can't start with 0 and have repeating digits.

Let's say the computer computes 2658 as a secret number to be guessed by the player.
When player enters her guess such as 1234, the computer would display -1 meaning that only one digit of 1234 exist in the secret number and its position is wrong.
When the player enters 5678 the similarly the computer displays +2-1.

And the game goes on until the player correctly guess the secret number and the computer displays +4.
The game also keeps track of all guesses entered by the players so far and lists them when it displays its feedback to the player
so that the player can compute her next guess correctly.
*/
func main() {
	//Generating number from computer via rand function
	rand.Seed(time.Now().UnixNano())
	generatedNumber := rand.Intn(9999-1000+1) + 1000
	fmt.Println("Generated Number : ", generatedNumber)
	strGeneratedNumber := strconv.Itoa(generatedNumber)

	var guessedNum string
	var existCount int
	var positionCount int

	for {
		existCount, positionCount = 0, 0

		fmt.Println("Please Enter 4-Digit Number")
		fmt.Scanln(&guessedNum)

		//Must enter 4-Digit number
		if len(guessedNum) > 4 {
			fmt.Println("Error, you should've 4-Digit number")
			continue
		}

		//Comparing numbers digit by digit
		for i := 0; i < len(strGeneratedNumber); i++ {
			for j := 0; j < len(guessedNum); j++ {

				if i == j && strGeneratedNumber[i] == guessedNum[j] {
					positionCount++
				}
				if strGeneratedNumber[i] == guessedNum[j] && i != j {
					existCount++
				}
			}
		}

		//Displaying feedback to the player
		if positionCount != 0 {
			fmt.Printf("+%d", positionCount)
		}
		if existCount != 0 {
			fmt.Printf("-%d", existCount)
		}
		fmt.Println()

		//Term condition
		if positionCount == 4 {
			fmt.Println("Congrats!!")
			break
		}
	}
}

/*
Generated Number :  8730
Please Enter 4-Digit Number
8563
+1-1
Please Enter 4-Digit Number
5432
+1
Please Enter 4-Digit Number
6582
-1
Please Enter 4-Digit Number
0378
-4
Please Enter 4-Digit Number
8630
+3
Please Enter 4-Digit Number
8730
+4
Congrats!!
*/
