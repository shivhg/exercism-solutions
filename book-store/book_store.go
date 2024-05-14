package bookstore

func Cost(books []int) int {
	totalBooks := 0
	discountBuckets := make(map[int][]int)

	for i, newBookId := range books {
		totalBooks += 1
		handled := false
		for key, value := range discountBuckets {
			found := false
			for _, bookId := range value {
				if bookId == newBookId {
					found = true
					break
				}
			}

			if !found {
				discountBuckets[key] = append(discountBuckets[i], newBookId)
				handled = true
			}
		}

		if !handled {
			discountBuckets[i] = append(discountBuckets[i], newBookId)
		}
	}

	totalCost := 0
	for key, value := range discountBuckets {
		totalCost +=
	}
	return totalBooks * 800
}
