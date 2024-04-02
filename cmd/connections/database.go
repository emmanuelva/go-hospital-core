package connections

import (
	"database/sql"
	"log"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	//conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectToDB(DSN string) (*sql.DB, error) {
	connection, err := openDB(DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to postgres")
	return connection, nil
}
