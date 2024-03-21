const orderContainer = document.getElementById("container");
window.addEventListener("load", function (event) {
  event.preventDefault();
  fetch("/user/order/all", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      orderContainer.innerHTML = "";
      var i = 0;
      data.forEach((order) => {
        i++;
        var name = document.createElement("p");
        name.textContent = "Order " + i + " with id = " + order.id;
        orderContainer.appendChild(name);
      });
    })
    .catch((error) => console.error("Ошибка", error));
});
