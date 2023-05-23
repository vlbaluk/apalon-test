package interfaces

import . "math/big"

type Calculate func(x *Float, y *Float) string
type CachedCalculate func(x *Float, y *Float, keyHead string, fn Calculate) (string, bool)
