package request

import (
	"main/internal/core/model/quant"
	"main/internal/pkg/objconv"
	"time"
)

type QuantOptU struct {
	MainSectors         []string         `json:"main_sectors,omitempty" example:"IT,소재"`
	NetRevenue          quant.IntPair    `json:"net_revenue,omitempty"`
	NetRevenueRate      quant.DoublePair `json:"net_revenue_rate,omitempty"`
	NetProfit           quant.IntPair    `json:"net_profit,omitempty"`
	NetProfitRate       quant.DoublePair `json:"net_profit_rate,omitempty"`
	DERatio             quant.DoublePair `json:"de_ratio,omitempty"`
	Per                 quant.DoublePair `json:"per,omitempty"`
	Psr                 quant.DoublePair `json:"psr,omitempty"`
	Pbr                 quant.DoublePair `json:"pbr,omitempty"`
	Pcr                 quant.DoublePair `json:"pcr,omitempty"`
	Operating           quant.DoublePair `json:"operating,omitempty"`
	Investing           quant.DoublePair `json:"investing,omitempty"`
	Financing           quant.DoublePair `json:"financing,omitempty"`
	DividendYield       quant.DoublePair `json:"dividend_yield,omitempty"`
	DividendPayoutRatio quant.DoublePair `json:"dividend_payout_ratio,omitempty"`
	Roa                 quant.DoublePair `json:"roa,omitempty"`
	Roe                 quant.DoublePair `json:"roe,omitempty"`
	MarketCap           quant.IntPair    `json:"market_cap,omitempty"`
	StartDate           time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"start_date,omitempty" example:"2016-03-31T00:00:000.Z"`
	EndDate             time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"end_date,omitempty" example:"2021-03-31T00:00:000.Z"`
}

func (r *QuantOptU) ToMap() map[string]interface{} {
	return objconv.ToMap(r)
}
