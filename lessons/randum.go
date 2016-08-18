package main

import "fmt"
import "math"
import "math/rand"
import "time"

func main() {

t := time.Now()
var s int = t.Second()

//fmt.Println(s)

for i := 0; i < 15; i++ {


var a float64 = (float64(s) + float64(rand.Intn(15))) * 0.25

var c = math.Ceil(a)

	fmt.Println(c)
}

}