package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FunctionsSuggestions struct {
	Ti [4]float64
	Fi [4]float64
	Ni [4]float64
	Si [4]float64
	Te [4]float64
	Se [4]float64
	Fe [4]float64
	Ne [4]float64
}

const (
	QID_TI_DOM = 11
	QID_TI_AUX = 12
	QID_TI_TER = 13
	QID_TI_INF = 14
	QID_TE_DOM = 111
	QID_TE_AUX = 121
	QID_TE_TER = 131
	QID_TE_INF = 141

	QID_FI_DOM = 21
	QID_FI_AUX = 22
	QID_FI_TER = 23
	QID_FI_INF = 24
	QID_FE_DOM = 211
	QID_FE_AUX = 221
	QID_FE_TER = 231
	QID_FE_INF = 241

	QID_NI_DOM = 31
	QID_NI_AUX = 32
	QID_NI_TER = 33
	QID_NI_INF = 34
	QID_NE_DOM = 321
	QID_NE_AUX = 322
	QID_NE_TER = 323
	QID_NE_INF = 324

	QID_SI_DOM = 41
	QID_SI_AUX = 42
	QID_SI_TER = 43
	QID_SI_INF = 44
	QID_SE_DOM = 411
	QID_SE_AUX = 412
	QID_SE_TER = 413
	QID_SE_INF = 414
)

func ProcessResults(qpath string) {
	// prepare file with questions
	jsonFile, err := os.Open(qpath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened json file")
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	// obsolete
	byteEmpValue, _ := ioutil.ReadAll(jsonFile)
	var questions []Question
	err = json.Unmarshal(byteEmpValue, &questions)
	if err != nil {
		return
	}
	// struct filled
	results := FunctionsSuggestions{}
	for _, question := range questions {
		switch question.TargetId {
		case "ni":
			// If rules disabled
			// results[question.TargetPosition] += question.Answer
			// Else
			// get scope + count of changes. return 0?
			
		}
	}
}
