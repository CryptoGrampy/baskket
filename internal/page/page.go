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

package page

import (
	"html/template"
	"log"
	"os"

	"gitlab.com/moneropay/baskket/internal/product"
)

var Payment *template.Template

// Render templates and write to files in /public.
func Render(products []product.Product) {
	item, err := template.ParseFiles("template/product.html")
	if err != nil {
		log.Fatal(err)
	}

	if stat, err := os.Stat("public"); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir("public", 0755); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else if !stat.IsDir() {
		log.Fatal("/public exists and it is not a directory.")
	}

	for _, p := range products {
		f, err := os.Create("public/" + p.Location)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}

		if err = item.Execute(f, p); err != nil {
			log.Fatal(err)
		}
	}

	renderPage("index.html", products)
        renderPage("cart.html", nil)

	// This template will be executed many times with different data.
	Payment, err = template.ParseFiles("template/payment.html")
	if err != nil {
		log.Fatal(err)
	}
}

func renderPage(s string, p []product.Product) {
	index, err := template.ParseFiles("template/" + s)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("public/" + s)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err = index.Execute(f, p); err != nil {
		log.Fatal(err)
	}
}
