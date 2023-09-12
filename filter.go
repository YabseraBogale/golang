package main

import "fmt"


func filter(a []int, cond func(int) bool) []int{
    result :=[]int{}
    for _,n:=range a{
            if cond(n){
                // result := append(result,n) cann't work b/c := is for new assigement and it becomes new variable
                result=append(result,n)
        }
    }

    return result
}

func main(){
    a:= []int{1,2,3,4,5,6}
    even:=filter(a,
                 func(n int) bool{
                    return n%2==0
        })
    fmt.Println(even)
}
