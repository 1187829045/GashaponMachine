package awards

func addPrefix(prefix string, brands []string) []string {
	for i := range clothingBrands {
		clothingBrands[i] = prefix + clothingBrands[i]
	}
	return clothingBrands
}
