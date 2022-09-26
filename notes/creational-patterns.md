# Creational patterns

## Builder pattern

The Builder pattern, once applied correctly enables the separation of the construction of a complex struct (object)
from its representation, thus allowing multiple representations from the same construction process.

In the day-to-day activities of a programmer, the Builder pattern is particularly useful when a struct (object)
can be in a valid state with all or just some attributes.

## Factory pattern

The Factory pattern enables the abstraction of concrete structs (classes) through a factory method responsible to 
create the right struct given the input it receives. 

This pattern is useful in the following scenarios:

- When a struct (class) can't anticipate which struct should be created. E.g: Payment service: CardPayment, CashPayment

## Singleton pattern

This pattern restricts the creation (instantiation) of a struct (class) to a single instance.

In Go, a singleton can be thread-safe if sensitive blocks are wrapped under `once.Do(func(){})` from the `sync` package.

```go
package main

import ("sync")

type SomeService struct {}

var once sync.Once
var connectionToSomething *SomeService

func GetInstanceOfSomething() *SomeService {
    once.Do(func() {
        // Run whatever shouldn't run twice here
        connectionToSomething = &SomeService{}
    })
    
    return connectionToSomething
} 
```

### References

- Design Patterns - Elements of Reusable Object-Oriented Software [Book]
- Go Design Pattern [Course]