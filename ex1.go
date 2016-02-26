package main

import "fmt"

// BikeOne is the bike object. It will keep track of its positions
type BikeOne struct {
	x       int32
	y       int32
	prevPos []*BikeOne
}

func (b *BikeOne) start() {
	b.prevPos = append(b.prevPos, &BikeOne{x: b.x, y: b.y})
}

func (b *BikeOne) moveUp() {
	b.y++
	b.prevPos = append(b.prevPos, &BikeOne{x: b.x, y: b.y})
}

func (b *BikeOne) moveDown() {
	b.y--
	b.prevPos = append(b.prevPos, &BikeOne{x: b.x, y: b.y})
}

func (b *BikeOne) moveRight() {
	b.x++
	b.prevPos = append(b.prevPos, &BikeOne{x: b.x, y: b.y})
}

func (b *BikeOne) moveLeft() {
	b.x--
	b.prevPos = append(b.prevPos, &BikeOne{x: b.x - 1, y: b.y})
}

func main() {
	bike := &BikeOne{x: 0, y: 0}

	bike.start()

	for i := 0; i < 10; i++ {
		bike.moveRight()
		bike.moveUp()
	}

	for i := 0; i < 20; i++ {
		bike.moveRight()
		bike.moveDown()
	}

	for i := 0; i < 30; i++ {
		bike.moveUp()
	}

	for i := 0; i < len(bike.prevPos); i++ {
		fmt.Println(bike.prevPos[i])
	}
}
