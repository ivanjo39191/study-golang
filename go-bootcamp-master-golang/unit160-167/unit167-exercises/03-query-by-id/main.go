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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
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

	games_map := make(map[int]game)
	for _, g := range games {
		games_map[g.item.id] = g
	}

	fmt.Printf("Ivan's game store has %d games.\n", len(games))
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
>  list: list games
>  id index: querues a game by id
>  quit: quit
`)
		if !in.Scan() {
			break
		}

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		fmt.Println()
		switch cmd[0] {
		case "quit":
			fmt.Println("bye!bye!")
			return
		case "list":
			for _, g := range games {
				fmt.Println(g.item.id, g.item.name, g.item.price, g.genre)
			}
		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}
			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}
			g, ok := games_map[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}
			fmt.Println(g.item.id, g.item.name, g.item.price, g.genre)
		}
	}
}
