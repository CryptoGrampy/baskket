/*
 * Copyright (C) 2021 İrem Kuyucu <siren@kernal.eu>
 *
 * This file is part of Baskket.
 *
 * Baskket is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Baskket is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Baskket.  If not, see <https://www.gnu.org/licenses/>.
 */

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Connect(host string, port uint, user, pass, dbname string) {
	u := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, dbname)
	var err error
	DB, err = pgxpool.Connect(context.Background(), u)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}
}

func Close() {
	DB.Close()
}

func Migrate() {
	err := ExecWithTimeout(context.Background(), 3 * time.Second, `
	    CREATE TABLE IF NOT EXISTS products (
	        id		serial PRIMARY KEY,
	        title		character varying(256),
	        description	character varying(2048),
	        price		numeric NOT NULL,
		quantity	integer,
		image		text
	    );
	    CREATE TABLE IF NOT EXISTS orders (
	        id		serial PRIMARY KEY,
		total		bigint NOT NULL CHECK (total >= 0),
		received	bigint DEFAULT 0,
		email		text,
		address		text
		)`)
	if err != nil {
		log.Fatal(err)
	}
}
