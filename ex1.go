package main

import "fmt"

// Bike is the bike object. It will keep track of its positions
type Bike struct {
	x       int32
	y       int32
	prevPos []*Bike
}

func (b *Bike) start() {
	b.prevPos = append(b.prevPos, &Bike{x: b.x, y: b.y})
}

func (b *Bike) moveUp() {
	b.y++
	b.prevPos = append(b.prevPos, &Bike{x: b.x, y: b.y})
}

func (b *Bike) moveDown() {
	b.y--
	b.prevPos = append(b.prevPos, &Bike{x: b.x, y: b.y})
}

func (b *Bike) moveRight() {
	b.x++
	b.prevPos = append(b.prevPos, &Bike{x: b.x, y: b.y})
}

func (b *Bike) moveLeft() {
	b.x--
	b.prevPos = append(b.prevPos, &Bike{x: b.x - 1, y: b.y})
}

func main() {
	bike := &Bike{x: 0, y: 0}

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
