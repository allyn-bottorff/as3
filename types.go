package as3parse

import (
	"fmt"
	"encoding/json"
)

type Monitor struct {
	Name        string `mapstructure:"name,omitempty"`
	Ciphers     string `mapstructure:"ciphers"`
	Class       string `mapstructure:"class"`
	Interval    int    `mapstructure:"interval"`
	MonitorType string `mapstructure:"monitorType"`
	Receive     string `mapstructure:"receive"`
	ReceiveDown string `mapstructure:"receiveDown"`
	Send        string `mapstructure:"send"`
	Timeout     string `mapstructure:"timeout"`
}

type Pool struct {
	Name              string              `mapstructure:"name,omitempty"`
	Class             string              `mapstructure:"class"`
	LoadBalancingMode string              `mapstructure:"loadBalancingMode"`
	Members           []Member            `mapstructure:"members"`
	Monitors          []map[string]string `mapstructure:"monitors"`
}

type Member struct {
	AddressDiscovery string `mapstructure:"addressDiscovery"`
	ExternalId       string `mapstructure:"externalId"`
	Hostname         string `mapstructure:"hostname"`
	ServicePort      string `mapstructure:"servicePort"`
}

type VirtualServer struct {
	Pool               Pool
	Name               string
	AllowVlans         []map[string]string `mapstructure:"allowVlans"`
	Class              string              `mapstructure:"class"`
	ClientTLS          map[string]string   `mapstructure:"clientTLS"`
	ProfileTCP         string              `mapstructure:"profileTCP"`
	ProfileHTTP        map[string]string   `mapstructure:"profileHTTP,omitempty"`
	Redirect80         bool                `mapstructure:"redirect80"`
	ServerTLS          map[string]string   `mapstructure:"serverTLS"`
	VirtualAddresses   []string            `mapstructure:"virtualAddresses"`
	VirtualPort        int                 `mapstructure:"virtualPort"`
	PersistenceMethods []string            `mapstructure:"persistenceMethods,omitempty"`
}

// AS3 Application. This is a container for virtual servers and related load
// balancing objects
type Application struct {
	Name           string
	Monitors       []Monitor
	Pools          []Pool
	VirtualServers []VirtualServer
	Template       string `mapstructure:"template"`
}

func (a *Application) CountVS() int {
	return len(a.VirtualServers)
}
func (a *Application) CountMons() int {
	return len(a.Monitors)
}
func (a *Application) CountPools() int {
	return len(a.Pools)
}

// AS3 Tenant, reformatted to use lists of object types instead of individual
// named objects
type Tenant struct {
	Name               string
	Applications       []Application
	DefaultRouteDomain int    `mapstructure:"defaultRouteDomain"`
	Enable             bool   `mapstructure:"enable"`
	OptimisticLockKey  string `mapstructure:"optimisticLockKey"`
}

func (t *Tenant) Summarize() {
	fmt.Printf("  Tenant:\n")
	fmt.Printf("    Name: %s\n", t.Name)
	fmt.Printf("    Apps: %d\n", len(t.Applications))

}

// AS3 Declaration, reformatted to use lists of object types instead of
// individual named objects
type Declaration struct {
	Tenants       []Tenant
	Label         string            `mapstructure:"label"`
	Remark        string            `mapstructure:"remark"`
	SchemaVersion string            `mapstructure:"schemaVersion"`
	Id            string            `mapstructure:"id"`
	UpdateMode    string            `mapstructure:"updateMode"`
	Controls      map[string]string `mapstructure:"controls"`
}

func (dec *Declaration) Summarize() {
	fmt.Printf("Declaration:\n")
	fmt.Printf("  Label: %s\n", dec.Label)
	fmt.Printf("  Remark: %s\n", dec.Remark)
	fmt.Printf("  SchemaVersion: %s\n", dec.SchemaVersion)
	fmt.Printf("  Id: %s\n", dec.Id)
	fmt.Printf("  Controls: %v\n", dec.Controls)
	fmt.Printf("  Tenants: %d\n", len(dec.Tenants))
	for _, t := range dec.Tenants {
		t.Summarize()
	}
}

func (dec *Declaration) PrintVSNames() {
	for _, t := range dec.Tenants {
		for _, a := range t.Applications {
			for _, v := range a.VirtualServers {
				fmt.Printf("%s\n", v.Name)
			}
		}
	}
}

// Print entire parsed declaration to the console in json format
func (dec *Declaration) PrintAll() {
	decBytes, err := json.Marshal(dec)
	if err != nil {
		fmt.Printf(string(decBytes))
	}
}
