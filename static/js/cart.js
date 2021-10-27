var c = localStorage.getItem('cart')
var form = document.getElementById('form')
var list = document.getElementById('listing')
var cart

if (c == null) {
	form.style.display = "none"
	list.innerHTML += "<p>Empty cart, can't find anything you like?</p>"
}
else {
	cart = JSON.parse(c)

	var total = 0.0

	for (i = 0; i < cart.length; i++) {
		h = '<p>' + cart[i].title + ' ' + cart[i].price + ' XMR</p>'
		list.innerHTML += h
		total += cart[i].price
	}
	f = `<input type="text" id="cart" name="cart" value='` + c.replace(/'/g, "&#39;") + `' hidden><br>`
	form.innerHTML += f

	h = '<h3>Total: ' + total + ' XMR</h3>'
	list.innerHTML += h
}
