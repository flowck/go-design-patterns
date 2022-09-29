# Behavioral Patterns

## The Observer Pattern (Pub/Sub or Event-Driven)

The Observer patterns allows one or many structs (class) to subscribe for data changes from another struct (class).

## How does it work

* This pattern expects all observers (structs listening for changes) to implement an interface that the observable will use to send the most updated data;
* The observable must maintain a reference to the observers;
* Observers must be able to unregister themselves at any point in the runtime;

