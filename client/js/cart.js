const productContainer = document.getElementById("container");
const makeOrderBtn = document.getElementById("makeOrder");
window.addEventListener("load", function (event) {
  event.preventDefault();

  fetch("/user/cart/get", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      data.forEach((product) => {
        var div = document.createElement("div");
        var name = document.createElement("p");
        var price = document.createElement("p");
        name.textContent = product.name;
        price.textContent = product.price;
        div.appendChild(name);
        div.appendChild(price);

        productContainer.appendChild(div);
      });
    })
    .catch((error) => console.error("Ошибка", error));
});

makeOrderBtn.addEventListener("click", function (event) {
  event.preventDefault();
  fetch("/user/order/make", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      window.location.href = "/catalog";
    })
    .catch((error) => console.error("Ошибка", error));
});
