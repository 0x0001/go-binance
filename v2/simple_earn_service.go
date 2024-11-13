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