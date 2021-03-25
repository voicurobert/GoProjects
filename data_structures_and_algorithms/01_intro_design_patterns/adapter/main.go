package main

import "fmt"

/*
The adapter pattern provides a wrapper with an interface required by the API client to link incompatible types and act as a translator between the two types.
The adapter uses the interface of a class to be a class with another compatible interface. When requirements change, there are scenarios where class
functionality needs to be changed because of incompatible interfaces.

The dependency inversion principle can be adhered to by using the adapter pattern, when a class defines its own interface to the next level module interface
implemented by an adapter class. Delegation is the other principle used by the adapter pattern. Multiple formats handling source-to-destination transformations
are the scenarios where the adapter pattern is applied.

The adapter pattern comprises the target, adaptee, adapter, and client:
	Target is the interface that the client calls and invokes methods on the adapter and adaptee.
	The client wants the incompatible interface implemented by the adapter.
	The adapter translates the incompatible interface of the adaptee into an interface that the client wants.

Let's say you have an IProcessor interface with a process method, the Adapter class implements the process method and has an Adaptee instance as an attribute.
The Adaptee class has a convert method and an adapterType instance variable. The developer while using the API client calls the process
interface method to invoke convert on Adaptee.
*/

// IProcess interface
type IProcess interface {
	process()
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// Adaptee struct
type Adaptee struct {
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process()")
	adapter.adaptee.convert()
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert()")
}

func main() {
	var procesor IProcess = Adapter{}
	procesor.process()
}
