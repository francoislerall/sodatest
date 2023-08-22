package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/msrocka/soda"
)

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
	unitGroups, _ := client.GetUnitGroups()
	
	for _, indicator := range profile.Indicators {
		fmt.Println(indicator.Name)
		for _, unitGroup := range unitGroups.UnitGroups {
			if indicator.UnitGroupUUID == unitGroup.UUID {
				if unitGroup.Name.Get("en") == "" {
					fmt.Println(" Unit group found: " + unitGroup.Name.Get("en"))
				} else {
					fmt.Println(" Unit group found: " + unitGroup.Name.Get("de"))
				}
			}
		}
	}
}
