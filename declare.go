package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//Multiplication of variables
func Multi(x, y int) (half int) {
	half = x * y
	return
}

//Summ of massive of variables
func SumVar(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}

//Letters index and true return
func Index(st string, a rune) (index int, ok bool) {
	for i, c := range st {
		if c == a {
			return i, true
		}
	}
	return // вернутся значения по умолчанию
}

//Fibonacci number
func Fib(n int) int {
	switch {
	case n <= 1: // терминальная ветка — то есть условие выхода из рекурсии
		return n
	default: // рекурсивная ветка
		return Fib(n-1) + Fib(n-2)
	}

}

//Iterative Fibonacci number
func IterFib(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return a
}

//Iterative Factorial
func Fact(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func main() {
	/*var x int
	fmt.Println("Введите число")
	fmt.Scanf("%d\n", &x)
	y := Fib(x)
	fmt.Println(y)
	*/
	PrintAllFiles(".")
}
