package main

import (
	"github.com/hoisie/web"
	"github.com/robmerrell/hipsterdb/datastore"
)

// the datastore that we interface with
var ds *datastore.Datastore

// Returns a key or abort with an error if it doesn't exist or the item
// has gone mainstream.
func getItem(ctx *web.Context, key string) string {
	item, err := ds.GetItem(key)
	if err != nil {
		var errorStatus int
		if err.Error() == datastore.ACCESS_MISSING {
			errorStatus = 404
		} else {
			errorStatus = 403
		}

		ctx.Abort(errorStatus, err.Error())
		return ""
	}

	return item.Value
}

// Upsert an item in the datastore.
func upsertItem(ctx *web.Context, key string) string {
	value := ctx.Params["value"]
	if value == "" {
		ctx.Abort(400, "Item creation is missing a value")
		return ""
	}

	// insert the item
	item := datastore.NewItem(key, value)
	ds.InsertItem(item)

	return "created " + key
}

// Delete an item from the datastore.
func deleteItem(ctx *web.Context, key string) string {
	ds.DeleteItem(key)
	datastore.RemoveFromMainstreamKeys(key)

	return "deleted " + key
}

func main() {
	// setup our datastore
	ds = &datastore.Datastore{OutOfStyleSeconds: 3, MainstreamThreshold: 3}
	ds.ProcessOutOfStyle()

	web.Get("/(.*)", getItem)
	web.Post("/(.*)", upsertItem)
	web.Delete("/(.*)", deleteItem)
	web.Run(":9999")
}
