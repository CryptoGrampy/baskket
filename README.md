# Baskket
Baskket is a minimal e-commerce application that uses MoneroPay to process payments in XMR.

Check out Kernal's instance at [baskket.kernal.eu](https://baskket.kernal.eu/).

## Usage
```
$ ./baskket -h
Usage of ./baskket:
  -bind="0.0.0.0:8080": Bind ip:port
  -callback-address="http://0.0.0.0:3005": Bind ip:port
  -moneropay-host="http://moneropayd:5000/receive": MoneroPay http://ip:port/receive
  -postgres-database="baskket": PostgreSQL database
  -postgres-host="localhost": PostgreSQL cluster address
  -postgres-password="": PostgreSQL cluster password
  -postgres-port=5432: PostgreSQL cluster port
  -postgres-username="baskket": PostgreSQL cluster user
  -refresh-address="127.0.0.1:3010": Refresh ip:port
  -smtp-host="smtp:587": SMTP ip:port
  -smtp-password="": SMTP password
  -smtp-username="": SMTP user@domain
```
### Adding New Items
This is done by directly inserting a new row to the table `products` and moving the image of the product to `/static/img`.\
In order to re-render the pages with the newly added product, make a request to `/refreshProducts`.
```
baskket=> INSERT INTO products (title, description, price, quantity, image) VALUES ('Pixel art sticker 6x6', 'Kernal logo in pixel art with URL.', 10000000000, 45, 'static/img/pixel-sticker-6x6.png');

$ curl http://refresh-address/refreshProducts
```
### Modifying Templates
Make your changes and then make a request to `/refreshPages` to re-render.

## Extending Baskket
### Themes
HTML files go under `/template` and static files must be put in `/static`.\
See the table for the variables provided to template files:

| File         | Variable type             |
|--------------|---------------------------|
| index.html   | []product.Product         |
| product.html | product.Product           |
| payment.html | subaddress, amount string |

Template parsing is done in [`page.go`](https://gitlab.com/moneropay/baskket/-/blob/master/internal/page/page.go). The resulting HTML files are placed in `/public` and served. 

## Installation
```sh
$ git clone https://gitlab.com/moneropay/baskket.git && cd baskket
$ go build cmd/baskket.go # run go install to build to $GOPATH/bin instead
```
