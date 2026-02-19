package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// CopyToClipboard copies text to the system clipboard
func CopyToClipboard(text string) error {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("pbcopy")
	case "linux":
		// Try xclip first, then xsel
		if exec.Command("which", "xclip").Run() == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if exec.Command("which", "xsel").Run() == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return fmt.Errorf("clipboard not available (install xclip or xsel)")
		}
	case "windows":
		cmd = exec.Command("clip")
	default:
		return fmt.Errorf("clipboard not supported on %s", runtime.GOOS)
	}
	
	if cmd == nil {
		return fmt.Errorf("clipboard command not available")
	}
	
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

// TryCopyToClipboard attempts to copy to clipboard, returns success status
func TryCopyToClipboard(text string) bool {
	err := CopyToClipboard(text)
	return err == nil
}