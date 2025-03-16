package service

var nextGen = 0

func IdGenerator() int {
	nextGen++
	return nextGen
}
