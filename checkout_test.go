package checkout_test

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/pauldambra/checkoutkata"
)

/*
Item   Unit      Special
         Price     Price
--------------------------
  A     50       3 for 130
  B     30       2 for 45
  C     20
  D     15
*/
func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("when scanning", func() {
		g.Describe("a single C", func() {
			g.It("should make the total 20", func() {
				checkout := Checkout{}
				checkout.Scan("C")
				g.Assert(checkout.Total).Equal(20)
			})
		})

		g.Describe("a single D", func() {
			g.It("should make the total 15", func() {
				checkout := Checkout{}
				checkout.Scan("D")
				g.Assert(checkout.Total).Equal(15)
			})
		})

		g.Describe("two Cs", func() {
			g.It("should make the total 40", func() {
				checkout := Checkout{}
				checkout.Scan("C")
				checkout.Scan("C")
				g.Assert(checkout.Total).Equal(40)
			})
		})

		g.Describe("two Ds", func() {
			g.It("should make the total 30", func() {
				checkout := Checkout{}
				checkout.Scan("C")
				checkout.Scan("C")
				g.Assert(checkout.Total).Equal(40)
			})
		})

		g.Describe("a C and a D", func() {
			g.It("should make the total 35", func() {
				checkout := Checkout{}
				checkout.Scan("C")
				checkout.Scan("D")
				g.Assert(checkout.Total).Equal(35)
			})
		})
		// g.It("Should add two numbers ", func() {
		// 	g.Assert(1 + 1).Equal(2)
		// })
		// g.It("Should match equal numbers", func() {
		// 	g.Assert(2).Equal(4)
		// })
		// g.It("Should substract two numbers")
	})
}
