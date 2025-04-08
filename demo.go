package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Block struct {
	Timestamp    time.Time
	URL          string
	IP           string
	Hash         string
	PreviousHash string
}

// ANSI color codes
const (
	Reset  = "\033[0m"
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Red    = "\033[31m"
	Bold   = "\033[1m"
)

var blockchain []Block // Global blockchain variable

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Welcome message
	fmt.Printf("%s╔════════════════════════════════════╗%s\n", Cyan, Reset)
	fmt.Printf("%s║   Welcome to the Blockchain CLI    ║%s\n", Cyan, Reset)
	fmt.Printf("%s╚════════════════════════════════════╝%s\n", Cyan, Reset)
	fmt.Printf("%s✓ Track URLs with style! Type 'exit' to quit.%s\n\n", Green, Reset)

	var prevHash string

	for {
		// Cool prompt
		fmt.Printf("%s➤ %sEnter URL:%s ", Blue, Bold, Reset)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Printf("\n%s⚡ Shutting down the blockchain node... Goodbye!%s\n", Yellow, Reset)
			break
		}

		// Parse and validate URL
		parsedURL, err := url.Parse(input)
		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			fmt.Printf("%s✗ Invalid URL! Include scheme (e.g., https://) and try again.%s\n", Red, Reset)
			continue
		}

		// Open browser with animation
		fmt.Printf("%s🚀 Launching %s%s", Yellow, Cyan, parsedURL.String())
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Print(".")
		}
		fmt.Println(Reset)
		err = openBrowser(parsedURL.String())
		if err != nil {
			fmt.Printf("%s✗ Failed to open browser: %v%s\n", Red, err, Reset)
			continue
		}

		// Trace IP with a spinner effect
		fmt.Printf("%s🔍 Tracing IP", Blue)
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Print(".")
		}
		ips, err := net.LookupIP(parsedURL.Hostname())
		if err != nil || len(ips) == 0 {
			fmt.Printf("\n%s✗ IP trace failed: %v%s\n", Red, err, Reset)
			continue
		}
		ipStr := ips[0].String()
		fmt.Printf("\n%s✓ Resolved IP: %s%s\n", Green, ipStr, Reset)

		// Log to blockchain with style
		block := createBlock(parsedURL.String(), ipStr, prevHash)
		blockchain = append(blockchain, block) // This is line 94 or near it
		prevHash = block.Hash

		// Fancy block display with safe hash slicing
		hashDisplay := block.Hash
		prevHashDisplay := block.PreviousHash
		if len(block.Hash) > 16 {
			hashDisplay = block.Hash[:16]
		}
		if len(block.PreviousHash) > 16 {
			prevHashDisplay = block.PreviousHash[:16]
		}

		fmt.Printf("%s┌──[ New Blockchain Event ]──────────────%s\n", Cyan, Reset)
		fmt.Printf("│ %sURL:%s %s%s\n", Bold, Reset, Green, block.URL)
		fmt.Printf("│ %sIP:%s %s%s\n", Bold, Reset, Green, block.IP)
		fmt.Printf("│ %sTimestamp:%s %s%s\n", Bold, Reset, Yellow, block.Timestamp.Format("Mon, 02 Jan 2006 15:04:05 MST"))
		fmt.Printf("│ %sHash:%s %s%s...%s\n", Bold, Reset, Blue, hashDisplay, Reset)
		fmt.Printf("│ %sPrev Hash:%s %s%s...%s\n", Bold, Reset, Blue, prevHashDisplay, Reset)
		fmt.Printf("%s└───────────────────────────────────────%s\n", Cyan, Reset)
		fmt.Printf("%s✓ Block added successfully! (Chain length: %d)%s\n\n", Green, len(blockchain), Reset)
	}
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func createBlock(url, ip, prevHash string) Block {
	timestamp := time.Now()
	data := fmt.Sprintf("%s|%s|%s|%s", timestamp.String(), url, ip, prevHash)
	hash := sha256.Sum256([]byte(data))
	return Block{
		Timestamp:    timestamp,
		URL:          url,
		IP:           ip,
		Hash:         hex.EncodeToString(hash[:]),
		PreviousHash: prevHash,
	}
}
