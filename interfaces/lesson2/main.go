package main

import "fmt"

type animal struct {
}

// Eat , Sleep and Breed are true for any Animal - polymorphic behavior
func (a *animal) Eat()   { fmt.Println(" eating") }
func (a *animal) Sleep() { fmt.Println(" sleeping") }
func (a *animal) Breed() { fmt.Println(" breeding") }

type dog struct {
	pommy animal
}
type cat struct {
	polly animal
}
type bird struct {
	mitto animal
}

func (a *bird) Breed() {
	fmt.Print("I am a bird")
	a.mitto.Breed()
}

func (a *dog) Eat() { a.pommy.Eat() }

func (a *dog) Sleep() {
	fmt.Print("I am a Dog")
	a.pommy.Sleep()
}
func (a *dog) Breed() {
	fmt.Print("I am a Dog")
	a.pommy.Breed()
}

func (a *cat) Eat() { a.polly.Eat() }

func (a *cat) Sleep() {
	fmt.Print("I am a Cat")
	a.polly.Sleep()
}
func (a *cat) Breed() {
	fmt.Print("I am a Cat")
	a.polly.Breed()
}

type SleepingDen interface {
	Sleep()
}
type BreedingDen interface {
	Breed()
}
type EatingDen interface {
	Breed()
}

func main() {
	fmt.Println("Lesson 2")
	animals := []SleepingDen{new(cat), new(dog)}
	for _, x := range animals {
		x.Sleep()
	}
	
	pets := []BreedingDen{new(cat), new(dog), new(bird)}
	for _, x := range pets {
		x.Breed()
	}
}
