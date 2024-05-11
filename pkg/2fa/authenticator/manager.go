package authenticator

import (
	"fmt"
	"io"
	"net/http"
)

type TwoFaManager interface {
	GeneratePairingLink(appInfo string, secretCode string) string
	СheckTwoFactorPin(pin int, secretCode string) (bool, error)
}

type Manager struct {
	appName string
}

func NewManager(appName string) *Manager {
	return &Manager{
		appName: appName,
	}
}

func (m *Manager) GeneratePairingLink(appInfo string, secretCode string) string {
	link := fmt.Sprintf("https://www.authenticatorapi.com/pair.aspx?AppName=%s&AppInfo=%s&SecretCode=%s",
		m.appName, appInfo, secretCode)
	return link
}

func (m *Manager) СheckTwoFactorPin(pin int, secretCode string) (bool, error) {
	if err := validatePin(pin); err != nil {
		return false, err
	}

	link := fmt.Sprintf("https://www.authenticatorApi.com/Validate.aspx?Pin=%v&SecretCode=%v",
		pin, secretCode)

	resp, err := http.Get(link)
	if err != nil {
		return false, fmt.Errorf("sending a request: %v", err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("read data: %v", err)
	}

	if string(data) == "True" {
		return true, nil
	} else {
		return false, nil
	}
}

func validatePin(pin int) error {
	if pin < 100000 && pin > 999999 {
		return fmt.Errorf("invalid pin")
	}

	return nil
}
