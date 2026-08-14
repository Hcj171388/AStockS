package scoring

import (
	"math"
	"sort"
)

type StockInput struct {
	Code                  string  `json:"code"`
	Name                  string  `json:"name"`
	FloatMarketCap        float64 `json:"floatMarketCap"`
	AvgTurnover20d        float64 `json:"avgTurnover20d"`
	TotalTurnover         float64 `json:"totalTurnover"`
	ChangePct             float64 `json:"changePct"`
	TrendScore            float64 `json:"trendScore"`
	MomentumScore         float64 `json:"momentumScore"`
	TurnoverRate          float64 `json:"turnoverRate"`
	BuySellScore          float64 `json:"buySellScore"`
	ChipConcentration     float64 `json:"chipConcentration"`
	ExchangeHealth        float64 `json:"exchangeHealth"`
	MarketScore           float64 `json:"marketScore"`
	ConsecutiveInflowDays int     `json:"consecutiveInflowDays"`
	InflowSlope           float64 `json:"inflowSlope"`
	SuperOrderNetInflow   float64 `json:"superOrderNetInflow"`
	LargeOrderNetInflow   float64 `json:"largeOrderNetInflow"`
	MediumOrderNetInflow  float64 `json:"mediumOrderNetInflow"`
	SmallOrderNetInflow   float64 `json:"smallOrderNetInflow"`
}

type StockScore struct {
	StockInput
	NetInflow       float64  `json:"netInflow"`
	FundStrength    float64  `json:"fundStrength"`
	FundEfficiency  float64  `json:"fundEfficiency"`
	FundConsistency float64  `json:"fundConsistency"`
	FundScore       float64  `json:"fundScore"`
	PriceScore      float64  `json:"priceScore"`
	ChipScore       float64  `json:"chipScore"`
	Score           float64  `json:"score"`
	Reasons         []string `json:"reasons"`
}

func Rank(stocks []StockInput) []StockScore {
	out := make([]StockScore, 0, len(stocks))
	for _, s := range stocks {
		if !passesBaseFilters(s) {
			continue
		}
		scored := Score(s)
		out = append(out, scored)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}

func Score(s StockInput) StockScore {
	net := s.SuperOrderNetInflow + s.LargeOrderNetInflow + s.MediumOrderNetInflow + s.SmallOrderNetInflow
	strength := safeDiv(net, s.FloatMarketCap)
	consistency := safeDiv(net, s.TotalTurnover)
	efficiency := safeDiv(s.ChangePct, math.Abs(strength)+0.000001)
	fund := clamp01(strength*18)*35 + clamp01(efficiency/8)*20 + clamp01(float64(s.ConsecutiveInflowDays)/5)*20 + clamp01(consistency*10)*15 + clamp01(s.InflowSlope*10)*10
	price := clamp01(s.TrendScore)*35 + clamp01(efficiency/10)*30 + clamp01(s.MomentumScore)*35
	chip := clamp01(s.TurnoverRate/8)*25 + clamp01(s.BuySellScore)*25 + clamp01(s.ChipConcentration)*25 + clamp01(s.ExchangeHealth)*25
	total := fund*0.35 + price*0.25 + chip*0.2 + clamp01(s.MarketScore)*100*0.2
	reasons := []string{}
	if net > 0 {
		reasons = append(reasons, "资金净流入为正")
	}
	if s.ConsecutiveInflowDays >= 3 {
		reasons = append(reasons, "资金连续流入")
	}
	if s.TrendScore >= 0.65 {
		reasons = append(reasons, "价格趋势向上")
	}
	if s.MomentumScore >= 0.65 {
		reasons = append(reasons, "动能较强")
	}
	return StockScore{StockInput: s, NetInflow: net, FundStrength: strength, FundEfficiency: efficiency, FundConsistency: consistency, FundScore: round(fund), PriceScore: round(price), ChipScore: round(chip), Score: round(total), Reasons: reasons}
}

func passesBaseFilters(s StockInput) bool {
	return s.FloatMarketCap >= 2_000_000_000 && s.FloatMarketCap <= 50_000_000_000 && s.AvgTurnover20d >= 30_000_000 && s.Name != ""
}
func safeDiv(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
func round(v float64) float64 { return math.Round(v*100) / 100 }
