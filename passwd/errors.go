package passwd

import (
  "errors"
)

var InvalidFormat = errors.New("incorrect format")
var FieldNotInt = errors.New("unable to convert field to integer")

