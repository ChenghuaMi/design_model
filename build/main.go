package main

import "fmt"

//建造着模式

type House struct {
	Door   string
	Window string
}

type HouseBuilder interface {
	BuildDoor(val string)
	BuildWindow(val string)
	GetHouse() *House
}
type Bao struct {
	house *House
}

func NewBao() *Bao {
	return &Bao{
		house: &House{},
	}
}
func (b *Bao) BuildDoor(val string) {
	b.house.Door = val
	fmt.Println("door")
}
func (b *Bao) BuildWindow(val string) {
	b.house.Window = val
	fmt.Println("window")
}
func (b *Bao) GetHouse() *House {
	return b.house
}

type Boss struct {
	builder HouseBuilder
}

func NewBoss(builder HouseBuilder) *Boss {
	return &Boss{builder: builder}
}
func (b *Boss) buildHose() *House {
	b.builder.BuildWindow("win11")
	b.builder.BuildDoor("dodddd")
	fmt.Println(b.builder.GetHouse())
	return b.builder.GetHouse()
}
func main() {
	bao := NewBao()
	boss := NewBoss(bao)

	boss.buildHose()
}
