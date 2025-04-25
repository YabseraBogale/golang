fetch("https://127.0.0.1:8080/department").then((response)=>{
        return response.json()
    }).then(data =>{
       
        console.log(data);
        
    }).catch(err =>{
        console.error(err)
    })
