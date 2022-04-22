package main

func main() {
	var (
		mobydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	book.print(mobydick)
	game.print(minecraft)
	game.print(tetris)
}
