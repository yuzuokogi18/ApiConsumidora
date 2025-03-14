package infrastructure

import (
	"apiConsumer/src/reservation/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) Save(reservation *domain.Reservation) error {
	query := `INSERT INTO reservations (customer_name, room_type, start_date, end_date, price) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := repo.db.QueryRow(query, reservation.CustomerName, reservation.RoomType, reservation.StartDate, reservation.EndDate, reservation.Price).Scan(&reservation.Id)
	if err != nil {
		return fmt.Errorf("Error al guardar la reserva: %v", err)
	}

	return nil
}

func (repo *PostgresRepository) GetById(id int32) (*domain.Reservation, error) {
	query := `SELECT id, customer_name, room_type, start_date, end_date, price 
              FROM reservations WHERE id = $1`
	row := repo.db.QueryRow(query, id)

	var reservation domain.Reservation
	if err := row.Scan(&reservation.Id, &reservation.CustomerName, &reservation.RoomType, &reservation.StartDate, &reservation.EndDate, &reservation.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Reserva no encontrada")
		}
		return nil, err
	}

	return &reservation, nil
}

func (repo *PostgresRepository) GetByCustomerName(name string) ([]domain.Reservation, error) {
	query := `SELECT id, customer_name, room_type, start_date, end_date, price 
              FROM reservations WHERE customer_name = $1`
	rows, err := repo.db.Query(query, name)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener reservas por nombre del cliente: %v", err)
	}
	defer rows.Close()

	var reservations []domain.Reservation
	for rows.Next() {
		var reservation domain.Reservation
		if err := rows.Scan(&reservation.Id, &reservation.CustomerName, &reservation.RoomType, &reservation.StartDate, &reservation.EndDate, &reservation.Price); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (repo *PostgresRepository) GetAll() ([]domain.Reservation, error) {
	query := `SELECT id, customer_name, room_type, start_date, end_date, price FROM reservations`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var reservations []domain.Reservation
	for rows.Next() {
		var reservation domain.Reservation
		if err := rows.Scan(&reservation.Id, &reservation.CustomerName, &reservation.RoomType, &reservation.StartDate, &reservation.EndDate, &reservation.Price); err != nil {
			return nil, fmt.Errorf("Error al escanear los resultados: %v", err)
		}
		reservations = append(reservations, reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error durante la iteración de resultados: %v", err)
	}

	return reservations, nil
}

func (repo *PostgresRepository) Update(id int32, reservation domain.Reservation) error {
	query := `UPDATE reservations 
              SET customer_name = $1, room_type = $2, start_date = $3, end_date = $4, price = $5 
              WHERE id = $6`
	_, err := repo.db.Exec(query, reservation.CustomerName, reservation.RoomType, reservation.StartDate, reservation.EndDate, reservation.Price, id)
	if err != nil {
		return fmt.Errorf("Error al actualizar la reserva: %v", err)
	}
	return nil
}

func (repo *PostgresRepository) Delete(id int32) error {
	query := `DELETE FROM reservations WHERE id = $1`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar la reserva: %v", err)
	}
	return nil
}
