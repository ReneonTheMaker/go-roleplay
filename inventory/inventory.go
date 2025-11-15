package inventory

type Inventory interface {
	GetId() string
	GetItems() []Item
	AddItem(item Item) error
	RemoveItem(itemID string) error
	FindItemByID(itemID string) (Item, error)
}

type Weighted interface {
	Inventory
	GetTotalWeight() float64
	GetMaxWeight() float64
	SetMaxWeight(weight float64) error
}

type Categorized interface {
	Inventory
	GetCategories() []string
	AddCategory(category string) error
	RemoveCategory(category string) error
}

type Sortable interface {
	Inventory
	SortByName() error
	SortByWeight() error
	SortByTag() error
}

func GetItemIndex(row, col, columns int) int {
	return row*columns + col
}

func GetItemCoords(index, columns int) (int, int) {
	row := index / columns
	col := index % columns
	return row, col
}
