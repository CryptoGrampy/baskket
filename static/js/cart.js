var c = localStorage.getItem('cart')
var form = document.getElementById('form')
var list = document.getElementById('listing')
var cart
const items = new Map()

if (c == null) {
	form.style.display = "none"
	list.innerHTML += "<p>Empty cart, can't find anything you like?</p>"
}
else {
	cart = JSON.parse(c)
	var total = 0
	
	cart.forEach((element, index) => {
		if (items.has(element.id)) {
			items.set(element.id, [element.title, element.price, items.get(element.id)[2] + 1, index])
		} else {
			items.set(element.id, [element.title, element.price, 1, index])
		}
		total += element.atomic
	})
	
	var i
	items.forEach((value, key, map) => {
		i++
		console.log(value, key)
		h = '<div id="item"><div id="title">' + value[2] + ' x '+ value[0]
			+ ' ' + value[1] + ' XMR</div><div class="remove" id="' 
			+ i + '">x</div></div>'
		list.innerHTML += h
	})

	f = `<input type="text" id="cart" name="cart" value='` + c.replace(/'/g, "&#39;") + `' hidden><br>`
	form.innerHTML += f

	h = '<h3>Total: ' + parseFloat(total) / 1e12 + ' XMR</h3>'
	list.innerHTML += h
}
