package gildedroseR

type Item struct {
	Name            string
	SellIn, Quality int
}

func (I *Item) UpdateItem() {
	// Helpers
	isAgedBrie := I.Name == "Aged Brie"
	isBackstagePass := I.Name == "Backstage passes to a TAFKAL80ETC concert"
	isSulfuras := I.Name == "Sulfuras, Hand of Ragnaros"
	isConjured := I.Name == "Conjured Mana Cake"

	if isSulfuras {
		// Do nothing for this item
		return
	}

	// Process quality change
	if isAgedBrie || isBackstagePass {
		// Items that get better by the time
		I.Quality = I.Quality + 1
		if isBackstagePass {
			if I.SellIn < 11 {
				I.Quality = I.Quality + 1
			}
			if I.SellIn < 6 {
				I.Quality = I.Quality + 1
			}
		}
	} else {
		// Items that get worse by the time
		I.Quality = I.Quality - 1
		if isConjured {
			I.Quality = I.Quality - 1
		}
	}

	// Reduce SellIn (= age items)
	I.SellIn = I.SellIn - 1

	// Treat "expired" items
	if I.SellIn < 0 {
		if isAgedBrie {
			I.Quality = I.Quality + 1
		} else if isBackstagePass {
			I.Quality = 0
		} else {
			I.Quality = I.Quality - 1
			if isConjured {
				I.Quality = I.Quality - 1
			}
		}
	}

	// Limit Quality between 0 and 50
	if I.Quality > 50 {
		I.Quality = 50
	} else if I.Quality < 0 {
		I.Quality = 0
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		item.UpdateItem()
	}
}
