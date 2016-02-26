package main

import "fmt"

// BikeOne is the bike object. It will keep track of its positions
type BikeOne struct {
	X       int32
	Y       int32
	prevPos []*BikeOne
}

func (b *BikeOne) start() {
	b.prevPos = append(b.prevPos, &BikeOne{X: b.X, Y: b.Y})
}

func (b *BikeOne) moveUp() {
	b.Y++
	b.prevPos = append(b.prevPos, &BikeOne{X: b.X, Y: b.Y})
}

func (b *BikeOne) moveDown() {
	b.Y--
	b.prevPos = append(b.prevPos, &BikeOne{X: b.X, Y: b.Y})
}

func (b *BikeOne) moveRight() {
	b.X++
	b.prevPos = append(b.prevPos, &BikeOne{X: b.X, Y: b.Y})
}

func (b *BikeOne) moveLeft() {
	b.X--
	b.prevPos = append(b.prevPos, &BikeOne{X: b.X, Y: b.Y})
}

func main() {
	bike := &BikeOne{X: 0, Y: 0}

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
