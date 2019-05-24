package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){


	arr := make([]int,0,10)
	
	fmt.Println("Enter single space separated integers upto 10 ")
	var s string

	reader := bufio.NewReader(os.Stdin)

	s, _ = reader.ReadString('\n')

	nums:=strings.TrimSpace(s)

	data := strings.Split(nums," ")


	for _,el:= range data {


		n,_:=strconv.Atoi(el)
		arr=append(arr,n)
	}

	fmt.Println("Before sorting ",arr)
	
	bubbleSort(arr)

	fmt.Println("After sorting : ",arr)
}

func bubbleSort( arr []int){
	
	for i,_ := range(arr){
		
		for j := 0;j<len(arr)-1-i;j++ {

			if arr[j] > arr[j+1]{

				swap(arr,j)
		
			}
		
		}
	}
}

func swap(arr []int,index int){

	arr[index],arr[1+index] = arr[index+1],arr[index]

}
