package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"text/tabwriter"
)

func list(c *cli.Context) error {
	manifest, _ := requestVersionManifest()

	listReleases := c.Bool("release")
	listSnapshots := c.Bool("snapshot")
	listBetaVersions := c.Bool("beta")
	listAlphaVersions := c.Bool("alpha")

	if !listReleases && !listSnapshots && !listBetaVersions && !listAlphaVersions {
		listReleases = true
	}

	fmt.Printf("Latest versions:\n")
	fmt.Printf("\tRelease: %s\n", manifest.Latest.Release)
	fmt.Printf("\tSnapshot: %s\n", manifest.Latest.Snapshot)
	fmt.Print("\nAll versions:\n")

	typeFilter := map[string]bool{
		"release":   listReleases,
		"snapshot":  listSnapshots,
		"old_beta":  listBetaVersions,
		"old_alpha": listAlphaVersions,
	}

	var versions []VersionEntry
	for _, val := range manifest.Versions {
		if typeFilter[val.Type] {
			versions = append(versions, val)
		}
	}

	if c.Bool("description") {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		_, _ = fmt.Fprintln(w, "ID\tRELEASE TIME\tTYPE")

		for _, val := range versions {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\n", val.Id, val.ReleaseTime, val.Type)
		}

		err := w.Flush()
		if err != nil {
			return err
		}
	} else {
		for _, val := range versions {
			fmt.Printf("\t%s\n", val.Id)
		}
	}

	return nil
}
