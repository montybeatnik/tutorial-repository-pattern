package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/montybeatnik/tutorials/repository-pattern/devices/models"

	_ "github.com/lib/pq"
)

// PGRepo represntes a postgres data store. It has a db, which is a
// handle representing a pool of zero or more underlying connections.
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
func (pr *PGRepo) NewDevice(newDevice models.NewDeviceRequest) (models.Device, error) {
	query := `INSERT INTO devices (hostname, ip, clli) VALUES ($1, $2, $3)
	RETURNING id;`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	device := mapDeviceAttrs(newDevice)
	if err := pr.db.QueryRowContext(ctx, query, device.Hostname, device.IP, device.CLLI).Scan(&device.ID); err != nil {
		return device, err
	}
	return device, nil
}

// GetDeviceByIP queries the DB for the given IP.
func (pr *PGRepo) GetDeviceByIP(ip string) (models.Device, error) {
	query := `SELECT id, hostname, ip, clli FROM devices WHERE ip = $1`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var dev models.Device
	if err := pr.db.QueryRowContext(ctx, query, ip).Scan(
		&dev.ID,
		&dev.Hostname,
		&dev.IP,
		&dev.CLLI); err != nil {
		return dev, err
	}
	return dev, nil
}
