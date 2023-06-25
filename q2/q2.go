package q2

import "errors"

type Employee struct {
	ID         int
	Name       string
	Title      string
	BaseSalary float64
	Bonuses    []float64
}

func CalculateTotalSalary(employee *Employee) (float64, error) {
	totalBonus := 0.0
	for _, bonus := range emp.MonthlyBonus {
		totalBonus += bonus
	}

	totalSalary := emp.BaseSalary + totalBonus

	if totalBonus > 1500.0 {
		emp.Position = "Senior " + emp.Position
	}

	return totalSalary, nil
}
