package sttp

import (
	"time"

	"github.com/google/uuid"
)

type Guid uuid.UUID

var EmptyGuid Guid = Guid(uuid.Nil)

func NewGuid() Guid {
	return Guid(uuid.New())
}

type Ticks int64

const TicksMin Ticks = 0                   // 01/01/0001 00:00:00.000
const TicksMax Ticks = 3155378975999999999 // 12/31/1999 11:59:59.999
const TicksPerSecond Ticks = 10000000
const UnixBaseOffset Ticks = 621355968000000000

func ToTime(ticks Ticks) time.Time {
	return time.Unix(0, int64(ticks-UnixBaseOffset)*100)
}

func ToTicks(time time.Time) Ticks {
	return Ticks(time.UnixNano()/100) + UnixBaseOffset
}
