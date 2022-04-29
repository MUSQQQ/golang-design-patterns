package main

type Australia struct{}

func (a *Australia) createAnt() Ant {
	return &Oecophylla{}
}
