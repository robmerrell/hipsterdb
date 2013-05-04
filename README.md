# hipsterDB

### the webscaliest database that you have probably never heard of

hipsterDB is the best way to make your web application hipster without draping your server in flannel and taping mustache hair to your motherboard.

hipsterDB is a key/value store that only returns data as long as it isn't mainstream. The more often that you access a key the more mainstream it becomes. After data has gone mainstream you will have to wait for it to go out of style before using it again.

# Get/Build

1. hipsterDB is written in go, so make sure you have go setup.
2. go get github.com/robmerrell/hipsterdb.git
3. pull on your skinny jeans, don your thick-rimmed glasses and run hipsterDB

# Use

hipsterDB uses a RESTful interface to upsert, get, and delete data.

## upsert
create or update the key _keyname_ with the value in the post data.

```bash
$ curl --data "value=thevalue" "localhost:9999/keyname"
created keyname
```

## get
get the value stored in _keyname_. If the key has been accessed so much that it is considered mainstream an error will be returned.

```bash
$ curl "localhost:9999/keyname"
Sorry, this item has gone mainstream
```

## delete
delete the key from the database

```bash
$ curl -X DELETE "localhost:9999/keyname"
deleted keyname
```

# Limitations
hipsterDB is limited in pretty much every way imaginable. Don't use this in any real life scenario. Really. Don't.

# License

```
Copyright (c) 2013 Rob Merrell

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
