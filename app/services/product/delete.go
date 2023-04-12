package product

func (service ProductService) Delete(productID int) (err error) {
	err = service.ProductRepository.Delete(service.DB, productID)
	if err != nil {
		return err
	}
	return nil
	// or just `return err`
}