package application

import "apiConsumer/src/reservation/domain"

type ViewReservationByIdUseCase struct {
	postgresRepository domain.IReservationPostgres
}

func NewViewReservationByIdUseCase(postgresRepository domain.IReservationPostgres) *ViewReservationByIdUseCase {
	return &ViewReservationByIdUseCase{postgresRepository: postgresRepository}
}

func (uc *ViewReservationByIdUseCase) Run(id int32) (*domain.Reservation, error) {
	return uc.postgresRepository.GetById(id)
}
