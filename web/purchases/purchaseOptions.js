
document.addEventListener('DOMContentLoaded', displayUserMail);

function displayUserMail() {
    const userInfo = document.getElementById('user-info');
    const userMail = document.createElement('p');
    userMail.textContent = localStorage.getItem('user-mail');
    userMail.classList.add('green-color', 'bold-words', 'user-mail');
    userInfo.appendChild(userMail);
}

   
