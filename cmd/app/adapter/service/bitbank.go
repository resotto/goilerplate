package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/resotto/goilerplate/cmd/app/domain/valueobject"
	"github.com/spf13/viper"
)

type bitbankticker struct {
	Sell      string
	Buy       string
	High      string
	Low       string
	Last      string
	Vol       string
	Timestamp int
}

// bitbanktickerresponse is response of bitbank ticker api
type bitbanktickerresponse struct {
	Success int
	Data    bitbankticker
}

type bitbankohlc struct {
	Type  string
	Ohlcv [][]interface{}
}

type bitbankcandlestick struct {
	Candlestick []bitbankohlc
}

// bitbankohlcresponse is response of bitbank ohlc api
type bitbankohlcresponse struct {
	Success int
	Data    bitbankcandlestick
}

// Bitbank is an bitcoin exchange
type Bitbank struct{}

// Ticker gets ticker via bitbank public api
func (b Bitbank) Ticker(p valueobject.Pair) valueobject.Ticker {
	pair := b.convertPair(p)
	host := viper.Get("BITBANK_HOST")
	url := fmt.Sprintf("%v/%v/ticker", host, pair)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data bitbanktickerresponse
	json.Unmarshal(bytes, &data)
	return valueobject.Ticker{
		Sell:      data.Data.Sell,
		Buy:       data.Data.Buy,
		High:      data.Data.High,
		Low:       data.Data.Low,
		Last:      data.Data.Last,
		Vol:       data.Data.Vol,
		Timestamp: strconv.Itoa(data.Data.Timestamp / 1000),
	}
}

// Ohlc gets open, high, low, and close via bitbank public api
// NOTICE: This works from 0AM (UTC) due to its api constraints
func (b Bitbank) Ohlc(p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick {
	pair := b.convertPair(p)
	timeunit := b.convertTimeunit(t)
	yyyy := b.yyyy(t)
	host := viper.Get("BITBANK_HOST")
	url := fmt.Sprintf("%v/%v/candlestick/%v/%v", host, pair, timeunit, yyyy)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data bitbankohlcresponse
	json.Unmarshal(bytes, &data)
	return convertToCandlestick(data)
}

func convertToCandlestick(res bitbankohlcresponse) []valueobject.CandleStick {
	ohlcs := res.Data.Candlestick[0].Ohlcv
	cs := make([]valueobject.CandleStick, 0)
	for _, v := range ohlcs {
		timestamp := strconv.FormatInt(int64(v[5].(float64)/1000), 10)
		cs = append(cs, valueobject.CandleStick{
			Open:      v[0].(string),
			High:      v[1].(string),
			Low:       v[2].(string),
			Close:     v[3].(string),
			Volume:    v[4].(string),
			Timestamp: timestamp,
		})
	}
	return cs
}

// OneMin is a timeunit
const (
	OneMin valueobject.Timeunit = iota
	FiveMin
	FifteenMin
	ThirtyMin
	OneHour
	FourHour
	EightHour
	TweleveHour
	OneDay
	OneWeek
)

func (b Bitbank) convertTimeunit(t valueobject.Timeunit) string {
	return b.timeunitNames()[t]
}

func (b Bitbank) timeunitNames() []string {
	return []string{
		"1min",
		"5min",
		"15min",
		"30min",
		"1hour",
		"4hour",
		"8hour",
		"12hour",
		"1day",
		"1week",
	}
}

func (b Bitbank) convertPair(p valueobject.Pair) string {
	return b.pairNames()[p]
}

func (b Bitbank) pairNames() []string {
	return []string{
		"btc_jpy",
	}
}

func (b Bitbank) yyyy(t valueobject.Timeunit) string {
	day := time.Now()
	var layout string
	switch t {
	case OneMin, FiveMin, FifteenMin, ThirtyMin, OneHour:
		layout = "20060102"
	default:
		layout = "2006"
	}
	return day.Format(layout)
}
