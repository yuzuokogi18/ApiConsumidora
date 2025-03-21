package infrastructure

import (
	"apiConsumer/src/hotel/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type HotelPostgresRepository struct {
    db *sql.DB
}

func NewHotelPostgresRepository(db *sql.DB) *HotelPostgresRepository {
    return &HotelPostgresRepository{db: db}
}

func (repo *HotelPostgresRepository) Save(hotel *domain.Hotel) error {
    query := "INSERT INTO hotels (name, location) VALUES ($1, $2)"
    _, err := repo.db.Exec(query, hotel.Name, hotel.Location)
    if err != nil {
        return fmt.Errorf("error al guardar el hotel: %v", err)
    }
    return nil
}

func (repo *HotelPostgresRepository) GetById(id int32) (*domain.Hotel, error) {
	query := `SELECT id, name, location, stars, price 
              FROM hotels WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	var hotel domain.Hotel
	if err := row.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Stars, &hotel.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Hotel no encontrado")
		}
		return nil, err
	}

	return &hotel, nil
}

func (repo *HotelPostgresRepository) GetAll() ([]domain.Hotel, error) {
	query := `SELECT id, name, location, stars, price FROM hotels`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var hotels []domain.Hotel
	for rows.Next() {
		var hotel domain.Hotel
		if err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Location, &hotel.Stars, &hotel.Price); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (repo *HotelPostgresRepository) Update(id int32, hotel domain.Hotel) error {
	query := `UPDATE hotels 
              SET name = $1, location = $2, stars = $3, price = $4 
              WHERE id = $5`
	_, err := repo.db.Exec(query, hotel.Name, hotel.Location, hotel.Stars, hotel.Price, id)
	if err != nil {
		return fmt.Errorf("Error al actualizar el hotel: %v", err)
	}
	return nil
}

func (repo *HotelPostgresRepository) Delete(id int32) error {
	query := `DELETE FROM hotels WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar el hotel: %v", err)
	}
	return nil
}
