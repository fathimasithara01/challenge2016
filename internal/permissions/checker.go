package permissions

type Decision int

const (
	NoDecision Decision = iota
	Allow
	Deny
)

type Action int

const (
	Include Action = iota
	Exclude
)

type Line struct {
	Distributor string
	Region      string
}

type Checker struct {
	svc *Service
}

func NewChecker(s *Service) *Checker {
	return &Checker{svc: s}
}

func (c *Checker) Evaluate(l Line) bool {
	return c.svc.Check(l)
}
