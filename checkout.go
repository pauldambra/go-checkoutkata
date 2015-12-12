package checkout

// Checkout allows scanning of products and totalling of bills
type Checkout struct {
	Total int // scanned total
}

var prices = map[string]int{
	"C": 20,
	"D": 15,
}

//Scan scans a product
func (c *Checkout) Scan(code string) {
	c.Total += prices[code]
}
