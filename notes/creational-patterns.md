# Creational patterns

## Builder pattern

The Builder pattern, once applied correctly enables the separation of the construction of a complex struct (object)
from its representation, thus allowing multiple representations from the same construction process.

In the day-to-day activities of a programmer, the Builder pattern is particularly useful when a struct (object)
can be in a valid state with all or just some attributes.

## Factory pattern

The Factory pattern enables the abstraction of concrete structs (classes) through a factory method responsible to 
create the right struct given the input it receives. 

A common example would be a factory method responsible to choose and return a payment method struct.

### References

- Design Patterns - Elements of Reusable Object-Oriented Software [Book]
- Go Design Pattern [Course]