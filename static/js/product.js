document.getElementById("button").addEventListener("click", function(){
    var local = localStorage.getItem("cart");
    var cart = local ? JSON.parse(local) : [];
    cart.push({id: parseInt(document.getElementById("id").getAttribute("value")),
        title: document.getElementById("title").getAttribute("value"),
        price: parseFloat(document.getElementById("price").getAttribute("value")),
	atomic: parseInt(document.getElementById("atomic").getAttribute("value"), 10)
    });
    localStorage.setItem("cart", JSON.stringify(cart));
});
