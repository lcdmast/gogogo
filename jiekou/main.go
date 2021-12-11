package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpence(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Println(expense)
}

func main() {
	a := Permanent{
		1,
		5000,
		30,
	}
	b := Permanent{
		2,
		6000,
		40,
	}
	c := Permanent{
		3,
		8000,
		50,
	}
	d := Contract{
		4,
		5000,
	}

	emp := []SalaryCalculator{a, b, c, d}

	totalExpence(emp)
}
