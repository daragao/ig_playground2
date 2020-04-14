package ig

type Balance struct {
	Available  float64 `json:"available"`
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
}

type Account struct {
	AccountAlias    string  `json:"accountAlias"`
	AccountId       string  `json:"accountId"`
	AccountName     string  `json:"accountName"`
	AccountType     string  `json:"accountType"`
	Balance         Balance `json:"balance"`
	CanTransferFrom bool    `json:"canTransferFrom"`
	CanTransferTo   bool    `json:"canTransferTo"`
	Currency        string  `json:"currency"`
	Preferred       bool    `json:"preferred"`
	Status          string  `json:"status"`
}

type OauthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
}

type Session struct {
	ClientId              string `json:"clientId"`
	AccountId             string `json:"accountId"`
	TimezoneOffset        int64  `json:"timezoneOffset"`
	LightstreamerEndpoint string `json:"lightstreamerEndpoint"`
	OauthToken            OauthToken
}

type MarketData struct {
	Bid                      float64 `json:"bid"`
	DelayTime                float64 `json:"delayTime"`
	Epic                     string  `json:"epic"`
	Expiry                   string  `json:"expiry"`
	High                     float64 `json:"high"`
	InstrumentName           string  `json:"instrumentName"`
	InstrumentType           string  `json:"instrumentType"`
	LotSize                  float64 `json:"lotSize"`
	MarketStatus             string  `json:"marketStatus"`
	NetChange                float64 `json:"netChange"`
	Offer                    float64 `json:"offer"`
	OtcTradeable             bool    `json:"otcTradeable"`
	PercentageChange         float64 `json:"percentageChange"`
	ScalingFactor            float64 `json:"scalingFactor"`
	StreamingPricesAvailable bool    `json:"streamingPricesAvailable"`
	UpdateTime               string  `json:"updateTime"`
	UpdateTimeUTC            string  `json:"updateTimeUTC"`
}

type MarketNode struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Unit struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type DealingRules struct {
	MarketOrderPreference         string `json:"marketOrderPreference"`
	MaxStopOrLimitDistance        Unit   `json:"maxStopOrLimitDistance"`
	MinControlledRiskStopDistance Unit   `json:"minControlledRiskStopDistance"`
	MinDealSize                   Unit   `json:"minDealSize"`
	MinNormalStopOrLimitDistance  Unit   `json:"minNormalStopOrLimitDistance"`
	MinStepDistance               Unit   `json:"minStepDistance"`
	TrailingStopsPreference       string `json:"minStepDistance"`
}

type Currency struct {
	BaseExchangeRate float64 `json:"baseExchangeRate"`
	Code             string  `json:"code"`
	ExchangeRate     float64 `json:"exchangeRate"`
	IsDefault        bool    `json:"isDefault"`
	Symbol           string  `json:"symbol"`
}

type MarketExpiryDetails struct {
	LastDealingDate string `json:"lastDealingDate"`
	SettlementInfo  string `json:"settlementInfo"`
}

type DepositBand struct {
	Currency string  `json:"currency"`
	Margin   float64 `json:"margin"`
	Max      float64 `json:"max"`
	Min      float64 `json:"min"`
}

type RolloverDetails struct {
	lastRolloverTime string `json:"lastRolloverTime"`
	rolloverInfo     string `json:"rolloverInfo"`
}

type OpeningHours struct {
	MarketTimes []struct {
		CloseTime string `json:"closeTime"`
		OpenTime  string `json:"openTime"`
	} `json:"marketTimes"`
}

type InstrumentDetails struct {
	ChartCode                      string              `json:"chartCode"`
	ContractSize                   string              `json:"contractSize"`
	ControlledRiskAllowed          bool                `json:"controlledRiskAllowed"`
	Country                        string              `json:"country"`
	Currencies                     []Currency          `json:"currencies"`
	Epic                           string              `json:"epic"`
	Expiry                         string              `json:"expiry"`
	ExpiryDetails                  MarketExpiryDetails `json:"expiryDetails"`
	ForceOpenAllowed               bool                `json:"forceOpenAllowed"`
	LimitedRiskPremium             Unit                `json:"limitedRiskPremium"`
	LotSize                        float64             `json:"lotSize"`
	MarginDepositBands             []DepositBand       `json:"marginDepositBands"`
	MarginFactor                   float64             `json:"marginFactor"`
	MarginFactorUnit               string              `json:"marginFactorUnit"`
	MarketId                       string              `json:"marketId"`
	Name                           string              `json:"name"`
	NewsCode                       string              `json:"newsCode"`
	OnePipMeans                    string              `json:"onePipMeans"`
	OpeningHours                   OpeningHours        `json:"openingHours"`
	RolloverDetails                RolloverDetails     `json:"rolloverDetails"`
	SlippageFactor                 Unit                `json:"slippageFactor"`
	SpecialInfo                    []string            `json:"specialInfo"`
	SprintMarketsMaximumExpiryTime float64             `json:"sprintMarketsMaximumExpiryTime"`
	SprintMarketsMinimumExpiryTime float64             `json:"sprintMarketsMinimumExpiryTime"`
	StopsLimitsAllowed             bool                `json:"stopsLimitsAllowed"`
	StreamingPricesAvailable       bool                `json:"streamingPricesAvailable"`
	Type                           string              `json:"type"`
	Unit                           string              `json:"unit"`
	ValueOfOnePip                  string              `json:"valueOfOnePip"`
}

type MarketSnapshotData struct {
	Bid                       float64 `json:"bid"`
	BinaryOdds                float64 `json:"binaryOdds"`
	ControlledRiskExtraSpread float64 `json:"controlledRiskExtraSpread"`
	DecimalPlacesFactor       float64 `json:"decimalPlacesFactor"`
	DelayTime                 float64 `json:"delayTime"`
	High                      float64 `json:"high"`
	Low                       float64 `json:"low"`
	MarketStatus              string  `json:"marketStatus"`
	NetChange                 float64 `json:"netChange"`
	Offer                     float64 `json:"offer"`
	PercentageChange          float64 `json:"percentageChange"`
	ScalingFactor             float64 `json:"scalingFactor"`
	UpdateTime                string  `json:"updateTime"`
}

type MarketDetail struct {
	DealingRules DealingRules       `json:"dealingRules"`
	Instrument   InstrumentDetails  `json:"instrument"`
	Snapshot     MarketSnapshotData `json:"snapshot"`
}

type PriceList struct {
	InstrumentType string          `json:"instrumentType"`
	Metadata       Metadata        `json:"metadata"`
	Prices         []PriceSnapshot `json:"prices"`
}

type PageMetadata struct {
	PageNumber int64 `json:"pageNumber"`
	PageSize   int64 `json:"pageSize"`
	TotalPages int64 `json:"totalPages"`
}

type Allowance struct {
	AllowanceExpiry    int64 `json:"allowanceExpiry"`
	RemainingAllowance int64 `json:"remainingAllowance"`
	TotalAllowance     int64 `json:"totalAllowance"`
}

type Metadata struct {
	PageData  PageMetadata `json:"pageData"`
	Size      int64        `json:"size"`
	Allowance Allowance    `json:"allowance"`
}

type Price struct {
	Ask        float64 `json:"ask"`
	Bid        float64 `json:"bid"`
	LastTraded float64 `json:"lastTraded"`
}

type PriceSnapshot struct {
	ClosedPrice      Price   `json:"closedPrice"`
	HighPrice        Price   `json:"highPrice"`
	LastTradedVolume float64 `json:"lastTradedVolume"`
	LowPrice         Price   `json:"lowPrice"`
	OpenPrice        Price   `json:"openPrice"`
	SnapshotTime     string  `json:"snapshotTime"`
	SnapshotTimeUTC  string  `json:"snapshotTimeUTC"`
}
