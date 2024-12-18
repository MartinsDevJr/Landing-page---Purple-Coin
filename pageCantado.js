
//função para o menu mobile (hamburger)
const navBar = document.getElementById ("hamburger");
const hamburger = document.querySelector('.hamburger')
const navLinks = document.querySelector('.nav-links')

hamburger.addEventListener('click', () => {
  hamburger.classList.toggle('active');
  navLinks.classList.toggle('active');
})

navLinks.querySelectorAll('a').forEach(link =>{
  link.addEventListener('click', () => {
    hamburger.classList.remove('active');
    navLinks.classList.remove('active');
  })
})




const phoneInput = document.getElementById("phone");

phoneInput.addEventListener("input", (e) => {
  let value = e.target.value.replace(/\D/g, "");
  value = value.replace(/^(\d{2})(\d)/, "($1) $2");
  value = value.replace(/(\d{5})(\d)/, "$1-$2");
  e.target.value = value;
});
