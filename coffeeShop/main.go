package main
import (
 "fmt"
 "sync"
)

type DataStore struct {
 sync.Mutex // ← this mutex protects the cache below
 cache map[string]string
}

func New() *DataStore{
 return &DataStore{
   cache: make(map[string]string),
 }
}

func (ds *DataStore) set(key string, value string) {
 ds.Lock()
 defer ds.Unlock()
 ds.cache[key] = value
}

func (ds *DataStore) get(key string) string {
 ds.Lock()
 defer ds.Unlock()
 if ds.count() > 0 { //<-- count() also takes a lock!
  item := ds.cache[key]
  return item
 }
 return ""
}

func (ds *DataStore) count() int {
 ds.Lock()
 defer ds.Unlock()
 return len(ds.cache)
}

func main() {
/* Running this below will deadlock because the get() method will       take a lock and will call the count() method which will also take a  lock before the set() method unlocks()
*/
 store := New()
 store.set("Go", "Lang")
 result := store.get("Go")
 fmt.Println(result)
}