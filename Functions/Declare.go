package functions

func Multi(x, y int) (half int) {
	half = x * y
	return
}

func SumVar(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}

func main() {

}
