package Coincheck

import (
	"encoding/json"
	"log"
	"time"
)

type JsonTicker struct {
	Last      float64 `json:"last"`
	Bid       float64 `json:bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}

type UnmarshalTicker JsonTicker

func (client CoinCheck) GetTicker() *UnmarshalTicker {
	var js UnmarshalTicker
	tmp := []byte(client.ticker.all())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	return &js
}

type JsonTrade struct {
	Id        int64       `json:"id"`
	Amount    json.Number `json:"amount"`
	Rate      int64       `json:"rate"`
	OrderType string      `json:"order_type"`
	CreatedAt time.Time   `json:"created_at"`
}
type UnmarshalTrade struct {
	Id        int64
	Amount    float64
	Rate      float64
	OrderType string
	CreatedAt time.Time
}

func (client *CoinCheck) GetTrades() []UnmarshalTrade {
	var js []JsonTrade
	var um []UnmarshalTrade
	tmp := []byte(client.trade.all())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um = make([]UnmarshalTrade, len(js))
	for i := range js {
		um[i].Id = js[i].Id
		um[i].Amount, _ = js[i].Amount.Float64()
		um[i].Rate = float64(js[i].Rate)
		um[i].OrderType = js[i].OrderType
		um[i].CreatedAt = js[i].CreatedAt
	}
	return um
}

type JsonOrderBook struct {
	Asks [][]json.Number `json:"asks"`
	Bids [][]json.Number `json:"bids"`
}

type UnmarshalOrderBook struct {
	Asks []Pair
	Bids []Pair
}

type Pair struct {
	Rate   float64
	Amount float64
}

func (client *CoinCheck) GetOrderBooks() *UnmarshalOrderBook {
	var js JsonOrderBook
	var um UnmarshalOrderBook
	tmp := []byte(client.trade.all())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	for i := range js.Asks {
		um.Asks[i].Rate, _ = js.Asks[i][0].Float64()
		um.Asks[i].Amount, _ = js.Asks[i][1].Float64()
		um.Bids[i].Rate, _ = js.Bids[i][0].Float64()
		um.Bids[i].Amount, _ = js.Bids[i][1].Float64()
	}
	return &um
}

func (client *CoinCheck) PostOrderCreate(rate int64, amount float64, order_type string) {
}

type JsonOrderOpens struct {
	Success bool `json:"success"`
	Orders  []struct {
		Id                     int64       `json:"id"`
		OrderType              string      `json:"order_type"`
		Rate                   int64       `json:"rate"`
		Pair                   string      `json:"pair"`
		PendingAmount          json.Number `json:"pending_amount"`
		PendingMarketBuyAmount json.Number `json:"pending_market_buy_amount"`
		StopLossRate           json.Number `json:"stop_loss_rate"`
		CreatedAt              time.Time   `json:"created_at"`
	} `json:"orders"`
}

type UnmarshalOrderOpens struct {
	Success bool
	Orders  []struct {
		Id                     int64
		OrderType              string
		Rate                   int64
		Pair                   string
		PendingAmount          float64
		PendingMarketBuyAmount float64
		StopLossRate           float64
		CreatedAt              time.Time
	}
}

func (client *CoinCheck) GetOrderOpens() *UnmarshalOrderOpens {
	var js JsonOrderOpens
	var um UnmarshalOrderOpens
	tmp := []byte(client.order.opens())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	for i := range js.Orders {
		um.Orders[i].Id = js.Orders[i].Id
		um.Orders[i].OrderType = js.Orders[i].OrderType
		um.Orders[i].Rate = js.Orders[i].Rate
		um.Orders[i].Pair = js.Orders[i].Pair
		um.Orders[i].PendingAmount, _ = js.Orders[i].PendingAmount.Float64()
		um.Orders[i].PendingMarketBuyAmount, _ = js.Orders[i].PendingMarketBuyAmount.Float64()
		um.Orders[i].StopLossRate, _ = js.Orders[i].StopLossRate.Float64()
		um.Orders[i].CreatedAt = js.Orders[i].CreatedAt
	}
	return &um
}

type JsonTransactions struct {
	Success      bool `json:"success"`
	Transactions []struct {
		Id        int64     `json:"id"`
		OrderId   int64     `json:"order_id"`
		CreatedAt time.Time `json:"created_at"`
		Funds     struct {
			Btc json.Number `json:"btc"`
			Jpy json.Number `json:"jpy"`
		} `json:"funds"`
		Pair        string      `json:"pair"`
		Rate        json.Number `json:"rate"`
		FeeCurrency string      `json:"fee_currency"`
		Fee         json.Number `json:"fee"`
		Liquidity   string      `json:"liquidity"`
		Side        string      `json:"side"`
	} `json:"transactions"`
}

type UnmarshalTransactions struct {
	Success      bool
	Transactions []struct {
		Id        int64
		OrderId   int64
		CreatedAt time.Time
		Funds     struct {
			Btc float64
			Jpy float64
		}
		Pair        string
		Rate        float64
		FeeCurrency string
		Fee         float64
		Liquidity   string
		Side        string
	}
}

func (client *CoinCheck) GetTransactions() *UnmarshalTransactions {
	var js JsonTransactions
	var um UnmarshalTransactions
	tmp := []byte(client.order.transactions())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	for i := range js.Transactions {
		um.Transactions[i].Id = js.Transactions[i].Id
		um.Transactions[i].OrderId = js.Transactions[i].OrderId
		um.Transactions[i].CreatedAt = js.Transactions[i].CreatedAt
		um.Transactions[i].Funds.Btc, _ = js.Transactions[i].Funds.Btc.Float64()
		um.Transactions[i].Funds.Jpy, _ = js.Transactions[i].Funds.Jpy.Float64()
		um.Transactions[i].Pair = js.Transactions[i].Pair
		um.Transactions[i].Rate, _ = js.Transactions[i].Rate.Float64()
		um.Transactions[i].FeeCurrency = js.Transactions[i].FeeCurrency
		um.Transactions[i].Fee, _ = js.Transactions[i].Fee.Float64()
		um.Transactions[i].Liquidity = js.Transactions[i].Liquidity
		um.Transactions[i].Side = js.Transactions[i].Side
	}
	return &um
}

/*
type JsonTransactionsPagination struct {
	Success    bool `json:"success"`
	Pagination struct {
		Limit         int64  `json:"limit"`
		Order         string `json:"order"`
		StartingAfter string `json:"starting_after"`
		EndingBefore  string `json:"ending_before"`
	} `json:"pagination"`
	Data []struct {
		Id        int64     `json:"id"`
		OrderId   int64     `json:"order_id"`
		CreatedAt time.Time `json:"created_at"`
		Funds     struct {
			Btc json.Number `json:"btc"`
			Jpy json.Number `json:"jpy"`
		} `json:"funds"`
		Pair        string      `json:"pair"`
		Rate        json.Number `json:"rate"`
		FeeCurrency string      `json:"fee_currency"`
		Fee         json.Number `json:"fee"`
		Liquidity   string      `json:"liquidity"`
		Side        string      `json:"side"`
	} `json:"data"`
}

type UnmarshalTransactionsPagination struct {
	Success    bool
	Pagination struct {
		Limit         int64
		Order         string
		StartingAfter string
		EndingBefore  string
	}
	Data []struct {
		Id        int64
		OrderId   int64
		CreatedAt time.Time
		Funds     struct {
			Btc float64
			Jpy float64
		}
		Pair        string
		Rate        float64
		FeeCurrency string
		Fee         float64
		Liquidity   string
		Side        string
	}
}

func (client *CoinCheck) GetTransactionsPagination() *UnmarshalTransactionsPagination {
	var js JsonTransactionsPagination
	var um UnmarshalTransactionsPagination
	tmp := []byte(client.order.transactionsi_pagination())
	if len(tmp) == 0 {
		return nil
	}
	fmt.Println(string(tmp))
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	um.Pagination.Limit = js.Pagination.Limit
	um.Pagination.Order = js.Pagination.Order
	um.Pagination.StartingAfter = js.Pagination.StartingAfter
	um.Pagination.EndingBefore = js.Pagination.EndingBefore
	for i := range js.Data {
		um.Data[i].Id = js.Data[i].Id
		um.Data[i].OrderId = js.Data[i].OrderId
		um.Data[i].CreatedAt = js.Data[i].CreatedAt
		um.Data[i].Funds.Btc, _ = js.Data[i].Funds.Btc.Float64()
		um.Data[i].Funds.Jpy, _ = js.Data[i].Funds.Jpy.Float64()
		um.Data[i].Pair = js.Data[i].Pair
		um.Data[i].Rate, _ = js.Data[i].Rate.Float64()
		um.Data[i].FeeCurrency = js.Data[i].FeeCurrency
		um.Data[i].Fee, _ = js.Data[i].Fee.Float64()
		um.Data[i].Liquidity = js.Data[i].Liquidity
		um.Data[i].Side = js.Data[i].Side
	}
	return &um
}
*/

type JsonLeveragePositions struct {
	Success    bool `json:"success"`
	Pagination struct {
		Limit         int64  `json:"limit"`
		Order         string `json:"order"`
		StartingAfter string `json:"starting_after"`
		EndingBefore  string `json:"ending_before"`
	} `json:"pagination"`
	Data []struct {
		Id         int64       `json:"id"`
		Pair       string      `json:"pair"`
		Status     string      `json:"status"`
		CreatedAt  time.Time   `json:"created_at"`
		ClosedAt   time.Time   `json:"closed_at"`
		OpenRate   json.Number `json:"open_rate"`
		ClosedRate json.Number `json:"closed_rate"`
		Amount     json.Number `json:"amount"`
		AllAmount  json.Number `json:"all_amount"`
		Side       string      `json:"side"`
		Pl         json.Number `json:"pl"`
		NewOrder   struct {
			Id            int64       `json:"id"`
			Side          string      `json:"side"`
			Rate          json.Number `json:"rate"`
			Amount        json.Number `json:"amount"`
			PendingAmount json.Number `json:"pending_amount"`
			Status        string      `json:"status"`
			CreatedAt     time.Time   `json:"created_at"`
		} `json:"new_order"`
		CloseOrders []struct {
			Id            int64       `json:"id"`
			Side          string      `json:"side"`
			Rate          json.Number `json:"rate"`
			Amount        json.Number `json:"amount"`
			PendingAmount json.Number `json:"pending_amount"`
			Status        string      `json:"status"`
			CreatedAt     time.Time   `json:"created_at"`
		} `json:"close_orders"`
	} `json:"data"`
}

type UnmarshalLeveragePositions struct {
	Success    bool
	Pagination struct {
		Limit         int64
		Order         string
		StartingAfter string
		EndingBefore  string
	}
	Data []struct {
		Id         int64
		Pair       string
		Status     string
		CreatedAt  time.Time
		ClosedAt   time.Time
		OpenRate   float64
		ClosedRate float64
		Amount     float64
		AllAmount  float64
		Side       string
		Pl         float64
		NewOrder   struct {
			Id            int64
			Side          string
			Rate          float64
			Amount        float64
			PendingAmount float64
			Status        string
			CreatedAt     time.Time
		}
		CloseOrders []struct {
			Id            int64
			Side          string
			Rate          float64
			Amount        float64
			PendingAmount float64
			Status        string
			CreatedAt     time.Time
		}
	}
}

func (client *CoinCheck) GetLeveragePositions() *UnmarshalLeveragePositions {
	var js JsonLeveragePositions
	var um UnmarshalLeveragePositions
	tmp := []byte(client.leverage.positions())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	um.Pagination.Limit = js.Pagination.Limit
	um.Pagination.Order = js.Pagination.Order
	um.Pagination.StartingAfter = js.Pagination.StartingAfter
	um.Pagination.EndingBefore = js.Pagination.EndingBefore
	for i := range js.Data {
		um.Data[i].Id = js.Data[i].Id
		um.Data[i].Pair = js.Data[i].Pair
		um.Data[i].Status = js.Data[i].Status
		um.Data[i].CreatedAt = js.Data[i].CreatedAt
		um.Data[i].ClosedAt = js.Data[i].ClosedAt
		um.Data[i].OpenRate, _ = js.Data[i].OpenRate.Float64()
		um.Data[i].ClosedRate, _ = js.Data[i].ClosedRate.Float64()
		um.Data[i].Amount, _ = js.Data[i].Amount.Float64()
		um.Data[i].AllAmount, _ = js.Data[i].AllAmount.Float64()
		um.Data[i].Side = js.Data[i].Side
		um.Data[i].Pl, _ = js.Data[i].Pl.Float64()
		um.Data[i].NewOrder.Id = js.Data[i].NewOrder.Id
		um.Data[i].NewOrder.Side = js.Data[i].NewOrder.Side
		um.Data[i].NewOrder.Rate, _ = js.Data[i].NewOrder.Rate.Float64()
		um.Data[i].NewOrder.Amount, _ = js.Data[i].NewOrder.Amount.Float64()
		um.Data[i].NewOrder.PendingAmount, _ = js.Data[i].NewOrder.PendingAmount.Float64()
		um.Data[i].NewOrder.Status = js.Data[i].NewOrder.Status
		um.Data[i].NewOrder.CreatedAt = js.Data[i].NewOrder.CreatedAt
		for j := range js.Data[i].CloseOrders {
			um.Data[i].CloseOrders[j].Id = js.Data[i].CloseOrders[j].Id
			um.Data[i].CloseOrders[j].Side = js.Data[i].CloseOrders[j].Side
			um.Data[i].CloseOrders[j].Rate, _ = js.Data[i].CloseOrders[j].Rate.Float64()
			um.Data[i].CloseOrders[j].Amount, _ = js.Data[i].CloseOrders[j].Amount.Float64()
			um.Data[i].CloseOrders[j].PendingAmount, _ = js.Data[i].CloseOrders[j].PendingAmount.Float64()
			um.Data[i].CloseOrders[j].Status = js.Data[i].CloseOrders[j].Status
			um.Data[i].CloseOrders[j].CreatedAt = js.Data[i].CloseOrders[j].CreatedAt
		}
	}
	return &um
}

type JsonBalance struct {
	Success      bool        `json:"success"`
	Jpy          json.Number `json:"jpy"`
	Btc          json.Number `json:"btc"`
	JpyReserved  json.Number `json:"jpy_reserved"`
	BtcReserved  json.Number `json:"btc_reserved"`
	JpyLendInUse json.Number `json:"jpy_lend_in_use"`
	BtcLendInUse json.Number `json:"btc_lend_in_use"`
	JpyLent      json.Number `json:"jpy_lent"`
	BtcLent      json.Number `json:"btc_lent"`
	JpyDebt      json.Number `json:"jpy_debt"`
	BtcDebt      json.Number `json:"btc_debt"`
}

type UnmarshalBalance struct {
	Success      bool
	Jpy          float64
	Btc          float64
	JpyReserved  float64
	BtcReserved  float64
	JpyLendInUse float64
	BtcLendInUse float64
	JpyLent      float64
	BtcLent      float64
	JpyDebt      float64
	BtcDebt      float64
}

func (client *CoinCheck) GetBalance() *UnmarshalBalance {
	var js JsonBalance
	var um UnmarshalBalance
	tmp := []byte(client.account.balance())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	um.Jpy, _ = js.Jpy.Float64()
	um.Btc, _ = js.Btc.Float64()
	um.JpyReserved, _ = js.JpyReserved.Float64()
	um.BtcReserved, _ = js.BtcReserved.Float64()
	um.JpyLendInUse, _ = js.JpyLendInUse.Float64()
	um.BtcLendInUse, _ = js.BtcLendInUse.Float64()
	um.JpyLent, _ = js.JpyLent.Float64()
	um.BtcLent, _ = js.BtcLent.Float64()
	um.JpyDebt, _ = js.JpyDebt.Float64()
	um.BtcDebt, _ = js.BtcDebt.Float64()
	return &um
}

type JsonLeverageBalance struct {
	Success bool `json:"success"`
	Margin  struct {
		Jpy json.Number `json:"jpy"`
	} `json:"margin"`
	MarginAvailale struct {
		Jpy json.Number `json:"jpy"`
	} `json:"margin_available"`
	MarginLevel json.Number `json:"margin_level"`
}

type UnmarshalLeverageBalance struct {
	Success bool
	Margin  struct {
		Jpy float64
	}
	MarginAvailale struct {
		Jpy float64
	}
	MarginLevel float64
}

func (client *CoinCheck) GetLeverageBalance() *UnmarshalLeverageBalance {
	var js JsonLeverageBalance
	var um UnmarshalLeverageBalance
	tmp := []byte(client.account.leverage_balance())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	um.Margin.Jpy, _ = js.Margin.Jpy.Float64()
	um.MarginAvailale.Jpy, _ = js.MarginAvailale.Jpy.Float64()
	um.MarginLevel, _ = js.MarginLevel.Float64()
	return &um
}

type JsonAccountInfo struct {
	Success         bool        `json:"success"`
	Id              int64       `json:"id"`
	Email           string      `json:"email"`
	IdentityStatus  string      `json:"identity_status"`
	BitcoinAddress  string      `json:"bitcoin_address"`
	LendingLeverage json.Number `json:"lending_leverage"`
	TakerFee        json.Number `json:"taker_fee"`
	MakerFee        json.Number `json:"maker_fee"`
}

type UnmarshalAccountInfo struct {
	Success         bool
	Id              int64
	Email           string
	IdentityStatus  string
	BitcoinAddress  string
	LendingLeverage float64
	TakerFee        float64
	MakerFee        float64
}

func (client *CoinCheck) GetAccountInfo() *UnmarshalAccountInfo {
	var js JsonAccountInfo
	var um UnmarshalAccountInfo
	tmp := []byte(client.account.info())
	if len(tmp) == 0 {
		return nil
	}
	if err := json.Unmarshal(tmp, &js); err != nil {
		log.Fatal(err)
	}
	um.Success = js.Success
	um.Id = js.Id
	um.Email = js.Email
	um.IdentityStatus = js.IdentityStatus
	um.BitcoinAddress = js.BitcoinAddress
	um.LendingLeverage, _ = js.LendingLeverage.Float64()
	um.TakerFee, _ = js.TakerFee.Float64()
	um.MakerFee, _ = js.MakerFee.Float64()
	return &um
}
