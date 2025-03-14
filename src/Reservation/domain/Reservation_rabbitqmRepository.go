package domain

// IReservationRabbitmq define la operación para enviar mensajes a RabbitMQ
type IReservationRabbitmq interface {
	Save(reservation *Reservation) error
}
