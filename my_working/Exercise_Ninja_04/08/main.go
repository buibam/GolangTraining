package main

import (
	"fmt"
)
func main(){
	m := map[string] []string {`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`}, `no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`}}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println("This is the record for: ", k)
		for i, val := range v {
			fmt.Printf("%v:\t%v\n", i, val)
		}
	}
}
