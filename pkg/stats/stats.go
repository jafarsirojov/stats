package stats

import (
	"github.com/jafarsirojov/bank/v2/pkg/types"
)

//Avg -
func Avg(payments []types.Payment) (money types.Money) {
	k := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			money += payment.Amount
			k++
		}
	}
	return money / types.Money(k)
}

// TotalInCategory -
func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			money += payment.Amount
		}
	}
	return
}

//CategoriesAvg -
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	moneys := map[types.Category]types.Money{}
	quantity := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			moneys[payment.Category] += payment.Amount
			quantity[payment.Category]++
		}
	}

	for key := range moneys {
		moneys[key] /= quantity[key]
	}
	return moneys
}

//PeriodsDynamic - 
func PeriodsDynamic(first, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for key, amount := range second {
		result[key] += amount
	}
	for key, amount := range first {
		result[key] -= amount
	}
	return result
}