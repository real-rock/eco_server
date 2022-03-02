package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"main/internal/core/model/quant"
	"main/internal/core/model/request"
	"main/internal/core/pb"
	"main/internal/pkg/objconv"
	"time"
)

type QuantOption struct {
	gorm.Model          `json:"-"`
	QuantID             uint               `gorm:"column:quant_id" json:"quant_id"`
	MainSectors         []quant.MainSector `gorm:"constraint:OnDelete:CASCADE;" json:"main_sectors"`
	NetRevenue          quant.IntPair      `gorm:"embedded;embeddedPrefix:net_revenue_" json:"net_revenue"`
	NetRevenueRate      quant.DoublePair   `gorm:"embedded;embeddedPrefix:net_revenue_rate_" json:"net_revenue_rate"`
	NetProfit           quant.IntPair      `gorm:"embedded;embeddedPrefix:net_profit_" json:"net_profit"`
	NetProfitRate       quant.DoublePair   `gorm:"embedded;embeddedPrefix:net_profit_rate_" json:"net_profit_rate"`
	DERatio             quant.DoublePair   `gorm:"embedded;embeddedPrefix:de_ratio_" json:"de_ratio"`
	Per                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:per_" json:"per"`
	Psr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:psr_" json:"psr"`
	Pbr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:pbr_" json:"pbr"`
	Pcr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:pcr_" json:"pcr"`
	Activities          quant.Activities   `gorm:"embedded;embeddedPrefix:activities_" json:"activities"`
	DividendYield       quant.DoublePair   `gorm:"embedded;embeddedPrefix:dividend_yield_" json:"dividend_yield"`
	DividendPayoutRatio quant.DoublePair   `gorm:"embedded;embeddedPrefix:dividend_payout_ratio_" json:"dividend_payout_ratio"`
	Roa                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:roa_" json:"roa"`
	Roe                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:roe_" json:"roe"`
	MarketCap           quant.IntPair      `gorm:"embedded;embeddedPrefix:market_cap_" json:"market_cap"`
	StartDate           time.Time          `time_format:"2006-01-02T15:04:05.000Z" json:"start_date" example:"2016-03-31T00:00:000.Z"`
	EndDate             time.Time          `time_format:"2006-01-02T15:04:05.000Z" json:"end_date" example:"2021-03-31T00:00:000.Z"`
}

func NewQuantOption(req *request.QuantC) *QuantOption {
	return &QuantOption{
		QuantID:        req.QuantID,
		MainSectors:    quant.NewMainSectors(req.QuantID, req.MainSectors),
		NetRevenue:     req.NetRevenue,
		NetRevenueRate: req.NetRevenueRate,
		NetProfit:      req.NetProfit,
		NetProfitRate:  req.NetProfitRate,
		DERatio:        req.DERatio,
		Per:            req.Per,
		Psr:            req.Psr,
		Pbr:            req.Pbr,
		Pcr:            req.Pcr,
		Activities: quant.Activities{
			Operating: req.Operating,
			Investing: req.Investing,
			Financing: req.Financing,
		},
		DividendYield:       req.DividendYield,
		DividendPayoutRatio: req.DividendPayoutRatio,
		Roa:                 req.Roa,
		Roe:                 req.Roe,
		MarketCap:           req.MarketCap,
		StartDate:           req.StartDate,
		EndDate:             req.EndDate,
	}
}

func (q *QuantOption) ToRequest() *pb.QuantRequest {
	var sectors []string
	for _, ms := range q.MainSectors {
		sectors = append(sectors, ms.ToString())
	}
	return &pb.QuantRequest{
		MainSector:          sectors,
		NetRevenue:          q.NetRevenue.ToPB(),
		NetRevenueRate:      q.NetRevenueRate.ToPB(),
		NetProfit:           q.NetProfit.ToPB(),
		NetProfitRate:       q.NetProfitRate.ToPB(),
		DeRatio:             q.DERatio.ToPB(),
		Per:                 q.Per.ToPB(),
		Psr:                 q.Psr.ToPB(),
		Pbr:                 q.Pbr.ToPB(),
		Pcr:                 q.Pcr.ToPB(),
		Activities:          q.Activities.ToPB(),
		DividendYield:       q.DividendYield.ToPB(),
		DividendPayoutRatio: q.DividendPayoutRatio.ToPB(),
		Roa:                 q.Roa.ToPB(),
		Roe:                 q.Roe.ToPB(),
		MarketCap:           q.MarketCap.ToPB(),
		StartDate:           timestamppb.New(q.StartDate),
		EndDate:             timestamppb.New(q.EndDate),
	}
}

func (q *QuantOption) ToMap() map[string]interface{} {
	return objconv.ToMap(q)
}

func (q *QuantOption) TableName() string {
	return "quant_options"
}
