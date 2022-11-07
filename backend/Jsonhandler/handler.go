package jsonhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type JsonData struct {
	Id    int    `json: "id"`
	Date  string `json: "date"`
	Input string `json: "Input"`
	Time  string `json: "time"`
	Rows  int    `json: "rows"`
}

type QualificationFormReturn struct {
	Title string
	State bool
}

type StatusJson struct {
	Date   string `json: "date"`
	Status string `json: "status"`
}

type Form struct {
	Status string
}

var destinationPath, _ = os.UserHomeDir()

func SaveAsJson(id int, date, input, time string, rows int) {
	m := []JsonData{}
	obj := JsonData{Id: id, Date: date, Input: input, Time: time, Rows: rows}
	m = append(m, obj)
	data, err := json.MarshalIndent(m, " ", " ")
	if err != nil {
		fmt.Println("Something went wrong saving json")
	}

	if _, err := os.ReadDir(destinationPath + "/Documents/TextFieldOutput/"); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+"/Documents/TextFieldOutput/", 0755)
	}

	if _, err := os.ReadFile(destinationPath + "/Documents/TextFieldOutput/" + date + ".json"); err != nil {
		fmt.Println("Creating File...")
		os.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", data, 0645)
	} else {
		local := []JsonData{}
		output, _ := ReadFromJson(date)
		allowed := true
		for i, v := range output {
			if v.Id == obj.Id {
				if obj.Time == "00:00" {
					output[i].Input = obj.Input
					output[i].Rows = obj.Rows
				} else if obj.Input == "" {
					output[i].Time = obj.Time
				} else if obj.Time != "00:00" && obj.Input != "" {
					output[i].Input = obj.Input
					output[i].Time = obj.Time
				}
				allowed = false
				break
			}
		}
		local = append(local, output[:]...)
		if allowed {
			local = append(local, obj)
		}
		// fmt.Println(local)
		appendData, _ := json.MarshalIndent(local, " ", " ")

		os.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", appendData, 0755)
	}

}

func ReadFromJson(date string) ([]JsonData, error) {
	file, err := os.ReadFile(destinationPath + "/Documents/TextFieldOutput/" + date + ".json")
	if err != nil {
		return nil, err
	}

	output := []JsonData{}

	_ = json.Unmarshal([]byte(file), &output)

	return output, nil
}

func RemoveFromJson(date string, index int) {
	output, err := ReadFromJson(date)
	if err != nil {
		fmt.Println(err)
	}
	var local []JsonData
	if len(output) >= index {
		output = append(output[:index], output[index+1:]...)
	}
	local = append(local, output[:]...)
	//fmt.Println(local)
	for i, _ := range local {
		local[i].Id = i
	}
	rawData, _ := json.MarshalIndent(local, " ", " ")

	os.WriteFile(destinationPath+"/Documents/TextFieldOutput/"+date+".json", rawData, 0755)

}

func WriteStatusJson(date, status string) {
	if _, err := os.ReadDir(destinationPath + "/Documents/AbsenceStatus/"); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+"/Documents/AbsenceStatus/", 0755)
	}
	rec := Form{Status: status}
	data, _ := json.MarshalIndent(rec, " ", " ")

	os.WriteFile(destinationPath+"/Documents/AbsenceStatus/"+date+".json", data, 0755)
}

func ReadStatusJson(date string) (StatusJson, error) {
	data, err := os.ReadFile(destinationPath + "/Documents/AbsenceStatus/" + date + ".json")
	if err != nil {
		return StatusJson{}, err
	}
	output := StatusJson{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		log.Fatalf("error unmarshal json: %v", err)
	}

	return output, nil
}

func WriteQualificationsJson(qualis []interface{}, date, endpoint string) {
	if _, err := os.ReadDir(destinationPath + endpoint); err != nil {
		fmt.Println("Creating Directory...")
		os.Mkdir(destinationPath+endpoint, 0755)
	}

	data, _ := json.MarshalIndent(qualis, " ", " ")
	os.WriteFile(destinationPath+endpoint+date+".json", data, 0755)
}

func ReadQualificationsJson(date, endpoint string) []QualificationFormReturn {
	data, _ := os.ReadFile(destinationPath + endpoint + date + ".json")

	var qualis []QualificationFormReturn
	_ = json.Unmarshal(data, &qualis)

	return qualis
}

func ReadJsonMonth(month string) ([]JsonData, error) {
	var dataStore []JsonData
	files, _ := os.ReadDir(destinationPath + "/Documents/TextFieldOutput/")
	for _, val := range files {
		character := strings.Split(val.Name(), ".")
		if month == character[1] {
			data, err := ReadFromJson(strings.ReplaceAll(val.Name(), ".json", ""))
			if err != nil {
				return nil, err
			}
			dataStore = append(dataStore, data...)
		}
	}
	return dataStore, nil
}
