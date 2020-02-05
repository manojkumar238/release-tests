package steps

import (
	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/clients"
	"github.com/openshift-pipelines/release-tests/pkg/tkn"
)

func Namespace() string {
	return gauge.GetScenarioStore()["namespace"].(string)

}

func Clients() *clients.Clients {
	switch cs := gauge.GetScenarioStore()["clients"].(type) {
	case *clients.Clients:
		return cs
	default:
		return nil
	}
}

func OperatorClient() *clients.Clients {
	switch c := gauge.GetSuiteStore()["opclient"].(type) {
	case *clients.Clients:
		return c
	default:
		return nil
	}
}

func Tkn() tkn.Cmd {
	switch n := gauge.GetSuiteStore()["tkn"].(type) {
	case tkn.Cmd:
		return n
	default:
		panic("Error")
	}
}
