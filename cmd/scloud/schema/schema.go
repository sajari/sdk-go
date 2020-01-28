package schema

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	sajari "code.sajari.com/sdk-go"
)

var (
	topLevelCommands = []string{"stats", "add", "get"}
)

type Field struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Type        sajari.FieldType `json:"type"`
	Repeated    bool             `json:"repeated"`
	Mode        sajari.FieldMode `json:"mode"`
}

type FieldIndex struct {
	Spec        string `json:"spec"`
	Description string `json:"description"`
}

type Schema struct {
	Fields []Field `json:"fields"`
}

func Run(client *sajari.Client, args []string) {

	iflags := flag.NewFlagSet("stats", flag.ExitOnError)
	ignoreFields := iflags.String("ignore-fields", "", "list of comma separated fields `field1,field2,...` to ignore")

	if len(args) == 0 {
		fmt.Printf("\nusage: scloud schema <cmd> [options...]\n\n")
		iflags.Usage()
		return
	}
	iflags.Parse(args[1:])

	ignoreFieldsMap := map[string]bool{}
	if *ignoreFields != "" {
		for _, field := range strings.Split(*ignoreFields, ",") {
			ignoreFieldsMap[field] = true
		}
	}

	switch args[0] {
	case "stats":
		if len(args) <= 1 {
			fmt.Printf("usage: scloud schema stats <path> [options...]\n\npath: to file to read JSON schema from")
			return
		}
		fs := getFields(args[1], ignoreFieldsMap)
		fmt.Printf("Stats for schema: %v\n", args[1])
		fmt.Printf("Total fields: %v\n", len(fs))

		fieldsByType := map[string]int{}
		var fixedWidthFields int
		for _, f := range fs {
			fieldsByType[fmt.Sprintf("%v:repeated=%v", f.Type, f.Repeated)]++
			if !f.Repeated {
				switch f.Type {
				case sajari.TypeInteger, sajari.TypeFloat, sajari.TypeBoolean, sajari.TypeDouble, sajari.TypeTimestamp:
					fixedWidthFields++
				}
			}
		}
		fmt.Printf("Fields by type: %v\n", fieldsByType)
		fmt.Printf("Total fixed width fields: %v\n", fixedWidthFields)

	case "add":
		if len(args) <= 1 {
			fmt.Printf("usage: scloud schema add <path> [options...]\n\npath: `path` to file to read JSON schema from\n")
			return
		}
		schema := client.Schema()
		fs := getFields(args[1], ignoreFieldsMap)
		for _, f := range fs {
			if err := schema.CreateField(context.Background(), f); err != nil {
				fmt.Printf("error adding field: %v", err)
				return
			}
		}

	case "get":
		var path string
		if len(args) >= 2 {
			path = args[1]
		}
		schema := client.Schema()
		fields := schema.Fields(context.Background())

		var count int
		var flds []Field
		for {
			f, err := fields.Next()
			if err != nil {
				break
			}
			count++
			if !ignoreFieldsMap[f.Name] {
				flds = append(flds, Field{
					Name:        f.Name,
					Description: f.Description,
					Type:        f.Type,
					Repeated:    f.Repeated,
					Mode:        f.Mode,
				})
			}
		}
		if count == 0 {
			fmt.Printf("No fields for collection: %v/%v\n", client.Project, client.Collection)
		}

		sch := Schema{
			Fields: flds,
		}

		b, err := json.MarshalIndent(sch, "", "  ")
		if err != nil {
			fmt.Printf("error marshalling JSON: %v", err)
			return
		}

		var out io.Writer = os.Stdout
		if path != "" {
			f, err := os.Create(path)
			if err != nil {
				fmt.Printf("error creating file for schema: %v", err)
				return
			}
			out = f
			defer f.Close()
		}
		fmt.Fprintf(out, "%s\n", b)
		return

	default:
		fmt.Printf("usage: scloud schema <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
		return
	}

}

func getFields(path string, ignoreFieldsMap map[string]bool) []sajari.Field {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading JSON schema file: %v", err)
		return []sajari.Field{}
	}

	s := Schema{}
	if err := json.Unmarshal(b, &s); err != nil {
		fmt.Printf("error unmarshalling JSON schema file: %v", err)
		return []sajari.Field{}
	}

	var fields []sajari.Field
	for _, f := range s.Fields {
		if !ignoreFieldsMap[f.Name] {

			fields = append(fields, sajari.Field{
				Name:        f.Name,
				Description: f.Description,
				Type:        f.Type,
				Repeated:    f.Repeated,
				Mode:        f.Mode,
			})
		}
	}
	return fields
}
