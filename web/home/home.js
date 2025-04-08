document.addEventListener('DOMContentLoaded', () => {
    let token = isUserLogged();
    if (token == false) {
        window.location.href = 'web/login/login.html';
    }
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
})