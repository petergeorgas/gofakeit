package gofakeit

import (
	"fmt"
	"testing"
)

func TestHTMLStub(t *testing.T) {
	Seed(14)

	fmt.Println(HTML(&HTMLOptions{}))
	// Output: <button></button>
}
func TestHTMLFields(t *testing.T) {
	Seed(12)

	res, _ := HTML(&HTMLOptions{
		Fields: []Field{
			{Name: "name", Function: "name"},
		},
	})
	fmt.Println(res)
	// Output: <h3>Abdullah Roob</h3>
}
