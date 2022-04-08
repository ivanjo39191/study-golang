// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item  item
		genre string
	}

	games := []game{
		{item{1, "god of war", 50}, "action adventure"},
		{item{2, "x-com 2", 30}, "strategy"},
		{item{3, "minecraft", 20}, "sandbox"},
	}

	fmt.Printf("Ivan's game store has %d games.\n", len(games))
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
>  list: list games
>  quit: quit
`)
		if !in.Scan() {
			break
		}

		fmt.Println()
		switch in.Text() {
		case "quit":
			fmt.Println("bye!bye!")
			return
		case "list":
			for _, g := range games {
				fmt.Println(g.item.id, g.item.name, g.item.price, g.genre)
			}
		}
	}
}
