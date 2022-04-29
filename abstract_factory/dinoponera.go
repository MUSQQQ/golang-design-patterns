package main

import "fmt"

type Dinoponera struct{}

func (d *Dinoponera) sting() {
	fmt.Println("Dinoponera has strong venom. Stinging can be very painful.")
}

func (d *Dinoponera) describe() {
	fmt.Println("Dinoponera are among the largest ants in the world. They can reach up to 4cm in length.")
}
