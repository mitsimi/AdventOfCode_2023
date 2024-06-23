package main

import "fmt"

type gameRound struct {
	red   int
	green int
	blue  int
}

func (r *gameRound) SetCubes(red, green, blue int) {
	r.SetRed(red)
	r.SetGreen(green)
	r.SetBlue(blue)
}

func (r *gameRound) GetCubes() (red int, green int, blue int) {
	return r.GetRed(), r.GetGreen(), r.GetBlue()
}

func (r *gameRound) SetRed(count int) {
	r.red = count
}

func (r *gameRound) GetRed() int {
	return r.red
}

func (r *gameRound) SetGreen(count int) {
	r.green = count
}

func (r *gameRound) GetGreen() int {
	return r.green
}

func (r *gameRound) SetBlue(count int) {
	r.blue = count
}

func (r *gameRound) GetBlue() int {
	return r.blue
}

func (r *gameRound) ToString() string {
	return fmt.Sprintf("{r: %d, g: %d, b: %d", r.red, r.green, r.blue)
}

type cubeGame struct {
	id     int
	rounds []gameRound
}

func (g *cubeGame) SetID(id int) {
	g.id = id
}

func (g *cubeGame) GetID() int {
	return g.id
}

func (g *cubeGame) AddRound(round gameRound) {
	g.rounds = append(g.rounds, round)
}

func (g *cubeGame) GetRounds() []gameRound {
	return g.rounds
}

func (g *cubeGame) ToString() string {
	str := fmt.Sprintf("%d : [", g.id)
	for _, r := range g.rounds {
		str = fmt.Sprintf("%s %s", str, r.ToString())
	}
	str = fmt.Sprintf("%s]", str)
	return str
}