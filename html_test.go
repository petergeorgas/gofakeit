package gofakeit

import (
	"fmt"
	"testing"
)

func TestHTMLStub(t *testing.T) {
	Seed(14)

	fmt.Println(HTMLStub())
	// Output: <button></button>
}
func TestHTMLStub2(t *testing.T) {
	Seed(10)

	fmt.Println(HTMLStub())
	// Output: <h1></h1>
}
