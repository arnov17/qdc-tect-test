package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Player struct {
	pemain int
	dadu   []int
	one    []int
	point  int
}

func main() {
	var nPemain int
	var Mdadu int

	fmt.Println("Input jumlah Pemain & Jumlah Dadu... ")
	fmt.Println("contoh = 3 4 (with space)")
	fmt.Scanf("%d", &nPemain)
	fmt.Scanf("%d", &Mdadu)
	fmt.Printf("Jumah Pemain = %d , Jumlah Dadu = %d \n", nPemain, Mdadu)
	fmt.Printf("\n")

	initPlayer := setPlayDice(nPemain, Mdadu)
	round := 0
	keepPlay := true

	for keepPlay {
		round += 1

		fmt.Printf("Giliran lempar dadu ke = %d\n", round)

		randDice := randDice(initPlayer)

		print(randDice)
		fmt.Printf("\nSetelah Evaluasi\n")

		newDice, sumPlayerPlay := evaluateDice(randDice)

		print(newDice)
		fmt.Println("================= \n")

		if sumPlayerPlay == 1 {
			getWinner(newDice)
			break
		}
	}
}

func setPlayDice(nPemain int, mDadu int) []Player {
	initPlayer := []Player{}
	for j := 0; j < nPemain; j++ {
		a := Player{pemain: j + 1}
		initPlayer = append(initPlayer, a)

		for k := 1; k <= mDadu; k++ {
			initPlayer[j].dadu = append(initPlayer[j].dadu, 0)
		}
	}

	return initPlayer
}

func randDice(initPlayer []Player) []Player {
	for j := 0; j < len(initPlayer); j++ {
		dice := []int{}
		for k := 1; k <= len(initPlayer[j].dadu); k++ {
			numRand := rand.Intn(6) + 1
			dice = append(dice, numRand)
		}
		initPlayer[j].dadu = dice
	}
	return initPlayer
}

func evaluateDice(initPlayer []Player) ([]Player, int) {
	for l := 0; l < len(initPlayer); l++ {
		s := []int{}
		dadu := initPlayer[l].dadu
		for i := 0; i < len(dadu); i++ {
			if dadu[i] != 6 && dadu[i] != 1 {
				s = append(s, dadu[i])
			}
			if dadu[i] == 6 {
				initPlayer[l].point += 1
			}
			if dadu[i] == 1 {
				if l == len(initPlayer)-1 {
					initPlayer[0].one = append(initPlayer[0].one, 1)
				} else {
					initPlayer[l+1].one = append(initPlayer[l+1].one, 1)
				}
			}

		}
		initPlayer[l].dadu = s
	}

	var sumPlayerPlay = len(initPlayer)
	for x := 0; x < len(initPlayer); x++ {
		var y = initPlayer[x].one

		initPlayer[x].dadu = append(initPlayer[x].dadu, y...)
		initPlayer[x].one = []int{}

		if len(initPlayer[x].dadu) == 0 {
			sumPlayerPlay -= 1
		}
	}

	return initPlayer, sumPlayerPlay

}

func getWinner(initPlayer []Player) {
	var winner int
	score := 0
	for z := 0; z < len(initPlayer); z++ {
		if initPlayer[z].point > score {
			score = initPlayer[z].point
			winner = initPlayer[z].pemain
		}
	}

	moreHighScore := 0
	for y := 0; y < len(initPlayer); y++ {
		if initPlayer[y].point == score {
			moreHighScore += 1
		}
	}
	if moreHighScore > 1 {
		fmt.Printf("Theres is no winner, score result is draw with score %d", score)
	} else {
		fmt.Printf("The winner is Player #: %d with score : %d \n", winner, score)
	}

}

func print(initPlayer []Player) {
	for _, index := range initPlayer {
		daduString := arrayToString(index.dadu, ", ")
		fmt.Printf("Pemain # %d ( %d ) : %s \n", index.pemain, index.point, daduString)
	}
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
