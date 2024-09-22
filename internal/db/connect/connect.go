package connect

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

// Функция для подключения к базе данных
func ConnectDB() *pgxpool.Pool {
	databaseUrl := "postgres://postgres:120789@localhost:1207/e-commerce"
	ctx := context.Background()

	var err error
	dbpool, err = pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Проверяем подключение
	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	return dbpool
}

// Функция для получения экземпляра пула соединений
func GetDBPool() *pgxpool.Pool {
	return dbpool
}

// Функция для закрытия соединения
func CloseDB() {
	if dbpool != nil {
		dbpool.Close()
	}
}
