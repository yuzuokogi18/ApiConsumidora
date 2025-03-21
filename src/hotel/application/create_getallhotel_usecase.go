package application

import "apiConsumer/src/hotel/domain"

type GetAllHotelsUseCase struct {
    hotelRepository domain.IHotelPostgres
}

func NewGetAllHotelsUseCase(hotelRepository domain.IHotelPostgres) *GetAllHotelsUseCase {
    return &GetAllHotelsUseCase{hotelRepository: hotelRepository}
}

func (uc *GetAllHotelsUseCase) Run() ([]domain.Hotel, error) {
    return uc.hotelRepository.GetAll()
}
