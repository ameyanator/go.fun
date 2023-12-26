package main

func main() {
	shirtItem := newItem("Nike Shirt")
	ameya := NewCustomer(shirtItem, "Ameya")
	soumya := NewCustomer(shirtItem, "Soumya")

	shirtItem.register(ameya)
	shirtItem.register(soumya)

	shirtItem.updateAvailability()
}
