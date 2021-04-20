package prototype

import "errors"

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

func GetShirtCloner() ShirtCloner {
	return &ShirtCache{}
}

type ShirtCache struct {
}

func (*ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		item := *whitePrototype
		return &item, nil
	case Black:
		item := *blackPrototype
		return &item, nil
	case Blue:
		item := *bluePrototype
		return &item, nil
	}
	return nil, errors.New("Not implemented yet")
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return ""
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: Blue,
}

func (shirt *Shirt) GetPrice() float32 {
	return shirt.Price
}
