package dtcpv1

import (
	"strings"

	"github.com/shopspring/decimal"
)

const (
	//
	License_Fields_Derivation = "derivation"
	License_Fields_Commercial = "commercial"

	// Whether Derivation is allowed.
	License_Derivative_y  = "y"
	License_Derivative_n  = "n"
	License_Derivative_sa = "sa" // for share-alike

	// Whether commercial usage is allowed
	License_Commercial_y = "y"
	License_Commercial_n = "n"
)

type License struct {
	Name       string           `json:"name"`
	Version    string           `json:"version"`
	Parameters []LicenseSubject `json:"parameters"`
}

type LicenseSubject struct {
	Name  string
	Value string
}

func NewLicenseDerivative(value string) *LicenseSubject {
	value = strings.ToLower(value)
	if value != License_Derivative_y &&
		value != License_Derivative_n &&
		value != License_Derivative_sa {
		return nil
	}

	return &LicenseSubject{
		Name:  "derivative",
		Value: value,
	}
}

func NewLicenseDerivative_y() *LicenseSubject {
	return &LicenseSubject{
		Name:  "derivative",
		Value: License_Derivative_y,
	}
}

func NewLicenseDerivative_n() *LicenseSubject {
	return &LicenseSubject{
		Name:  "derivative",
		Value: License_Derivative_n,
	}
}

func NewLicenseDerivative_sa() *LicenseSubject {
	return &LicenseSubject{
		Name:  "derivative",
		Value: License_Derivative_sa,
	}
}

func NewLicenseCommercial(value string) *LicenseSubject {
	value = strings.ToLower(value)
	if value != License_Commercial_y &&
		value != License_Commercial_n {
		return nil
	}

	return &LicenseSubject{
		Name:  "commercial",
		Value: value,
	}
}

func NewLicenseCommercial_y() *LicenseSubject {
	return &LicenseSubject{
		Name:  "commercial",
		Value: License_Commercial_y,
	}
}

func NewLicenseCommercial_n(value string) *LicenseSubject {
	return &LicenseSubject{
		Name:  "commercial",
		Value: License_Commercial_n,
	}
}

type LicensePrice struct {
	Name  string
	Value decimal.Decimal
}

func NewLicensePST(value decimal.Decimal) *LicensePrice {
	if value.Cmp(decimal.NewFromFloat(0)) == -1 {
		return nil
	}

	return &LicensePrice{
		Name:  "currency",
		Value: value,
	}
}

func NewLicensePrice(value decimal.Decimal) *LicensePrice {
	if value.Cmp(decimal.NewFromFloat(0)) == -1 {
		return nil
	}

	return &LicensePrice{
		Name:  "price",
		Value: value,
	}
}
