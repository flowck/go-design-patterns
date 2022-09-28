package main

import "go-design-patterns/patterns/structural/facade"

func main() {
	// Two façades that wrap the coffee machine's API
	facade.MakeAmericano(8)
	facade.MakeLatte(9)
}
