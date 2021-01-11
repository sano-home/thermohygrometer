package model

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Queryer interface implements QueryRowContext, QueryContext and ExecContext
// that has same signature in database/sql package.
type Queryer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// TXer interface implements QueryRowContext, QueryContext, ExecContext,
// Commit and Rollback that has same signature in database/sql package.
type TXer interface {
	Queryer
	Commit() error
	Rollback() error
}

// DBer interface implements QueryRowContext, QueryContext, ExecContext,
// BeginTx and Close that has same signature in database/sql package.
type DBer interface {
	Queryer
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
}

// NewSQLite3 returns DBer.
func NewSQLite3(dbPath string) (DBer, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, err
}

// TemperatureAndHumidity is TemperatureAndHumidity.
type TemperatureAndHumidity struct {
	ID            int64
	Temperature   float32
	Humidity      float32
	Unixtimestamp int64
}

// Create creates TemperatureAndHumidity.
func (t *TemperatureAndHumidity) Create(ctx context.Context, q Queryer) error {
	sqlStmt := `INSERT INTO temperature_and_humidity
					(temperature, humidity, unixtimestamp)
				VALUES
					($1, $2, $3)`
	result, err := q.ExecContext(ctx,
		sqlStmt, &t.Temperature, &t.Humidity, &t.Unixtimestamp)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

// GetLatestTemperatureAndHumidity returns the latest temperature and humidity.
func GetLatestTemperatureAndHumidity(ctx context.Context, q Queryer) (*TemperatureAndHumidity, error) {
	sqlStmt := `SELECT id, temperature, humidity, unixtimestamp
				FROM temperature_and_humidity
				ORDER BY unixtimestamp DESC
				LIMIT 1`

	var th TemperatureAndHumidity
	err := q.QueryRowContext(ctx,
		sqlStmt).Scan(&th.ID, &th.Temperature, &th.Humidity, &th.Unixtimestamp)
	if err != nil {
		return nil, err
	}
	return &th, nil
}

// GetTemperatureAndHumidities returns temperature and humidity history.
func GetTemperatureAndHumidities(ctx context.Context, q Queryer, since, before time.Time) ([]*TemperatureAndHumidity, error) {
	sqlStmt := `SELECT id, temperature, humidity, unixtimestamp
				FROM temperature_and_humidity
				WHERE unixtimestamp BETWEEN $1 AND $2
				ORDER BY unixtimestamp`

	rows, err := q.QueryContext(ctx,
		sqlStmt, since.Unix(), before.Unix())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ths []*TemperatureAndHumidity
	for rows.Next() {
		var th TemperatureAndHumidity
		err := rows.Scan(&th.ID, &th.Temperature, &th.Humidity, &th.Unixtimestamp)
		if err != nil {
			return nil, err
		}
		ths = append(ths, &th)
	}
	return ths, nil
}
