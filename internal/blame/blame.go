package blame

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// BlameInfo holds author and date information from git blame.
type BlameInfo struct {
	Author string
	Date   string
}

// GetBlameData executes git blame to retrieve author name and date for a specific line.
// Returns author name, author date, and an error if git fails or the file is not tracked.
func GetBlameData(filePath string, lineNumber int) (*BlameInfo, error) {
	if lineNumber <= 0 {
		return nil, fmt.Errorf("invalid line number: %d", lineNumber)
	}

	lineStr := fmt.Sprintf("%d", lineNumber)
	rangeStr := fmt.Sprintf("%s,%s", lineStr, lineStr)

	cmd := exec.Command("git", "blame", "-L", rangeStr, "--porcelain", filePath)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git blame failed for %s:%d: %v", filePath, lineNumber, err)
	}

	info, err := parsePortcelainOutput(string(output))
	if err != nil {
		return nil, err
	}

	return info, nil
}

// parsePortcelainOutput parses the --porcelain output from git blame.
// Expected format: multi-line with author, author-time, author-tz
func parsePortcelainOutput(output string) (*BlameInfo, error) {
	var author string
	var authorTime int64
	var authorTz string

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "author ") {
			author = strings.TrimPrefix(line, "author ")
		} else if strings.HasPrefix(line, "author-time ") {
			fmt.Sscanf(strings.TrimPrefix(line, "author-time "), "%d", &authorTime)
		} else if strings.HasPrefix(line, "author-tz ") {
			authorTz = strings.TrimPrefix(line, "author-tz ")
		}
	}

	if author == "" {
		return nil, fmt.Errorf("no author found in porcelain output")
	}

	// Convert Unix timestamp to a readable date format
	t := time.Unix(authorTime, 0).UTC()
	dateStr := t.Format("2006-01-02 15:04:05") + " " + authorTz

	return &BlameInfo{
		Author: author,
		Date:   dateStr,
	}, nil
}

// FormatBlameInfo returns a formatted string suitable for console output.
func FormatBlameInfo(info *BlameInfo) string {
	if info == nil {
		return "Unknown"
	}
	return fmt.Sprintf("%s (%s)", info.Author, info.Date)
}
