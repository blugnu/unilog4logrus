module github.com/blugnu/unilog4logrus

go 1.20

retract [v1.0.0, v1.1.2] // released with incorrect module name and/or incorrect retractions

require (
	github.com/blugnu/go-logspy v0.1.1
	github.com/blugnu/unilog v1.1.4
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/blugnu/errorcontext v0.2.2 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
