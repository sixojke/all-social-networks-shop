package payok

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Service interface {
	GetLink(pay *Payment) (string, error)
}

type Client struct {
	shop *shopInfo
}

func NewClient(shopId, successUrl, secretKey string) *Client {
	return &Client{&shopInfo{
		ShopId:     shopId,
		SuccessUrl: successUrl,
		SecretKey:  secretKey,
	}}
}

type Payment struct {
	PaymentId   string
	Amount      float64
	Description string
	Currency    string
}

var paymentIdRegex = regexp.MustCompile("^[a-z0-9_-]{1,36}$")

func (p *Payment) Validate() error {
	if !paymentIdRegex.MatchString(p.PaymentId) {
		return errors.New("invalid payment id")
	}

	return nil
}

type shopInfo struct {
	ShopId     string
	SuccessUrl string
	SecretKey  string
}

type signInp struct {
	PaymentId   string
	Amount      string
	Description string
	Currency    string
	Shop        *shopInfo
}

func (s *Client) GetLink(pay *Payment) (string, error) {
	if err := pay.Validate(); err != nil {
		return "", err
	}

	sign := getSign(&signInp{
		PaymentId:   pay.PaymentId,
		Amount:      fmt.Sprintf("%v", pay.Amount),
		Description: pay.Description,
		Currency:    pay.Currency,
		Shop:        s.shop,
	})

	return fmt.Sprintf("https://payok.io/pay?amount=%v&payment=%v&shop=%v&desc=%v&currency=%v&sign=%v&success_url=%v",
		pay.Amount, pay.PaymentId, s.shop.ShopId, pay.Description, pay.Currency, sign, s.shop.SuccessUrl), nil
}

func getSign(inp *signInp) string {
	var sign []string
	sign = append(sign, inp.Amount, inp.PaymentId, inp.Shop.ShopId, inp.Currency, inp.Description, inp.Shop.SecretKey)
	data := []byte(strings.Join(sign, "|"))
	fmt.Println(strings.Join(sign, "|"))

	hash := md5.Sum(data)
	hashInBytes := hash[:]
	return hex.EncodeToString(hashInBytes)
}
