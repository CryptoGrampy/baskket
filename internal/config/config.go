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

package config

import "github.com/namsral/flag"

var BindAddr string
var CallbackAddr string

var MoneroPayHost string

var SmtpHost string
var SmtpUser string
var SmtpPassword string

var PostgresHost string
var PostgresPort uint
var PostgresUser string
var PostgresPassword string
var PostgresDatabase string

func Init() {
	flag.StringVar(&BindAddr, "bind", "0.0.0.0:8080",
	    "Bind ip:port")
	flag.StringVar(&CallbackAddr, "callback-address", "http://0.0.0.0:3005",
	    "Bind ip:port")
	flag.StringVar(&MoneroPayHost, "moneropay-host",
	"http://moneropayd:5000/receive", "MoneroPay http://ip:port/receive")
	flag.StringVar(&SmtpHost, "smtp-host", "smtp:587", "SMTP ip:port")
	flag.StringVar(&SmtpUser, "smtp-username", "", "SMTP user@domain")
	flag.StringVar(&SmtpPassword, "smtp-password", "", "SMTP password")
	flag.StringVar(&PostgresHost, "postgres-host", "localhost",
	    "PostgreSQL cluster address")
	flag.UintVar(&PostgresPort, "postgres-port", 5432,
	    "PostgreSQL cluster port")
	flag.StringVar(&PostgresUser, "postgres-username", "baskket",
	    "PostgreSQL cluster user")
	flag.StringVar(&PostgresPassword, "postgres-password", "",
	    "PostgreSQL cluster password")
	flag.StringVar(&PostgresDatabase, "postgres-database", "baskket",
	    "PostgreSQL database")

	flag.Parse()
}
