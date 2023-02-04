package main

import (
	"fmt"
	"math/rand"
	"time"
)

type position struct {
	x int
	y int
}

type life struct {
	pos   position
	power int
	ziga  string
}

const dim = 8

func main() {
	var field [dim][dim]life

	for {
		me := born(field)

		if field[me.pos.x][me.pos.y].ziga != "" {
			if field[me.pos.x][me.pos.y].power <= me.power {
				field[me.pos.x][me.pos.y] = me
			} //elseは何もしない。
		} else {
			field[me.pos.x][me.pos.y] = me
		}

		fmt.Println("@")
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				if field[i][j].ziga != "" {
					fmt.Print(field[i][j].ziga)
				} else {
					fmt.Print("_")
				}
			}
			fmt.Print("\n")
		}

		time.Sleep(time.Millisecond * 500)
	}

}

func born(field [dim][dim]life) life {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(dim)
	y := rand.Intn(dim)
	p := rand.Intn(10)

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	me := life{position{x, y}, p, string(letterRunes[rand.Intn(len(letterRunes))])}

	return me
}
