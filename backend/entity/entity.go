package entity

import(
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Book struct {
	gorm.Model
	Title	string		`valid:"stringlength(3|100)~Title Must long 3-100 character"`
	Price	float64		`valid:"range(50|5000)~Price must be between 50 and 5000"`
	Code	string		`valid:"matches(^BK[0-9]{6}$)~Code must start with BK followed by 6 digits (0-9)"`
}

func (b *Book) ValidateBook() error{
	ok, err := govalidator.ValidateStruct(b)
	if !ok {
		return err
	}
	return nil
}