function showAlert(message) {
    const alert = document.getElementById('alert-modal')
    const icon = document.createElement("i")
    icon.classList.add("fa-solid" ,"fa-triangle-exclamation")
    const alertMessage = document.createElement('p')
    const alertButton = document.createElement('button')
    alertButton.id = 'alert-button'
    alertMessage.textContent = message
    alertButton.textContent = 'Accept'
    alert.appendChild(icon)
    alert.appendChild(alertMessage)
    alert.appendChild(alertButton)
    alert.showModal()

    alertButton.addEventListener(('click'), () => {
        alert.innerHTML = ''
        alert.close()
    })
}
