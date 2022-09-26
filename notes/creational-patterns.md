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

### Steps to implement a factory

Let me take as an example the use-case of Payment service, where all the available payment methods must perform the same operations, but differently. The operations are the following:
- `Pay(amount float64)`
- ``

1 - Define an interface (Payment) which has to be implemented by

### References

- Design Patterns - Elements of Reusable Object-Oriented Software [Book]
- Go Design Pattern [Course]