package main

import "fmt"

// Bike is the bike object. It will keep track of its positions
type Bike struct {
	x       int32
	y       int32
	prevPos []*Bike
}

func (b *Bike) moveUp() {
	b.y++
	b.prevPos = append(b.prevPos, b)
}

func (b *Bike) moveDown() {
	b.y--
	b.prevPos = append(b.prevPos, b)
}

func (b *Bike) moveRight() {
	b.x++
	b.prevPos = append(b.prevPos, b)
}

func (b *Bike) moveLeft() {
	b.x--
	b.prevPos = append(b.prevPos, b)
}

func main() {
	bike := &Bike{x: 0, y: 0}

	bike.moveUp()
	bike.moveRight()
	bike.moveLeft()
	bike.moveDown()

	fmt.Print("Current Bike Position: ")
	fmt.Println(bike)
}
