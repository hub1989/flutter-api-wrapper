package utils

import (
	"errors"
	"fmt"
	"math/rand"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// GenerateRef Generate transaction reference
func GenerateRef(referencePrefix string) string {
	length := 10
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return referencePrefix + string(bytes)
}

// VerifyTransactionReference Checks that the transaction reference(TxRef) match
func VerifyTransactionReference(apiTransactionRef, transactionRef interface{}) error {
	if apiTransactionRef != transactionRef {
		return fmt.Errorf(
			"transaction not verified because the transaction reference doesn't match: '%s' != '%s'",
			apiTransactionRef, transactionRef,
		)
	}

	return nil
}

// VerifySuccessMessage The status should equal "success" for a successful transaction
func VerifySuccessMessage(status string) error {
	if status != "success" {
		return errors.New("transaction not verified because status is not equal to 'success'")
	}

	return nil
}

func VerifyChargeResponse(chargeResponse string) error {
	if chargeResponse != "00" && chargeResponse != "0" {
		return errors.New("transaction not verified because the charged response is not equal to '00' or '0'")
	}

	return nil
}

// VerifyCurrencyCode The Currency code must match
func VerifyCurrencyCode(apiCurrencyCode, currencyCode interface{}) error {
	if apiCurrencyCode != currencyCode {
		return fmt.Errorf(
			"transaction not verified because the currency code doesn't match: '%s' != '%s'",
			apiCurrencyCode, currencyCode,
		)
	}

	return nil
}

// VerifyChargedAmount The Charged Amount must be greater than or equal to the paid amount
func VerifyChargedAmount(apiChargedAmount, chargedAmount float64) error {
	if chargedAmount < apiChargedAmount {
		return errors.New("transaction not verified, incorrect amount: charged amount should be greater or equal amount to be paid")
	}

	return nil
}
