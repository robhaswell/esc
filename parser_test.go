package esc

import (
    "testing"
)

func testEmpty(t *testing.T) {
    cart, err = parseShoppingList("")
    if err == nil {
        t.Fail("No error raised")
    }
}

func TestCSVItemFirst(t *testing.T) {
}
