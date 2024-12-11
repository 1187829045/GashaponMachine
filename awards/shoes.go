package awards

var shoeBrands []string

func init_shoe() {

	shoeBrands = []string{
		"Nike", "Adidas", "Puma", "Converse", "Vans",
		"Reebok", "New Balance", "ASICS", "Sketchers", "Under Armour",
	}

	shoeBrands = addPrefix("hat", hatBrands)
}
