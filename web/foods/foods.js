const baseUrl = 'http://localhost:8080/foods/';
const iconMap = {
  Vegetable: ['fa-solid', 'fa-carrot', 'orange-color'],
  Fruit: ['fa-solid', 'fa-apple-whole', 'red-color'],
  Cheese: ['fa-solid', 'fa-cheese', 'yellow-color'],
  Dairy: ['fa-solid', 'fa-beer-mug-empty', 'black-color'],
  Meat: ['fa-solid', 'fa-drumstick-bite', 'brown-color'],
};

document.addEventListener('DOMContentLoaded', fetchAllFoods);

function fetchAllFoods() {
  // let token = isUserLogged();
  // if (token == false) {
  //   window.location.href = '/pages/login/login.html';
  // }
  const userInfo = document.getElementById('user-info');
  const userMail = document.createElement('p');
  userMail.textContent = localStorage.getItem('user-mail');
  userMail.classList.add('green-color', 'bold-words', 'user-mail');
  userInfo.appendChild(userMail);
  const query = '?filter_all=true';

  makeRequest(
    baseUrl + query,
    'GET',
    '',
    'application/json',
    true,
    showFoods,
    failedGet
  );
}

function failedGet(response) {
  console.log('Failed to get foods:', response);
}

function showFoods(data) {
  if (data.message) {
    showAlert(data.message)
    return
  }
  const main = document.getElementById('main-container');
  data.forEach((food) => {
    let foodContainer = document.createElement('div');
    let foodName = document.createElement('p');
    let foodType = document.createElement('p');
    let foodCurrentQuantity = document.createElement('p');
    let foodMinimumQuantity = document.createElement('p');
    let foodUnitPrice = document.createElement('p');
    let foodMoments = document.createElement('p');
    let icon = document.createElement('i');
    icon.classList.add(
      iconMap[food.type][0],
      iconMap[food.type][1],
      iconMap[food.type][2]
    );

    foodContainer.classList.add('food-container');
    foodName.textContent = food.name;
    foodName.classList.add(
      'food-name',
      'food-text',
      'green-color',
      'big-font-size'
    );
    foodType.textContent = food.type;
    foodType.classList.add('food-type', 'food-text');
    foodCurrentQuantity.textContent = `Quantity: ${food.current_quantity}`;
    foodCurrentQuantity.classList.add('food-current-quantity', 'food-text');
    foodMinimumQuantity.textContent = `Minimum: ${food.minimum_quantity}`
    foodMinimumQuantity.classList.add('food-minimum-quantity', 'food-text');
    foodUnitPrice.textContent = `Unit Price: ${food.unit_price}`;
    foodUnitPrice.classList.add('food-unit-price', 'food-text');
    foodMoments.textContent = food.moments
    foodMoments.classList.add('food-moments', 'food-text');

    foodContainer.appendChild(icon);
    foodContainer.appendChild(foodName);
    foodContainer.appendChild(foodType);
    foodContainer.appendChild(foodMoments);
    foodContainer.appendChild(foodUnitPrice);
    foodContainer.appendChild(foodCurrentQuantity);
    foodContainer.appendChild(foodMinimumQuantity);

    let deleteButton = document.createElement('button');
    let deleteIcon = document.createElement('i');
    deleteButton.appendChild(deleteIcon);
    deleteButton.setAttribute('food-id', food._id);
    deleteIcon.classList.add('fa-solid', 'fa-trash-can');
    deleteButton.classList.add('delete-btn-food');

    let editButton = document.createElement('button');
    let editIcon = document.createElement('i');
    editButton.setAttribute('food-id', food._id);
    editButton.appendChild(editIcon);
    editIcon.classList.add('fa-solid', 'fa-pencil');
    editButton.classList.add('edit-btn-food');

    let divButtons = document.createElement('div');
    divButtons.classList.add('flex-buttons');
    divButtons.appendChild(editButton);
    divButtons.appendChild(deleteButton);
    foodContainer.appendChild(divButtons);
    main.appendChild(foodContainer);
  });
  addDeleteEvents();
  addEditEvents();
}

function addDeleteEvents() {
  document.querySelectorAll('.food-container').forEach((container) => {
    const deleteButton = container.querySelector('.delete-btn-food');
    deleteButton.addEventListener('click', (event) => {
      const foodId = deleteButton.getAttribute('food-id');
      makeRequest(
        baseUrl + foodId,
        'DELETE',
        '',
        'application/json',
        true,
        successDelete,
        failedDelete
      );
      document.getElementById('main-container').removeChild(container);
    });
  });
}

function addEditEvents() {
  document.querySelectorAll('.food-container').forEach((container) => {
    const editButton = container.querySelector('.edit-btn-food');
    editButton.addEventListener('click', () => {
      const modal = document.getElementById('edit-modal');
      modal.classList.add('display-flex-column');
      modal.setAttribute('_id', editButton.getAttribute('food-id'))
      document.getElementById('foodname-edit').value = container.querySelector('.food-name').textContent;
      document.getElementById('foodtype-edit').value = container.querySelector('.food-type').textContent;
      document.getElementById('foodunitprice-edit').value = parseFloat(container.querySelector('.food-unit-price').textContent.split(" ")[2]);
      document.getElementById('foodquantity-edit').value = parseInt(container.querySelector('.food-current-quantity').textContent.split(" ")[1]);
      document.getElementById('foodminimum-edit').value = parseInt(container.querySelector('.food-minimum-quantity').textContent.split(" ")[1]);
      modal.showModal();
    });
  });
}

document.getElementById('new-btn-food').addEventListener('click', () => {
  const modal = document.getElementById('create-modal');
  modal.classList.add('display-flex-column');
  modal.showModal();
});

document.getElementById('create-close-modal').addEventListener('click', () => {
  const modal = document.getElementById('create-modal');
  modal.classList.remove('display-flex-column');
  modal.close();
});

document.getElementById('edit-close-modal').addEventListener('click', () => {
  const modal = document.getElementById('edit-modal');
  modal.classList.remove('display-flex-column');
  modal.close();
});

document.getElementById('dialog-create-form').addEventListener('submit', (event) => modalSubmit(event, 'create'));
document.getElementById('dialog-edit-form').addEventListener('submit', (event) => modalSubmit(event, 'edit'));

function modalSubmit(event, method) {
  event.preventDefault();
  if (method === 'edit') {
    const foodId = document.getElementById('edit-modal').getAttribute('_id')
    const data = {
      type: document.getElementById('foodtype-edit').value,
      name: document.getElementById('foodname-edit').value,
      unit_price: parseFloat(document.getElementById('foodunitprice-edit').value),
      current_quantity: parseInt(document.getElementById('foodquantity-edit').value),
      minimum_quantity: parseInt(document.getElementById('foodminimum-edit').value),
    };

    makeRequest(
      baseUrl + foodId,
      'PUT',
      data,
      'application/json',
      true,
      successEdit,
      failedEdit
    );
  } else {
    const data = {
      type: document.getElementById('foodtype').value,
      moments: Array.from(
        document
          .getElementById('foodmoments')
          .querySelectorAll('input[type="checkbox"]:checked')
      ).map((checkbox) => checkbox.value),
      name: document.getElementById('foodname').value,
      unit_price: parseFloat(document.getElementById('foodunitprice').value),
      current_quantity: parseInt(document.getElementById('foodquantity').value),
      minimum_quantity: parseInt(document.getElementById('foodminimum').value),
    };

    makeRequest(
      baseUrl,
      'POST',
      data,
      'application/json',
      true,
      successCreate,
      failedCreate
    );
  }
}

function successEdit(response) {
  console.log('Success editing a food:', response);
  window.location.href = '../foods/foods.html';
}

function failedEdit(response) {
  console.log('Failed to edit a food:', response);
}

function successCreate(response) {
  console.log('Success creating a food:', response);
  window.location.href = '../foods/foods.html';
}

function failedCreate(response) {
  console.log('Failed to create a food:', response);
}

function successDelete(response) {
  console.log(response);
}

function failedDelete(response) {
  alert('Failed deleting the food');
  console.log(response);
}