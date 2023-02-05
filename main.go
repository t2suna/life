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

func (pos *position) SetXY(x int, y int) {
	if x > (dim - 1) {
		pos.x = dim - 1
	} else if x < 0 {
		pos.x = 0
	} else {
		pos.x = x
	}
	if y > (dim - 1) {
		pos.y = dim - 1
	} else if y < 0 {
		pos.y = 0
	} else {
		pos.y = y
	}

}

type life struct {
	pos   position
	power int
	ziga  string
	moved bool
}

const dim = 8

func eliminatedAnnounce(loser life, winner life) {
	fmt.Println(loser.ziga + " is eliminated by " + winner.ziga)
}

func battle(a life, b life) life {
	if a.ziga != "" {
		if a.power <= b.power {
			eliminatedAnnounce(a, b)
			return b
		} else {
			eliminatedAnnounce(b, a)
			return a
		}
	} else {
		return b
	}
}

func printField(field [dim][dim]life) {
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
}

func cleanFlag(field [dim][dim]life) [dim][dim]life {
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if field[i][j].ziga != "" && field[i][j].moved {
				field[i][j].moved = false
			}
		}
	}
	return field
}

func cleanTerminal() {
	for i := 0; i < dim; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\n")

}

func move(field [dim][dim]life) [dim][dim]life {
	//Move and Battle
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if field[i][j].ziga != "" && !field[i][j].moved {
				rand.Seed(time.Now().UnixNano())
				x := rand.Intn(3) - 1
				y := rand.Intn(3) - 1
				tmp := field[i][j]
				field[i][j] = life{}
				var nextPos position
				nextPos.SetXY(x+i, y+j)
				tmp.moved = true
				field[nextPos.x][nextPos.y] = tmp

			}
		}
	}
	return field
}

func main() {
	var field [dim][dim]life

	for {

		field = move(field)
		field = cleanFlag(field)

		printField(field)
		cleanTerminal()

		//Born new Alphabet
		field = born(field)

		for i := 0; i < dim; i++ {
			fmt.Print(" ")
		}
		fmt.Print("\n")

		printField(field)
		cleanTerminal()

		time.Sleep(time.Millisecond * 1000)
	}

}

func born(field [dim][dim]life) [dim][dim]life {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(dim)
	y := rand.Intn(dim)
	p := rand.Intn(10)

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := string(letterRunes[rand.Intn(len(letterRunes))])
	me := life{position{x, y}, p, str, false}

	field[me.pos.x][me.pos.y] = battle(field[me.pos.x][me.pos.y], me)
	return field
}
