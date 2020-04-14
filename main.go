package main

import (
	"encoding/json"
	"log"

	"./client/ig"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ig := ig.IGClient{URL: "https://demo-api.ig.com/gateway/deal", APIKey: "API_KEY"}
	err := ig.Login("USERNAME", "PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	defer ig.Logout()
	b, _ := json.MarshalIndent(ig.Session, "\t", "\t")
	log.Printf("\n\t%s", string(b))

	//ig.GetSessionStreamTokens()

	// stream
	ig.Subscribe()
	/*
		// marketData, marketNode, err := ig.GetMarketNavigation("0")
		marketData, marketNode, err := ig.GetMarketNavigation("267252")
		if err != nil {
			log.Fatal(err)
		}
		b, _ = json.MarshalIndent(marketData, "\t", "\t")
		log.Printf("\n\t%s", string(b))
		b, _ = json.MarshalIndent(marketNode, "\t", "\t")
		log.Printf("\n\t%s", string(b))

	*/
	// epic := "IX.D.SPTRD.DAILY.IP"
	/*
		marketDetail, err := ig.GetMarketDetails("IX.D.SPTRD.DAILY.IP")
		if err != nil {
			log.Fatal(err)
		}
		b, _ = json.MarshalIndent(marketDetail, "\t", "\t")
		log.Printf("\n\t%s", string(b))
	*/
	/*
		epic := "IX.D.SPTRD.DAILY.IP"
		priceList, err := ig.GetPrices(epic, "MINUTE_30", "", "", "", "", "")
		if err != nil {
			log.Fatal(err)
		}
		b, _ = json.MarshalIndent(priceList, "\t", "\t")
		log.Printf("\n\t%s", string(b))
	*/
}
