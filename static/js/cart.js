var c = localStorage.getItem('cart')
var form = document.getElementById('form')
var list = document.getElementById('listing')
var total = document.getElementById('total')
var cart
var sum

function registerEvents() {
	[...document.querySelectorAll('.remove')].forEach(element => {
		element.addEventListener('click', deleteItem)
	})
}

function updateItem(element) {
	var listing = document.getElementById("item-" + element.id)
	var quantity = parseInt(element.getAttribute('quantity'), 10) - 1
	if (quantity == 0) {
		listing.remove()
		return
	}

	var title = element.getAttribute('title')
	var price = element.getAttribute('price')
	var atomic = element.getAttribute('atomic')

	listing.innerHTML = '<div class="title">' 
		+ quantity + ' x ' + title + ' ' + price
		+ ' XMR</div><div class="remove" id="'
		+ element.id + '" title="' + title 
		+ '" quantity="' + quantity 
		+ '" price="' + price 
		+ '" atomic="' + atomic
		+ '">x</div>'
	
	registerEvents()
}

function updateTotal(val) {
	total.innerHTML = 'Total: ' + parseFloat(val) / 1e12 + ' XMR'
}

function deleteItem() {
	sum -= parseInt(this.getAttribute('atomic'), 10)
	updateTotal(sum)
	updateItem(this)
}

if (c == null) {
	form.style.display = "none"
	list.innerHTML += "<p>Empty cart, can't find anything you like?</p>"
} else {
	cart = JSON.parse(c)
	sum = 0
	
	cart.forEach((element, index) => {
		h = '<div class="cartitem" id="item-' + element.id 
			+ '"><div class="title">' + element.quantity 
			+ ' x ' + element.title + ' ' + element.price
			+ ' XMR</div><div class="remove" id="'
			+ element.id + '" title="' + element.title
			+ '" quantity="' + element.quantity
			+ '" price="' + element.price
			+ '" atomic="' + element.atomic + '">x</div></div>'

		list.innerHTML += h
		sum += element.atomic * element.quantity
	})

	registerEvents()
	
	f = `<input type="text" id="cart" name="cart" value='` + c.replace(/'/g, "&#39;") + `' hidden><br>`
	form.innerHTML += f

	updateTotal(sum)
}
