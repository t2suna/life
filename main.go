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
	if x > (dimY - 1) {
		pos.x = dimY - 1
	} else if x < 0 {
		pos.x = 0
	} else {
		pos.x = x
	}
	if y > (dimX - 1) {
		pos.y = dimX - 1
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

const dimX = 50
const dimY = 10

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

func printField(field [dimY][dimX]life) {
	for i := 0; i < dimY; i++ {
		for j := 0; j < dimX; j++ {
			if field[i][j].ziga != "" {
				fmt.Print(field[i][j].ziga)
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")
	}
}

func cleanFlag(field [dimY][dimX]life) [dimY][dimX]life {
	for i := 0; i < dimY; i++ {
		for j := 0; j < dimX; j++ {
			if field[i][j].ziga != "" && field[i][j].moved {
				field[i][j].moved = false
			}
		}
	}
	return field
}

func move(field [dimY][dimX]life) [dimY][dimX]life {
	//Move and Battle
	for i := 0; i < dimY; i++ {
		for j := 0; j < dimX; j++ {
			if field[i][j].ziga != "" && !field[i][j].moved {
				x := rand.Intn(3) - 1
				y := rand.Intn(3) - 1
				tmp := field[i][j]
				field[i][j] = life{}
				var nextPos position
				nextPos.SetXY(x+i, y+j)
				tmp.moved = true
				field[nextPos.x][nextPos.y] = tmp
				//fmt.Println(x, y)
			}
		}
	}
	return field
}

func main() {
	var field [dimY][dimX]life

	for {

		field = move(field)
		field = cleanFlag(field)

		fmt.Print("\n")

		//Born new Alphabet
		field = born(field)

		for i := 0; i < dimY; i++ {
			fmt.Print(" ")
		}
		fmt.Print("\n")

		printField(field)
		fmt.Print("\n")

		time.Sleep(time.Millisecond * 100)
	}

}

func born(field [dimY][dimX]life) [dimY][dimX]life {

	x := rand.Intn(dimY)
	y := rand.Intn(dimX)
	p := rand.Intn(10)

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//var letterRunes = []rune("@")
	str := string(letterRunes[rand.Intn(len(letterRunes))])
	me := life{position{x, y}, p, str, false}
	fmt.Println(me.ziga + " was borned.")
	field[me.pos.x][me.pos.y] = battle(field[me.pos.x][me.pos.y], me)
	return field
}
