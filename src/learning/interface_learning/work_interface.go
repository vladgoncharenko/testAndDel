package main

import "fmt"

type Worker interface {
	Work()
}

type SomeWorker struct {
	Profession string
}

func (s SomeWorker) Work() {
	fmt.Println(s.Profession)
}

func main() {
	some := SomeWorker{Profession: "Builder"}
	w := Worker(some)

	w.Work()

	myFl := MyFloat(22.3)
	mySt := MyString("string")

	myChan := make(chan Renderer, 2)

	myChan <- myFl
	myChan <- mySt

	close(myChan)

	Display(myChan)
}

// тоже интересная штука

type (
	MyString string
	MyFloat  float64
)

type Renderer interface {
	Render()
}

func (f MyFloat) Render() {
	fmt.Println(f)
}

func (s MyString) Render() {
	fmt.Println(s)
}

func Display(ch chan Renderer) {
	for v := range ch {
		v.Render()
	}
}
