fetch("http://127.0.0.1:8080/department").then((response)=>{
        return response.json()
    }).then(data =>{
       
       
        const department=document.getElementById("department")
        let option=""
        for(let i=0;i<data.length;i++){
            option+="<option value='"+data[i]+"'>"+data[i]+"</option>"
        }
        department.innerHTML=option
        
    }).catch(err =>{
        console.error(err)
})

fetch("http://127.0.0.1:8080/job_title").then((response)=>{
        return response.json()
    }).then(data =>{
       console.log(data);
       
       
        const job_title=document.getElementById("job_title")
        let option=""
        for(let i=0;i<data.length;i++){
            option+="<option value='"+data[i]+"'>"+data[i]+"</option>"
        }
        job_title.innerHTML=option
        
    }).catch(err =>{
        console.error(err)})
