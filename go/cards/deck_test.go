package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestNewDeck(t *testing.T) {

	g := Goblin(t)

	g.Describe("newDeck", func() {

		a := len(newDeck())
		g.It(" nous desirons un nouveau deck ", func() {
			g.Assert(20).Equal(a)
		})


	})

	
}
