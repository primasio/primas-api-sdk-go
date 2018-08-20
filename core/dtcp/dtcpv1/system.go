package dtcpv1

import "github.com/shopspring/decimal"

type SystemParametersGet struct {
	LockAmountContent     decimal.Decimal // Token lock amount for posting content.
	LockPeriodContent     int             // Token lock period in seconds for posting content.
	LockAmountGroupJoin   decimal.Decimal // Token lock amount for joining group.
	LockAmountGroupCreate decimal.Decimal // Token lock amount for creating group.
	ConsumeAmountReport   decimal.Decimal // Token consumed amount for report share.
}
