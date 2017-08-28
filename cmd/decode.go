package main

import (
	"fmt"
	"encoding/json"
	"github.com/pawski/go-xchange/walutomat"
)

func main() {
	byt := []byte(`{"offers":[{"buy":"4.2516","sell":"4.2573","buy_old":"4.2515","sell_old":"4.2573","count_sell":13,"count_buy":50,"pair":"EURPLN","avg":"4.2604","avg_old":"4.2608"},{"buy":"1.1852","sell":"1.1903","buy_old":"1.1851","sell_old":"1.1903","count_sell":78,"count_buy":28,"pair":"EURUSD","avg":"1.1922","avg_old":"1.1922"},{"buy":"1.1360","sell":"1.1400","buy_old":"1.1360","sell_old":"1.1400","count_sell":18,"count_buy":68,"pair":"EURCHF","avg":"1.1407","avg_old":"1.1404"},{"buy":"3.5799","sell":"3.5833","buy_old":"3.5799","sell_old":"3.5835","count_sell":29,"count_buy":51,"pair":"USDPLN","avg":"3.5734","avg_old":"3.5740"},{"buy":"0.9551","sell":"0.9552","buy_old":"0.9551","sell_old":"0.9552","count_sell":57,"count_buy":55,"pair":"USDCHF","avg":"0.9568","avg_old":"0.9567"},{"buy":"4.5831","sell":"4.5999","buy_old":"4.5830","sell_old":"4.5998","count_sell":89,"count_buy":25,"pair":"GBPPLN","avg":"4.6020","avg_old":"4.6038"},{"buy":"1.0787","sell":"1.0811","buy_old":"1.0787","sell_old":"1.0811","count_sell":50,"count_buy":48,"pair":"GBPEUR","avg":"1.0801","avg_old":"1.0803"},{"buy":"1.2856","sell":"1.2880","buy_old":"1.2856","sell_old":"1.2880","count_sell":68,"count_buy":69,"pair":"GBPUSD","avg":"1.2878","avg_old":"1.2879"},{"buy":"1.2204","sell":"1.2319","buy_old":"1.2204","sell_old":"1.2319","count_sell":55,"count_buy":21,"pair":"GBPCHF","avg":"1.2321","avg_old":"1.2322"},{"buy":"3.7421","sell":"3.7437","buy_old":"3.7401","sell_old":"3.7437","count_sell":61,"count_buy":75,"pair":"CHFPLN","avg":"3.7352","avg_old":"3.7357"}]}`)

	var offersResponse walutomat.OffersResponse

	if err := json.Unmarshal(byt, &offersResponse); err != nil {
		panic(err)
	}
	fmt.Println(offersResponse)
}
