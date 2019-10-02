package main
import (
	// "fmt"
	"sort"
	"strings"
	"net/http"
	"encoding/json"
	"log"
)
// Data structure to contain the input data
type RuneSlice []rune
func (r RuneSlice)Len() int { 
	return len(r)
}
func (r RuneSlice)Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r RuneSlice)Less(i, j int) bool {
	return r[i] < r[j]
}
func GetKey(s string) string {
	r := (RuneSlice)(strings.ToLower(s))
	sort.Sort(r)
	return string(r)
}
type multimap map[string] []string
func (data multimap)Insert(val string) {
	key := GetKey(val)
	lst, ok := data[key]
	if ok {
		for _, el := range lst {
			if el == val {
				return
			}
		}
		data[key] = append(lst, val)
	} else {
		data[key] = append(make([]string, 0), val)
	}
}
func (data multimap)Get(val string) (ret []string) {
	ret, _ = data[GetKey(val)]
	return
}
// server part
type handler struct {
	data multimap
}
func (h *handler)handleGet(w http.ResponseWriter, r *http.Request) { 
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	val := r.FormValue("word")
	b, err := json.Marshal(h.data.Get(val))
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
func (h *handler)handleLoad(w http.ResponseWriter, r *http.Request) { 
	h.data = make(multimap)   
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	var m []string
	for key, _ := range r.PostForm {
		err := json.Unmarshal([]byte(key), &m)
		if err != nil {
			log.Fatal(err)
		}
		for _, val := range m {
			h.data.Insert(val)
		}
	}
}
// main
func main() {
	port := "8080"
	host := "localhost"
	h := handler{nil} 
	
	http.HandleFunc("/get", (&h).handleGet)
	http.HandleFunc("/load", (&h).handleLoad)

	http.ListenAndServe(host + ":" + port, nil)
}