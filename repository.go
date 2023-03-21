package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/montybeatnik/tutorials/repository-pattern/models"

	_ "github.com/lib/pq"
)

type Repository interface {
	StoreDevice(device models.Device) error
	GetDeviceByIP(ip string) (models.Device, error)
}

type InMemRepo struct {
	store map[int]models.Device
}

var count = 0

func NewInMemRepo() *InMemRepo {
	store := make(map[int]models.Device)
	store[count] = models.Device{Hostname: "hostname1", IP: "1.1.1.1"}
	return &InMemRepo{store: store}
}

func (mr *InMemRepo) StoreDevice(device models.Device) error {
	count++
	mr.store[count] = device
	log.Println(mr.store)
	return nil
}

func (mr *InMemRepo) GetDeviceByIP(ip string) (models.Device, error) {
	var match int
	for pk, dev := range mr.store {
		if dev.IP == ip {
			match = pk
		}
	}
	return mr.store[match], nil
}

type PGRepo struct {
	db *sql.DB
}

func NewPGRepo(dsn string) *PGRepo {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
	}
	return &PGRepo{db: db}
}

func (pr *PGRepo) Ping() error {
	return pr.db.Ping()
}

func (pr *PGRepo) StoreDevice(device models.Device) error {
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
