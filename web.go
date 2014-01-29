package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  c "github.com/hstove/gender/classifier"
  "log"
  "strconv"
)

var classifier = c.Classifier()

func classify(res http.ResponseWriter, req *http.Request) {
  name := req.URL.Path[len("/classify/"):]
  gender, prob := c.Classify(classifier, name)
  prob = prob * 100
  probability := strconv.FormatFloat(prob, 'f', 6, 64)
  if probability == "59.369936" {
    probability = "?"
    gender = "unknown"
  }
  jsonMap := map[string]string {
    "name": name,
    "gender": gender,
    "probability": probability,
  }
  json, _ := json.Marshal(jsonMap)
  fmt.Println(jsonMap)
  res.Header().Set("Content-Type", "application/json; charset=utf-8")
  res.Write(json)
}

func main() {
  http.HandleFunc("/classify/", classify)
  http.Handle("/", http.FileServer(http.Dir("./public")))
  err := http.ListenAndServe(":5000", nil)
  if err != nil {
      log.Fatal("Error: %v", err)
  }
}
