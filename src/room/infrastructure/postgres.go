package infrastructure

import (
	"database/sql"
	"fmt"
	"apiConsumer/src/room/domain"
)

type RoomPostgresRepository struct {
	db *sql.DB
}

func NewRoomPostgresRepository(db *sql.DB) *RoomPostgresRepository {
	return &RoomPostgresRepository{db: db}
}

func (repo *RoomPostgresRepository) Save(room *domain.Room) error {
	query := `INSERT INTO rooms (hotel_id, type, capacity, price) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.db.QueryRow(query, room.HotelId, room.Type, room.Capacity, room.Price).Scan(&room.Id)
	if err != nil {
		return fmt.Errorf("error al guardar la habitación: %v", err)
	}
	return nil
}

func (repo *RoomPostgresRepository) GetById(id int32) (*domain.Room, error) {
	query := `SELECT id, hotel_id, type, capacity, price FROM rooms WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	var room domain.Room
	if err := row.Scan(&room.Id, &room.HotelId, &room.Type, &room.Capacity, &room.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("habitación no encontrada")
		}
		return nil, err
	}

	return &room, nil
}

func (repo *RoomPostgresRepository) GetAll() ([]domain.Room, error) {
	query := `SELECT id, hotel_id, type, capacity, price FROM rooms`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las habitaciones: %v", err)
	}
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var room domain.Room
		if err := rows.Scan(&room.Id, &room.HotelId, &room.Type, &room.Capacity, &room.Price); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (repo *RoomPostgresRepository) GetByHotelId(hotelId int32) ([]domain.Room, error) {
	query := `SELECT id, hotel_id, type, capacity, price FROM rooms WHERE hotel_id = $1`
	rows, err := repo.db.Query(query, hotelId)
	if err != nil {
		return nil, fmt.Errorf("error al obtener habitaciones: %v", err)
	}
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		var room domain.Room
		if err := rows.Scan(&room.Id, &room.HotelId, &room.Type, &room.Capacity, &room.Price); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (repo *RoomPostgresRepository) Update(id int32, room domain.Room) error {
	query := `UPDATE rooms SET type = $1, capacity = $2, price = $3 WHERE id = $4`
	_, err := repo.db.Exec(query, room.Type, room.Capacity, room.Price, id)
	if err != nil {
		return fmt.Errorf("error al actualizar la habitación: %v", err)
	}
	return nil
}

func (repo *RoomPostgresRepository) Delete(id int32) error {
	query := `DELETE FROM rooms WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la habitación: %v", err)
	}
	return nil
}
