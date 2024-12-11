package bytefuzz

type Interface interface {
	Distance(text, target []byte) float64
	DistanceString(text, target string) float64
}
