const iconMap = {
  Vegetable: ["fa-solid", "fa-carrot"],
  Fruit: ["fa-solid", "fa-apple-whole"],
  Cheese: ["fa-solid", "fa-cheese"],
  Dairy: ["fa-solid", "fa-beer-mug-empty"],
  Meat: ["fa-solid", "fa-drumstick-bite"],
};

document.addEventListener("DOMContentLoaded", () => {
  // let token = isUserLogged();
  // if (token == false) {
  //   window.location.href = "/pages/login/login.html";
  // }
  const userInfo = document.createElement("div");
  userInfo.id = "user-info";
  userInfo.classList.add("user-info");
  const userLogo = document.createElement("i");
  userLogo.classList.add("bi", "bi-person-circle", "green-color");
  const userMail = document.createElement("p");
  userMail.textContent = localStorage.getItem("user-mail");
  userMail.classList.add("green-color", "bold-words", "user-mail");
  userInfo.appendChild(userLogo);
  userInfo.appendChild(userMail);

  const login = document.getElementById("login");
  const navBar = document.getElementById("navbar");
  navBar.removeChild(login);
  navBar.appendChild(userInfo);
  const baseUrl = "http://localhost:8080/foods/?filter_all=true";
  makeRequest(
    baseUrl,
    "GET",
    "",
    "application/json",
    true,
    successGetFoods,
    failedGetFoods
  );
});

function showFoods(data) {
  const main = document.getElementById("list-food");
  main.innerHTML = "";
  if (data.message) {
    showAlert(data.message + ' to create a recipe')
    document.getElementById('alert-button').addEventListener(('click'), () => {
      window.location.href = "../foods/foods.html"
    })
  }
  data.forEach((food) => {
    let foodContainer = document.createElement("div");
    foodContainer.classList.add("food-container");
    foodContainer.setAttribute("_id", food._id);
    let icon = document.createElement("i");
    icon.classList.add(
      iconMap[food.type][0],
      iconMap[food.type][1],
    );
    let foodName = document.createElement("p");
    foodName.classList.add("name", "big-font-size");
    foodName.textContent = food.name;

    let quantity = document.createElement("div");
    quantity.classList.add("quantity");

    const decreaseButton = document.createElement("button");
    decreaseButton.classList.add("fa-solid", "fa-minus", "btns");
    decreaseButton.id = "decrease";

    const numberInput = document.createElement("input");
    numberInput.classList.add("numQuantity");
    numberInput.id = "numberInput";
    numberInput.value = 0;
    numberInput.min = 0;
    numberInput.step = 1;

    const increaseButton = document.createElement("button");
    increaseButton.classList.add("fa-solid", "fa-plus", "btns");
    increaseButton.id = "increase";

    quantity.appendChild(decreaseButton);
    quantity.appendChild(numberInput);
    quantity.appendChild(increaseButton);

    increaseButton.addEventListener("click", () => {
      numberInput.value = parseInt(numberInput.value) + 1;
    });

    decreaseButton.addEventListener("click", () => {
      if (parseInt(numberInput.value) > 0) {
        numberInput.value = parseInt(numberInput.value) - 1;
      }
    });

    foodContainer.appendChild(icon);
    foodContainer.appendChild(foodName);
    foodContainer.appendChild(quantity);
    main.appendChild(foodContainer);
  });
}

document.getElementById("createBtn").onclick = () => {
  const foodContainers = document.querySelectorAll(".food-container");
  const foodQuantity = [];
  foodContainers.forEach((conteiner) => {
    const quantity = parseInt(conteiner.querySelector(".numQuantity").value);
    if (quantity > 0) {
      const foodId = conteiner.getAttribute("_id");
      const foodName = conteiner.querySelector(".name").textContent;

      foodQuantity.push({
        _id: foodId,
        name: foodName,
        quantity: quantity,
      });
    }
  });
  const recipeName = document.getElementById("recipeName").value;
  const recipeMoment = document.getElementById("recipeMoment").value;
  const description = document.getElementById("recipeDescription").value;
  const bool = false;
  if (recipeName == "" || recipeMoment == "" || description == "") {
    showAlert("Data is required");
    bool = true;
  }
  if (foodQuantity.length == 0 && !bool) {
    showAlert("You must select foods to create a recipe");
    return
  }
  const data = {
    recipe_name: recipeName,
    recipe_ingredients: foodQuantity,
    recipe_moment: recipeMoment,
    recipe_description: description,
  };
  makeRequest(
    "http://localhost:8080/recipes/",
    "POST",
    data,
    "application/json",
    true,
    successCreate,
    failedCreate
  );
};
function successGetFoods(response) {
  showFoods(response);
  console.log("Éxito:", response);
}

function failedGetFoods(response, responseBody) {
  console.log("Falla:", response);
}

function successCreate(response) {
  const modal = document.getElementById("modalSucces");
  modal.showModal();
  console.log("Éxito:", response);
}
document.getElementById("btnno").onclick = () => {
  window.location = "AllRecipes.html";
};
document.getElementById("btnyes").onclick = () => {
  window.location = "newRecipe.html";
};
function failedCreate(response, responseBody) {
   showAlert(responseBody.error)
  console.log("Falla:", response);
}
