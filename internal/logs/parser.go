package logs

import (
    "bufio"
    "os"
    "strings"

    "github.com/taalt/api-gravekeeper/internal/utils"
)

// ParseAccessLogs reads an access log and returns a map of unique accessed paths.
func ParseAccessLogs(logPath string) (map[string]bool, error) {
    f, err := os.Open(logPath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    accessed := make(map[string]bool)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()

        // Find the first quoted request section: "METHOD /path?qs HTTP/1.1"
        first := strings.Index(line, "\"")
        if first == -1 {
            // fallback: try to split fields and find token starting with GET/POST
            fields := strings.Fields(line)
            for i, t := range fields {
                if strings.HasPrefix(t, "GET") || strings.HasPrefix(t, "POST") || strings.HasPrefix(t, "PUT") || strings.HasPrefix(t, "DELETE") || strings.HasPrefix(t, "PATCH") {
                    if i+1 < len(fields) {
                        path := fields[i+1]
                        path = strings.SplitN(path, "?", 2)[0]
                        path = utils.NormalizeLogPath(path)
                        accessed[path] = true
                    }
                    break
                }
            }
            continue
        }
        second := strings.Index(line[first+1:], "\"")
        if second == -1 {
            continue
        }
        req := line[first+1 : first+1+second]
        parts := strings.SplitN(req, " ", 3)
        if len(parts) < 2 {
            continue
        }
        path := parts[1]
        path = strings.SplitN(path, "?", 2)[0]
        path = utils.NormalizeLogPath(path)
        accessed[path] = true
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return accessed, nil
}
