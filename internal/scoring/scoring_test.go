package scoring

import "testing"

func TestScoreCombinesOrderNetInflow(t *testing.T) {
	stock := StockInput{Name: "测试", FloatMarketCap: 10_000_000_000, AvgTurnover20d: 50_000_000, TotalTurnover: 100_000_000, SuperOrderNetInflow: 10, LargeOrderNetInflow: 20, MediumOrderNetInflow: -3, SmallOrderNetInflow: 1, TrendScore: .5, MomentumScore: .5, MarketScore: .5}
	score := Score(stock)
	if score.NetInflow != 28 {
		t.Fatalf("expected net inflow 28, got %v", score.NetInflow)
	}
}

func TestRankFiltersAndSorts(t *testing.T) {
	stocks := []StockInput{
		{Name: "低成交", FloatMarketCap: 10_000_000_000, AvgTurnover20d: 1, TotalTurnover: 100_000_000},
		{Name: "A", FloatMarketCap: 10_000_000_000, AvgTurnover20d: 50_000_000, TotalTurnover: 100_000_000, SuperOrderNetInflow: 2_000_000, LargeOrderNetInflow: 1_000_000, TrendScore: .9, MomentumScore: .9, MarketScore: .9},
		{Name: "B", FloatMarketCap: 10_000_000_000, AvgTurnover20d: 50_000_000, TotalTurnover: 100_000_000, SuperOrderNetInflow: 1, TrendScore: .1, MomentumScore: .1, MarketScore: .1},
	}
	ranked := Rank(stocks)
	if len(ranked) != 2 {
		t.Fatalf("expected 2 ranked stocks, got %d", len(ranked))
	}
	if ranked[0].Score < ranked[1].Score {
		t.Fatalf("expected descending scores")
	}
}
