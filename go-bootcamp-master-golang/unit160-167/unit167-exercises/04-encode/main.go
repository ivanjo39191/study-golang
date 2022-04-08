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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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
    > list   : lists all the games
    > id N   : queries a game by id
    > save   : exports the data to json and quits
    > quit   : quits
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
		case "save":
			type jsonGame struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Genre string `json:"genre"`
				Price int    `json:"price"`
			}
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.item.id, g.item.name, g.genre, g.item.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}
			fmt.Println(string(out))
			return

		}
	}
}
