package binance

import (
	"context"
	"net/http"
)

type SimpleEarnGetAccountService struct {
	c *Client
}

type SimpleEarnAccount struct {
	TotalAmountInBTC          string `json:"totalAmountInBTC"`
	TotalAmountInUSDT         string `json:"totalAmountInUSDT"`
	TotalFlexibleAmountInBTC  string `json:"totalFlexibleAmountInBTC"`
	TotalFlexibleAmountInUSDT string `json:"totalFlexibleAmountInUSDT"`
	TotalLockedInBTC          string `json:"totalLockedInBTC"`
	TotalLockedInUSDT         string `json:"totalLockedInUSDT"`
}

func (s *SimpleEarnGetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SimpleEarnAccount)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, err
}

type SimpleEarnListFlexibleProductService struct {
	c       *Client
	asset   string
	current int
	size    int
}

func (s *SimpleEarnListFlexibleProductService) Asset(asset string) *SimpleEarnListFlexibleProductService {
	s.asset = asset
	return s
}

func (s *SimpleEarnListFlexibleProductService) Current(current int) *SimpleEarnListFlexibleProductService {
	s.current = current
	return s
}

func (s *SimpleEarnListFlexibleProductService) Size(size int) *SimpleEarnListFlexibleProductService {
	s.size = size
	return s
}

type SimpleEarnFlexibleProductResp struct {
	Rows  []SimpleEarnFlexibleProduct `json:"rows"`
	Total int                         `json:"total"`
}

type SimpleEarnFlexibleProduct struct {
	Asset                      string            `json:"asset"`
	LatestAnnualPercentageRate string            `json:"latestAnnualPercentageRate"`
	TierAnnualPercentageRate   map[string]string `json:"tierAnnualPercentageRate"`
	AirDropPercentageRate      string            `json:"airDropPercentageRate"`
	CanPurchase                bool              `json:"canPurchase"`
	CanRedeem                  bool              `json:"canRedeem"`
	IsSoldOut                  bool              `json:"isSoldOut"`
	Hot                        bool              `json:"hot"`
	MinPurchaseAmount          string            `json:"minPurchaseAmount"`
	ProductId                  string            `json:"productId"`
	SubscriptionStartTime      int64             `json:"subscriptionStartTime"`
	Status                     string            `json:"status"`
}

func (s *SimpleEarnListFlexibleProductService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleProductResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/list",
		secType:  secTypeSigned,
	}

	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	if s.current != 0 {
		r.setParam("current", s.current)
	}
	if s.size != 0 {
		r.setParam("size", s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnFlexibleProductResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnListLockedProductService struct {
	c       *Client
	asset   string
	current int
	size    int
}

func (s *SimpleEarnListLockedProductService) Asset(asset string) *SimpleEarnListLockedProductService {
	s.asset = asset
	return s
}

func (s *SimpleEarnListLockedProductService) Current(current int) *SimpleEarnListLockedProductService {
	s.current = current
	return s
}

func (s *SimpleEarnListLockedProductService) Size(size int) *SimpleEarnListLockedProductService {
	s.size = size
	return s
}

type SimpleEarnLockedProductResp struct {
	Rows  []SimpleEarnLockedProduct `json:"rows"`
	Total int                       `json:"total"`
}

type SimpleEarnLockedProduct struct {
	ProjectId string `json:"projectId"`
	Detail    struct {
		Asset                 string `json:"asset"`
		RewardAsset           string `json:"rewardAsset"`
		Duration              int    `json:"duration"`
		Renewable             bool   `json:"renewable"`
		IsSoldOut             bool   `json:"isSoldOut"`
		Apr                   string `json:"apr"`
		Status                string `json:"status"`
		SubscriptionStartTime int64  `json:"subscriptionStartTime"`
		ExtraRewardAsset      string `json:"extraRewardAsset"`
		ExtraRewardAPR        string `json:"extraRewardAPR"`
	} `json:"detail"`
	Quota struct {
		TotalPersonalQuota string `json:"totalPersonalQuota"`
		Minimum            string `json:"minimum"`
	} `json:"quota"`
}

func (s *SimpleEarnListLockedProductService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnLockedProductResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/locked/list",
		secType:  secTypeSigned,
	}

	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	if s.current != 0 {
		r.setParam("current", s.current)
	}
	if s.size != 0 {
		r.setParam("size", s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnLockedProductResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnGetFlexiblePositionService struct {
	c         *Client
	asset     string
	productId string
	current   int
	size      int
}

func (s *SimpleEarnGetFlexiblePositionService) Asset(asset string) *SimpleEarnGetFlexiblePositionService {
	s.asset = asset
	return s
}

func (s *SimpleEarnGetFlexiblePositionService) ProductId(productId string) *SimpleEarnGetFlexiblePositionService {
	s.productId = productId
	return s
}

func (s *SimpleEarnGetFlexiblePositionService) Current(current int) *SimpleEarnGetFlexiblePositionService {
	s.current = current
	return s
}

func (s *SimpleEarnGetFlexiblePositionService) Size(size int) *SimpleEarnGetFlexiblePositionService {
	s.size = size
	return s
}

type SimpleEarnFlexiblePositionResp struct {
	Rows  []SimpleEarnFlexiblePosition `json:"rows"`
	Total int                          `json:"total"`
}

type SimpleEarnFlexiblePosition struct {
	TotalAmount                    string            `json:"totalAmount"`
	TierAnnualPercentageRate       map[string]string `json:"tierAnnualPercentageRate"`
	LatestAnnualPercentageRate     string            `json:"latestAnnualPercentageRate"`
	YesterdayAirdropPercentageRate string            `json:"yesterdayAirdropPercentageRate"`
	Asset                          string            `json:"asset"`
	AirDropAsset                   string            `json:"airDropAsset"`
	CanRedeem                      bool              `json:"canRedeem"`
	CollateralAmount               string            `json:"collateralAmount"`
	ProductId                      string            `json:"productId"`
	YesterdayRealTimeRewards       string            `json:"yesterdayRealTimeRewards"`
	CumulativeBonusRewards         string            `json:"cumulativeBonusRewards"`
	CumulativeRealTimeRewards      string            `json:"cumulativeRealTimeRewards"`
	CumulativeTotalRewards         string            `json:"cumulativeTotalRewards"`
	AutoSubscribe                  bool              `json:"autoSubscribe"`
}

func (s *SimpleEarnGetFlexiblePositionService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexiblePositionResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/position",
		secType:  secTypeSigned,
	}

	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	if s.productId != "" {
		r.setParam("productId", s.productId)
	}
	if s.current != 0 {
		r.setParam("current", s.current)
	}
	if s.size != 0 {
		r.setParam("size", s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnFlexiblePositionResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnGetLockedPositionService struct {
	c          *Client
	asset      string
	positionId int
	projectId  string
	current    int64
	size       int64
}

func (s *SimpleEarnGetLockedPositionService) Asset(asset string) *SimpleEarnGetLockedPositionService {
	s.asset = asset
	return s
}

func (s *SimpleEarnGetLockedPositionService) PositionId(positionId int) *SimpleEarnGetLockedPositionService {
	s.positionId = positionId
	return s
}

func (s *SimpleEarnGetLockedPositionService) ProjectId(projectId string) *SimpleEarnGetLockedPositionService {
	s.projectId = projectId
	return s
}

func (s *SimpleEarnGetLockedPositionService) Current(current int64) *SimpleEarnGetLockedPositionService {
	s.current = current
	return s
}

func (s *SimpleEarnGetLockedPositionService) Size(size int64) *SimpleEarnGetLockedPositionService {
	s.size = size
	return s
}

type SimpleEarnLockedPositionResp struct {
	Rows  []SimpleEarnLockedPosition `json:"rows"`
	Total int                        `json:"total"`
}

type SimpleEarnLockedPosition struct {
	PositionId            int    `json:"positionId"`
	ParentPositionId      int    `json:"parentPositionId"`
	ProjectId             string `json:"projectId"`
	Asset                 string `json:"asset"`
	Amount                string `json:"amount"`
	PurchaseTime          int64  `json:"purchaseTime"`
	Duration              string `json:"duration"`
	AccrualDays           string `json:"accrualDays"`
	RewardAsset           string `json:"rewardAsset"`
	APY                   string `json:"APY"`
	RewardAmt             string `json:"rewardAmt"`
	ExtraRewardAsset      string `json:"extraRewardAsset"`
	ExtraRewardAPR        string `json:"extraRewardAPR"`
	EstExtraRewardAmt     string `json:"estExtraRewardAmt"`
	NextPay               string `json:"nextPay"`
	NextPayDate           int64  `json:"nextPayDate"`
	PayPeriod             string `json:"payPeriod"`
	RedeemAmountEarly     string `json:"redeemAmountEarly"`
	RewardsEndDate        int64  `json:"rewardsEndDate"`
	DeliverDate           int64  `json:"deliverDate"`
	RedeemPeriod          string `json:"redeemPeriod"`
	RedeemingAmt          string `json:"redeemingAmt"`
	RedeemTo              string `json:"redeemTo"`
	PartialAmtDeliverDate int64  `json:"partialAmtDeliverDate"`
	CanRedeemEarly        bool   `json:"canRedeemEarly"`
	CanFastRedemption     bool   `json:"canFastRedemption"`
	AutoSubscribe         bool   `json:"autoSubscribe"`
	Type                  string `json:"type"`
	Status                string `json:"status"`
	CanReStake            bool   `json:"canReStake"`
}

func (s *SimpleEarnGetLockedPositionService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnLockedPositionResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/locked/position",
		secType:  secTypeSigned,
	}

	if s.asset != "" {
		r.setParam("asset", s.asset)
	}
	if s.positionId != 0 {
		r.setParam("positionId", s.positionId)
	}
	if s.projectId != "" {
		r.setParam("projectId", s.projectId)
	}
	if s.current != 0 {
		r.setParam("current", s.current)
	}
	if s.size != 0 {
		r.setParam("size", s.size)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnLockedPositionResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnGetFlexibleQuotaService struct {
	c         *Client
	productId string
}

func (s *SimpleEarnGetFlexibleQuotaService) ProductId(productId string) *SimpleEarnGetFlexibleQuotaService {
	s.productId = productId
	return s
}

type SimpleEarnFlexiblePersonalLeftQuotaResp struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}

func (s *SimpleEarnGetFlexibleQuotaService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexiblePersonalLeftQuotaResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/flexible/personalLeftQuota",
		secType:  secTypeSigned,
	}

	if s.productId != "" {
		r.setParam("productId", s.productId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnFlexiblePersonalLeftQuotaResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnGetLockedQuotaService struct {
	c         *Client
	projectId string
}

func (s *SimpleEarnGetLockedQuotaService) ProjectId(projectId string) *SimpleEarnGetLockedQuotaService {
	s.projectId = projectId
	return s
}

type SimpleEarnLockedPersonalLeftQuotaResp struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}

func (s *SimpleEarnGetLockedQuotaService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnLockedPersonalLeftQuotaResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/locked/personalLeftQuota",
		secType:  secTypeSigned,
	}

	if s.projectId != "" {
		r.setParam("projectId", s.projectId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnLockedPersonalLeftQuotaResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnSubscribeSourceAccount string

const (
	SourceAccountSpot SimpleEarnSubscribeSourceAccount = "SPOT"
	SourceAccountFund SimpleEarnSubscribeSourceAccount = "FUND"
	SourceAccountAll  SimpleEarnSubscribeSourceAccount = "ALL"
)

type SimpleEarnSubscribeFlexibleProductService struct {
	c             *Client
	productId     string
	amount        string
	autoSubscribe bool
	sourceAccount SimpleEarnSubscribeSourceAccount
}

func (s *SimpleEarnSubscribeFlexibleProductService) ProductId(productId string) *SimpleEarnSubscribeFlexibleProductService {
	s.productId = productId
	return s
}

func (s *SimpleEarnSubscribeFlexibleProductService) Amount(amount string) *SimpleEarnSubscribeFlexibleProductService {
	s.amount = amount
	return s
}

func (s *SimpleEarnSubscribeFlexibleProductService) AutoSubscribe(autoSubscribe bool) *SimpleEarnSubscribeFlexibleProductService {
	s.autoSubscribe = autoSubscribe
	return s
}

func (s *SimpleEarnSubscribeFlexibleProductService) SourceAccount(sourceAccount SimpleEarnSubscribeSourceAccount) *SimpleEarnSubscribeFlexibleProductService {
	s.sourceAccount = sourceAccount
	return s
}

type SimpleEarnSubscribeFlexibleProductResp struct {
	PurchaseId int  `json:"purchaseId"`
	Success    bool `json:"success"`
}

func (s *SimpleEarnSubscribeFlexibleProductService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnSubscribeFlexibleProductResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/simple-earn/flexible/subscribe",
		secType:  secTypeSigned,
	}

	if s.productId != "" {
		r.setParam("productId", s.productId)
	}
	if s.amount != "" {
		r.setParam("amount", s.amount)
	}
	if s.autoSubscribe {
		r.setParam("autoSubscribe", s.autoSubscribe)
	}
	if s.sourceAccount != "" {
		r.setParam("sourceAccount", string(s.sourceAccount))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnSubscribeFlexibleProductResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnSubscribeLockedProductService struct {
	c             *Client
	projectId     string
	amount        string
	autoSubscribe bool
	sourceAccount SimpleEarnSubscribeSourceAccount
	redeemTo      string
}

func (s *SimpleEarnSubscribeLockedProductService) ProjectId(projectId string) *SimpleEarnSubscribeLockedProductService {
	s.projectId = projectId
	return s
}

func (s *SimpleEarnSubscribeLockedProductService) Amount(amount string) *SimpleEarnSubscribeLockedProductService {
	s.amount = amount
	return s
}

func (s *SimpleEarnSubscribeLockedProductService) AutoSubscribe(autoSubscribe bool) *SimpleEarnSubscribeLockedProductService {
	s.autoSubscribe = autoSubscribe
	return s
}

func (s *SimpleEarnSubscribeLockedProductService) SourceAccount(sourceAccount SimpleEarnSubscribeSourceAccount) *SimpleEarnSubscribeLockedProductService {
	s.sourceAccount = sourceAccount
	return s
}

func (s *SimpleEarnSubscribeLockedProductService) RedeemTo(redeemTo string) *SimpleEarnSubscribeLockedProductService {
	s.redeemTo = redeemTo
	return s
}

type SimpleEarnSubscribeLockedProductResp struct {
	PurchaseId int  `json:"purchaseId"`
	PositionId int  `json:"positionId"`
	Success    bool `json:"success"`
}

func (s *SimpleEarnSubscribeLockedProductService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnSubscribeLockedProductResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/simple-earn/locked/subscribe",
		secType:  secTypeSigned,
	}

	if s.projectId != "" {
		r.setParam("projectId", s.projectId)
	}
	if s.amount != "" {
		r.setParam("amount", s.amount)
	}
	if s.autoSubscribe {
		r.setParam("autoSubscribe", s.autoSubscribe)
	}
	if s.sourceAccount != "" {
		r.setParam("sourceAccount", string(s.sourceAccount))
	}
	if s.redeemTo != "" {
		r.setParam("redeemTo", s.redeemTo)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SimpleEarnSubscribeLockedProductResp)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
