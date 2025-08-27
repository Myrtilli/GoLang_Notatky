
package main

import (
 "log"
 "net/http"
 "fmt"
)


func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
        // Просто виводимо що отримали запит
 log.Printf("http request received from %s\nMethod type : %s", r.URL.Path, r.Method)
        
       // Пишемо у відповідь, що статус ОК
 w.WriteHeader(http.StatusOK)
       // Пишемо у Response Body текст
 fmt.Fprintf(w, "YEEEES, IT IS WORKING!!!!")
       // Виводимо на консоль про відправку response назад до клієнта
 log.Printf("Sent StatusOK back to %s", r.URL.Path)
}


func main() {
 http.HandleFunc("/", httpRequestHandler)
 log.Fatal(http.ListenAndServe(":8080", nil))
}
