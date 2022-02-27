package main

import "fmt"

func isElementInArray(val int,ar []int) bool{
	in := false
	for i := 0; i < len(ar); i++ {
		if ar[i] == val {
			in = true
		}
	}
	return in
}

func getUniqueCombinations(dn []int) []int{
	var used []int
	var comb []int
	
	for i := 0; i < len(dn); i++ {
		used = append(used, dn[i])
		if isElementInArray(dn[i],comb) == false {	
			comb = append(comb,dn[i])
		}
		
		for x := 0;x < len(dn);x++ {
			if isElementInArray(dn[x],used) == false {	
			    newValue := dn[i]+dn[x]
				if isElementInArray(newValue,comb) == false {	
					comb = append(comb,newValue)
				}
			}
		}
		
	}
	fmt.Println(comb)
	return comb
}

func getLowestAmtCantMakeCash(ar []int) int {
	canMake := true
	amt := 1
	comb := getUniqueCombinations(ar)
	
	for canMake == true {
		canMake = (isElementInArray(amt,comb) == true)
		amt++
	}
	
	return amt
}

func payAmtInLeastBills(amt int,ar []int, count int) int {
	lowestBillCount := -1
	
	for i := 0; i < len(ar); i++ {
		if amt-ar[i] == 0 {
			return count+1;
		}
		
		if amt-ar[i] > 0 {
			a := payAmtInLeastBills(amt-ar[i],ar,count+1)
			if a < lowestBillCount || lowestBillCount == -1 {
				lowestBillCount = a
			}
		}
	}
	
	return lowestBillCount

}

func main() {
    fmt.Println("Currency Exchange")
	
	denominations := []int{1, 2, 3, 4, 5}
	
	fmt.Println(getLowestAmtCantMakeCash(denominations))
	fmt.Println(payAmtInLeastBills(20,denominations,0))
}