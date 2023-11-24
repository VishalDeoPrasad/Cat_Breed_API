package repo

import (
	"fmt"
	"math/big"
	"math/rand"
	"net/smtp"
	"time"
)

// GenerateOTP generates a random 6-digit OTP with a provided seed.
func GenerateOTP(seed int64) (string, error) {
	rand.Seed(seed)
	otp := rand.Intn(1000000)
	return fmt.Sprintf("%06d", otp), nil
}



// GenerateOTP generates a random 6-digit OTP.
func GenerateOTP() (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp, nil
}
func main() {
	// Specify a seed (for example, based on the current Unix time)
	seed := time.Now().UnixNano()

	// Generate an OTP with the specified seed
	otp, err := repo.GenerateOTP(seed)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}

	fmt.Printf("Generated OTP with seed %d: %s\n", seed, otp)
}

// SendOTPByEmail sends the generated OTP to the specified email address.
func SendOTPByEmail(email, otp string) error {
	// Sender's email address and password
	from := "vishal.prasad2009@gmail.com"
	password := "oxwr hxxl uegs oqyt"

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Message content with the generated OTP
	message := []byte(fmt.Sprintf("Subject: OTP for Login\n\nYour OTP is: %s", otp))

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{email}, message)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func main() {
	// Example usage
	otp, err := GenerateOTP()
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}

	err = SendOTPByEmail("tyagivikalp99@gmail.com", otp)
	if err != nil {
		fmt.Println("Error sending OTP:", err)
		return
	}

	fmt.Println("OTP sent successfully!")
}
