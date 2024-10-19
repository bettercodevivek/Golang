const Form = document.getElementById('form')

Form.addEventListener('submit', async function(event){
    event.preventDefault();
    const Username = document.getElementById('username').value;
    const Password = document.getElementById('password').value;
    const Date = document.getElementById('date').value;
    
    const response = await fetch("http://localhost:8080/register",{
        method:'POST',
        headers:{
            'content-type':'application/json'
        },
        body:JSON.stringify({Username,Password,Date})
    })
    
    const result = await response.json()

   document.getElementById('responseMessage').textContent=result.message;

});