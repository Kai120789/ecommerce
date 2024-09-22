package create

import (
	"context"
	"fmt"
	"log"
	"shop/internal/db/connect"
)

func CreateTables() {

	dbpool := connect.ConnectDB()
	ctx := context.Background()

	// create table User
	_, err := dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS "User" (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
		second_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100),
        password VARCHAR(255) NOT NULL
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table User created successfully!")

	// create table Category
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Category (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		photo TEXT,
		slug VARCHAR(255) UNIQUE
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Category created successfully!")

	// create table Brand
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Brand (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		photo TEXT
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Brand created successfully!")

	// create table Product
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Product (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		photo TEXT,
		slug VARCHAR(255) UNIQUE,
		category_id INT REFERENCES Category(id) ON DELETE SET NULL,
		description TEXT,
        brand_id INT REFERENCES Brand(id) ON DELETE SET NULL,
		default_price DOUBLE PRECISION
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Product created successfully!")

	// create table Review
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Review (
        id SERIAL PRIMARY KEY,
		rating DOUBLE PRECISION,
        name VARCHAR(255) NOT NULL,
		product_id INT REFERENCES Product(id) ON DELETE CASCADE,
		description TEXT
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Review created successfully!")

	// create table Color
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Color (
        id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		color VARCHAR(7),
		image TEXT
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Color created successfully!")

	// create table Product_variant
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Product_variant (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
		photo TEXT,
		product_id INT REFERENCES Product(id) ON DELETE CASCADE,
		slug VARCHAR(255) UNIQUE,
		color_id INT REFERENCES Color(id),
		category_id INT REFERENCES Category(id),
		description TEXT,
		images TEXT[],
		brand_id INT REFERENCES Brand(id),
		price DOUBLE PRECISION NOT NULL
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Product_variant created successfully!")

	// create table Basket
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Basket (
        id SERIAL PRIMARY KEY,
		user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
		count INT NOT NULL,
		product_id INT REFERENCES Product(id) ON DELETE CASCADE,
		product_variant_id INT REFERENCES Product_variant(id) ON DELETE CASCADE
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Basket created successfully!")

	// create table Order
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS "Order" (
        id SERIAL PRIMARY KEY,
    	user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
		full_price DOUBLE PRECISION
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Order created successfully!")

	// create table Product_order
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Product_order (
        id SERIAL PRIMARY KEY,
		product_id INT REFERENCES Product(id) ON DELETE CASCADE,
		product_variant_id INT REFERENCES Product_variant(id) ON DELETE CASCADE,
    	order_id INT REFERENCES "Order"(id) ON DELETE CASCADE
    );
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Product_order created successfully!")

	// create table Product_color
	_, err = dbpool.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS Product_color (
        id SERIAL PRIMARY KEY,
		product_id INT REFERENCES Product(id) ON DELETE CASCADE,
		color_id INT REFERENCES Color(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}
	fmt.Println("Table Product_color created successfully!")
}
