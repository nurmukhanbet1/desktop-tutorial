package main

import "fmt"

func main() {

	samsungPhone := &PhoneCase{}

	//Add cheese topping
	phoneWithBuds := &buds{
		samsungPhone: samsungPhone,
	}

	//Add tomato topping
	samsungWithCaseAndTypeC := &typeC{
		samsungPhone: phoneWithBuds,
	}

	fmt.Printf("Price of samsung with buds and TypeC %d\n", samsungWithCaseAndTypeC.getPrice())
}

