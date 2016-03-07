package main

import (
    "testing"
    "flag"
    "os"
)

var testSSO = false

func TestMain(m *testing.M) {
    flag.BoolVar(&testSSO, "sso", false, "Run SSO tests")
    flag.Parse()
    os.Exit(m.Run())
}

func TestGetSSOToken(t *testing.T) {
    // Validate that an SSO token can be acquired
    if testSSO == false {
        t.SkipNow()
    }
    token := getSSOToken()
    if token == "" {
        t.Fail()
    }
}
