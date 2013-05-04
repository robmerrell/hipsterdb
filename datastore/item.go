package datastore

import (
	"log"
)

// Items are the values stored and retrieved from the database
type Item struct {
	// Key is the location key of the item
	Key string

	// Value is the value of the item
	Value string

	// MainstreamScore is the score given to an item to determine how mainstream it is
	MainstreamScore uint

	// TTOutOfStyle is the time in seconds until this item goes out of style and the MainstreamScore is reset
	TTOutOfStyle uint
}

// NewItem creates a new item for the datastore
func NewItem(key, value string) *Item {
	return &Item{Key: key, Value: value}
}

// Increase the mainstream score of an item. If the item has gone mainstream then also
// set the TTOutOfStyle countdown to use the item again. Returns a boolean stating whether
// the item went mainstream or not.
func (self *Item) IncrementMainstreamScore(mainstreamThreshold, outOfStyleSeconds uint) bool {
	self.MainstreamScore++

	isMainstream := self.MainstreamScore >= mainstreamThreshold
	if isMainstream {
		mainstreamKeys.PushBack(self.Key)
		self.TTOutOfStyle = outOfStyleSeconds
	}

	return isMainstream
}

// Decrement the time to out of style seconds. If it reaches 0 then reset the mainstream score.
func (self *Item) DecrementOutOfStyle() {
	self.TTOutOfStyle--

	if self.TTOutOfStyle <= 0 {
		RemoveFromMainstreamKeys(self.Key)
		log.Printf("%s is now out of style", self.Key)
		self.MainstreamScore = 0
	}
}
