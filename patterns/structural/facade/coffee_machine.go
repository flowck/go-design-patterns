package facade

import "fmt"

type CoffeeMachine struct {
	beanAmount   float32
	grinderLevel int
	waterTemp    float32
	waterAmount  float32
	milkAmount   float32
	addFoam      bool
}

//
// The coffee_machine represents a sub-system that contains complex APIs for simple-like operations
// such as: making an americano coffee
//

func (c *CoffeeMachine) startCoffee(beanAmount float32, grind int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grind
	fmt.Println("Starting coffee order with beans", beanAmount, "and grind level", grind)
}

func (c *CoffeeMachine) endCoffee() {
	fmt.Println("Ending coffee order")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding the beans:", c.beanAmount, "beans at", c.grinderLevel)
	return true
}

func (c *CoffeeMachine) useMilk(amount float32) float32 {
	fmt.Println("Adding milk", amount, "oz")
	c.milkAmount = amount
	return amount
}

func (c *CoffeeMachine) setWaterTemp(temp float32) {
	fmt.Println("Setting water temp", temp)
	c.waterTemp = temp
}

func (c *CoffeeMachine) useHotWater(amount float32) float32 {
	fmt.Println("Adding hot water", amount)
	c.waterAmount = amount
	return amount
}

func (c *CoffeeMachine) doFoam() {
	fmt.Println("Making foam")
}
