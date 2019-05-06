package main

type Fooer interface {
	Foo()
	ImplementsFooer()
}

type Bar struct{}

func (b Bar) ImplementsFooer() {}

func (b Bar) Foo() {}

func main() {

	println("test check em tempo de compilacao!")
}
