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
package ferris

import (
	"fmt"
	"strings"
)

func GuessGame() {
	title := "Welcome to Ferris: Rye's Game of Guesses."
	fmt.Println(strings.Title(title))
}