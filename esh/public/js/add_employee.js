fetch("http://127.0.0.1:8080/office_department_job_title").then((response)=>{
        return response.json()
    }).then(data =>{
       const job_title=document.getElementById("job_title")
       const office=document.getElementById("office")
        const department=document.getElementById("department")
        let option_office=""
        let option_department=""
        let option_job_title=""
        for(let i=0;i<data.length;i++){
            option_office+="<option value='"+data[i][0]+"'>"+data[i]+"</option>"
            option_department+="<option value='"+data[i][1]+"'>"+data[i]+"</option>"
            option_job_title+="<option value='"+data[i][2]+"'>"+data[i]+"</option>"
        }
        office.innerHTML=option_office
        job_title.innerHTML=option_job_title
        department.innerHTML=option_department
        
    }).catch(err =>{
        console.error(err)
})