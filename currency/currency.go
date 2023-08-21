// curl 'https://coinlayer.com/coin_api.php' \
//   -H 'authority: coinlayer.com' \
//   -H 'accept: application/json, text/javascript, */*; q=0.01' \
//   -H 'accept-language: en-US,en;q=0.9' \
//   -H 'dnt: 1' \
//   -H 'referer: https://coinlayer.com/' \
//   -H 'sec-ch-ua: "Not)A;Brand";v="24", "Chromium";v="116"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
//   -H 'x-requested-with: XMLHttpRequest' \
//   --compressed

// RESPONSE
// {
// 	   "timestamp": 1692656441,
//     "long_date": "2023-08-21 22:20:41",
//     "current_BTC": 26142.189625,
//     "current_BTC_high": 29022.99,
//     "current_BTC_low": 25807.3,
//     "current_BTC_vol": 41405544.574615,
//     "current_BTC_cap": 508600534189.09424,
//     "current_BCH": 187.7025,
//     "current_EOS": 0.583315,
//     "current_ETH": 1671.175857,
//     "current_LTC": 66.944014,
//     "current_XMR": 149.4638,
//     "current_XRP": 0.524524,
// }
package currency

import (
	"encoding/json"
	"io"
	"net/http"
)

type Currency struct {
	Timestamp      int64   `json:"timestamp"`
	LongDate       string  `json:"long_date"`
	CurrentBTC     float64 `json:"current_BTC"`
	CurrentBTCHigh float64 `json:"current_BTC_high"`
	CurrentBTCLow  float64 `json:"current_BTC_low"`
	CurrentBTCVol  float64 `json:"current_BTC_vol"`
	CurrentBTCCap  float64 `json:"current_BTC_cap"`
	CurrentBCH     float64 `json:"current_BCH"`
	CurrentEOS     float64 `json:"current_EOS"`
	CurrentETH     float64 `json:"current_ETH"`
	CurrentLTC     float64 `json:"current_LTC"`
	CurrentXMR     float64 `json:"current_XMR"`
	CurrentXRP     float64 `json:"current_XRP"`
}

func GetCurrency() (Currency, error) {
	var c Currency

	req, err := http.NewRequest("GET", "https://coinlayer.com/coin_api.php", nil)
	if err != nil {
		return c, err
	}

	req.Header.Add("authority", "coinlayer.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	// req.Header.Add("accept-language", "en-US,en;q=0.9")
	// req.Header.Add("dnt", "1")
	// req.Header.Add("referer", "https://coinlayer.com/")
	// req.Header.Add("sec-ch-ua", "\"Not)A;Brand\";v=\"24\", \"Chromium\";v=\"116\"")
	// req.Header.Add("sec-ch-ua-mobile", "?0")
	// req.Header.Add("sec-ch-ua-platform", "\"macOS\"")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}

	return c, nil

}
