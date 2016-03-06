package esc

import (
    "encoding/json"
    "io"
    "log"
    "fmt"
    "net"
    "net/http"
    "net/url"
)
import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "github.com/skratchdot/open-golang/open"
    "github.com/parnurzeal/gorequest"
)

const escClientId = "8a9aa2811b064f8cb1d07d37ac519696"
const escClientSecret = "FQx7pUyFhDzLIK8vxhcKAUznTExC2XpPWaQ7YGfU" // This is OK: http://pastebin.com/TfWg1ywi
const ssoServerPort = 6462

var token string
var messages chan(string)

type MyMainWindow struct {
    *walk.MainWindow
}

func main() {

    mw := new(MyMainWindow)

    var inTE *walk.TextEdit

    MainWindow{
        Title:   "EVE Shopping Cart",
        AssignTo: &mw.MainWindow,
        MinSize: Size{400, 400},
        Layout:  VBox{},
        Children: []Widget{
            TextEdit{AssignTo: &inTE},
            PushButton{
                Text: "Create fitting",
                OnClicked: func() {
                    createAction_Clicked(inTE.Text())
                },
            },
        },
    }.Run()
}

func createAction_Clicked(input string) {
    if token == "" {
        token = getSSOToken()
    }
}

func getSSOToken() string {
    u, err := url.Parse("https://login.eveonline.com/oauth/authorize")
    if err != nil {
        log.Fatal(err)
    }
    q := u.Query()
    q.Set("response_type", "code")
    q.Set("redirect_uri", fmt.Sprintf("http://localhost:%d", ssoServerPort))
    q.Set("client_id", escClientId)
    q.Set("scope", "characterFittingsWrite")
    u.RawQuery = q.Encode()
    open.Run(u.String())

    messages := make(chan string)

    go func() {
        listener, err := net.Listen("tcp",
            fmt.Sprintf(":%d", ssoServerPort))
        if err != nil {
            log.Print(err)
        }
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            // Receive the code information
            values := r.URL.Query()
            messages <- values.Get("code")
            close(messages)

            // Say thank you
            w.WriteHeader(http.StatusOK)
            io.WriteString(w, "You may close this window or tab.")
            listener.Close()
        })
        http.Serve(listener, nil)
    }()

    code := <-messages

    // Go grab the token

    req := gorequest.New().SetBasicAuth(escClientId, escClientSecret)
    req.Post("https://login.eveonline.com/oauth/token")
    req.Send("grant_type=authorization_code&code="+code)
    _, body, errs := req.End()
    if errs != nil {
        log.Print(errs)
    }

    var dat map[string]interface{}
    if err := json.Unmarshal([]byte(body), &dat); err != nil {
        log.Fatal(err)
    }
    log.Print(dat)
    return dat["access_token"].(string)
}
