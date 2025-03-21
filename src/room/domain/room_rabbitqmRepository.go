package domain

type IRoomRabbitmq interface {
	Save(room *Room) error
}
