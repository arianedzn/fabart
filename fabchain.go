package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}
type Art struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	ArtistName string `json:"artistname"`
	Medium     string `json:"medium"`
	Status     string `json:"status"`
	Size       string `json:"size"`
	Value      string `json:"value"`
	Location   string `json:"location"`
	Owner      string `json:"owner"`
}
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Art
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	arts := []Art{
		{Name: "LED Poi", Image: "https://www.homeofpoi.com/img/Ignis-pixel-led-poi.jpg", ArtistName: "John Smith",
			Medium: "Canvas, giclee", Status: "available", Size: "40x65", Value: "3750", Location: "Manila, Philippines", Owner: "John Smith"},
		{Name: "Kissing The Stars", Image: "https://i.pinimg.com/736x/ff/b5/11/ffb51152019f00459dc573d2acd9738b.jpg", ArtistName: "John Smith",
			Medium: "oil on b mounted panel", Status: "sold", Size: "48x72", Value: "7200", Location: "Manila, Philippines", Owner: "John Smith"},
		{Name: "Abstact Neon Shapes", Image: "https://i.pinimg.com/originals/16/bb/d8/16bbd8602c69a97c65ed902a2895f6c2.jpg", ArtistName: "John Smith",
			Medium: "Digital Painting Aluminum Chromaluxe Print", Status: "available", Size: "24x30", Value: "5000", Location: "Manila, Philippines", Owner: "John Smith"},
	}

	for i, art := range arts {
		artAsBytes, _ := json.Marshal(art)
		err := ctx.GetStub().PutState("ART"+strconv.Itoa(i), artAsBytes)

		if err != nil {
			return fmt.Errorf("failed to put to world state. %s", err.Error())
		}
	}
	return nil
}
func (s *SmartContract) QueryAllArt(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		art := new(Art)
		_ = json.Unmarshal(queryResponse.Value, art)

		quertResult := QueryResult{Key: queryResponse.Key, Record: art}
		results = append(results, quertResult)
	}

	return results, nil
}

type Profile struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Award    string `json:"award"`
	Bio      string `json:"bio"`
}

type QueryProfileResult struct {
	Key    string `json:"Key"`
	Record *Profile
}

func (s *SmartContract) InitLedgerProfile(ctx contractapi.TransactionContextInterface) error {
	profiles := []Profile{
		{Name: "Keena Villarin", Image: "https://blog.kakaocdn.net/dn/bdU3kw/btqJDXNqFJa/mhTRcJ9H83KIb6rRJBt6k1/img.jpg", Location: "Laguna, Philippines",
			Phone: "09123456789", Email: "keen23@email.com", Award: "Philippine Art Awards fetes winner", Bio: "A Found Technical Master..."},
		{Name: "Ariane Dizon", Image: "https://cdn.bhdw.net/im/twice-s-dahyun-in-alcohol-free-mv-shoot-2021-wallpaper-72674_w635.jpg", Location: "Apalit, Pampanga, Philippines",
			Phone: "09212345000", Email: "ariane21@email.com", Award: "N/A", Bio: "Ariane's current work..."},
		{Name: "Dhayo Villanueva", Image: "https://cdn.seat42f.com/wp-content/uploads/2020/11/30120457/lee-minho.jpg", Location: "Victoria, Oriental Mindoro, Philippines",
			Phone: "09345678901", Email: "dhayo01@email.cexitom", Award: "N/A", Bio: "He is a film and video ..."},
		{Name: "Jennie Kim", Image: "https://filmdaily.co/wp-content/uploads/2020/03/Blackpink-Jennie-lede-1536x960.jpg", Location: "Gangnam, South Korea",
			Phone: "01434452245", Email: "kim.jennie@email.com", Award: "South Korea Art Awards fetes winner", Bio: "A Found Technical Master..."},
	}

	for i, profile := range profiles {
		profileAsBytes, _ := json.Marshal(profile)
		err := ctx.GetStub().PutState("PROFILE"+strconv.Itoa(i), profileAsBytes)

		if err != nil {
			return fmt.Errorf("failed to put to world state. %s", err.Error())
		}
	}
	return nil
}
func (s *SmartContract) QueryAllProfile(ctx contractapi.TransactionContextInterface) ([]QueryProfileResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryProfileResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		profile := new(Profile)
		_ = json.Unmarshal(queryResponse.Value, profile)

		quertProfileResult := QueryProfileResult{Key: queryResponse.Key, Record: profile}
		results = append(results, quertProfileResult)
	}

	return results, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabart chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabart chaincode: %s", err.Error())
	}
}
