package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func shuffle(list []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i)
		list[i-1], list[j] = list[j], list[i-1]
	}
}

func getCards() (cards []int) {
	cards = make([]int, 52)
	for i := range make([]int, 52) {
		if ((i % 13) + 1) > 10 {
			cards[i] = 10
		} else {
			cards[i] = (i % 13) + 1
		}
	}
	return
}

func main() {
	// initialize state
	cards := getCards()
	reader := bufio.NewReader(os.Stdin)
	player := 0
	dealer := 0
	shuffle(cards)
	// log.Println(cards)

	// set first player points
	player = cards[0]
	cards = cards[1:]
	fmt.Print("Your points: ", player, "\n")
	player = player + cards[0]
	cards = cards[1:]
	fmt.Print("Your points: ", player, "\n")

game:
	for true {
		fmt.Print("Plese press key. [H(Hit)/S(Stand)]: ")
		text, err := reader.ReadString('\n')
		handleError(err)

		switch strings.TrimSpace(text) {
		case "H":
			hit := cards[0]
			fmt.Print("Hit!: ", hit, "\n")
			player = player + cards[0]
			fmt.Print("your points: ", player, "\n")
			cards = cards[1:]
		case "S":
			fmt.Println("Stand.")
			break game
		default:
			fmt.Println("dontmatch", text)
		}

		// log.Println(cards)
		fmt.Println()
	}

	for dealer < 17 {
		dealer = dealer + cards[0]
		cards = cards[1:]
		fmt.Print("Dealer points: ", dealer, "\n")
	}

	fmt.Println()
	fmt.Print("Dealer final points: ", dealer, "\n")
	fmt.Print("Your final points: ", player, "\n")
}
