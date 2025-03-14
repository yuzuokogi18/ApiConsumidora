package domain
// IReservationPostgres define las operaciones con PostgreSQL
type IReservationPostgres interface {
	Save(reservation *Reservation) error
	GetById(id int32) (*Reservation, error)
	GetAll() ([]Reservation, error)
	Update(id int32, reservation Reservation) error
	Delete(id int32) error
}
