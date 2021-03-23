package main

import "fmt"

/*
The proxy pattern forwards to a real object and acts as an interface to others. The proxy pattern controls access to an object and provides additional functionality.
The additional functionality can be related to authentication, authorization, and providing rights of access to the resource-sensitive object.
The real object need not be modified while providing additional logic. Remote, smart, virtual, and protection proxies are the scenarios where this pattern is
applied. It is also used to provide an alternative to extend functionality with inheritance and object composition.
A proxy object is also referred to as a surrogate, handle, or wrapper.

The proxy pattern comprises the subject interface, the RealSubject class, and the Proxy class:
	Subject is an interface for the RealObject and Proxy class.
	The RealSubject object is created and maintained as a reference in the Proxy class.
	RealSubject is resource sensitive, required to be protected, and expensive to create.
	RealObject is a class that implements the IRealObject interface. It has a performAction method.
	VirtualProxy is used to access RealObject and invoke the performAction method.
*/

//IRealObject interface
type IRealObject interface {
	performAction()
}

//RealObject struct
type RealObject struct {
}

//RealObject class method performAction
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

//VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}

//VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
