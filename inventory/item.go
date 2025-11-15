package inventory

import "time"

type Item interface {
	GetID() 							string
	GetName() 						string
	GetTags() 				    []string
	AddTag(tag string) 		error
	RemoveTag(tag string)	error
}

type Consumable interface {
	Item
	Use() 						error
}

type Equippable interface {
	Item
	Equip() 	bool
	Unequip() bool
}

type Stackable interface {
	Item
	GetQuantity()        int
	SetQuantity(qty int) error
}

type Perishable interface {
	Item
	GetExpirationDate() time.Time
	IsExpired() 		    bool
}

type Describable interface {
	Item
	GetDescription() string
}
