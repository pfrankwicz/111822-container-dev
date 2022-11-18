package main

import (

        "net/http"
        "github.com/go-git/go-git/v5"

_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
    URL:      "https://github.com/schollz/httpfileserver.git",
    Progress: os.Stdout,
})
 
)

      //using httpnewserver.new for file server. github.com/schollz

func main() {

          // Any request to /static/somefile.txt will serve the file at the location ./somefile.txt

        http.HandleFunc("/static/", httpfileserver.New("/static/", ".").Handle())

        http.ListenAndServe(":8080", nil)

}
