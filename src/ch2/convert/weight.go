package convert

import "fmt"

type Pound float64
type Kilogram float64
type Stone float64

func (p Pound) String() string    { return fmt.Sprintf("%g Pounds", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g Kilogram", k) }
func (s Stone) String() string    { return fmt.Sprintf("%g Stone", s) }

func PoundToKilogram(p Pound) Kilogram { return Kilogram(p / 2.2) }
func KilogramToPound(k Kilogram) Pound { return Pound(k * 2.2) }
func PoundToStone(p Pound) Stone       { return Stone(p / 14) }
func KilogramToStone(k Kilogram) Stone { return PoundToStone(KilogramToPound(k)) }
