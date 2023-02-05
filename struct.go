package main

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
