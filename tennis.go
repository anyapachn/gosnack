package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CurrentScore(score int) string {
	word := []string{"Love", "Fifteen", "Thirty", "Forty"}
	return word[score]
}

func main() {
	var scoreA, scoreB int
	var wordA, wordB string

	fmt.Println("Start new game")
	wordA = CurrentScore(scoreA)
	wordB = CurrentScore(scoreB)
	fmt.Println(scoreA, " - ", scoreB, " : ", wordA, " - ", wordB)

	for i := 0; i < 9; i++ {
		rand.Seed(time.Now().UnixNano())
		score := rand.Intn(100)
		//fmt.Println(score)
		if score%2 != 0 {
			scoreA++
			//fmt.Println("A ", scoreA)
			if scoreA < 4 {
				wordA = CurrentScore(scoreA)
			}
		} else {
			scoreB++
			//fmt.Println("B ", scoreB)
			if scoreB < 4 {
				wordB = CurrentScore(scoreB)
			}
		}
		switch {
		case scoreA == 4:
			fmt.Println("Game A Win")
			i = 9
		case scoreB == 4:
			fmt.Println("Game B Win")
			i = 9
		default:
			fmt.Println(scoreA, " - ", scoreB, " : ", wordA, " - ", wordB)
		}

	}

}
