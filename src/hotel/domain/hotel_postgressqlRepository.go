package domain

type IHotelPostgres interface {
	Save(hotel *Hotel) error
	GetById(id int32) (*Hotel, error)
	GetAll() ([]Hotel, error)
	Update(id int32, hotel Hotel) error
	Delete(id int32) error
}
