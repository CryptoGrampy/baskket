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

package product

type Product struct {
	Id uint
	Title string
	Description string
	PriceFloat float64
	PriceAtomic uint64
	Quantity uint
	Image string
	Location string
}
