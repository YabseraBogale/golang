package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,avPrepTime int) int{
    switch{
        case avPrepTime==0: return len(layers)*2
    	default: return len(layers)*avPrepTime
    }
}
// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int,float64){
    var sauce float64
    var noodles int
    for _,layer:=range layers{
        if layer=="sauce"{
    		sauce+=1
        } else if layer=="noodles"{
        	noodles+=1
        }
    }
	sauce=sauce*0.2
    noodles=noodles*50
	return noodles,sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1]=friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64,cook int) []float64{
     portions:=[]float64{}
     for _,values:=range quantities {
     		portions=append(portions,(values/2)*float64(cook))
     }
	return portions
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
