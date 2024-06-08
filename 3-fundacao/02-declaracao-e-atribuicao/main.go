package main

const a = "hello, world"

var (
	b bool    = true
	c int     = 10
	d string  = "Diego Ramos"
	e float32 = 0.3333333333333333333333333333333333333333333333333333333
)

func main() {
	b = !b

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	f := "Nathália"
	println(f)
}
