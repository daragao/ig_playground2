# IG Playground 2
_IG Playgorund_  was another repo which I don't want to delete yet

Trying to make a library for the IG API.

Main challenge seems to be to connect to the streaming API without using the Lightstreamer.

Reference for the IG API: https://labs.ig.com/

Example of usage:
```go
	ig := ig.IGClient{URL: "https://demo-api.ig.com/gateway/deal", APIKey: "API_KEY"}
	err := ig.Login("USERNAME", "PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	defer ig.Logout()

    // print session values
	b, _ := json.MarshalIndent(ig.Session, "\t", "\t")
	log.Printf("\n\t%s", string(b))
```

The _ig_ object is defined in the `client/ig/ig.go` file, and also are most of its functions.
