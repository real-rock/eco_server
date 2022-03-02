package response

import (
	"fmt"
	"main/internal/core/data"
	"main/internal/core/pb"
	"main/internal/pkg/logger"
	"time"
)

type ChartData struct {
	StartDate       time.Time `time_format:"2006-01-02T15:04:05.000Z" json:"start_date" bson:"start_date" example:"2016-03-31T00:00:000.Z"`
	ProfitRateData  []float32 `json:"profit_rate_data" bson:"profit_rate_data" example:"[8.31201046811529,15.13554790878776,-1.336521221573761,-1.42408166715555,10.420784591586559,8.305691643668455,17.68356243256443,9.407034979656027,-4.15162926200139,5.542443496088845,6.654446258518339]"`
	ProfitKospiData []float32 `json:"profit_kospi_data" bson:"profit_kospi_data" example:"[1995.85,1994.15,1983.4,1970.35,2016.19,2034.65,2043.63,2008.19,1983.48,2026.46,2067.57]"`
}

type QuantResponse struct {
	QuantID             uint      `json:"quant_id" bson:"quant_id" example:"5"`
	CumulativeReturn    float64   `json:"cumulative_return" bson:"cumulative_return" example:"15.95"`
	AnnualAverageReturn float64   `json:"annual_average_return" bson:"annual_average_return" example:"-2.21"`
	WinningPercentage   float64   `json:"winning_percentage" bson:"winning_percentage" example:"45.45"`
	MaxLossRate         float64   `json:"max_loss_rate" bson:"max_loss_rate" example:"-26.46"`
	HoldingsCount       int32     `json:"holdings_count" bson:"holdings_count" example:"7"`
	ChartData           ChartData `json:"chart_data" bson:"chart_data,inline"`
}

func NewQuantResultFromPB(pb *pb.QuantResult) *QuantResponse {
	return &QuantResponse{
		CumulativeReturn:    pb.CumulativeReturn,
		AnnualAverageReturn: pb.AnnualAverageReturn,
		WinningPercentage:   pb.WinningPercentage,
		MaxLossRate:         pb.MaxLossRate,
		HoldingsCount:       pb.HoldingsCount,
		ChartData: ChartData{
			StartDate:      pb.ChartData.StartDate.AsTime(),
			ProfitRateData: pb.ChartData.ProfitRateData,
		},
	}
}

func (qr *QuantResponse) AddKospiData() error {
	var kospiVal []float32

	dataLen := len(qr.ChartData.ProfitRateData)
	startTime := qr.ChartData.StartDate
	idx := data.KospiData.Date[startTime]
	if idx < dataLen {
		return fmt.Errorf("error in AddKospiData got wrong data: 'idx':%d, 'dataLen':%d", idx, dataLen)
	}

	for i := idx; i >= idx-dataLen; i-- {
		if i < 0 {
			logger.Logger.Errorf("error in AddKospiData: got negative index: 'i': %d", idx)
			return fmt.Errorf("error in AddKospiData: got negative index: 'i': %d", idx)
		}
		kospiVal = append(kospiVal, data.KospiData.IndexVal[i])
	}
	qr.ChartData.ProfitKospiData = kospiVal

	return nil
}
