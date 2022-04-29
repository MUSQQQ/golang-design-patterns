package main

import "fmt"

type Ant interface {
	sting()
	describe()
}

type Continent interface {
	createAnt() Ant
}

func main() {
	australianFactory := Australia{}
	southAmericanFactory := SouthAmerica{}

	australianAnt := australianFactory.createAnt()
	southAmericanAnt := southAmericanFactory.createAnt()

	fmt.Println("First ant from Australia")
	australianAnt.sting()
	australianAnt.describe()

	fmt.Println("Next ant from South America")
	southAmericanAnt.sting()
	southAmericanAnt.describe()
}
