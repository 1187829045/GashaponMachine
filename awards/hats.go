package awards

var hatBrands []string

func init_hat() {
	hatBrands = []string{
		"New Era", "Nike", "Adidas", "Under Armour", "Puma",
		"Columbia", "The North Face", "Patagonia", "Carhartt", "Champion",
	}
	hatBrands = addPrefix("hat", hatBrands)
}
