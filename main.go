package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	responseEnv = "RESPONSE_TEXT"
	userdata    = `
#cloud-config
users:
- name: root
  ssh_authorized_keys:
  - ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEHKykavZESjDCpxxI7xd30BU2XfVe75Trqj/WMxn/2f cyrill@mbp
`
	metadata = `{"uuid": "fake"}`
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request on path %s\n", r.Method, r.URL.Path)
		if strings.Contains(r.URL.Path, "user-data") {
			io.WriteString(w, userdata)
			return
		}
		if strings.Contains(r.URL.Path, "meta-data") {
			io.WriteString(w, userdata)
			return
		}

		io.WriteString(w, os.Getenv(responseEnv)+"\n")
	})

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.Fatal(err)
	}
}
