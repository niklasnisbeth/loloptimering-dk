package main

import "net/http"
import "io/ioutil"
import "log"

func main() {
  filePaths := []string{"avokade.jpg", "spacer.bmp"};
  
  mkHandler := func (s string) func(w http.ResponseWriter, r *http.Request) {
    file,_ := ioutil.ReadFile(s);
    return func(w http.ResponseWriter, r *http.Request) {
      w.Write(file)
    }
  }

  http.HandleFunc("/", mkHandler("index.html"));

  for i, _ := range filePaths {
    http.HandleFunc("/" + filePaths[i], mkHandler(filePaths[i]));
  }

  log.Fatal(http.ListenAndServe(":10001", nil))
}
