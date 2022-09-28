package facade

import "fmt"

//
// The Façade which wraps more complex APIs from the coffee_machine sub-system
//

func MakeAmericano(size float32) {
	fmt.Println("Making an Americano")
	machine := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	machine.startCoffee(beansAmount, 2)
	machine.grindBeans()
	machine.useHotWater(size)
	machine.endCoffee()

	fmt.Println("Americano is ready")
}

func MakeLatte(size float32) {
	fmt.Println("Making a Latte")
	machine := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	machine.startCoffee(beansAmount, 4)
	machine.grindBeans()
	machine.useHotWater(size)

	milkAmount := (size / 8.0) * 2.0
	machine.useMilk(milkAmount)
	machine.doFoam()
	machine.endCoffee()

	fmt.Println("Latte is ready")
}
