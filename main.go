package main

import (
	"flag"
	"fmt"
	"github.com/hoisie/web"
	"github.com/robmerrell/hipsterdb/datastore"
	"os"
)

// the datastore that we interface with
var ds *datastore.Datastore

// command line flags
var flagOutOfStyleSeconds, flagMainstreamThreshold uint
var flagPort string

// command line usage
func usage() {
	fmt.Fprintf(os.Stderr, "usage: hipsterdb [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	// setup the command line processing
	flag.UintVar(&flagOutOfStyleSeconds, "o", 100, "seconds it takes for mainstream data to go out of style")
	flag.UintVar(&flagMainstreamThreshold, "m", 20, "how many times data can be accessed before it is mainstream")
	flag.StringVar(&flagPort, "p", ":9999", "port number to run the DB interface on")

	flag.Usage = usage
	flag.Parse()

	// setup our datastore
	ds = &datastore.Datastore{OutOfStyleSeconds: flagOutOfStyleSeconds, MainstreamThreshold: flagMainstreamThreshold}
	datastore.ProcessOutOfStyle()
}

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
	web.Get("/(.*)", getItem)
	web.Post("/(.*)", upsertItem)
	web.Delete("/(.*)", deleteItem)
	web.Run(flagPort)
}
