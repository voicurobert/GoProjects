package main

import "fmt"

/*
A composite is a group of similar objects in a single object. Objects are stored in a tree form to persist the whole hierarchy.
The composite pattern is used to change a hierarchical collection of objects. The composite pattern is modeled on a heterogeneous collection.
New types of objects can be added without changing the interface and the client code. You can use the composite pattern, for example, for UI layouts on the web,
for directory trees, and for managing employees across departments. The pattern provides a mechanism to access the individual objects and groups in a similar manner.

The composite pattern comprises the component interface, component class, composite, and client:
	The component interface defines the default behavior of all objects and behaviors for accessing the components of the composite.
	The composite and component classes implement the component interface.
	The client interacts with the component interface to invoke methods in the composite.

Let's say there is an IComposite interface with the perform method and BranchClass that implements IComposite and has the addLeaf, addBranch, and perform methods.
The Leaflet class implements IComposite with the perform method. BranchClass has a one-to-many relationship with leafs and branches.
*/

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet perform - name:", leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch Branch) perform() {
	fmt.Println("Branch perform - name:", branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch2.add(leaf2)

	branch.addBranch(branch2)

	branch.perform()
}
