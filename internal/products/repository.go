package products

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

var (
	prodList = []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
)

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) (list []Product, err error) {
	for _, p := range prodList {
		if p.SellerID == sellerID {
			list = append(list, p)
		}
	}
	return list, nil
}
