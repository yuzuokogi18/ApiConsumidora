package application

import "apiConsumer/src/reservation/domain"

type UpdateReservationUseCase struct {
	postgresRepository domain.IReservationPostgres
}

func NewUpdateReservationUseCase(postgresRepository domain.IReservationPostgres) *UpdateReservationUseCase {
	return &UpdateReservationUseCase{postgresRepository: postgresRepository}
}

func (uc *UpdateReservationUseCase) Run(id int32, reservation domain.Reservation) error {
	return uc.postgresRepository.Update(id, reservation)
}
