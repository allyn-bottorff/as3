package as3parse

import (
	"encoding/json"
	"testing"
)

func TestParseTenant(t *testing.T) {
	expectedTenant := Tenant{
		DefaultRouteDomain: 0,
		Enable:             true,
		OptimisticLockKey:  "",
	}

	jsonBytes := []byte(`
	{"class": "tenant",
	"defaultRouteDomain": 0,
	"enable": true,
	"optimisticLockKey": ""
	}
	`)

	jsonMap := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &jsonMap)

	tenant := ParseTenant(jsonMap, "test")

	if tenant.Enable != expectedTenant.Enable {
		t.Fatalf("Failed to match enable: %v -> %v", tenant.Enable, expectedTenant.Enable)
	}

}

func TestParseDec(t *testing.T) {
	controlsMap := make(map[string]string)
	controlsMap["archiveTimestamp"] = "some timestamp"
	expectedDec := Declaration{
		Label:         "AS3 direct deploy",
		Remark:        "HTTP with custom persistence",
		SchemaVersion: "3.12.0",
		UpdateMode:    "selective",
		Controls:      controlsMap,
		Id:            "autogen_2f98bc55-0f8b-4ec2-ade4-efa4e3ef30a5",
	}

	jsonBytes := []byte(`
	{"class": "ADC",
"controls": {
"archiveTimestamp": "some timestamp"
},
"label": "AS3 direct deploy",
"remark": "HTTP with custom persistence",
"schemaVersion": "3.12.0",
"id": "autogen_2f98bc55-0f8b-4ec2-ade4-efa4e3ef30a5",
"updateMode": "selective"}
`)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &jsonMap)

	dec := ParseDec(jsonMap)

	if dec.Label != expectedDec.Label {
		t.Fatalf("Failed to match label: %s -> %s", dec.Label, expectedDec.Label)
	}
	if dec.Remark != expectedDec.Remark {
		t.Fatalf("Failed to match remark")
	}
	if dec.SchemaVersion != expectedDec.SchemaVersion {
		t.Fatalf("Failed to match SchemaVersion")
	}
	if dec.Id != expectedDec.Id {
		t.Fatalf("Failed to match Id")
	}
	if dec.UpdateMode != expectedDec.UpdateMode {
		t.Fatalf("Failed to match UpdateMode")
	}
	if dec.Controls["archiveTimestamp"] != expectedDec.Controls["archiveTimestamp"] {
		t.Fatalf("Failed to match Controls archiveTimestamp: %s -> %s", dec.Controls["archiveTimestamp"], expectedDec.Controls["archiveTimestamp"])
	}

}


func TestPrintDec(t *testing.T) {

	jsonBytes := []byte(`
	{"class": "ADC",
"controls": {
"archiveTimestamp": "some timestamp"
},
"label": "AS3 direct deploy",
"remark": "HTTP with custom persistence",
"schemaVersion": "3.12.0",
"id": "autogen_2f98bc55-0f8b-4ec2-ade4-efa4e3ef30a5",
"updateMode": "selective"}
`)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &jsonMap)
	dec := ParseDec(jsonMap)
	dec.PrintAll()

}
