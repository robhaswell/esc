package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func main() {
    var inTE *walk.TextEdit

    MainWindow{
        Title:   "EVE Shopping Cart",
        MinSize: Size{400, 400},
        Layout:  VBox{},
        Children: []Widget{
            TextEdit{AssignTo: &inTE},
            PushButton{
                Text: "Create fitting",
                OnClicked: func() {
                },
            },
        },
    }.Run()
}
