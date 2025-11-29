package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"

	"github.com/taalt/api-gravekeeper/internal/blame"
	"github.com/taalt/api-gravekeeper/internal/logs"
	"github.com/taalt/api-gravekeeper/internal/scanner"
)

var (
	codeDir string
	logFile string
)

// rootCmd is the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-gravekeeper",
	Short: "Find zombie and shadow API routes",
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan code and logs to detect unused and undocumented routes",
	RunE: func(cmd *cobra.Command, args []string) error {
		if codeDir == "" {
			return fmt.Errorf("--code-dir is required")
		}
		if logFile == "" {
			return fmt.Errorf("--log-file is required")
		}

		absCodeDir, err := filepath.Abs(codeDir)
		if err != nil {
			return err
		}

		definedRoutes, err := scanner.FindDefinedRoutes(absCodeDir)
		if err != nil {
			return err
		}

		accessedRoutes, err := logs.ParseAccessLogs(logFile)
		if err != nil {
			return err
		}

		// Convert defined routes to map for efficient lookup by path
		definedMap := make(map[string]*scanner.RouteInfo)
		for i := range definedRoutes {
			definedMap[definedRoutes[i].Path] = &definedRoutes[i]
		}

		var zombies []*scanner.RouteInfo
		for i := range definedRoutes {
			if !accessedRoutes[definedRoutes[i].Path] {
				zombies = append(zombies, &definedRoutes[i])
			}
		}

		var shadows []string
		for r := range accessedRoutes {
			if _, found := definedMap[r]; !found {
				shadows = append(shadows, r)
			}
		}

		// Sort zombies by path for consistent output
		sort.Slice(zombies, func(i, j int) bool {
			return zombies[i].Path < zombies[j].Path
		})
		sort.Strings(shadows)

		fmt.Println("Zombies (defined but never accessed):")
		if len(zombies) == 0 {
			fmt.Println("  None found")
		} else {
			for _, z := range zombies {
				blameInfo, err := blame.GetBlameData(z.FilePath, z.LineNumber)
				blameStr := "Unknown"
				if err == nil && blameInfo != nil {
					blameStr = blame.FormatBlameInfo(blameInfo)
				}
				fmt.Printf("  %s (%s:%d) - %s\n", z.Path, z.FilePath, z.LineNumber, blameStr)
			}
		}

		fmt.Println()
		fmt.Println("Shadows (accessed but not defined):")
		if len(shadows) == 0 {
			fmt.Println("  None found")
		} else {
			for _, s := range shadows {
				fmt.Printf("  %s\n", s)
			}
		}

		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&codeDir, "code-dir", "c", "", "Root directory to scan for code (required)")
	scanCmd.Flags().StringVarP(&logFile, "log-file", "l", "", "Path to access log file (required)")
	scanCmd.MarkFlagRequired("code-dir")
	scanCmd.MarkFlagRequired("log-file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
