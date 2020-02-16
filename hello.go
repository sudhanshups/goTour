package main

import ("fmt"
		"math/cmplx"
		"math"
)

var (//types 
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X, Y float64
}


type Abser interface {
	Abss() float64
}

type MyFloat float64


func main() {
	fmt.Printf("hello, world\n")

	fmt.Printf("a-- Start\n")
	fmt.Println(add(42, 13), "\na-- End\n")


	a, b := swap("hello", "world")
	fmt.Println(a, b, "\nb-- End\n")

	c, d := split(10)
	fmt.Println(c, d, "\nc-- End\n")
	
	var i int
	var python, java bool
	fmt.Println(i, c, python, java,"\nd-- End\n")

	var ii, jj int = 1, 2
	k := 3 //short declaration, only works within function
	fmt.Println(ii, jj,k , "\ne-- End\n")


	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\nf-- End\n\n", z, z)

	//type conversion 
	iii := 42.893489 //iii becomes type of float automatically
	f := float64(iii)
	u := uint(f)
	fmt.Println(iii,f,u, "\ng-- End\n")


	const Pi = 3.14 //direct equall  no := 
	fmt.Println(Pi, "\nh-- End\n")


	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum, "\ng-- End\n")

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum, "\ng-- End\n")

	fmt.Println(pow(3, 2, 10),pow(3, 3, 20), "\nh-- End\n")


	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println(sum, "\ni-- End\n")


	//var m map[string]Vertex
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["sud"] = Vertex{
		10,20,
	}

	fmt.Println(m)
	delete(m,"sud")
	elem, ok := m["sud"]

	fmt.Println(m )
	fmt.Println(elem, ok, "\n j-- End\n")
	

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i),neg(-2*i),)
	}
	fmt.Println("\n k-- End\n")

	v := Vertex{3, 4}
	fmt.Println(v.Abs(), "\n l-- End\n")

	v = Vertex{3, 4}
	v.Scale(10) //original values are modified
	fmt.Println(v.Abs(), "\n-- End\n")


	v = Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p),"\nn-- End\n")


	//interface 
	var aa Abser
	ff := MyFloat(-math.Sqrt2)
	v = Vertex{3, 4}

	aa = ff  // aa MyFloat implements Abser
	fmt.Println(aa.Abss())
	
	aa = &v // aa *Vertex implements Abser
	fmt.Println(aa.Abss(), "\nm -- End\n")


	var iiii interface{} = "hello"

	s := iiii.(string)
	fmt.Println(s)

	s, ok = iiii.(string)
	fmt.Println(s, ok)

	f, ok = iiii.(float64)
	fmt.Println(f, ok, "\nn -- End\n")

	//f = iii.(float64) // panic
	//fmt.Println(f)

}


func (f MyFloat) Abss() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abss() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func split(sum int) (x, y int) { //naked/named  return used for very small fuction 
	x = sum * 4
	y = sum - 4
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}
