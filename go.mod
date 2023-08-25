module github.com/blugnu/unilog4logrus

go 1.18

retract [v1.0.0, v1.1.2] // released with incorrect module name and/or incorrect retractions

require (
	github.com/blugnu/go-logspy v0.1.1
	github.com/blugnu/unilog v1.1.3
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/blugnu/go-errorcontext v0.1.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
)
