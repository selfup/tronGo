package main

import "fmt"

type Bike struct {
	x, y int32
}

func (b *Bike) moveUp() {
	bb := b.y
	b.y = bb + 1
}

func (b *Bike) moveDown() {
	bb := b.y
	b.y = bb - 1
}

func (b *Bike) moveRight() {
	bb := b.x
	b.x = bb + 1
}

func (b *Bike) moveLeft() {
	bb := b.x
	b.x = bb - 1
}

func main() {
	bike := &Bike{x: 0, y: 0}

	bike.moveUp()
	bike.moveRight()
	bike.moveRight()
	fmt.Println(bike)
}
