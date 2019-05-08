// Go in action
// @jeffotoni
// 2019-05-06

package main

import (
	"encoding/csv"
	"io"
	//"errors"
	"fmt"
	//"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Estados struct {
	Nome     string
	Extensao float64
}

func main() {
	fmt.Println(os10maioresEstadosDoBrasil())
}

func openCsv(src string) (e []Estados, err error) {

	file, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return e, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
			// handle the error...
			// break? continue? neither?
		}

		extensao, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			// panic(err)
			// log.Fatal(err)
			fmt.Println(err)
			continue
		}

		e = append(e, Estados{Nome: strings.ToLower(record[0]), Extensao: extensao})
	}

	fmt.Println(e)
	return e, err
}

func sortStructDesc(e []Estados) {
	sort.Slice(e, func(i, j int) bool {
		return e[i].Extensao > e[j].Extensao
	})
}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var data []string
	e, err := openCsv("./dados.csv")
	if err != nil {
		return data, err
	}
	sortStructDesc(e)
	for i := 0; i < 10; i++ {
		data = append(data, e[i].Nome)
	}
	return data, nil
}
