package main

import "fmt"

// Bike is the bike object. It will keep track of its positions
type Bike struct {
	x, y int32
}

func addPositions(ar []*Bike, bike *Bike) {
	ar = append(ar, bike)
	fmt.Println(ar)
	fmt.Println(ar[0])
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
	var previousPositions []*Bike
	bike := &Bike{x: 0, y: 0}

	bike.moveUp()
	addPositions(previousPositions, bike)
	bike.moveRight()
	addPositions(previousPositions, bike)
	bike.moveRight()
	addPositions(previousPositions, bike)
	fmt.Print("Current Bike Position: ")
	fmt.Println(bike)
}
