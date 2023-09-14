package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index <0{
        return -1
    } else if len(slice)-1>=index{        
        return slice[index]
    }
    return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index<0{
        slice=append(slice,value)
        return slice
    } else if len(slice)-1>=index{
    	slice[index]=value
        return slice
    }
	slice=append(slice,value)
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newArray:=[]int{}
    for _,num:=range values{
        newArray=append(newArray,num)
    }
	for _,num:=range slice{
        newArray=append(newArray,num)
    }
	return newArray
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index<0{
        return slice
    } else if len(slice)-1>=index{
    	return append(slice[:index],slice[index+1:]...)
    }
	return slice
}
