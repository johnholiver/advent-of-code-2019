package material

import (
	"strconv"
	"strings"
)

type MaterialTransform struct {
	Produces MaterialCounter
	Consumes []MaterialCounter
}

func NewMaterialTransformation(input string) MaterialTransform {
	equation := strings.Split(input, " => ")
	sConsumed := strings.Split(equation[0], ", ")
	mProduced := NewMaterialCount(equation[1])

	mTrans := MaterialTransform{
		Produces: mProduced,
		Consumes: make([]MaterialCounter, len(sConsumed)),
	}

	for i, s := range sConsumed {
		mTrans.Consumes[i] = NewMaterialCount(s)
	}
	return mTrans
}

type MaterialCounter struct {
	Material string
	Count    int
}

func NewMaterialCount(input string) MaterialCounter {
	mc := strings.Split(input, " ")
	c, _ := strconv.Atoi(mc[0])
	return MaterialCounter{
		Material: mc[1],
		Count:    c,
	}
}
