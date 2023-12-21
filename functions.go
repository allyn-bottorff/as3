// Copyright 2023 Allyn L. Bottorff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package as3parse

import (
	"github.com/mitchellh/mapstructure"
)

// Parse declaration-specific fields
func ParseDec(rawDec map[string]interface{}) Declaration {

	var dec Declaration

	mapstructure.Decode(rawDec, &dec)

	for k := range rawDec {
		if val, typeChk := rawDec[k].(map[string]interface{}); typeChk {
			if class, ok := val["class"]; ok {
				if class == "Tenant" {
					dec.Tenants = append(dec.Tenants, ParseTenant(val, k))
				}
			}
		}
	}

	return dec

}

// Parse tenant-specific fields (no nested objects)
func ParseTenant(rawTenant map[string]interface{}, name string) Tenant {

	var tenant Tenant

	mapstructure.Decode(rawTenant, &tenant)
	tenant.Name = name

	for k := range rawTenant {
		if val, typeChk := rawTenant[k].(map[string]interface{}); typeChk {
			if class, ok := val["class"]; ok {
				if class == "Application" {
					tenant.Applications = append(tenant.Applications, ParseApp(val, k))
				}
			}
		}
	}

	return tenant
}

// Parse application-specific fields (no nested objects)
func ParseApp(rawApp map[string]interface{}, name string) Application {
	var app Application

	mapstructure.Decode(rawApp, &app)
	app.Name = name

	for k := range rawApp {
		if val, typeChk := rawApp[k].(map[string]interface{}); typeChk {
			if class, ok := val["class"]; ok {
				switch class {
				case "Service_HTTPS":
					app.VirtualServers = append(app.VirtualServers, ParseVS(val, k))
				case "Service_HTTP":
					app.VirtualServers = append(app.VirtualServers, ParseVS(val, k))
				case "Service_TCP":
					app.VirtualServers = append(app.VirtualServers, ParseVS(val, k))
				case "Service_UDP":
					app.VirtualServers = append(app.VirtualServers, ParseVS(val, k))
				case "Service_L4":
					app.VirtualServers = append(app.VirtualServers, ParseVS(val, k))
				case "Monitor":
					app.Monitors = append(app.Monitors, ParseMon(val, k))
				case "Pool":
					app.Pools = append(app.Pools, ParsePool(val, k))
				}
			}
		}
	}

	return app
}

func ParseVS(rawVS map[string]interface{}, name string) VirtualServer {
	var vs VirtualServer

	mapstructure.Decode(rawVS, &vs)
	vs.Name = name

	return vs
}

func ParseMon(rawMon map[string]interface{}, name string) Monitor {
	var mon Monitor

	mapstructure.Decode(rawMon, &mon)
	mon.Name = name

	return mon
}

func ParsePool(rawPool map[string]interface{}, name string) Pool {
	var pool Pool

	mapstructure.Decode(rawPool, &pool)
	pool.Name = name

	return pool
}
