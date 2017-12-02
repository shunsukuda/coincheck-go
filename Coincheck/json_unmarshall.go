package Coincheck

import (
	"encoding/json"
	"strconv"
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

// type UnmarshalTicker JsonTicker

func (client CoinCheck) GetTicker(unmarshal int) (js *JsonTicker, raw string) {
	raw = client.ticker.all()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	return
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

func (client *CoinCheck) GetTradeLog(unmarshal int) (um []UnmarshalTrade, raw string) {
	var js []JsonTrade
	raw = client.trade.all()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	um = make([]UnmarshalTrade, len(js))
	for i := range js {
		um[i].Id = js[i].Id
		um[i].Amount, _ = js[i].Amount.Float64()
		um[i].Rate = float64(js[i].Rate)
		um[i].OrderType = js[i].OrderType
		um[i].CreatedAt = js[i].CreatedAt
	}
	return
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

func (client *CoinCheck) GetOrderBook(unmarshal int) (um *UnmarshalOrderBook, raw string) {
	var js JsonOrderBook
	raw = client.order_book.all()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	um = &UnmarshalOrderBook{}
	um.Asks = make([]Pair, len(js.Asks))
	um.Bids = make([]Pair, len(js.Bids))
	for i := range js.Asks {
		um.Asks[i].Rate, _ = js.Asks[i][0].Float64()
		um.Asks[i].Amount, _ = js.Asks[i][1].Float64()
		um.Bids[i].Rate, _ = js.Bids[i][0].Float64()
		um.Bids[i].Amount, _ = js.Bids[i][1].Float64()
	}
	return
}

type JsonOrder struct {
	Success      bool        `json:"success"`
	ID           int64       `json:"id"`
	Rate         json.Number `json:"rate"`
	Amount       json.Number `json:"amount"`
	OrderType    string      `json:"order_type"`
	StopLossRate json.Number `json:"stop_loss_rate"`
	Pair         string      `json:"pair"`
	CreatedAt    time.Time   `json:"created_at"`
}

type UnmarshalOrder struct {
	Success      bool
	ID           int64
	Rate         float64
	Amount       float64
	OrderType    string
	StopLossRate float64
	Pair         string
	CreatedAt    time.Time
}

func (client *CoinCheck) PostOrderNew(rate int64, amount float64, order_type string, unmarshal int) (um *UnmarshalOrder, raw string) {
	order := `{` +
		`"rate":"` + strconv.FormatInt(rate, 10) + `",` +
		`"amount":"` + strconv.FormatFloat(amount, 'f', -1, 64) + `",` +
		`"order_type":"` + order_type + `",` +
		`"pair":"btc_jpy"}`
	var js JsonOrder
	raw = client.order.create(order)
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	return
}

func (client *CoinCheck) PostOrderMarketNew(yen int64, amount float64, unmarshal int) (um *UnmarshalOrder, raw string) {
	order := `{`
	order_type := "market_"

	if yen > 0 {
		order_type += "buy"
		order += `"market_buy_amount":"` + strconv.FormatInt(yen, 10) + `",`
	} else if amount > 0.0 {
		order_type += "sell"
		order += `"amount":"` + strconv.FormatFloat(amount, 'f', -1, 64) + `",`
	}
	order += `"order_type":"` + order_type + `","pair":"btc_jpy"}`

	var js JsonOrder
	raw = client.order.create(order)
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	return
}

func (client *CoinCheck) PostOrderLeverageNew(rate int64, amount float64, order_type string, unmarshal int) (um *UnmarshalOrder, raw string) {
	order := `{` +
		`"rate":"` + strconv.FormatInt(rate, 10) + `",` +
		`"amount":"` + strconv.FormatFloat(amount, 'f', -1, 64) + `",` +
		`"order_type":"leverage_` + order_type + `",` +
		`"pair":"btc_jpy"}`
	var js JsonOrder
	raw = client.order.create(order)
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	return
}

func (client *CoinCheck) PostOrderCloseNew(rate int64, amount float64, id int64, order_type string, unmarshal int) (um *UnmarshalOrder, raw string) {
	order := `{` +
		`"rate":"` + strconv.FormatInt(rate, 10) + `",` +
		`"amount":"` + strconv.FormatFloat(amount, 'f', -1, 64) + `",` +
		`"position_id":"` + strconv.FormatInt(id, 10) + `",` +
		`"order_type":"close_` + order_type + `",` +
		`"pair":"btc_jpy"}`
	var js JsonOrder
	raw = client.order.create(order)
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	return
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

func (client *CoinCheck) GetOrderList(unmarshal int) (um *UnmarshalOrderOpens, raw string) {
	var js JsonOrderOpens
	raw = client.order.opens()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	um.Success = js.Success
	// XXX:Ordersを確保しなければいけない->別に構造体を定義する
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
	return
}

type JsonOrderCancel struct {
	Success bool  `json:"success"`
	Id      int64 `json:"id"`
}

func (client *CoinCheck) PostOrderCancel(id int64, unmarshal int) (success bool, raw string) {
	var js JsonOrderCancel
	raw = client.order.cancel(strconv.FormatInt(id, 10))
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	success = js.Success
	return
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

func (client *CoinCheck) GetTransactions(unmarshal int) (um *UnmarshalTransactions, raw string) {
	var js JsonTransactions
	raw = client.order.transactions()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
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
	return
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
	raw := []byte(client.order.transactionsi_pagination())
	if len(raw) == 0 {
		return nil
	}
	fmt.Println(string(raw))
	if err := json.Unmarshal(raw, &js); err != nil {
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

/*
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
	raw := []byte(client.leverage.positions())
	if len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, &js); err != nil {
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
*/

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

func (client *CoinCheck) GetBalance(unmarshal int) (um *UnmarshalBalance, raw string) {
	var js JsonBalance
	raw = client.account.balance()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
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
	return
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

func (client *CoinCheck) GetLeverageBalance(unmarshal int) (um *UnmarshalLeverageBalance, raw string) {
	var js JsonLeverageBalance
	raw = client.account.leverage_balance()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	um.Success = js.Success // XXX.エラー
	um.Margin.Jpy, _ = js.Margin.Jpy.Float64()
	um.MarginAvailale.Jpy, _ = js.MarginAvailale.Jpy.Float64()
	um.MarginLevel, _ = js.MarginLevel.Float64()
	return
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

func (client *CoinCheck) GetAccountInfo(unmarshal int) (um *UnmarshalAccountInfo, raw string) {
	var js JsonAccountInfo
	raw = client.account.info()
	if len(raw) == 0 || unmarshal == 0 {
		return
	}
	if err := json.Unmarshal([]byte(raw), &js); err != nil {
		return
	}
	um.Success = js.Success
	um.Id = js.Id
	um.Email = js.Email
	um.IdentityStatus = js.IdentityStatus
	um.BitcoinAddress = js.BitcoinAddress
	um.LendingLeverage, _ = js.LendingLeverage.Float64()
	um.TakerFee, _ = js.TakerFee.Float64()
	um.MakerFee, _ = js.MakerFee.Float64()
	return
}
