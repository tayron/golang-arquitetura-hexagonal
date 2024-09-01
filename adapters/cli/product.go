package cli

import (
	"fmt"

	"github.com/tayron/golang-arquitetura-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action, productID, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID # %s, with the name %s has been created with the price R$ %.2f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		productResult, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", productResult.GetName())

	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		productResult, err := service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", productResult.GetName())

	default:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: R$ %.2f\nStatus:: %s\n ",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	}

	return result, nil
}
