package product

import (
	"challenge294/app/models"
)

func (service ProductService) Create(p *models.Product, userID uint) (product *models.Product, err error) {
	p.UserID = userID // Product.UserID yang akan dibuat diset dengan ID user yang telah login/ID yang tersimpan di userData

	product, err = service.ProductRepository.Create(service.DB, p)
	if err != nil {
		return nil, err
	}
	return product, nil
}
