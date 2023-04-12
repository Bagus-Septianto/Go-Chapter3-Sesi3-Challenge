package product

import (
	"challenge294/app/models"
)

func (service ProductService) Read(productID int) (product *models.Product, err error) {
	product, err = service.ProductRepository.Read(service.DB, productID)
	if err != nil {
		return nil, err
	}

	// cek if userID == product.UserID. update: baru inget sudah diimplementasi di middleware
	return product, nil
}
