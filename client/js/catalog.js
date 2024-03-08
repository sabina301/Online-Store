const addProductBtn = document.getElementById("addProduct");
const productContainer = document.getElementById("container");
addProductBtn.style.display = "none";

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
      if (data.role == "admin") {
        addProductBtn.style.display = "block";
      } else {
        addProductBtn.style.display = "none";
      }
    })
    .catch((error) => console.error("Ошибка", error));

  fetch("/catalog/get/products/all", {
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
      console.log(data);
    })
    .catch((error) => console.error("Ошибка", error));
});

addProductBtn.addEventListener("click", function (event) {
  event.preventDefault();
  window.location.href = "/admin/catalog/edit";
});
