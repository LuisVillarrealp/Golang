package conversor

func Euros(usd float64) float64 {
	euros := 0.8782
	return usd * euros
}
func Libras(usd float64) float64 {
	return usd * 0.7557
}
func Won(usd float64) float64 {
	return usd * 1365.97
}
func Btc(usd float64) float64 {
	return usd * 0.00001185
}
