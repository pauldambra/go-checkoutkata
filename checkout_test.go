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

var testBaskets = []struct {
	codes         []string
	expectedTotal int
}{
	{[]string{"A"}, 50},
	{[]string{"B"}, 30},
	{[]string{"C"}, 20},
	{[]string{"D"}, 15},
	{[]string{"C", "C"}, 40},
	{[]string{"D", "D"}, 30},
	{[]string{"C", "D"}, 35},
	{[]string{"A", "D"}, 65},
}

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("when scanning some known baskets", func() {
		g.It("should correctly calculate the totals", func() {

			for _, testBasket := range testBaskets {
				checkout := Checkout{}
				for _, code := range testBasket.codes {
					checkout.Scan(code)
				}
				g.Assert(checkout.Total).Equal(testBasket.expectedTotal)
			}
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
