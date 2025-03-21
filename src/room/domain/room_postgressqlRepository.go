package domain

type IRoomPostgres interface {
	Save(room *Room) error
	GetById(id int32) (*Room, error)
	GetAll() ([]Room, error)
	GetByHotelId(hotelId int32) ([]Room, error)
	Update(id int32, room Room) error
	Delete(id int32) error
}
