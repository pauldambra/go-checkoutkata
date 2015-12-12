package checkout

// Checkout allows scanning of products and totalling of bills
type Checkout struct {
	basket map[string]int //private scanned items
}

type specialOffer struct {
	quantity int
	newPrice int
}

var prices = map[string]int{
	"A": 50,
	"B": 30,
	"C": 20,
	"D": 15,
}

var offers = map[string]specialOffer{
	"A": {3, 130},
	"B": {2, 45},
}

//Scan scans a product
func (c *Checkout) Scan(code string) {
	if c.basket == nil {
		c.basket = make(map[string]int)
	}
	current := c.basket[code]
	c.basket[code] = current + 1
}

func calculatePrice(code string, quantity int) int {
	return prices[code] * quantity
}

func calculateOfferPrice(code string, basketQuantity int, offer specialOffer) int {
	price := 0
	for basketQuantity >= offer.quantity {
		price += offer.newPrice
		basketQuantity -= offer.quantity
	}
	return price + calculatePrice(code, basketQuantity)
}

//Total adds up the current basket and returns the total
func (c Checkout) Total() int {
	total := 0
	for code, quantity := range c.basket {
		offer, exists := offers[code]
		if exists {
			total += calculateOfferPrice(code, quantity, offer)
		} else {
			total += calculatePrice(code, quantity)
		}
	}
	return total
}
