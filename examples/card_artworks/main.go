package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List Card Artwork
	clp := rize.CardArtworkListParams{
		ProgramUID: "DbxJUHVuqt3C7hGK",
		Limit:      100,
		Offset:     0,
	}
	cl, err := rc.CardArtworks.List(context.Background(), &clp)
	if err != nil {
		log.Fatal("Error fetching Card Artwork\n", err)
	}
	output, _ := json.MarshalIndent(cl, "", "\t")
	log.Println("List Card Artwork:", string(output))

	// Get Card Artwork
	cg, err := rc.CardArtworks.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching CardArtwork\n", err)
	}
	output, _ = json.MarshalIndent(cg, "", "\t")
	log.Println("Get CardArtwork:", string(output))
}
