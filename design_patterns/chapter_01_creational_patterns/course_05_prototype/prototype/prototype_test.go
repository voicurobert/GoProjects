package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("Cache is nil")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type asserion failed")
	}

	shirt1.SKU = "aabbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type asserion failed")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("shirt1 cannot be equal to shirt 2")
	}

	t.Log(shirt1.GetInfo())
	t.Log(shirt2.GetInfo())
}
