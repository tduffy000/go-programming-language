package convert

import "fmt"

type Mile float64
type Foot float64
type Yard float64
type Meter float64

func (m Mile) String() string  { return fmt.Sprintf("%g Miles", m) }
func (f Foot) String() string  { return fmt.Sprintf("%g Feet", f) }
func (y Yard) String() string  { return fmt.Sprintf("%g Yards", y) }
func (m Meter) String() string { return fmt.Sprintf("%g Meter", m) }

func FootToMile(f Foot) Mile   { return Mile(f / 5280) }
func MileToFoot(m Mile) Foot   { return Foot(m * 5280) }
func FootToMeter(f Foot) Meter { return Meter(f / 3.28) }
func MeterToFoot(m Meter) Foot { return Foot(m * 3.28) }
func FootToYard(f Foot) Yard   { return Yard(f / 3) }
func YardToFoot(y Yard) Foot   { return Foot(y * 3) }
