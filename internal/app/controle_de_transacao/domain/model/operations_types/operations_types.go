package operations_types

const (
	CASH_PUCRCHASE       = 1
	INSTALLMENT_PURCHASE = 2
	WITHDRAW             = 3
	PAYMENT              = 4
)

func IsValidOperationType(operationID int) bool {
	switch operationID {
	case CASH_PUCRCHASE:
		return true
	case INSTALLMENT_PURCHASE:
		return true
	case WITHDRAW:
		return true
	case PAYMENT:
		return true
	default:
		return false
	}
}
