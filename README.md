# Baskket
Baskket is a minimal e-commerce application that uses MoneroPay to process payments in XMR.

## Using Baskket
### Adding New Items
This is done by directly inserting a new row to the table `products`.

## Extending Baskket
### Themes
HTML files go under `/template` and static files must be put in `/static`.
See the table for the variables provided to template files:

| File         | Variable type             |
|--------------|---------------------------|
| index.html   | []product.Product         |
| product.html | product.Product           |
| payment.html | subaddress, amount string |

Template parsing is done in [`pages.go`](https://gitlab.com/moneropay/baskket/-/blob/master/internal/pages/pages.go). The resulting HTML files are placed in `/public` and served.
