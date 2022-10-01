# Behavioral Patterns

## The Observer Pattern (Pub/Sub or Event-Driven)

The Observer patterns allows one or many structs (class) to subscribe for data changes from another struct (class).

## How does it work

* This pattern expects all observers (structs listening for changes) to implement an interface that the observable will use to send the most updated data;
* The observable must maintain a reference to the observers;
* Observers must be able to unregister themselves at any point in the runtime;


## The Iterator Pattern

This pattern abstracts the logic needed to access a collection of elements in a struct (object).

## What strategies can be used to implement The Iterator Pattern?

- Callback (push model)
- Interface that exposes two methods: `hasNext()` `getNext()` (pull model)

## Callback implementation

- Ideally the callback function should return an `error` to allow the struct (object) to handle any potential error thrown inside it (callback).

## 

