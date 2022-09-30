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

	// List Adjustments
	alp := rize.AdjustmentListParams{
		CustomerUID:            "uKxmLxUEiSj5h4M3",
		AdjustmentTypeUID:      "2Ej2tsFbQT3S1HYd",
		ExternalUID:            "PT3sH7oxxQPwchrf",
		USDAdjustmentAmountMax: 10,
		USDAdjustmentAmountMin: 5,
		Sort:                   "adjustment_type_name_asc",
	}
	al, err := rc.Adjustments.List(&alp)
	if err != nil {
		log.Fatal("Error fetching Adjustments\n", err)
	}
	output, _ := json.Marshal(al)
	log.Println("List Adjustments:", string(output))

	// Create Adjustment
	acp := rize.AdjustmentCreateParams{
		ExternalUID:         "partner-generated-id",
		CustomerUID:         "kaxHFJnWvJxRJZxq",
		USDAdjustmentAmount: 2.43,
		AdjustmentTypeUID:   "KM2eKbR98t4tdAyZ",
	}
	ac, err := rc.Adjustments.Create(&acp)
	if err != nil {
		log.Fatal("Error creating Adjustment\n", err)
	}
	output, _ = json.Marshal(ac)
	log.Println("Create Adjustment:", string(output))

	// Get Adjustment
	ag, err := rc.Adjustments.Get("exMDShw6yM3NHLYV")
	if err != nil {
		log.Fatal("Error fetching Adjustment\n", err)
	}
	output, _ = json.Marshal(ag)
	log.Println("Get Adjustment:", string(output))

	// List Adjustment Types
	atp := rize.AdjustmentTypeListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ProgramUID:  "DbxJUHVuqt3C7hGK",
	}
	lat, err := rc.Adjustments.ListAdjustmentTypes(&atp)
	if err != nil {
		log.Fatal("Error fetching Adjustment Types\n", err)
	}
	output, _ = json.Marshal(lat)
	log.Println("List Adjustment Types:", string(output))

	// Get Adjustment Type
	gat, err := rc.Adjustments.GetAdjustmentType("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Adjustment Type\n", err)
	}
	output, _ = json.Marshal(gat)
	log.Println("Get Adjustment Type:", string(output))
}