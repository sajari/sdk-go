package csvcmd

import (
	"context"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	sajari "code.sajari.com/sdk-go"
)

var (
	topLevelCommands = []string{"export", "import"}
)

// Run executes several different csv related options
func Run(client *sajari.Client, args []string) error {
	switch args[0] {
	case "export":
		return Export(client, args)
	case "import":
		return errors.New("csv import not supported yet")
	}
	return fmt.Errorf("usage: scloud csv <%v> [options...]\n", strings.Join(topLevelCommands, "|"))
}

// Export exports a collection to CSV
func Export(client *sajari.Client, args []string) error {
	iflags := flag.NewFlagSet("csv", flag.ExitOnError)
	fields := iflags.String("fields", "", "comma delimited `fields` to export. Empty exports all fields")
	field := iflags.String("field", "", "`field` to scan for keys")
	limit := iflags.Int("limit", 100, "number of records to sample. Use -1 for no limit")
	inputpath := iflags.String("inputpath", "", "`inputpath` to load the keys from")
	path := iflags.String("path", "", "`path` to save the export")
	workers := iflags.Int("workers", 10, "num `workers` to use concurrently")

	if len(args) == 0 {
		defer iflags.Usage()
		return fmt.Errorf("\nusage: scloud csv <%v> [options...]\n\n", strings.Join(topLevelCommands, "|"))
	}
	iflags.Parse(args[1:])

	if *path == "" {
		return fmt.Errorf("`path` cannot be empty")
	}

	f, err := os.Create(*path)
	if err != nil {
		return err
	}
	defer f.Close()
	cw := csv.NewWriter(f)

	ctx := context.Background()

	fs := []string{*field}
	fs = append(fs, strings.Split(*fields, ",")...)

	ch := make(chan *sajari.Key, 10)
	wch := make(chan []string, 10)

	wg := sync.WaitGroup{}
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			for k := range ch {
				rec, err := client.GetRecord(ctx, k)
				if err != nil {
					fmt.Printf("failed to get record (%v): %v\n", k, err)
					continue
				}
				out := make([]string, len(fs))
				out[0] = k.String()
				for i, f := range fs {
					if v, ok := rec[f]; ok {
						out[i] = fmt.Sprintf("%v", v)
					}
				}
				wch <- out
			}
			wg.Done()
		}()
	}
	defer wg.Wait()

	// Write the records to CSV
	go func() {
		cw.Write(fs) // Header
		for {
			for o := range wch {
				cw.Write(o)
			}
			cw.Flush()
			break
		}
	}()

	if *inputpath != "" {
		fmt.Printf("exporting CSV using record keys from: %v\n", *inputpath)
		return csvIterator(ctx, client, *inputpath, *field, ch)
	}
	return keyIterator(ctx, client, *field, *limit, ch)
}

func keyIterator(ctx context.Context, client *sajari.Client, field string, limit int, ch chan *sajari.Key) error {
	it := client.Keys(ctx, field, limit)
	for {
		k, err := it.Next()
		if err == sajari.ErrDone {
			break
		}
		if err != nil {
			return fmt.Errorf("Could not get key: %v", err)
		}
		ch <- k
	}
	return nil
}

func csvIterator(ctx context.Context, client *sajari.Client, path string, field string, ch chan *sajari.Key) error {
	defer close(ch)
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	cr := csv.NewReader(f)
	row, err := cr.Read()
	if err != nil {
		return fmt.Errorf("error reading header row: %v", err)
	}

	if len(row) == 0 {
		return errors.New("csv header row is empty")
	}

	column := 0
	for i, f := range row {
		if f == field {
			fmt.Printf("found header column offset: %v\n", i)
			column = i
		}
	}

	ids := make(map[string]struct{}, 1000) // Detect duplicates

	for {
		row, err := cr.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error reading row: %v", err)
		}
		if len(row) < column {
			fmt.Printf("key column does not exist for this row: %v\n", row)
			continue
		}
		if _, ok := ids[row[column]]; ok {
			continue
		}
		ch <- sajari.NewKey(field, row[column])
		ids[row[column]] = struct{}{}
	}
	return nil
}
