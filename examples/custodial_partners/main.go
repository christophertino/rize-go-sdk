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

	// List Custodial Partner
	cl, err := rc.CustodialPartner.List()
	if err != nil {
		log.Fatal("Error fetching Custodial Partners\n", err)
	}
	output, _ := json.Marshal(cl)
	log.Println("List Custodial Partners:", string(output))

	// Get Custodial Partner
	cg, err := rc.CustodialPartner.Get("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching CustodialPartner\n", err)
	}
	output, _ = json.Marshal(cg)
	log.Println("Get Custodial Partner:", string(output))
}