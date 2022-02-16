module helloworld

go 1.14

require (
	github.com/deanwong/learngo/calculator v0.0.0
	github.com/deanwong/learngo/geometry v0.0.0
	rsc.io/quote v1.5.2
)

replace github.com/deanwong/learngo/calculator => ../calculator
replace github.com/deanwong/learngo/geometry => ../geometry
