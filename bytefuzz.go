package bytefuzz

type Interface interface {
	Distance(text, target []byte) int
	DistanceString(text, target string) int
}
