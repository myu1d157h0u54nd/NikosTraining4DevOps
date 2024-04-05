package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

// ******************************************************************************************************************
// API call
// ******************************************************************************************************************

func apiREST_ValidatorsMoniker() ([]byte, error) {

	// Request: Full node endpoints
	apiRestURL := "https://api-dydx.cosmos-spaces.cloud:443"
	apiRestModule := "/cosmos/staking/v1beta1/validators"

	apiRestFULL := apiRestURL + apiRestModule
	resp, err := http.Get(apiRestFULL)
	if err != nil {
		return nil, fmt.Errorf("Error in HTTP request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading HTTP response: %v", err)
	}

	return bodyBytes, nil
}

func apiRPC_ValidatorsVP() ([]byte, error) {

	// Request: Full node endpoints
	// https://rpc-dydx.cosmos-spaces.cloud/validators?height=11473919&page=1&per_page=100
	apiRestURL := "https://rpc-dydx.cosmos-spaces.cloud:443"
	apiRestModule := "/validators?height=11473919&page=1&per_page=100"

	apiRestFULL := apiRestURL + apiRestModule
	resp, err := http.Get(apiRestFULL)
	if err != nil {
		return nil, fmt.Errorf("Error in HTTP request: %v", err)
	}

	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading HTTP response: %v", err)
	}

	return rawBody, nil
}

// ******************************************************************************************************************
// helpers
// ******************************************************************************************************************

func dataToFile(data string, filename string) error {
	// create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// string -> byte
	dataBytes := []byte(data)

	// dataBytes -> filename
	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}

	fmt.Printf("data -> filename: %s\n", filename)
	return nil
}

// dataStructureValidatorsMoniker_
// https://transform.tools/json-to-go
type dataStructureValidatorsMoniker_ struct {
	Validators []struct {
		// OperatorAddress string `json:"operator_address"`
		ConsensusPubkey struct {
			// Type string `json:"@type"`
			Key string `json:"key"`
		} `json:"consensus_pubkey"`
		// Jailed          bool   `json:"jailed"`
		// Status          string `json:"status"`
		// Tokens          string `json:"tokens"`
		// DelegatorShares string `json:"delegator_shares"`
		Description struct {
			Moniker string `json:"moniker"`
			// Identity string `json:"identity"`
			// Website         string `json:"website"`
			// SecurityContact string `json:"security_contact"`
			// Details         string `json:"details"`
		} `json:"description"`
		// UnbondingHeight string    `json:"unbonding_height"`
		// UnbondingTime   time.Time `json:"unbonding_time"`
		Commission struct {
			CommissionRates struct {
				// Rate          string `json:"rate"`
				// MaxRate       string `json:"max_rate"`
				// MaxChangeRate string `json:"max_change_rate"`
			} `json:"commission_rates"`
			// UpdateTime time.Time `json:"update_time"`
		} `json:"commission"`
		// MinSelfDelegation       string   `json:"min_self_delegation"`
		// UnbondingOnHoldRefCount string   `json:"unbonding_on_hold_ref_count"`
		// UnbondingIds            []string `json:"unbonding_ids"`
	} `json:"validators"`
	Pagination struct {
		// NextKey string `json:"next_key"`
		// Total   string `json:"total"`
	} `json:"pagination"`
}

// dataStructureValidatorsVP_
// https://transform.tools/json-to-go
type dataStructureValidatorsVP_ struct {
	// Jsonrpc string `json:"jsonrpc"`
	// ID      int    `json:"id"`
	Result struct {
		// BlockHeight string `json:"block_height"`
		Validators []struct {
			// Address string `json:"address"`
			PubKey struct {
				// Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
			// ProposerPriority string `json:"proposer_priority"`
		} `json:"validators"`
		// Count string `json:"count"`
		// Total string `json:"total"`
	} `json:"result"`
}

// dataStructureCSVdata_
type dataStructureCSVdata_ struct {
	Moniker     string
	VotingPower int
	PubKey      string
}

// rawByte -> JSONpretty
func rawBodytoprettyJSONString(rawBody []byte) (string, error) {

	//
	// Indent appends to dst an indented form of the JSON-encoded src
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, rawBody, "", "\t")
	if err != nil {
		return "", err
	}

	// JSONpretty -> string
	return prettyJSON.String(), nil

}

func createCSVBuffer(ds []dataStructureCSVdata_) string {
	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)
	defer writer.Flush()

	writer.Write([]string{"Moniker", "VotingPower", "PubKey"})
	for _, validator := range ds {
		writer.Write([]string{validator.Moniker, strconv.Itoa(validator.VotingPower), validator.PubKey})
	}

	writer.Flush()
	return csvBuffer.String()
}

func main() {

	// ******************************************************************************************************************
	// ValidatorsVP:
	// ******************************************************************************************************************

	// HTTP Request + Response:
	ValidatorsVPrawBody, err := apiRPC_ValidatorsVP()
	if err != nil {
		log.Fatal(err)
	}

	ValidatorsVPprettyJSONString, err := rawBodytoprettyJSONString(ValidatorsVPrawBody)
	if err != nil {
		log.Fatal("Error traing convert rawBody to JSON:", err)
	}

	// output
	dataToFile(ValidatorsVPprettyJSONString, "ValidatorsVPprettyJSONString.json")

	// ******************************************************************************************************************
	// ValidatorsMoniker:
	// ******************************************************************************************************************

	// HTTP Request + Response:
	ValidatorsMonikerrawBody, err := apiREST_ValidatorsMoniker()
	if err != nil {
		log.Fatal(err)
	}

	ValidatorsMonikerprettyJSONString, err := rawBodytoprettyJSONString(ValidatorsMonikerrawBody)
	if err != nil {
		log.Fatal("Error traing convert rawBody to JSON:", err)
	}

	// output
	dataToFile(ValidatorsMonikerprettyJSONString, "ValidatorsMonikerprettyJSONString.json")

	// ******************************************************************************************************************
	// dataStructure: dataStructureValidatorsMoniker_ + dataStructureValidatorsVP_
	// ******************************************************************************************************************

	fullDataValidatorsMoniker := dataStructureValidatorsMoniker_{}
	json.Unmarshal(ValidatorsMonikerrawBody, &fullDataValidatorsMoniker)

	fullDataValidatorsVP := dataStructureValidatorsVP_{}
	json.Unmarshal(ValidatorsVPrawBody, &fullDataValidatorsVP)

	// ******************************************************************************************************************
	// Iteration:
	// ******************************************************************************************************************

	// Iteration to get: Moniker
	var ds4outputCSV []dataStructureCSVdata_
	for _, validator_Moniker := range fullDataValidatorsMoniker.Validators {
		// Iteration to get: VotingPower
		for _, validatorVP := range fullDataValidatorsVP.Result.Validators {
			if validatorVP.PubKey.Value == validator_Moniker.ConsensusPubkey.Key {
				votingPower, _ := strconv.Atoi(validatorVP.VotingPower)
				ds4outputCSV = append(ds4outputCSV, dataStructureCSVdata_{
					Moniker:     validator_Moniker.Description.Moniker,
					VotingPower: votingPower,
					PubKey:      validator_Moniker.ConsensusPubkey.Key,
				})
			}
		}
	}

	// Sort from largest to smallest
	sort.Slice(ds4outputCSV, func(i, j int) bool {
		return ds4outputCSV[i].VotingPower > ds4outputCSV[j].VotingPower
	})

	// top10
	var ds4outputCSVtop10 []dataStructureCSVdata_
	for i := 0; i < 10 && i < len(ds4outputCSV); i++ {
		ds4outputCSVtop10 = append(ds4outputCSVtop10, ds4outputCSV[i])
	}

	// output top10
	csvtop10string := createCSVBuffer(ds4outputCSVtop10)
	dataToFile(csvtop10string, "ValidatorsMoniker_and_VotingPower_top10.csv")

	// output full
	csvfullstring := createCSVBuffer(ds4outputCSV)
	dataToFile(csvfullstring, "ValidatorsMoniker_and_VotingPower_FULL.csv")

}
