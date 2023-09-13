package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/msrocka/soda"
)

func CheckIndicatorsUnitGroup(client *soda.Client, profile Profile) {
	unitGroups, _ := client.GetUnitGroups()

	var falsyIndicators []string
	for _, indicator := range profile.Indicators {
		exists := false
		for _, unitGroup := range unitGroups.UnitGroups {
			if indicator.UnitGroupUUID == unitGroup.UUID {
				exists = true
				break
			}
		}
		if (!exists) {
			falsyIndicators = append(falsyIndicators, indicator.Name)
		}
	}

	if (len(falsyIndicators) == 0) {
		fmt.Printf("All the unit groups of the profile indicators has been found (%d).", len(profile.Indicators))
	} else {
		fmt.Printf("The following indicator does not have a unit group in the DB (%d):\n - " + strings.Join(falsyIndicators,"\n - "), len(falsyIndicators))
	}
	return
}

func CheckIndicatorsUUID(client *soda.Client, profile Profile) {
	methods, _ := client.GetMethods()

	var falsyIndicators []string
	for _, indicator := range profile.Indicators {
		exists := false
		for _, method := range methods.Methods {
			if indicator.UUID == method.UUID {
				exists = true
				break
			}
		}
		if (!exists) {
			falsyIndicators = append(falsyIndicators, indicator.Name)
		}
	}

	if (len(falsyIndicators) == 0) {
		fmt.Printf("\nAll the indicators has been found (%d)", len(profile.Indicators))
	} else {
		fmt.Printf("\nThe following indicators could not be found (%d):\n - " + strings.Join(falsyIndicators,"\n - "), len(falsyIndicators))
	}
	return
}

func main()  {
	content, err := os.ReadFile("EN_15804_A2.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

	// Unmarshall the data into profile
    var profile Profile
    err = json.Unmarshal(content, &profile)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	fmt.Println("profile: " + profile.Id)

	client := soda.NewClient("https://oekobaudat.de/OEKOBAU.DAT/resource")
	
	CheckIndicatorsUnitGroup(client, profile)
	CheckIndicatorsUUID(client, profile)
}
