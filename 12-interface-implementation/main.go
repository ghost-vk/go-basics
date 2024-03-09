// INTERFACE IMPLEMENTATION Interfaces are implemented implicitly.
//
// A type never declares that it implements a given interface. If an interface
// exists and a type has the proper methods defined, then the type automatically
// fulfills that interface.
//
// ASSIGNMENT At Textio we have full-time employees and contract employees. We
// have been tasked with making a more general employee interface so that
// dealing with different employee types is simpler.
//
// Run the code. You should see an error indicating the contractor type does not
// fulfill the employee interface.
//
// Add the missing getSalary method to the contractor type so that it fulfills
// the employee interface.
//
// A contractor's salary is their hourly pay multiplied by how many hours they
// work per year.

package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func getInfo(e employee) {
	fmt.Printf("%s has year salary: %v\n", e.getName(), e.getSalary())
}

func main() {
	e1 := contractor{
		name:         "Bob",
		hourlyPay:    10,
		hoursPerYear: 100,
	}
	e2 := fullTime{
		name:   "Sue",
		salary: 20000,
	}
	getInfo(e1)
	getInfo(e2)
}
