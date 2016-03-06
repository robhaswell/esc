package esc

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "fmt"
    "net"
    "net/http"
    "net/http/httputil"
    "net/url"
)
import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "github.com/skratchdot/open-golang/open"
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
    log.Printf("Opening %s", u)
    open.Run(u.String())

    messages := make(chan string)

    go func() {
        listener, err := net.Listen("tcp",
            fmt.Sprintf(":%d", ssoServerPort))
        if err != nil {
            log.Print(err)
        }
        log.Printf("Listening on %s", fmt.Sprintf(":%d", ssoServerPort))
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            // Receive the code information
            log.Printf("Got hit on %s", r.URL)
            values := r.URL.Query()
            messages <- values.Get("code")
            log.Print("Sent message")
            close(messages)

            // Say thank you
            w.WriteHeader(http.StatusOK)
            io.WriteString(w, "You may close this window or tab.")
            listener.Close()
        })
        http.Serve(listener, nil)
    }()

    log.Print("Waiting for code")
    code := <-messages
    log.Printf("Got '%s'", code)

    // Go grab the token

    client := &http.Client{}
    var authReq = []byte("grant_type=authorization_code&code="+code)
    req, err := http.NewRequest("POST",
        "https://login.eveonline.com/oauth/token",
        bytes.NewBuffer(authReq))
    if err != nil {
        log.Fatal(err)
    }
    req.SetBasicAuth(escClientId, escClientSecret)
    debug(httputil.DumpRequestOut(req, true))
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    jsonBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(jsonBytes))
    var dat map[string]interface{}
    if err := json.Unmarshal(jsonBytes, &dat); err != nil {
        log.Fatal(err)
    }
    log.Print(dat)

    return ""
}


func debug(data []byte, err error) {
    if err == nil {
        log.Printf("%s\n\n", data)
    } else {
        log.Fatalf("%s\n\n", err)
    }
}
