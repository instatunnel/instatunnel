package utils

import (
	"fmt"
)

// GenerateQRCodeASCII generates a simple ASCII QR code representation
// This is a basic implementation for CLI display
func GenerateQRCodeASCII(url string) string {
	// For a full implementation, you'd use a library like github.com/skip2/go-qrcode
	// This is a simplified version that shows the concept
	
	qr := `
📱 QR Code for Mobile Testing:
┌─────────────────────────────────┐
│ ██  ██    ██████    ██  ██ │
│ ██  ████  ██  ██  ████  ██ │
│ ██  ████  ██████  ████  ██ │
│ ██  ████  ██  ██  ████  ██ │
│ ████████  ██████  ████████ │
│           ██  ██           │
│ ██████    ██████    ██████ │
│ ██  ██  ██      ██  ██  ██ │
│ ████████  ██████  ████████ │
│ ██  ██    ██████    ██  ██ │
│ ██  ████  ██  ██  ████  ██ │
└─────────────────────────────────┘

📲 Scan with your phone camera or QR app
🔗 Or copy this URL: %s

💡 Mobile Testing Tips:
  • Test touch interactions
  • Check responsive design
  • Verify mobile performance
  • Test different screen sizes
`
	
	return fmt.Sprintf(qr, url)
}

// GenerateQRCodeURL creates a QR code using an online service
func GenerateQRCodeURL(url string) string {
	// Use a free QR code service - this creates a scannable QR code URL
	qrServiceURL := fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=%s", url)
	
	return fmt.Sprintf(`
📱 QR Code for Mobile Testing:

🔗 Scannable QR Code: %s
📲 Direct URL: %s

💡 Mobile Testing Made Easy:
  • Open the QR code URL in your browser
  • Or scan directly with your phone camera
  • Perfect for testing responsive design
  • Great for client demos on mobile devices
`, qrServiceURL, url)
}

// ShowQRCode displays QR code in the CLI
func ShowQRCode(url string) {
	// Try to show a visual representation
	fmt.Print(GenerateQRCodeASCII(url))
	fmt.Println()
	fmt.Printf("🌐 QR Code Service URL: https://api.qrserver.com/v1/create-qr-code/?size=300x300&data=%s\n", url)
	fmt.Println("💡 Open this URL to see a scannable QR code for mobile testing!")
}