package main

import (  
	"fmt"
	"math"
	"math/rand"
	"time"
	"runtime"
)

func Sqrt(x float64) float64 {
	if  (x == 1 || x == 0 )  { 
		return x 
	}
	var i = 1.0000
	var t=1.0
	for t <= x {
		i+= .1
		t = float64(i) * float64(i)
	}
	return float64(i)-.1

}
func counting () {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
	fmt.Println("done")
}

func main() {
	counting () 


fmt.Println("\n------------------------------------------")
fmt.Println("Working for the weekend, switch, date string")
fmt.Println("-------------------------------------------")
today := time.Now().Weekday() 
fmt.Println(time.Saturday)
switch time.Saturday {
	case today + 0: 
		fmt.Println("Today", today)
	case today + 1: 
		fmt.Println("Tomorrow", today + 1)
	case today + 2: 
		fmt.Println("So close, but only ", today )
	default: 
	fmt.Println("Too far away")
}

fmt.Printf("it's currently %v\n", time.Now().Minute())


fmt.Println("------------------------------------------")
fmt.Println("Bool, while loop until random 0-10 gets 1")
fmt.Println("-------------------------------------------")
	var found bool = true 
	var count int = 0
	for count=0; found ; count++ { 
	rand.Seed(time.Now().UTC().UnixNano())
	switch a := rand.Intn(10); a {
	case 1:
		fmt.Printf("Winner Number 1,");
		found = false
	case 4:
		fmt.Printf("fore,");
	default:
		fmt.Printf("%v,",a);
	}
		fmt.Printf(" ");
		//count = i
}
		fmt.Printf("\nHad to try %v time(s) to find a 1\n \n",count);

fmt.Println("------------------------------------------")
fmt.Println("read runtime properties")
fmt.Println("-------------------------------------------")
    fmt.Printf("OS %v \n",runtime.GOOS)

	rand.Seed(time.Now().UTC().UnixNano())
	var a = rand.Intn(100)
	fmt.Printf("Rando number is: %v \n",a)

	 start := time.Now()
	fmt.Printf("My sqrt: %v ",Sqrt(float64(a)))
    elapsed := time.Since(start)
		fmt.Printf(" took: %v \n",elapsed)
	 start = time.Now()
	fmt.Printf("Golang sqrt: %v ", math.Sqrt(float64(a)))
    elapsed = time.Since(start)
	fmt.Printf(" took: %v \n",elapsed)
	sum := 1
	for sum < 1000 {
		sum += sum
		lim := 200
		fmt.Printf("%v " , math.Pow(float64(sum),2))
		if sum < lim {
			fmt.Printf("%v ",sum)
		} else {
			fmt.Printf("No:%v ", sum)
		}
		// can't use v here, though
	}
	fmt.Printf("\nSum: %v \n",sum)
}
