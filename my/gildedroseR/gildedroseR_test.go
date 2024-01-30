package gildedroseR

import "testing"

func TestUpdateQuality_normalItem(t *testing.T) {
	// define item
	item := []*Item{
		{"Foo", 10, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Foo" {
		t.Errorf("Expected item Name to be Foo, got %s", item[0].Name)
	}
	if item[0].SellIn != 9 {
		t.Errorf("Expected item SellIn to be 9, got %d", item[0].SellIn)
	}
	if item[0].Quality != 9 {
		t.Errorf("Expected item Quality to be 9, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_staleItem(t *testing.T) {
	// define item
	item := []*Item{
		{"Foo", 0, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Foo" {
		t.Errorf("Expected item Name to be Foo, got %s", item[0].Name)
	}
	if item[0].SellIn != -1 {
		t.Errorf("Expected item SellIn to be -1, got %d", item[0].SellIn)
	}
	if item[0].Quality != 8 {
		t.Errorf("Expected item Quality to be 8, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_minQuality(t *testing.T) {
	// define item
	item := []*Item{
		{"Foo", 0, 0},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Foo" {
		t.Errorf("Expected item Name to be Foo, got %s", item[0].Name)
	}
	if item[0].SellIn != -1 {
		t.Errorf("Expected item SellIn to be -1, got %d", item[0].SellIn)
	}
	if item[0].Quality != 0 {
		t.Errorf("Expected item Quality to be 0, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_AgedBrie(t *testing.T) {
	// define item
	item := []*Item{
		{"Aged Brie", 10, 0},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Aged Brie" {
		t.Errorf("Expected item Name to be Aged Brie, got %s", item[0].Name)
	}
	if item[0].SellIn != 9 {
		t.Errorf("Expected item SellIn to be 9, got %d", item[0].SellIn)
	}
	if item[0].Quality != 1 {
		t.Errorf("Expected item Quality to be 1, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_AgedBrieExpired(t *testing.T) {
	// define item
	item := []*Item{
		{"Aged Brie", 0, 0},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Aged Brie" {
		t.Errorf("Expected item Name to be Aged Brie, got %s", item[0].Name)
	}
	if item[0].SellIn != -1 {
		t.Errorf("Expected item SellIn to be -1, got %d", item[0].SellIn)
	}
	if item[0].Quality != 2 {
		t.Errorf("Expected item Quality to be 2, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_maxQuality(t *testing.T) {
	// define item
	item := []*Item{
		{"Aged Brie", 10, 50},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Aged Brie" {
		t.Errorf("Expected item Name to be Aged Brie, got %s", item[0].Name)
	}
	if item[0].SellIn != 9 {
		t.Errorf("Expected item SellIn to be 9, got %d", item[0].SellIn)
	}
	if item[0].Quality != 50 {
		t.Errorf("Expected item Quality to be 50, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_Sulfuras(t *testing.T) {
	// define item
	item := []*Item{
		{"Sulfuras, Hand of Ragnaros", 10, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Sulfuras, Hand of Ragnaros" {
		t.Errorf("Expected item Name to be Sulfuras, Hand of Ragnaros, got %s", item[0].Name)
	}
	if item[0].SellIn != 10 {
		t.Errorf("Expected item SellIn to be 10, got %d", item[0].SellIn)
	}
	if item[0].Quality != 10 {
		t.Errorf("Expected item Quality to be 10, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_SulfurasHQ(t *testing.T) {
	// define item
	item := []*Item{
		{"Sulfuras, Hand of Ragnaros", 10, 80},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != "Sulfuras, Hand of Ragnaros" {
		t.Errorf("Expected item Name to be Sulfuras, Hand of Ragnaros, got %s", item[0].Name)
	}
	if item[0].SellIn != 10 {
		t.Errorf("Expected item SellIn to be 10, got %d", item[0].SellIn)
	}
	if item[0].Quality != 80 {
		t.Errorf("Expected item Quality to be 80, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_BackstagePass(t *testing.T) {
	// helpers
	name := "Backstage passes to a TAFKAL80ETC concert"
	// define item
	item := []*Item{
		{name, 15, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != 14 {
		t.Errorf("Expected item SellIn to be 14, got %d", item[0].SellIn)
	}
	if item[0].Quality != 11 {
		t.Errorf("Expected item Quality to be 11, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_BackstagePass10d(t *testing.T) {
	// helpers
	name := "Backstage passes to a TAFKAL80ETC concert"
	// define item
	item := []*Item{
		{name, 10, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != 9 {
		t.Errorf("Expected item SellIn to be 9, got %d", item[0].SellIn)
	}
	if item[0].Quality != 12 {
		t.Errorf("Expected item Quality to be 12, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_BackstagePass5d(t *testing.T) {
	// helpers
	name := "Backstage passes to a TAFKAL80ETC concert"
	// define item
	item := []*Item{
		{name, 5, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != 4 {
		t.Errorf("Expected item SellIn to be 4, got %d", item[0].SellIn)
	}
	if item[0].Quality != 13 {
		t.Errorf("Expected item Quality to be 13, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_BackstagePassExpired(t *testing.T) {
	// helpers
	name := "Backstage passes to a TAFKAL80ETC concert"
	// define item
	item := []*Item{
		{name, 0, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != -1 {
		t.Errorf("Expected item SellIn to be -1, got %d", item[0].SellIn)
	}
	if item[0].Quality != 0 {
		t.Errorf("Expected item Quality to be 0, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_ConjuredItem(t *testing.T) {
	// helpers
	name := "Conjured Mana Cake"
	// define item
	item := []*Item{
		{name, 10, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != 9 {
		t.Errorf("Expected item SellIn to be 9, got %d", item[0].SellIn)
	}
	if item[0].Quality != 8 {
		t.Errorf("Expected item Quality to be 8, got %d", item[0].Quality)
	}
}

func TestUpdateQuality_ConjuredItemExpired(t *testing.T) {
	// helpers
	name := "Conjured Mana Cake"
	// define item
	item := []*Item{
		{name, -1, 10},
	}
	// call update quality
	UpdateQuality(item)

	// assert
	if item[0].Name != name {
		t.Errorf("Expected item Name to be %s, got %s", name, item[0].Name)
	}
	if item[0].SellIn != -2 {
		t.Errorf("Expected item SellIn to be -2, got %d", item[0].SellIn)
	}
	if item[0].Quality != 6 {
		t.Errorf("Expected item Quality to be 6, got %d", item[0].Quality)
	}
}
