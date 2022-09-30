package main

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk/internal"
	rize "github.com/rizefinance/rize-go-sdk/platform"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.RizeConfig{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewRizeClient(&config)
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
	cl, err := rc.CustodialAccounts.List(&cap)
	if err != nil {
		log.Fatal("Error fetching Custodial Accounts\n", err)
	}
	output, _ := json.Marshal(cl)
	log.Println("List Custodial Accounts:", string(output))

	// Get Custodial Account
	ca, err := rc.CustodialAccounts.Get("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Custodial Account\n", err)
	}
	output, _ = json.Marshal(ca)
	log.Println("Get Custodial Account:", string(output))
}