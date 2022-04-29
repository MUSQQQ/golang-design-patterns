package main

type SouthAmerica struct{}

func (s *SouthAmerica) createAnt() Ant {
	return &Dinoponera{}
}
