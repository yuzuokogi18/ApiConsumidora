package application

import "apiConsumer/src/reservation/domain"

type DeleteReservationUseCase struct {
	postgresRepository domain.IReservationPostgres
}

func NewDeleteReservationUseCase(postgresRepository domain.IReservationPostgres) *DeleteReservationUseCase {
	return &DeleteReservationUseCase{postgresRepository: postgresRepository}
}

func (uc *DeleteReservationUseCase) Run(id int32) error {
	return uc.postgresRepository.Delete(id)
}
