package builder

type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
}

type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	Build() Pizza
}

type ConcretePizzaBuilder struct {
	pizza Pizza
}

func (c *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	c.pizza.Size = size
	return c
}

func (c *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	c.pizza.Crust = crust
	return c
}
func (c *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	c.pizza.Cheese = true
	return c
}
func (c *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	c.pizza.Pepperoni = true
	return c
}
func (c *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	c.pizza.Mushrooms = true
	return c
}

func (c *ConcretePizzaBuilder) Build() Pizza {
	return c.pizza
}

type PizzaDirector struct{}

func (p *PizzaDirector) CreateMargherita(builder PizzaBuilder) Pizza {
	return builder.SetSize("M").SetCrust("Thin").AddCheese().Build()
}
