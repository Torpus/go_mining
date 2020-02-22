package main

import (
	"encoding/json"
    "fmt"
    "github.com/jaypipes/ghw"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)

var gpus = [22]string{"380","fury","470","480","570","580","vega56","vega64","5700xt","vii","1050ti","1060","1070ti","1070","1080ti","1080","1660ti","1660","2060","2070","2080ti","2080"}
var url = "https://whattomine.com/coins.json?adapt_q_"

type Detail struct {
	Tag string
	Profitability float64
}

func constainsString(str string) int {
    for index, element := range gpus {
        if (strings.Contains(str, element)) {
            return index
        }
    }
    return -1
}

func (n *Detail) UnmarshalJSON(buf []byte) error {
	var tmp interface{}
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	coins := tmp.(map[string]interface{})
	coinsMap := coins["coins"]

	for _, coin := range coinsMap.(map[string]interface{}) {
		c := coin.(map[string]interface{})
		if n.Profitability <= c["profitability"].(float64) {
			n.Tag = c["tag"].(string)
			n.Profitability = c["profitability"].(float64)
		}
	}

	return nil
}

func getGpu() string {
    re := regexp.MustCompile(`\[(.*)\]`)
    gpu, err := ghw.GPU()
    if err != nil {
        fmt.Printf("error getting GPU info: %v", err)
    }

    for _, card := range gpu.GraphicsCards {
        name := re.FindString(card.DeviceInfo.Product.Name)
        name = strings.Trim(name, "[")
        name = strings.Trim(name, "]")
        name = strings.ToLower(strings.Join(strings.Fields(name), ""))
        var index int = constainsString(name)
        if (index != -1) {
            return gpus[index]
        }
    }
    return ""
}

func callWhatToMine(url string) Detail {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    var r Detail
    if err := json.Unmarshal([]byte(body), &r); err != nil {
		fmt.Println(err)
    }
    return r
}

func main() {
    url := url + getGpu()
	fmt.Println(callWhatToMine(url))
}