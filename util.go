package main

import (
	"fmt"
	"math/rand"
)

/*
	mainで使用する、関数をまとめたもの
*/

// 排除のアナウンス
func eliminatedAnnounce(loser life, winner life) {
	fmt.Println(loser.ziga + " is eliminated by " + winner.ziga)
}

// aとbのpowerを比較し、大きい方をフィールドに残す
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

// フィールドを標準出力にprint
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

// movedフラグをクリアする。
func clearMoved(field [dimY][dimX]life) [dimY][dimX]life {
	for i := 0; i < dimY; i++ {
		for j := 0; j < dimX; j++ {
			if field[i][j].ziga != "" && field[i][j].moved {
				field[i][j].moved = false
			}
		}
	}
	return field
}

// フィールド内のアルファベットをランダムに一マス移動させる。
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

// アルファベットを生み出す。生まれたマスが被った場合はバトルさせる。
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
