package main

import (
	"context"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"gitee.com/quant1x/gotdx/quotes"
)

// App struct
type App struct {
	ctx context.Context
	api *quotes.StdApi
}

// NewApp creates a new App application struct
func NewApp() *App {
	api, err := quotes.NewStdApi()
	if err != nil {
		fmt.Print("err: ", err)
		return &App{}
	}
	return &App{api: api}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {

}

// beforeClose is called when the application is about to quit,
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

type Kline struct {
	Time   string  `json:"time"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Vol    float64 `json:"vol"`
	Amount float64 `json:"amount"`
}

type Index struct {
	Name      string  `json:"name"`
	Date      string  `json:"date"`
	Close     float64 `json:"close"`
	UpCount   uint16  `json:"upCount"`
	DownCount uint16  `json:"downCount"`
}

type IndexData struct {
	Name string
	Index
}

func (a *App) GetKline(code string, count uint16) []Kline {
	data, err := a.api.GetKLine(code, proto.KLINE_TYPE_RI_K, 0, count)
	if err != nil {
		return []Kline{}
	}

	var kData []Kline
	for _, v := range data.List {
		k := Kline{
			Time:   v.DateTime,
			Open:   v.Open,
			High:   v.High,
			Low:    v.Low,
			Close:  v.Close,
			Vol:    v.Vol,
			Amount: v.Amount,
		}
		kData = append(kData, k)
	}
	return kData
}

func (a *App) GetupDown() []IndexData {
	sh, _ := a.api.GetIndexBars("999999", proto.KLINE_TYPE_RI_K, 0, 1)
	sz, _ := a.api.GetIndexBars("399001", proto.KLINE_TYPE_RI_K, 0, 1)
	cyb, _ := a.api.GetIndexBars("399006", proto.KLINE_TYPE_RI_K, 0, 1)

	var data []IndexData

	shData := IndexData{
		Name: "sh",
		Index: Index{
			Name:      "上证指数",
			Date:      sh.List[0].DateTime,
			Close:     sh.List[0].Close,
			UpCount:   sh.List[0].UpCount,
			DownCount: sh.List[0].DownCount,
		},
	}

	szData := IndexData{
		Name: "sz",
		Index: Index{
			Name:      "深证指数",
			Date:      sz.List[0].DateTime,
			Close:     sz.List[0].Close,
			UpCount:   sz.List[0].UpCount,
			DownCount: sz.List[0].DownCount,
		},
	}

	cybData := IndexData{
		Name: "cyb",
		Index: Index{
			Name:      "创业板指",
			Date:      cyb.List[0].DateTime,
			Close:     cyb.List[0].Close,
			UpCount:   cyb.List[0].UpCount,
			DownCount: cyb.List[0].DownCount,
		},
	}

	data = append(data, shData)
	data = append(data, szData)
	data = append(data, cybData)
	return data
}
