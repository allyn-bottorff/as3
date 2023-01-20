# GoAS3Parse

Go AS3 parsing module

## Status

This project currently support parsing a raw AS3 declaration into a list of Go
structs. It can parse a single declaration into a list of Tenant, Application,
and Virtual Server objects, including as many sub properties/objects as
possible. Focus is on the AS3 objects typically seen in VU F5 configurations,
so some edge cases might not be parsed.


## Usage

```go
import (
	"encoding/json"
	"os"
	as3parse "gitlab.redchimney.com/publicpackages/goas3parse"

)


// Read in a JSON object file and store the contents as a map[string]interface{}
func readJson(s string) map[string]interface{} {
	f, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("Failed to read AS3 JSON file.")
	}

	jsonMap := make(map[string]interface{})

	json.Unmarshal(f, &jsonMap)

	return jsonMap

}

func main() {

	// Create a map[string]interface{} 
	jsonMap := readJson("./path-to-as3.json")

	// Declare an empty AS3 Declaration
	var dec as3parse.Declaration	

	// Fill the AS3 Declaration with the parsed contents of the AS3 json
	var dec as3parse.ParseDec(jsonMap)

	// Print out basic statistic about the declaration
	dec.Summarize()

}
```

