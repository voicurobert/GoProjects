package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nill")
		return
	}

	expectedCounter := counter
	_ = expectedCounter

	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling AddOne for the first time to count, the count must be 1 but it is %d \n", currentCount)
	}

	secondCounter := GetInstance()
	if secondCounter != expectedCounter {
		t.Error("Expected same instance, but got a different instance")
	}

	currentCount = secondCounter.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling AddOne using the second counter, the current count must be 2 but it is %d \n", currentCount)
	}
}
