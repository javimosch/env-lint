
package main
import ("bufio";"fmt";"os";"strings")
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	errors := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line=="" || strings.HasPrefix(line, "#") { continue }
		if !strings.Contains(line, "=") {
			fmt.Fprintf(os.Stderr,"Line %d: Missing = separator\n", lineNum)
			errors++
		} else {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			val := parts[1]
			if key=="" {
				fmt.Fprintf(os.Stderr,"Line %d: Empty key name\n", lineNum)
				errors++
			}
			if val=="" && strings.HasPrefix(scanner.Text(), strings.TrimSpace(scanner.Text())) {
				// warn about empty values
			}
		}
	}
	if errors==0 { fmt.Println("VALID") } else { fmt.Fprintf(os.Stderr,"Found %d error(s)\n", errors); os.Exit(1) }
}

