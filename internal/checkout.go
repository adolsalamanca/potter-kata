package checkout

const (
	HarryPotterFirst  = "HarryPotterFirst"
	HarryPotterSecond = "HarryPotterSecond"
	HarryPotterThird  = "HarryPotterThird"
	HarryPotterFourth = "HarryPotterFourth"
	HarryPotterFifth  = "HarryPotterFift"
)

type BookGroup []Book
type Book struct {
	Title string
	Price float64
}

func getDiscountTracker(books []Book) map[string]int {
	discountTracker := make(map[string]int)

	for _, b := range books {
		if b.isPotterBook() {
			discountTracker[b.Title] += 1
		}
	}
	return discountTracker
}

func calculatePrice(books []Book, discountTracker map[string]int) float64 {
	price := 0.0
	// bookGroups
	discount := getDiscount(discountTracker)

	for _, b := range books {
		price += b.Price * ((100 - float64(discount)) / 100)
	}

	return price
}

// getDiscount returns the discount based on the number of different books.
// < 2 books, 0 discount, means we just multiply by 100/100, we do nothing
// 2 books, 5 discount, means we multiply by 95/100, so * 0.95
// 3 books, 10 discount, means we multiply by 90/100, so * 0.90
// 4 books, 20 discount, means we multiply by 80/100, so * 0.80
// 5 books, 25 discount, means we multiply by 75/100, so * 0.75
func getDiscount(discountTracker map[string]int) int {
	if len(discountTracker) == 2 {
		return 5
	}

	if len(discountTracker) == 3 {
		return 10
	}

	if len(discountTracker) == 4 {
		return 20
	}

	if len(discountTracker) == 5 {
		return 25
	}

	return 0
}

func purchase(books []Book) float64 {
	discountTracker := getDiscountTracker(books)

	return calculatePrice(books, discountTracker)
}

func (b Book) isPotterBook() bool {
	switch b.Title {
	case HarryPotterFirst, HarryPotterSecond, HarryPotterThird, HarryPotterFourth, HarryPotterFifth:
		return true
	default:
		return false
	}
}

func (bg BookGroup) getDiscount() int {
	if len(bg) == 2 {
		return 5
	}

	if len(bg) == 3 {
		return 10
	}

	if len(bg) == 4 {
		return 20
	}

	if len(bg) == 5 {
		return 25
	}

	return 0
}
