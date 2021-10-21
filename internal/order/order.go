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

package order

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gitlab.com/moneropay/go-monero/walletrpc"
	"gitlab.com/moneropay/moneropay/pkg/models"

	"gitlab.com/moneropay/baskket/internal/config"
	"gitlab.com/moneropay/baskket/internal/database"
	"gitlab.com/moneropay/baskket/internal/pages"
	"gitlab.com/moneropay/baskket/pkg/cart"
	"gitlab.com/moneropay/baskket/pkg/product"
)

var Products []product.Product
var lastOrderId uint

func Init() {
        rows, err := database.QueryWithTimeout(context.Background(), 3 * time.Second,
            "SELECT id, title, description, price, quantity, image FROM products" +
            " ORDER BY id")
        if err != nil {
                log.Fatal(err)
        }

	reg, _ := regexp.Compile("[^a-zA-Z]+")
	for rows.Next() {
		var t product.Product
		if err = rows.Scan(&t.Id, &t.Title, &t.Description, &t.PriceAtomic,
		    &t.Quantity, &t.Image); err != nil {
			    log.Fatal(err)
		}
		t.PriceFloat = walletrpc.XMRToFloat64(t.PriceAtomic)
		t.Location = strings.ToLower(reg.ReplaceAllString(t.Title, "")) + ".html"
		f, err := os.Create("public/" + t.Location)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		Products = append(Products, t)
	}
}

func Place(w http.ResponseWriter, r *http.Request) {
	cart := r.FormValue("cart")
	address := r.FormValue("address")
	email := r.FormValue("email")

	// Calculate total cost of items in cart.
	total, err := getCost(cart)
	if err != nil {
		log.Println(err)
		return
	}
	row, err := database.QueryRowWithTimeout(r.Context(), 3 * time.Second,
	    "INSERT INTO orders (email, total, address) VALUES" +
	    " ($1, $2, $3) RETURNING id", email, total, address)
	if err != nil {
		log.Println(err)
		return
	}
	var lastOrderId uint
	if err = row.Scan(&lastOrderId); err != nil {
		log.Println(err)
		return
	}

	subaddr, err := getSubaddress(total, uint64(lastOrderId))
	if err != nil {
		log.Println(err)
		return
	}

	items := struct {
		Subaddress string
		Amount string
	}{
		Subaddress: subaddr,
		Amount: walletrpc.XMRToDecimal(total),
	}
	if err = pages.Payment.Execute(w, items); err != nil {
		log.Fatal(err)
	}
}

func getCost(cartJson string) (uint64, error) {
	var c []cart.CartItem
	if err := json.Unmarshal([]byte(cartJson), &c); err != nil {
		return 0, err
	}

	var total uint64
	for _, i := range c {
		total += Products[i.Id - 1].PriceAtomic
		Products[i.Id - 1].Quantity--
		// TODO: Decrement quantity in database at shutdown.
	}
	return total, nil
}

func getSubaddress(total, id uint64) (string, error){
	resp, err := http.PostForm(config.MoneroPayHost, url.Values{"amount": {strconv.FormatUint(total, 10)},
	"callback_url": {config.CallbackAddr + "/callback/" + strconv.FormatUint(id, 10)},
		"description": {"baskket"}})
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data models.ReceivePostResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return data.Address, nil
}
