package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

// RouteInfo holds metadata for a defined API route.
type RouteInfo struct {
	Path       string
	FilePath   string
	LineNumber int
}

// FindDefinedRoutes walks rootPath and returns routes with metadata defined with @app.route("...") in .py files.
func FindDefinedRoutes(rootPath string) ([]RouteInfo, error) {
	var routes []RouteInfo
	routeRe := regexp.MustCompile(`@app\.route\(\s*["']([^"']+)["']`)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".py" {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			matches := routeRe.FindStringSubmatch(line)
			if len(matches) >= 2 {
				routes = append(routes, RouteInfo{
					Path:       matches[1],
					FilePath:   path,
					LineNumber: lineNum,
				})
			}
		}
		return scanner.Err()
	})

	if err != nil {
		return nil, err
	}

	return routes, nil
}
