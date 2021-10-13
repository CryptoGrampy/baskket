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

package mail

import (
	"log"
	"net/smtp"
	"strings"

	"gitlab.com/moneropay/baskket/internal/config"
)

func Send(to, msg string) {
	m := []byte("From: " + config.SmtpUser +
		"\r\nTo: " + to +
		"\r\nSubject: Your order in Baskket\r\n\r\n" + msg)

	a := smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword,
		strings.Split(config.SmtpHost, ":")[0])

	if err := smtp.SendMail(config.SmtpHost, a, config.SmtpUser,
	[]string{to}, m); err != nil {
		log.Fatal(err)
	}
}
