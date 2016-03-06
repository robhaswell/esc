package esc
import (
    "strconv"
    "utf8"
    "errors"
)

type cartLine struct {
    Item    string
    Count   uint
}

func parseShoppingList(input string) ([]cartLine, error) {
    var cart []cartLine
    lines := string.Split(input, "\n")
    for _, line := range lines {
        parts = string.FieldsFunc(line, func(cp rune) {
            commaRune, _ = utf8.DecodeRune([]byte(","))
            tabRune, _ = utf8.DecodeRune([]byte("\t"))
            return cp == commaRune || cp == tabRune
        })
        count, err = strconv.ParseInt(parts[0])
        if err != nil {
            // First arg is not the count
            count, err = strconv.ParseInt(parts[1])
            if err != nil {
                // Neither arg is a count, skip this line
                continue
            } else {
                // Item, Count
                cart.Add({
                    Item:  string.Strip(parts[0]),
                    Count: count
                })
            }
        } else {
            // Count, Item
            cart.Add({
                Item:  string.Strip(parts[1]),
                Count: count
            })
        }
    }
    if len(cart) == 0 {
        return []cartLine, errors.New("No items found")
    }
    return cart
}
