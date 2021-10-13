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

package main

import (
	"gitlab.com/moneropay/baskket/internal/config"
	"gitlab.com/moneropay/baskket/internal/database"
	"gitlab.com/moneropay/baskket/internal/order"
	"gitlab.com/moneropay/baskket/internal/pages"
	"gitlab.com/moneropay/baskket/internal/router"
)

func main() {
	config.Init()

	database.Connect(
		config.PostgresHost, config.PostgresPort, config.PostgresUser,
		config.PostgresPassword, config.PostgresDatabase,
	)
	defer database.Close()
	database.Migrate()

	// Get available products.
	order.Init()

	// Pre-render templates.
	pages.Render(order.Products)

	// Serve the rendered pages.
	router.Route()
}
