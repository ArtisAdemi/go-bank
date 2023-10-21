package storage

import (
	"database/sql"
	"fmt"

	types "github.com/artisademi/go-bank/pkg/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccountByID(int) (*types.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=artis dbname=NewDB password=123123123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists accounts(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *types.Account) error {
	if s.db != nil {

		query := `INSERT INTO accounts (first_name, last_name, number, balance, created_at) values ($1, $2, $3, $4, $5)`

		resp, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", resp)
		return nil
	}
	return nil
}
func (s *PostgresStore) UpdateAccount(*types.Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*types.Account, error) {
	return nil, nil
}
