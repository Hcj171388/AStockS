package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"astocks/internal/scoring"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "web/index.html") })
	mux.HandleFunc("/api/rank", rankHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("AStockS listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func rankHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	stocks := sampleStocks()
	if r.Method == http.MethodPost {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&stocks); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(scoring.Rank(stocks))
}

func sampleStocks() []scoring.StockInput {
	return []scoring.StockInput{
		{Code: "600519", Name: "价值龙头", FloatMarketCap: 38_000_000_000, AvgTurnover20d: 180_000_000, TotalTurnover: 460_000_000, ChangePct: 3.2, TrendScore: .82, MomentumScore: .76, TurnoverRate: 3.5, BuySellScore: .8, ChipConcentration: .72, ExchangeHealth: .7, MarketScore: .78, ConsecutiveInflowDays: 4, InflowSlope: .08, SuperOrderNetInflow: 28_000_000, LargeOrderNetInflow: 16_000_000, MediumOrderNetInflow: 5_000_000, SmallOrderNetInflow: -2_000_000},
		{Code: "300750", Name: "成长先锋", FloatMarketCap: 22_000_000_000, AvgTurnover20d: 120_000_000, TotalTurnover: 320_000_000, ChangePct: 5.8, TrendScore: .9, MomentumScore: .88, TurnoverRate: 5.2, BuySellScore: .74, ChipConcentration: .68, ExchangeHealth: .65, MarketScore: .74, ConsecutiveInflowDays: 3, InflowSlope: .06, SuperOrderNetInflow: 18_000_000, LargeOrderNetInflow: 11_000_000, MediumOrderNetInflow: 3_000_000, SmallOrderNetInflow: -1_000_000},
		{Code: "002415", Name: "稳健制造", FloatMarketCap: 8_000_000_000, AvgTurnover20d: 55_000_000, TotalTurnover: 130_000_000, ChangePct: 1.6, TrendScore: .64, MomentumScore: .58, TurnoverRate: 2.2, BuySellScore: .66, ChipConcentration: .7, ExchangeHealth: .73, MarketScore: .7, ConsecutiveInflowDays: 5, InflowSlope: .04, SuperOrderNetInflow: 6_000_000, LargeOrderNetInflow: 4_000_000, MediumOrderNetInflow: 1_000_000, SmallOrderNetInflow: 500_000},
	}
}
