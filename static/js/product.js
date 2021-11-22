document.getElementById("button").addEventListener("click", function(){	
	var id = parseInt(document.getElementById("id").getAttribute("value"))
        var title = document.getElementById("title").getAttribute("value")
        var price = parseFloat(document.getElementById("price").getAttribute("value"))
	var atomic = parseInt(document.getElementById("atomic").getAttribute("value"), 10)

	var local = localStorage.getItem("cart");
	var cart = []
	var found = false

	if (local == null) {
		cart.push({id: id, title: title, price: price,
			atomic: atomic, quantity: 1})
    		localStorage.setItem("cart", JSON.stringify(cart));
		return
	}

	cart = JSON.parse(local)

	cart.forEach((element, index) => {
		if (element.id == id) {
                	element.quantity++
			found = true
		}
	})

	if (found == false) {
		cart.push({id: id, title: title, price: price,
			atomic: atomic, quantity: 1})
	}

	localStorage.setItem("cart", JSON.stringify(cart));
})
