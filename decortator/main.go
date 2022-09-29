package main

import "fmt"

func main() {

	samsungPhone := &PhoneCase{}

	
	phoneWithBuds := &buds{
		samsungPhone: samsungPhone,
	}


	samsungWithCaseAndTypeC := &typeC{
		samsungPhone: phoneWithBuds,
	}

	fmt.Printf("Price of samsung with buds and TypeC %d\n", samsungWithCaseAndTypeC.getPrice())
}

