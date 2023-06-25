package q3

import "errors"

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	if product == nil {
		return errors.New("Product is nil")
	}

	if sales == nil {
		return errors.New("Sales map is nil")
	}

	quantitySold, ok := sales[product.Code]
	if !ok {
		return errors.New("Product code not found in sales map")
	}

	if product.Quantity < quantitySold {
		return errors.New("Quantity of product that was sold is greater than the product's current stock quantity")
	}

	product.Quantity -= quantitySold
	
	return nil
}
