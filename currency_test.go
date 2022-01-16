package faker

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	c := New().Currency()
	NotExpect(t, "", c.Currency())
}

func TestCurrencyCode(t *testing.T) {
	c := New().Currency()
	NotExpect(t, "", c.Code())
}

func TestCurrencyNumber(t *testing.T) {
	c := New().Currency()
	NotExpect(t, "", c.Currency())
}

func TestCurrencyCountry(t *testing.T) {
	c := New().Currency()
	NotExpect(t, "", c.Country())
}

func TestCurrencyAndCode(t *testing.T) {
	c := New().Currency()
	currency, code := c.CurrencyAndCode()
	NotExpect(t, "", currency)
	NotExpect(t, "", code)
}
