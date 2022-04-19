package gofakeit

import (
	"fmt"
	"math/rand"
)

func HTMLStub() string { return htmlStubFunc(globalFaker.Rand) }

func htmlStubFunc(r *rand.Rand) string {
	elem := getRandValue(r, []string{"html", "element"})

	return fmt.Sprintf("<%s></%s>", elem, elem)
}
