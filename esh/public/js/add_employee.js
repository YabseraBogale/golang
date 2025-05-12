fetch("http://127.0.0.1:8080/office_department_job_title").then((response)=>{
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