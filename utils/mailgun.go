package utils

import (
	"context"
	"log"
	"os"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4" // Bibliothèque pour interagir avec l'API Mailgun
)

// SendOTPEmail : Envoie un email contenant un code OTP au destinataire spécifié
func SendOTPEmail(recipient, otp string) error {
	// Récupère le domaine et la clé API Mailgun depuis les variables d'environnement
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")

	// Initialise un nouvel objet Mailgun avec le domaine et la clé API
	mg := mailgun.NewMailgun(domain, apiKey)

	// Définit les informations de l'email
	sender := "noreply@" + domain           // Adresse email de l'expéditeur
	subject := "Votre code de vérification" // Sujet de l'email
	body := "Votre code OTP est : " + otp   // Corps de l'email contenant le code OTP
	message := mg.NewMessage(sender, subject, body, recipient)

	// Crée un contexte avec un délai d'expiration de 10 secondes pour l'envoi de l'email
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel() // Annule le contexte après l'exécution pour libérer les ressources

	// Envoie l'email via l'API Mailgun
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		log.Println("Erreur lors de l'envoi de l'email :", err) // Log l'erreur en cas d'échec
		return err
	}
	return nil // Retourne nil si l'envoi est réussi
}
