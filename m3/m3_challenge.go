// Desafio 3.1: Sistema de Notificações Multi-Canal

/*
*
* O que vais praticar:
*
✅ Definir interfaces
✅ Múltiplas implementações da mesma interface
✅ Polimorfismo (função aceita qualquer Notifier)
✅ Slices de interfaces
✅ Interface composition (usar outras interfaces como fmt.Stringer)
✅ Error handling com múltiplas fontes
✅ Dependency injection pattern
*
*/

package main

import (
	"fmt"
	"strings"
)

type Notifier interface {
	Send(recipient string, message string) error
}

type EmailNotifier struct {
	SMTPServer string
}

type SMSNotifier struct {
	Provider string // ex: "Vodafone
}

type PushNotifier struct {
	AppName string
}

// Send() imprime: "📧 Email para [recipient] via [SMTPServer]: [message]"
// Retorna erro se recipient não contém @
func (e EmailNotifier) Send(recipient string, message string) error {
	if !strings.Contains(recipient, "@") {
		return fmt.Errorf("email inválido: %s", recipient)
	}
	fmt.Printf("📧 Email para %s via %s: %s\n", recipient, e.SMTPServer, message)
	return nil
}

// Send() imprime: "📱 SMS para [recipient] via [Provider]: [message]"
// Retorna erro se recipient não começa com +
func (e SMSNotifier) Send(recipient string, message string) error {
	if !strings.HasPrefix(recipient, "+") {
		return fmt.Errorf("número de telefone inválido: %s", recipient)
	}
	fmt.Printf("📱 SMS para %s via %s: %s\n", recipient, e.Provider, message)
	return nil
}

// Send() imprime: "🔔 Push notification para [recipient] via app [AppName]: [message]"
// Sempre retorna nil (push nunca falha neste exemplo)
func (e PushNotifier) Send(recipient string, message string) error {
	fmt.Printf("🔔 Push notification para %s via app %s: %s\n", recipient, e.AppName, message)
	return nil
}

// Envia para todos os recipients
// Retorna slice com erros (nil se não houve erros)
// Não para se um falhar (continua para os próximos)
func SendBulk(notifier Notifier, recipients []string, message string) []error {
	var errors []error
	for _, recipient := range recipients {
		err := notifier.Send(recipient, message)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return errors
}

type NotificationService struct {
	notifiers []Notifier
}

func (s *NotificationService) AddNotifier(n Notifier) {
	s.notifiers = append(s.notifiers, n)
}

func (s *NotificationService) NotifyAll(recipient, message string) error {
	var errors []error
	for _, notifier := range s.notifiers {
		fmt.Printf("Enviando via %v...\n", notifier)
		err := notifier.Send(recipient, message)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Notificações enviadas com %d erros", len(errors))
	}
	return nil

}

//Nota: Usa value receiver, não pointer! (Convenção para String())

func (n EmailNotifier) String() string {
	return fmt.Sprintf("%T", n.SMTPServer)
}

func (n SMSNotifier) String() string {
	return fmt.Sprintf("%T", n.Provider)
}

func (n PushNotifier) String() string {
	return fmt.Sprintf("%T", n.AppName)
}

func main() {
	email := EmailNotifier{SMTPServer: "smtp.example.com"}
	sms := SMSNotifier{Provider: "Vodafone"}
	push := PushNotifier{AppName: "MyApp"}

	service := NotificationService{}
	service.AddNotifier(email)
	service.AddNotifier(sms)
	service.AddNotifier(push)

	// teste individual
	fmt.Println("Teste individual:")
	email.Send("user@example.com", "Olá, este é um teste de email!")
	sms.Send("+1234567890", "Olá, este é um teste de SMS!")
	push.Send("user123", "Olá, este é um teste de push!")

	// teste bulk
	fmt.Println("\nTeste bulk:")
	recipients := []string{"user@example.com", "+1234567890", "user123"}
	message := "Olá, este é um teste de notificação em massa!"
	for _, notifier := range service.notifiers {
		errors := SendBulk(notifier, recipients, message)
		if errors != nil {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}

	// notify all
	fmt.Println("\nTeste NotifyAll:")
	err := service.NotifyAll("user@example.com", "Olá, este é um teste de notificação para todos!")
	if err != nil {
		fmt.Println(err)
	}
}
