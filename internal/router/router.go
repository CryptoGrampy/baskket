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

package router

import (
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.com/moneropay/baskket/internal/config"
	"gitlab.com/moneropay/baskket/internal/order"
	"gitlab.com/moneropay/baskket/internal/page"
	"gitlab.com/moneropay/baskket/internal/payment"
	"gitlab.com/moneropay/baskket/internal/product"
)

func Route() {
	r := mux.NewRouter()
	r.HandleFunc("/order", order.Place)
	r.HandleFunc("/callback/{id}", payment.NotifyReceived)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	srv := &http.Server{
		Handler: r,
		Addr: config.BindAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/* Provide the endpoints for re-rendering templates
	 * and updating products on a seperate server.
	 * Because ideally, these endpoints shouldn't be accesible
	 * outside the local network.
	 */
	http.HandleFunc("/refreshPages", refreshPages)
	http.HandleFunc("/refreshProducts", refreshProducts)
	go http.ListenAndServe(config.RefreshAddr, nil)

	log.Fatal(srv.ListenAndServe())
}

func refreshPages(w http.ResponseWriter, r *http.Request) {
	page.Render(product.Products)
}

func refreshProducts(w http.ResponseWriter, r *http.Request) {
	product.Get()

	/* Re-render templates with the updated products array */
	page.Render(product.Products)
}
