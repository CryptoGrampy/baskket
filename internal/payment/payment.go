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

package payment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/moneropay/go-monero/walletrpc"
	"gitlab.com/moneropay/moneropay/pkg/models"

	"gitlab.com/moneropay/baskket/internal/database"
	"gitlab.com/moneropay/baskket/internal/mail"
)

func NotifyReceived(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data models.CallbackData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		return
	}

        row, err := database.QueryRowWithTimeout(r.Context(), 3 * time.Second,
	"SELECT total, received, email, address FROM orders WHERE id=$1", id)
        if err != nil {
                log.Println(err)
                return
        }
	var total, received uint64
	var email, address string
        if err = row.Scan(&total, &received, &email, &address); err != nil {
                log.Println(err)
                return
        }

	w.WriteHeader(http.StatusOK)

	xmr := received + data.Amount
	if (len(email) != 0) {
		if xmr >= total {
			mail.Send(email, fmt.Sprintf(
				"Thank you! We are preparing your order." +
				" Received: %s/%s XMR.",
				walletrpc.XMRToDecimal(xmr),
				walletrpc.XMRToDecimal(total)))
		} else {
			mail.Send(email, fmt.Sprintf(
				"Almost there! You need to transfer %s XMR more" +
				" to complete your order. Received %s/%s XMR.",
				walletrpc.XMRToDecimal(total - xmr),
				walletrpc.XMRToDecimal(xmr),
				walletrpc.XMRToDecimal(total)))
		}
	}

	err = database.ExecWithTimeout(r.Context(), 3 * time.Second,
	"UPDATE orders SET received = $1 WHERE id = $2", xmr, id)
	if err != nil {
		log.Println(err)
		return
	}
}
