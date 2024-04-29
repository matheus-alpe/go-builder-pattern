package main

import (
	"fmt"

	"github.com/matheus-alpe/go-builder-pattern/internal/builder"
)

func main() {
	b := &builder.ConcretePizzaBuilder{}
	director := builder.PizzaDirector{}

	margherita := director.CreateMargherita(b)
	fmt.Println("Margherita:", margherita)

	customPizza := b.SetSize("L").AddPepperoni().AddMushrooms().Build()
	fmt.Println("Custom Pizza:", customPizza)
}
