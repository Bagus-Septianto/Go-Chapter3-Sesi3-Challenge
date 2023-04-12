package product

import "challenge294/app/models"

func (service ProductService) Update(p *models.Product, productID int) (product *models.Product, err error) {
	product, err = service.ProductRepository.Update(service.DB, p, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
