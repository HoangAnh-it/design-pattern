package shoe

type IShoe interface {
	SetSize(size int)
	SetBrand(brand string)
	GetSize() int
	GetBrand() string
}

type shoe struct {
	size  int
	brand string
}

func (s *shoe) SetSize(size int) {
	s.size = size
}

func (s *shoe) GetSize() int {
	return s.size
}

func (s *shoe) SetBrand(brand string) {
	s.brand = brand
}

func (s *shoe) GetBrand() string {
	return s.brand
}
