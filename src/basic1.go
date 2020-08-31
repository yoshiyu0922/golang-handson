package src

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

// 戻り値に名前をつけられる
// 長い関数で使うと読みやすさに悪影響
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func main() {

	fmt.Printf("hello, world\n")

	fmt.Println("My Favorite Number is ", rand.Intn(10))
	fmt.Println("My Favorite Number is ", rand.Intn(100))

	message := fmt.Sprintf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Print(message)

	fmt.Println(math.Pi)

	addResult := add(42, 32)
	fmt.Printf("add(42, 32) : %d\n", addResult)

	addResult2 := add2(35, 23)
	fmt.Printf("add(35, 23) : %d\n", addResult2)

	//test := swap("hello", "world") // この書き方はエラーになる
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	x, y := split(19)
	fmt.Printf("%d = %d + %d\n", 19, x, y)

	var i int
	var c, python, java = true, false, "no!"
	i = 1
	num := 1 // 初期化 & 代入
	num = 2  // 再代入
	fmt.Println(i, num, c, python, java)

	var (
		ToBe   bool       = false
		MaxInt int        = math.MaxInt64
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i1 int
	var f1 float64
	var b1 bool
	var s1 string // 初期値はnullではなく、空文字
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

	// Type Conversion
	var x2, y2 = 3, 4
	var f2 float64 = math.Sqrt(float64(x2*x2 + y2*y2))
	var z2 uint = uint(f2)
	fmt.Println(x2, y2, z2)

	// Constant：先頭大文字
	const World = "世界"
	const Pi = 3.14
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
