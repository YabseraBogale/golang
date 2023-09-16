package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen":3,
        "half_of_a_dozen":	6,
        "dozen":12,
        "small_gross":120,
        "gross":144,
        "great_gross":1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	pieces, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += pieces
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	rpieces, rexists := units[unit]
	bpieces, bexists := bill[item]
	if !rexists || !bexists || rpieces > bpieces {
		return false
	} else if rpieces == bpieces {
		delete(bill, item)
	} else {
		bill[item] -= rpieces
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value,ok:=bill[item]
    return value,ok
}
