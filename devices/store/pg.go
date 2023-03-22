package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"

	_ "github.com/lib/pq"
)

type PGRepo struct {
	db *sql.DB
}

// NewPGRepo is a factory function that connects
// to a postgres RDBMS and drops into a device database
// using a DSN.
func NewPGRepo(dsn string) (*PGRepo, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to RDMBs: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldn't connect to DB: %w", err)
	}
	return &PGRepo{db: db}, nil
}

// NewDevice adds a device to DB.
func (pr *PGRepo) NewDevice(device models.Device) error {
	query := `INSERT INTO devices (hostname, ip) VALUES ($1, $2);`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := pr.db.ExecContext(ctx, query, device.Hostname, device.IP)
	if err != nil {
		return err
	}
	_ = res
	return nil
}

// GetDeviceByIP queries the DB for the given IP.
func (pr *PGRepo) GetDeviceByIP(ip string) (models.Device, error) {
	query := `SELECT id, hostname, ip FROM devices WHERE ip = $1`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var dev models.Device
	if err := pr.db.QueryRowContext(ctx, query, ip).Scan(&dev.ID, &dev.Hostname, &dev.IP); err != nil {
		return dev, err
	}
	return dev, nil
}
