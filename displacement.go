package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	fmt.Print("Enter values for acceleration, initial velocity, initial displacement (separated by space) ")

	var s string

	reader := bufio.NewReader(os.Stdin)

	s, _ = reader.ReadString('\n')

	nums:=strings.TrimSpace(s)

	data := strings.Split(nums," ")

	a, _ := strconv.ParseFloat(data[0],64)
	v0, _ := strconv.ParseFloat(data[1],64)
	s0, _ := strconv.ParseFloat(data[2],64)

	getdisp := GenDisplayFn(a,v0,s0)

	var t string



	for {
		fmt.Println("Enter time to get displacement or press X to exit")
		fmt.Scan(&t)

		if t =="X" {
			break
		}

		time, _ := strconv.ParseFloat(t,64)

		fmt.Println(getdisp(time))

	}

}

func GenDisplayFn(a,v0,s0 float64) func( t float64) float64 {

	return func(t float64) float64 {
		//fmt.Println(a,v0,s0,t)
		return (a*t*t)/2 + v0*t + s0
	}
}