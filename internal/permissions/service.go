package permissions

import (
	"strings"

	"github.com/fathimasithara01/realimage-challenge-2016/internal/geo"
)

type Rule struct {
	Action Action
	Region string
}

type Service struct {
	perms map[string][]Rule 
}

func NewService() *Service {
	return &Service{
		perms: make(map[string][]Rule),
	}
}

func (s *Service) AddRule(distributor string, action Action, region string) {
	distributor = strings.ToUpper(strings.TrimSpace(distributor))

	s.perms[distributor] = append(s.perms[distributor], Rule{
		Action: action,
		Region: region,
	})
}

func (s *Service) Inherit(child string, parent string) {
	parentRules, ok := s.perms[parent]
	if !ok {
		return
	}
	s.perms[child] = append(parentRules, s.perms[child]...)
}

func (s *Service) Check(line Line) bool {
	disc := strings.ToUpper(line.Distributor)
	rawRegion := line.Region

	rules, ok := s.perms[disc]
	if !ok {
		return false 
	}

	target := geo.ParseRegion(rawRegion)

	decision := NoDecision

	for _, r := range rules {
		cur := geo.ParseRegion(r.Region)

		if geo.Matches(target, cur) {
			if r.Action == Include {
				decision = Allow
			} else {
				decision = Deny
			}
		}
	}

	return decision == Allow
}
