const submit = document.getElementById("submit")

submit.addEventListener('click',(e)=>{
   e.preventDefault();
   console.log(e.offsetX);
})