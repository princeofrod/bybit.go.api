package main

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func main() {
	PlaceTrade()
}

func PlaceTrade() {
	client := bybit.NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"category": "linear", "settleCoin": "USDT", "limit": 10}
	orderResult, err := client.NewPositionService(params).GetPositionList(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(orderResult))
}
