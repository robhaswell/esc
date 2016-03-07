package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestEmpty(t *testing.T) {
    _, err := parseShoppingList("")
    if err == nil {
        t.Fail()
    }
}

func TestNotEnoughElements(t *testing.T) {
    _, err := parseShoppingList("a line\n")
    if err == nil {
        t.Fail()
    }
}

func TestNotValidElements(t *testing.T) {
    _, err := parseShoppingList("an item, another item")
    if err == nil {
        t.Fail()
    }
}

func TestNotValidDoubleCounts(t *testing.T) {
    _, err := parseShoppingList("1, 2\n3, 4")
    if err == nil {
        t.Fail()
    }
}

func TestCSVItemFirst(t *testing.T) {
    cart, _ := parseShoppingList("An item, 100\nMore items, 200")
    expected := []CartLine{
        CartLine{
            Item:   "An item",
            Count:  100,
        },
        CartLine{
            Item:   "More items",
            Count:  200,
        },
    }
    if !reflect.DeepEqual(cart, expected) {
        fmt.Printf("Got:\t\t%v\nExpected:\t%v\n", cart, expected)
        t.Fail()
    }
}

func TestCSVCountFirst(t *testing.T) {
    cart, _ := parseShoppingList("100, An item\n200, More items")
    expected := []CartLine{
        CartLine{
            Item:   "An item",
            Count:  100,
        },
        CartLine{
            Item:   "More items",
            Count:  200,
        },
    }
    if !reflect.DeepEqual(cart, expected) {
        fmt.Printf("Got:\t\t%v\nExpected:\t%v\n", cart, expected)
        t.Fail()
    }
}

func TestDTVItemFirst(t *testing.T) {
    cart, _ := parseShoppingList("An item\t100\nMore items\t200")
    expected := []CartLine{
        CartLine{
            Item:   "An item",
            Count:  100,
        },
        CartLine{
            Item:   "More items",
            Count:  200,
        },
    }
    if !reflect.DeepEqual(cart, expected) {
        fmt.Printf("Got:\t\t%v\nExpected:\t%v\n", cart, expected)
        t.Fail()
    }
}

func TestDTVCountFirst(t *testing.T) {
    cart, _ := parseShoppingList("100\tAn item\n200\tMore items")
    expected := []CartLine{
        CartLine{
            Item:   "An item",
            Count:  100,
        },
        CartLine{
            Item:   "More items",
            Count:  200,
        },
    }
    if !reflect.DeepEqual(cart, expected) {
        fmt.Printf("Got:\t\t%v\nExpected:\t%v\n", cart, expected)
        t.Fail()
    }
}

