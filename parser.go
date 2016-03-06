package esc
import (
    "errors"
    "strconv"
    "strings"
    "unicode/utf8"
)

type CartLine struct {
    Item    string
    Count   int
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

        parts[0] = strings.TrimSpace(parts[0])
        parts[1] = strings.TrimSpace(parts[1])

        count, err := strconv.Atoi(parts[0])
        if err != nil {
            // First arg is not the count
            count, err := strconv.Atoi(parts[1])
            if err != nil {
                // Neither arg is a count, skip this line
                continue
            } else {
                // Item, Count
                cart = append(cart, CartLine{
                    Item:  parts[0],
                    Count: count,
                })
            }
        } else {
            // First arg is a count
            _, err := strconv.Atoi(parts[1])
            if err != nil {
                // Count, Item
                cart = append(cart, CartLine{
                    Item:  parts[1],
                    Count: count,
                })
            } else {
                // Both items are a count, skip this line
                continue
            }
        }
    }
    if len(cart) == 0 {
        return cart, errors.New("No items found")
    }
    return cart, nil
}
