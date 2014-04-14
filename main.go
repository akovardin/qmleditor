package main

import (
    "fmt"
    "gopkg.in/qml.v0"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

type Editor struct {
    Text    string
    FileUrl string
}

func (e *Editor) SelectFile(fileUrl string) {
    fmt.Println("Selected file: ", fileUrl)
    e.FileUrl = strings.Replace(fileUrl, "file:/", "", 1)
    dat, err := ioutil.ReadFile(e.FileUrl)
    if err != nil {
        log.Println(err)
    }
    e.Text = string(dat)
    qml.Changed(e, &e.Text)
}

func (e *Editor) SaveFile(text string) {
    dat := []byte(text)
    err := ioutil.WriteFile(e.FileUrl, dat, 0644)
    if err != nil {
        log.Println(err)
    }
    fmt.Println("Save file")
}

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    qml.Init(nil)

    engine := qml.NewEngine()
    component, err := engine.LoadFile("Example.qml")
    if err != nil {
        return err
    }

    context := engine.Context()
    context.SetVar("editor", &Editor{})

    win := component.CreateWindow(nil)

    win.Show()
    win.Wait()

    return nil
}
