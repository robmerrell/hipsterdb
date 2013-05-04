package datastore

import (
	"container/list"
	"errors"
	"time"
)

type Datastore struct {
	// OutOfStyleSeconds is the number of seconds until an item that has gone
	// mainstream goes out of style and can be used again.
	OutOfStyleSeconds uint32

	// MainstreamThreshold is the number of accesses that make an item mainstream
	// and thus unusable until it has gone out of style.
	MainstreamThreshold uint8
}

// the key value store
var store map[string]*Item

// a reference to items that have gone mainstream
var mainstreamKeys *list.List

// errors for the datastore
const (
	ACCESS_MISSING        string = "Even I have never heard of that key before"
	ACCESS_ERR_MAINSTREAM string = "Sorry, this item has gone mainstream"
)

// prepare the datastore
func init() {
	store = make(map[string]*Item)
	mainstreamKeys = list.New()
}

// Insert an item into the datastore. If an item already exists at that key then it will be overwritten.
func (self *Datastore) InsertItem(item *Item) {
	store[item.Key] = item
}

// Delete an item from the datastore.
func (self *Datastore) DeleteItem(key string) {
	delete(store, key)
}

// Get an item from the datastore and inc its mainstream score. Returns an
// error if the key does not exist or the item in mainstream.
func (self *Datastore) GetItem(key string) (*Item, error) {
	item := store[key]

	// make sure the item exists
	if item == nil {
		return nil, errors.New(ACCESS_MISSING)
	}

	// make sure that it hasn't gone mainstream
	if item.MainstreamScore >= self.MainstreamThreshold {
		return nil, errors.New(ACCESS_ERR_MAINSTREAM)
	}

	item.IncrementMainstreamScore(self.MainstreamThreshold, self.OutOfStyleSeconds)

	return item, nil
}

// Start watching mainstream items to go out of style so that they can be used again.
func (self *Datastore) ProcessOutOfStyle() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for e := mainstreamKeys.Front(); e != nil; e = e.Next() {
					key := e.Value.(string)
					item := store[key]
					item.DecrementOutOfStyle()
				}
			}
		}
	}()
}
