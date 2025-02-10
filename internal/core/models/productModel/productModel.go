package productModel

type Product struct {
	ID          int64  `json:"id" gorm:"column:id;type:serial;" `
	ProductCode string `json:"product_code" gorm:"column:product_code;type:varchar(20);" `
	ProductName string `json:"product_name" gorm:"column:product_name;type:varchar(200);" `
}

func (b *Product) TableName() string {
	return "public.products"
}
