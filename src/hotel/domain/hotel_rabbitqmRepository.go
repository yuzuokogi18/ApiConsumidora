package domain

type IHotelRabbitmq interface {
	Save(hotel *Hotel) error
}
