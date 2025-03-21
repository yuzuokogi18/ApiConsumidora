package application

import "apiConsumer/src/room/domain"

type CreateRoomUseCase struct {
	postgresRepository domain.IRoomPostgres
}

func NewCreateRoomUseCase(postgresRepository domain.IRoomPostgres) *CreateRoomUseCase {
	return &CreateRoomUseCase{postgresRepository: postgresRepository}
}

func (uc *CreateRoomUseCase) Run(room *domain.Room) error {
	return uc.postgresRepository.Save(room)
}
