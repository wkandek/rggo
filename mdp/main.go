package main

import (
    "bytes"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday/v2"
)

const (
    header = `<HTML><HEAD><TITLE>Markdown Preview Tool</TITLE></HEAD><BODY>`
    footer = `</BODY></HTML>`
)


func parseContent(input []byte) []byte {
    output := blackfriday.Run(input)
    body := bluemonday.UGCPolicy().SanitizeBytes(output)

    var buffer bytes.Buffer
    buffer.WriteString(header)
    buffer.Write(body)
    buffer.WriteString(footer)

    return buffer.Bytes()
}


func saveHTML(outName string, data []byte) error {
    return ioutil.WriteFile(outName, data, 0644)
}


func run(filename string) error {
    input, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    htmlData := parseContent(input)
    outName := fmt.Sprintf("%s.html", filepath.Base(filename))
    fmt.Println(outName)

    return saveHTML(outName, htmlData)
}


func main() {
    filename := flag.String("file", "", "Markdownfile to preview")
    flag.Parse()

    if *filename == "" {
        flag.Usage()
        os.Exit(1)
    }
    err := run(*filename)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}




