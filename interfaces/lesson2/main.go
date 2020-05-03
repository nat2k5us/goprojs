package main

import "fmt"

type Animal struct {
}

func (a *Animal) Eat()   { fmt.Println(" eating") }
func (a *Animal) Sleep() { fmt.Println(" sleeping") }
func (a *Animal) Breed() { fmt.Println(" breeding") }

type Dog struct {
	pommy Animal
}

func (a *Dog) Eat() { a.pommy.Eat() }

func (a *Dog) Sleep() {
	fmt.Print("I am a Dog")
	a.pommy.Sleep()
}
func (a *Dog) Breed() {
	fmt.Print("I am a Dog")
	a.pommy.Breed()
}

type Cat struct {
	polly Animal
}

func (a *Cat) Eat() { a.polly.Eat() }

func (a *Cat) Sleep() {
	fmt.Print("I am a Cat")
	a.polly.Sleep()
}
func (a *Cat) Breed() {
	fmt.Print("I am a Cat")
	a.polly.Breed()
}

type SleepingDen interface {
	Sleep()
}
type BreedingDen interface {
	Breed()
}

func main() {
	fmt.Println("Lesson 2")
	animals := []SleepingDen{new(Cat), new(Dog)}
	for _, x := range animals {
		x.Sleep()
	}
	pets := []BreedingDen{new(Cat), new(Dog)}
	for _, x := range pets {
		x.Breed()
	}
}
