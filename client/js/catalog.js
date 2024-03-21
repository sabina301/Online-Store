const addProductBtn = document.getElementById("addProduct");
const productContainer = document.getElementById("container");
const cart = document.getElementById("cart");
const userOrders = document.getElementById("userOrders");
const adminOrders = document.getElementById("adminOrders");
addProductBtn.style.display = "none";
adminOrders.style.display = "none";
userOrders.style.display = "none";
var userId = 0;

window.addEventListener("load", function (event) {
  event.preventDefault();
  fetch("/catalog/getrole", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      userId = data.id;
      if (data.role == "admin") {
        addProductBtn.style.display = "block";
        cart.style.display = "none";
        adminOrders.style.display = "block";
      } else {
        addProductBtn.style.display = "none";
        cart.style.display = "block";
        userOrders.style.display = "block";
      }
    })
    .catch((error) => console.error("Ошибка", error));

  fetch("/catalog/product/get/all", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      data.forEach((product) => {
        console.log(product);
        var div = document.createElement("div");
        var name = document.createElement("p");
        var price = document.createElement("p");
        var addBtn = document.createElement("button");
        addBtn.textContent = "+";
        name.textContent = product.name;
        price.textContent = product.price;
        div.appendChild(name);
        div.appendChild(price);
        div.appendChild(addBtn);
        addBtn.addEventListener("click", function (event) {
          event.preventDefault();
          console.log("PROD ID = ", product.id);
          addThisProduct(product.id);
        });

        productContainer.appendChild(div);
      });
    })
    .catch((error) => console.error("Ошибка", error));
});

addProductBtn.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/admin/catalog/edit";
});

cart.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/user/cart";
});

adminOrders.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/admin/order";
});

userOrders.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/user/order";
});

function addThisProduct(productId) {
  fetch("/user/product/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      userId: userId,
      productId: productId,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.id < 0) {
        console.log("Ошибка при добавлении товара");
      } else {
        console.log("OK!");
      }
    })
    .catch((error) => console.error("Ошибка", error));
}
