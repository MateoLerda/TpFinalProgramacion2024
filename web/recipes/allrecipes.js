let baseUrl = "http://localhost:8080/recipes/?filter_all=";

let searchValue;
let foodType;
let mealTime;
let all = true;
const main = document.getElementById("recipe-list");

document.getElementById("toggleSwitch").addEventListener("change", function () {
  if (this.checked) {
    all = false;
  } else {
    all = true;
  }
  getRecipes();
});

function createUrl() {
  searchValue = document.getElementById("searchInput").value;
  foodType = document.getElementById("foodType").value;
  mealTime = document.getElementById("mealTime").value;
  baseUrl = `${baseUrl}${all}`;
  if (searchValue != "") {
    baseUrl = `${baseUrl}&filter_aproximation=${searchValue}`;
  }
  if (foodType != "all") {
    baseUrl = `${baseUrl}&filter_type=${foodType}`;
  }
  if (mealTime != "all") {
    baseUrl = `${baseUrl}&filter_moment=${mealTime}`;
  }
}

function getRecipes() {

  createUrl();
  makeRequest(
    baseUrl,
    "GET",
    "",
    "application/json",
    true,
    successCreate,
    failed
  );
  baseUrl = "http://localhost:8080/recipes/?filter_all=";
}

document.addEventListener("DOMContentLoaded", () => {
  // let token = isUserLogged();
  // if (token == false) {
  //     window.location.href = '/pages/login/login.html';
  // }
  const userInfo = document.createElement('div');
  userInfo.id = 'user-info'
  userInfo.classList.add('user-info')
  const userLogo = document.createElement('i')
  userLogo.classList.add('bi','bi-person-circle', 'green-color')
  const userMail = document.createElement('p');
  userMail.textContent = localStorage.getItem('user-mail');
  userMail.classList.add('green-color', 'bold-words', 'user-mail');
  userInfo.appendChild(userLogo);
  userInfo.appendChild(userMail);

  const login = document.getElementById('login')
  const navBar = document.getElementById('navbar')
  navBar.removeChild(login)
  navBar.appendChild(userInfo)
  getRecipes();
});

document.getElementById("foodType").onchange = () => {
  getRecipes();
};

document.getElementById("mealTime").onchange = () => {
  getRecipes();
};

const searchBar = document.getElementById("search-bar");

searchBar.addEventListener("submit", (e) => {
  e.preventDefault();
  getRecipes();
});

function showRecipes(data) {
  main.innerHTML = "";
  if (data.result == null) {
    showAlert("Not found any recipes")
    return
  }
  data.result.forEach((recipe) => {
    // Crear el contenedor principal para cada receta
    let recipeContainer = document.createElement("div");
    recipeContainer.classList.add("recipe-container");
    recipeContainer.setAttribute("_id", recipe._id);
    // Nombre de la receta
    let recipeName = document.createElement("p");
    recipeName.textContent = recipe.recipe_name;
    recipeName.classList.add("big-font-size");

    // Momento de la receta
    let recipeMoment = document.createElement("p");
    recipeMoment.textContent = `Momento: ${recipe.recipe_moment}`;
    recipeMoment.classList.add("big-font-size");

    // Botón para mostrar más detalles
    let showMoreBtn = document.createElement("button");
    showMoreBtn.classList.add("btnS");
    let showMoreIcon = document.createElement("i");
    showMoreIcon.classList.add("fa-solid", "fa-chevron-down");
    showMoreBtn.appendChild(showMoreIcon);

    let DeleteBtn = document.createElement("button");
    DeleteBtn.classList.add("btnS");
    let DeleteIcon = document.createElement("i");
    DeleteIcon.classList.add("fa-solid", "fa-trash");
    DeleteBtn.appendChild(DeleteIcon);

    let UpdateBtn = document.createElement("button");
    UpdateBtn.classList.add("btnS");
    let UpdateIcon = document.createElement("i");
    UpdateIcon.classList.add("fa-solid", "fa-pencil");
    UpdateBtn.appendChild(UpdateIcon);

    // Contenedor para la información adicional
    let recipeDetails = document.createElement("div");
    recipeDetails.classList.add("recipe-details");
    recipeDetails.style.display = "none"; // Ocultar detalles inicialmente

    // Información adicional de la receta
    let ingredients = document.createElement("p");
    ingredients.textContent = `Ingredients:  ${recipe.recipe_ingredients
      .map((ing) => ` ${ing.Name}: ${ing.quantity}`)
      .join(", ")}`;

    let description = document.createElement("p");
    description.textContent = `Descripction: ${recipe.recipe_description}`;

    // Agregar los detalles al contenedor de detalles
    recipeDetails.appendChild(ingredients);
    recipeDetails.appendChild(description);

    // Agregar el evento al botón para mostrar/ocultar detalles
    showMoreBtn.addEventListener("click", () => {
      if (recipeDetails.style.display === "none") {
        recipeDetails.style.display = "block";
        showMoreIcon.classList.remove("fa-chevron-down");
        showMoreIcon.classList.add("fa-chevron-up");
      } else {
        recipeDetails.style.display = "none";
        showMoreIcon.classList.remove("fa-chevron-up");
        showMoreIcon.classList.add("fa-chevron-down");
      }
    });
    DeleteBtn.addEventListener("click", () => {
      makeRequest(
        `http://localhost:8080/recipes/${recipe._id}`,
        "DELETE",
        "",
        "application/json",
        true,
        successDelete,
        failed
      );
    });
    UpdateBtn.addEventListener("click", () => {
      UpdateRecipe(recipe._id, recipe.recipe_name, recipe.recipe_description);
    });
    let recipeContainerFerst = document.createElement("div");
    recipeContainerFerst.classList.add("recipe-containerF");
    let buttonsContainer = document.createElement("div");
    buttonsContainer.classList.add("button-container");
    buttonsContainer.appendChild(showMoreBtn);
    buttonsContainer.appendChild(DeleteBtn);
    buttonsContainer.appendChild(UpdateBtn);

    if (!all) {
      let CookBtn = document.createElement("button");
      CookBtn.classList.add("btnS");
      let CookIcon = document.createElement("i");
      CookIcon.classList.add("fa-solid", "fa-utensils");
      CookBtn.appendChild(CookIcon);
      buttonsContainer.appendChild(CookBtn);
      
      CookBtn.addEventListener("click", () => {
        makeRequest(
          `http://localhost:8080/recipes/cook/${recipe._id}?cancel=false`, 
          "GET",
          "",
          "application/json",
          true,
          successCook,
          failed
        );
      });
    }
    
    // Agregar los elementos al contenedor de la receta
    recipeContainer.appendChild(recipeName);
    recipeContainer.appendChild(recipeMoment);
    recipeContainer.appendChild(buttonsContainer);
    recipeContainerFerst.appendChild(recipeContainer);
    recipeContainerFerst.appendChild(recipeDetails);

    main.appendChild(recipeContainerFerst);
  });
}

function UpdateRecipe(recipeId, name, description) {
  let id = recipeId;
  let model = document.createElement("div");
  model.classList.add("divUpdate");
  let exitbtn = document.createElement("button");
  exitbtn.classList.add("fa-solid", "fa-xmark", "btnS");
  let title = document.createElement("h2");
  title.textContent = "Edit Recipe";
  let divName = document.createElement("div");
  divName.classList.add("inputRecipe");
  let pname = document.createElement("p");
  pname.classList.add("big-font-size");
  pname.textContent = " Recipe name: ";
  let inputname = document.createElement("input");
  inputname.placeholder = "Enter new name";
  inputname.value = name;
  divName.appendChild(pname);
  divName.appendChild(inputname);
  let divDes = document.createElement("div");
  divDes.classList.add("inputRecipe");
  let pdes = document.createElement("p");
  pdes.classList.add("big-font-size");
  pdes.textContent = "Description: ";
  let inputdes = document.createElement("textarea");
  inputdes.classList.add("inputdes");
  inputdes.placeholder = "Enter new description";
  inputdes.rows = 5;
  inputdes.cols = 30;
  inputdes.value = description;
  divDes.appendChild(pdes);
  divDes.appendChild(inputdes);
  let buttonEdit = document.createElement("button");
  buttonEdit.classList.add("fa-solid", "fa-pencil", "edit-recipe");
  buttonEdit.textContent = "  Confirm   ";

  model.appendChild(exitbtn);
  model.appendChild(title);
  model.appendChild(divName);
  model.appendChild(divDes);
  model.appendChild(buttonEdit);

  const modalOverlay = document.getElementById("modalOverlay");
  modalOverlay.appendChild(model);
  modalOverlay.showModal();
  buttonEdit.addEventListener("click", () => {
    const data = {
      recipe_name: inputname.value,
      recipe_description: inputdes.value,
    };
    makeRequest(
      `http://localhost:8080/recipes/${recipeId}`,
      "PUT",
      data,
      "application/json",
      true,
      successUpdate,
      failed
    );
  });
  exitbtn.addEventListener("click", () => {
    modalClose();
  });
}

function modalClose() {
  modalOverlay.close();
  modalOverlay.innerHTML = "";
}

function successCreate(response) {
  showRecipes(response);
  console.log("Éxito:", response);
}

function failed(response, responseBody) {
  showAlert(responseBody.error)
  console.log("Falla:", response);
}

function successDelete(response) {
  showAlert("Successfully deleted recipe")
  document.getElementById("alert-button").addEventListener(('click'), () => {
    getRecipes()
});
  console.log("Éxito:", response);
}

function successCook(response) {
  showAlert("Successfully cooked recipe")
  document.getElementById("alert-button").addEventListener(('click'), () => {
    getRecipes()
});
  console.log("Éxito:", response);
}

function successUpdate(response) {
  modalClose();
  getRecipes();
  console.log("Éxito:", response);
}
