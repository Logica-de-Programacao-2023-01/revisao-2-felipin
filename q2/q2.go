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
	if employee == nil {
		return 0.0, errors.New("Funcionario invalido")
	} else {
		totalbonus := 0.0
		for _, bonus := range employee.Bonuses {
			totalbonus += bonus
		}
		totalsalary := employee.BaseSalary + totalbonus

		if totalbonus > 1500.0 {
			employee.Title = "Senior " + employee.Title
		}
		return totalsalary, nil
	}
}
