package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime/debug"
	"strings"
	"time"
	"my-first-go-project/user/hello/tempconv"
)

const c = "C"

var v int = 5

type T struct{}

var Pi float64

var a = "G"

func n() { println(a) }

func m() {
	a := "O"
	println(a)
}

func init() { // initialization of package
	Pi = 4 * math.Atan(1)
}

var aVar = 10

func main() {
	var a int
	Func1()
	fmt.Println(a)
	fmt.Println(valueConvert())
	fmt.Println(unkown)
	fmt.Println(man)
	fmt.Println(woman)
	fmt.Println(HOME)
	fmt.Println(USER)
	fmt.Println(GOROOT)
	fmt.Println(&USER)
	n()
	m()
	n()
	fmt.Println(Pi)
	fmt.Println(aVar == 5)
	fmt.Println(strings.HasPrefix("123456", "123"))
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format(time.RFC822))
	fmt.Println(time.Now().Format("2006-01-02"))
	controlConstruction()
	errTest()
	rs, err := Devide(12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rs)
	}

	Greeting("hello", "Enan", "Tom", "James")
	func1("func1")

	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	fmt.Println(arrtest())
	echo(os.Args)
	echo1(os.Args)


	fmt.Println(tempconv.CToF(tempconv.BoilingC))

}

func (t T) Method1() {

}

func Func1() {

}

// 类型转换
func valueConvert() int {
	a := 5.0
	b := int(a)
	return b
}

const (
	unkown = iota
	man
	woman
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func controlConstruction() {
	if 3 == 1 {
		fmt.Println(3 == 1)
	} else if 1 == 1 {
		fmt.Println(1 == 1)
	}
}

func errTest() error {
	f, err := os.Open("hahah")
	if err != nil {
		debug.PrintStack()
		fmt.Println(err)
		return err
	}
	fmt.Println(f)
	return nil
}

type ArithmeticError struct {
	error
}

func (this *ArithmeticError) Error() string {
	return "自定义的error，error名称为算数不合法"
}

func Devide(num1, num2 int) (rs int, err error) {
	if num2 == 0 {
		return 0, &ArithmeticError{}
	} else {
		return num1 / num2, nil
	}
}

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix + ":")
	for i := 0; i < len(who); i++ {
		fmt.Print(" " + who[i])
	}
	fmt.Println()
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func arrtest() string {
	str := "aaa12345678"
	return str[0:3]
}

func echo(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo1(args []string) {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
