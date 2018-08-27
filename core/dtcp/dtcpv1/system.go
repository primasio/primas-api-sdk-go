/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dtcpv1

import "github.com/shopspring/decimal"

type SystemParametersGet struct {
	LockAmountContent     decimal.Decimal // Token lock amount for posting content.
	LockPeriodContent     int             // Token lock period in seconds for posting content.
	LockAmountGroupJoin   decimal.Decimal // Token lock amount for joining group.
	LockAmountGroupCreate decimal.Decimal // Token lock amount for creating group.
	ConsumeAmountReport   decimal.Decimal // Token consumed amount for report share.
}
