package common

import "errors"

var (
	ErrBindingRequest       = errors.New("error binding request")
	ErrCurrencyNotSupported = errors.New("error currency not supported. supported currencies :: %v")
	ErrFiatNotSupported     = errors.New("error binding request")
	ErrProcessingRequest    = errors.New("error processing request")
)
