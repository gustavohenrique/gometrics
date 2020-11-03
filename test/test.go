package test

import (
	"testing"
)

type FN func(t *testing.T)

func It(ts *testing.T, name string, fn FN) {
	ts.Run(name, fn)
}
