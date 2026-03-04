package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	version             = "0.576.0"
	iconsURL            = "https://github.com/lucide-icons/lucide/releases/download/" + version + "/lucide-icons-" + version + ".zip"
	tempDir             = "./tmp"
	iconsOutputFilePath = "./pkg/ui/lucide/lucide.go"
)

func main() {
	if err := downloadAndExtract(); err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading/extracting: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	functions, err := generateFunctions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating functions: %v\n", err)
		os.Exit(1)
	}

	if err := writeOutput(functions); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated lucide.go with", len(functions), "icons")
}

func downloadAndExtract() error {
	resp, err := http.Get(iconsURL)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("failed to read zip: %w", err)
	}

	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}

	for _, file := range reader.File {
		if filepath.Ext(file.Name) == "" {
			continue
		}

		outFile, err := os.Create(filepath.Join(tempDir, filepath.Base(file.Name)))
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", file.Name, err)
		}

		rc, err := file.Open()
		if err != nil {
			rc.Close()
			return fmt.Errorf("failed to open zip entry %s: %w", file.Name, err)
		}

		_, err = io.Copy(outFile, rc)
		rc.Close()
		outFile.Close()
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", file.Name, err)
		}
	}

	return nil
}

func generateFunctions() ([]string, error) {
	var functions []string

	entries, err := os.ReadDir(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read temp dir: %w", err)
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".svg") {
			continue
		}

		svgPath := filepath.Join(tempDir, entry.Name())
		svgData, err := os.ReadFile(svgPath)
		if err != nil {
			continue
		}

		svgContent := string(svgData)
		children := extractSVGChildren(svgContent)
		if children == "" {
			continue
		}

		name := filepath.Base(svgPath)
		name = strings.TrimSuffix(name, ".svg")
		funcName := toCamelCase(name)

		functions = append(functions, fmt.Sprintf(`func %s(extraContent ...html.HTML) html.HTML {
	return svg(extraContent, html.Raw(%q))
}`, funcName, children))
	}

	return functions, nil
}

func extractSVGChildren(svg string) string {
	start := strings.Index(svg, "<svg")
	if start == -1 {
		return ""
	}
	end := strings.Index(svg[start:], ">")
	if end == -1 {
		return ""
	}
	start += end + 1

	closeIdx := strings.Index(svg[start:], "</svg>")
	if closeIdx == -1 {
		return ""
	}

	return svg[start : start+closeIdx]
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "-")
	var result string
	for _, part := range parts {
		result += strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
	}
	return result
}

func writeOutput(functions []string) error {
	dir := filepath.Dir(iconsOutputFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	var sb strings.Builder
	sb.WriteString(lucideFile)
	sb.WriteString("\n")
	for _, fn := range functions {
		sb.WriteString(fn)
		sb.WriteString("\n")
	}

	return os.WriteFile(iconsOutputFilePath, []byte(sb.String()), 0644)
}

var lucideFile = `
package lucide

import html "github.com/namzug16/gotags"

func svg(extraContent []html.HTML, paths html.HTML) html.HTML {
	return html.Svg(
		html.X.Attr("xmlns" ,"http://www.w3.org/2000/svg"),
		html.X.Width("24"),
		html.X.Height("24"),
		html.X.Attr("viewBox", "0 0 24 24"),
		html.X.Attr("fill","none"),
		html.X.Attr("stroke","currentColor"),
		html.X.Attr("stroke-width", "2"),
		html.X.Attr("stroke-linecap", "round"),
		html.X.Attr("stroke-linejoin", "round"),
		extraContent,
		paths,
	)
}
`
