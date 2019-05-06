package main

import "fmt"
import "os"

type MagicError struct{ Err error }

func (MagicError) Error() string {
	return "[Magic]"
}

func Generate() *os.PathError {
	return nil
}

func Test() error {
	return Generate()
}

func main() {

	_ = (*struct{})(nil)
	_ = []int(nil)

	var _ map[string]int = nil

	// _ = (*int)(nil) == (*bool)(nil) error
	_ = (*rune)(nil) == (*int32)(nil)
	_ = (*byte)(nil) == (*uint8)(nil)
	var bu = (*[]byte)(nil) == (*[]uint8)(nil)
	fmt.Println("[]byte == []uint8]:", bu)

	var ii = (interface{})(nil) == (*int)(nil)
	fmt.Println("interface:: ", ii)

	type IntPtr *int
	var pIntPtr = IntPtr(nil) == (*int)(nil)
	fmt.Println("pIntPtr:: ", pIntPtr)

	var a = (*struct{})(nil)
	fmt.Println(a)

	var n *struct{} = nil
	var i interface{} = nil

	fmt.Println(n == nil)
	fmt.Println(i == nil)

	bb := Generate()
	fmt.Println("Generate:", Generate() == nil, " bb: ", bb == nil)
	fmt.Println("Test:: ", Test() == nil)

}
