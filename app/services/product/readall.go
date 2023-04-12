package product

import (
	"challenge294/app/models"
)

func (service ProductService) ReadAll() (product *[]models.Product, err error) {
	product, err = service.ProductRepository.ReadAll(service.DB)
	if err != nil {
		return nil, err
	}

	// cek if userID == product.UserID. update: baru inget sudah diimplementasi di middleware
	return product, nil
}
