package esc
import (
    "errors"
    "strconv"
    "strings"
    "unicode/utf8"
)

type CartLine struct {
    Item    string
    Count   int64
}

func parseShoppingList(input string) ([]CartLine, error) {
    var cart []CartLine
    lines := strings.Split(input, "\n")
    for _, line := range lines {
        parts := strings.FieldsFunc(line, func(cp rune) bool {
            commaRune, _ := utf8.DecodeRune([]byte(","))
            tabRune, _ := utf8.DecodeRune([]byte("\t"))
            return cp == commaRune || cp == tabRune
        })
        if len(parts) < 2 {
            continue
        }
        count, err := strconv.ParseInt(parts[0], 10, 32)
        if err != nil {
            // First arg is not the count
            count, err := strconv.ParseInt(parts[1], 10, 32)
            if err != nil {
                // Neither arg is a count, skip this line
                continue
            } else {
                // Item, Count
                cart = append(cart, CartLine{
                    Item:  strings.TrimSpace(parts[0]),
                    Count: count,
                })
            }
        } else {
            _, err := strconv.ParseInt(parts[1], 10, 32)
            if err != nil {
                // Both items are a count, skip this line
                continue
            } else {
                // Count, Item
                cart = append(cart, CartLine{
                    Item:  strings.TrimSpace(parts[1]),
                    Count: count,
                })
            }
        }
    }
    if len(cart) == 0 {
        return cart, errors.New("No items found")
    }
    return cart, nil
}
