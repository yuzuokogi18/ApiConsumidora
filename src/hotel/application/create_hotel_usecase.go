// En el paquete hotel/application
package application

import "apiConsumer/src/hotel/domain"

type CreateHotelUseCase struct {
    hotelRepository domain.IHotelPostgres
}

func NewCreateHotelUseCase(hotelRepository domain.IHotelPostgres) *CreateHotelUseCase {
    return &CreateHotelUseCase{hotelRepository: hotelRepository}
}

func (uc *CreateHotelUseCase) Run(hotel *domain.Hotel) error {
    return uc.hotelRepository.Save(hotel)
}
