package application

import "apiConsumer/src/reservation/domain"

type ViewAllReservationsUseCase struct { 
	postgresRepository domain.IReservationPostgres
}

func NewViewAllReservationsUseCase(postgresRepository domain.IReservationPostgres) *ViewAllReservationsUseCase {
	return &ViewAllReservationsUseCase{postgresRepository: postgresRepository}
}

func (uc *ViewAllReservationsUseCase) Run() ([]domain.Reservation, error) {
	return uc.postgresRepository.GetAll()
}
