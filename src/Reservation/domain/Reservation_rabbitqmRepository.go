package domain

type IReservationRabbitmq interface {
	Save(reservation *Reservation) error
}
