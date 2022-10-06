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

	// List Custodial Accounts
	cap := rize.CustodialAccountListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      0,
		Liability:   true,
		Type:        "dda",
	}
	cl, err := rc.CustodialAccounts.List(context.Background(), &cap)
	if err != nil {
		log.Fatal("Error fetching Custodial Accounts\n", err)
	}
	output, _ := json.MarshalIndent(cl, "", "\t")
	log.Println("List Custodial Accounts:", string(output))

	// Get Custodial Account
	ca, err := rc.CustodialAccounts.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Custodial Account\n", err)
	}
	output, _ = json.MarshalIndent(ca, "", "\t")
	log.Println("Get Custodial Account:", string(output))
}
