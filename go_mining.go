package main

import (
	"encoding/json"
	"fmt"
)

const text =
`{
    "coins": {
        "Nicehash-Ethash": {
            "id": 15,
            "tag": "NICEHASH",
            "algorithm": "Ethash",
            "block_time": 1,
            "block_reward": 1,
            "block_reward24": 1,
            "last_block": 0,
            "difficulty": 1,
            "difficulty24": 1,
            "nethash": 6466835193590,
            "exchange_rate": 1.99893663379386,
            "exchange_rate24": 2.14490923336308,
            "exchange_rate_vol": 12.553093441446578,
            "exchange_rate_curr": "BTC",
            "market_cap": "$0.00",
            "estimated_rewards": "0.00018",
            "estimated_rewards24": "0.00019",
            "btc_revenue": "0.00017691",
            "btc_revenue24": "0.00018982",
            "profitability": 89,
            "profitability24": 101,
            "lagging": false,
            "timestamp": 1582085706
        },
        "Ethereum": {
            "id": 151,
            "tag": "ETH",
            "algorithm": "Ethash",
            "block_time": "13.2861",
            "block_reward": 2.0,
            "block_reward24": 2.0,
            "last_block": 9511398,
            "difficulty": 2.27050237419557e+15,
            "difficulty24": 2.24700837563304e+15,
            "nethash": 170893066753642,
            "exchange_rate": 0.02755656,
            "exchange_rate24": 0.0276427301004304,
            "exchange_rate_vol": 101149.608339618,
            "exchange_rate_curr": "BTC",
            "market_cap": "$30,559,008,668.45",
            "estimated_rewards": "0.00674",
            "estimated_rewards24": "0.00681",
            "btc_revenue": "0.00018561",
            "btc_revenue24": "0.00018755",
            "profitability": 100,
            "profitability24": 100,
            "lagging": false,
            "timestamp": 1582085653
        },
        "EthereumClassic": {
            "id": 162,
            "tag": "ETC",
            "algorithm": "Ethash",
            "block_time": "13.0387",
            "block_reward": 3.88,
            "block_reward24": 3.88000000000007,
            "last_block": 9821831,
            "difficulty": 163334790581410.0,
            "difficulty24": 160948322246079.0,
            "nethash": 12526922974024,
            "exchange_rate": 0.0009565,
            "exchange_rate24": 0.000969631707317073,
            "exchange_rate_vol": 950.74656188,
            "exchange_rate_curr": "BTC",
            "market_cap": "$1,124,059,247.59",
            "estimated_rewards": "0.18164",
            "estimated_rewards24": "0.18433",
            "btc_revenue": "0.00017374",
            "btc_revenue24": "0.00017631",
            "profitability": 84,
            "profitability24": 94,
            "lagging": false,
            "timestamp": 1582085651
			}
		}
	}
`

type Detail struct {
	Tag string
	Profitability float64
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
		var tmpN Detail
		tmpN.Tag = c["tag"].(string)
		tmpN.Profitability = c["profitability"].(float64)
		if n.Profitability <= tmpN.Profitability {
			n.Tag = tmpN.Tag
			n.Profitability = tmpN.Profitability
		}
	}

	return nil
}

func main() {
	var r Detail
	if err := json.Unmarshal([]byte(text), &r); err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}