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

	// List Synthetic Accounts
	lp := rize.SyntheticAccountListParams{
		CustomerUID:              "uKxmLxUEiSj5h4M3",
		ExternalUID:              "client-generated-id",
		PoolUID:                  "wTSMX1GubP21ev2h",
		Limit:                    100,
		Offset:                   0,
		SyntheticAccountTypeUID:  "q4mdMxMtjXfdbrjn",
		SyntheticAccountCategory: "general",
		Liability:                true,
		Status:                   "Active",
		Sort:                     "name_asc",
	}
	sl, err := rc.SyntheticAccounts.List(&lp)
	if err != nil {
		log.Fatal("Error fetching Synthetic Accounts\n", err)
	}
	output, _ := json.Marshal(sl)
	log.Println("List Synthetic Accounts:", string(output))

	// Create Synthetic Account
	cp := rize.SyntheticAccountCreateParams{
		ExternalUID:             "partner-generated-id",
		Name:                    "New Resource Name",
		PoolUID:                 "kaxHFJnWvJxRJZxq",
		SyntheticAccountTypeUID: "fRMwt6H14ovFUz1s",
		AccountNumber:           "123456789012",
		RoutingNumber:           "123456789",
		ExternalProcessorToken:  "processor-sandbox-96d86f35-ef58-4e4a-826f-4870b5d677f2",
	}
	sc, err := rc.SyntheticAccounts.Create(&cp)
	if err != nil {
		log.Fatal("Error creating Synthetic Account\n", err)
	}
	output, _ = json.Marshal(sc)
	log.Println("Create Synthetic Account:", string(output))

	// Get Synthetic Account
	sg, err := rc.SyntheticAccounts.Get("exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error fetching Synthetic Account\n", err)
	}
	output, _ = json.Marshal(sg)
	log.Println("Get Synthetic Account:", string(output))

	// Update Synthetic Account
	up := rize.SyntheticAccountUpdateParams{
		Name: "New Resource Name",
		Note: "note",
	}
	su, err := rc.SyntheticAccounts.Update("EhrQZJNjCd79LLYq", &up)
	if err != nil {
		log.Fatal("Error updating Synthetic Account\n", err)
	}
	output, _ = json.Marshal(su)
	log.Println("Update Synthetic Account:", string(output))

	// Delete Synthetic Account
	sd, err := rc.SyntheticAccounts.Delete("exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error deleting Synthetic Account\n", err)
	}
	output, _ = json.Marshal(sd)
	log.Println("Delete Synthetic Account:", string(output))

	// List Synthetic Account Types
	stp := rize.SyntheticAccountTypeListParams{
		ProgramUID: "EhrQZJNjCd79LLYq",
		Limit:      100,
		Offset:     0,
	}
	stl, err := rc.SyntheticAccounts.ListAccountTypes(&stp)
	if err != nil {
		log.Fatal("Error fetching Synthetic Account Types\n", err)
	}
	output, _ = json.Marshal(stl)
	log.Println("List Synthetic Account Types:", string(output))

	// Get Synthetic Account Type
	stg, err := rc.SyntheticAccounts.GetAccountType("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Synthetic Account Type\n", err)
	}
	output, _ = json.Marshal(stg)
	log.Println("Get Synthetic Account Type:", string(output))
}