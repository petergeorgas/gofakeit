package gofakeit

import (
	"errors"
	"fmt"
	"math/rand"
)

type HTMLOptions struct {
	Nested bool    `json:"nested" xml:"nested"`
	Fields []Field `json:"fields" xml:"fields"` // The fields to be generated
}

func HTML(ho *HTMLOptions) (string, error) { return htmlFunc(globalFaker.Rand, ho) }

func (f *Faker) HTML(ho *HTMLOptions) (string, error) { return htmlFunc(f.Rand, ho) }

func htmlFunc(r *rand.Rand, ho *HTMLOptions) (string, error) {

	if ho.Fields == nil { // Fields not provided, generate a stub.
		htmlElement := getRandValue(r, []string{"html", "element"})

		return fmt.Sprintf("<%s></%s>", htmlElement, htmlElement), nil
	}

	if len(ho.Fields) != 1 { // TOOD: Can we have multiple fields?
		return "", errors.New("must provide only one field")
	}

	funcInfo := GetFuncLookup(ho.Fields[0].Function)
	if funcInfo == nil {
		return "", errors.New("invalid function, " + ho.Fields[0].Function + " does not exist")
	}

	// Generate the value
	val, err := funcInfo.Generate(r, &ho.Fields[0].Params, funcInfo)
	if err != nil {
		return "", err
	}

	htmlElement := getRandValue(r, []string{"html", "element"})
	return fmt.Sprintf("<%s>%s</%s>", htmlElement, val, htmlElement), nil
}
