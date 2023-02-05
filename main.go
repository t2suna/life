package main

import (
	"fmt"
	"time"
)

const dimX = 50
const dimY = 10

func main() {
	var field [dimY][dimX]life

	for {

		field = move(field)
		field = clearMoved(field)

		fmt.Print("\n")

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
