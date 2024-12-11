package awards

var clothingBrands []string

func init_clothingBrands() {
	clothingBrands = []string{
		"Nike", "Adidas", "Uniqlo", "Zara", "H&M",
		"Levi's", "Gucci", "Prada", "Armani", "Supreme",
	}
	clothingBrands = addPrefix("衣服：", clothingBrands)
}
