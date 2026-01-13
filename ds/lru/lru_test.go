package lru

import "testing"

func TestLRU(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	if v := obj.Get(1); v != 1 {
		t.Errorf("expected 1, got %d", v)
	}

	obj.Put(3, 3) // evicts key 2

	if v := obj.Get(2); v != -1 {
		t.Errorf("expected -1, got %d", v)
	}
}
