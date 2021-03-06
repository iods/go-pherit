/*
Guess Game
1. Generate a random number from 1 to 100, store it as a target number
     for the player to guess.
2. Prompt the player to guess what the target number is, storing
     their response.
3. If the guess is less than the target, return "Your guess is too
     LOW." If the player's guess is greater than the target number,
     return, "Your guess is too HIGH."
4. Allow the player to guess up to 10 times. Before each guess, let
     the player know how many guesses they have left.
5. If the player's guess is equal to the target number, return "Your
     guess is correct!" and stop asking for guesses.
6. If the player ran out of turns without correctly guessing the target
     number, return, "Sorry, your attempts were unsuccessful. It was
     [X]" and return the target number.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GuessGame() {
	title := "\nWelcome to Ferris: Rye's Game of Guesses..."
	fmt.Println(title)
	fmt.Println("\nI'm thinking of a number from 0-100.")
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1 // what is the +1 ?
	fmt.Print("\nWhat is your guess?...\n")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("\nYou have", 10 - guesses, "attempts left.")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("\nYour guess is too LOW!")
			continue
		} else if guess > target {
			fmt.Println("\nYour guess is too HIGH!")
			continue
		} else  {
			success = true
			fmt.Println("\nYour guess is CORRECT!")
			break
		}
	}

	if !success {
		fmt.Println("Thanks for trying Ferris. You lose, the answer was:", target)
	}
}