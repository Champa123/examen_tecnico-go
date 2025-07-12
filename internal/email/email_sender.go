package email

import (
	"encoding/base64"
	"examen-tecnico-stori/internal/model"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(summary *model.Summary) {

	from := mail.NewEmail("Andres", "andres.fernandez.bina@gmail.com")
	subject := "Summary"
	to := mail.NewEmail("Andres", "afernandezbina@frba.utn.edu.ar") // Podria obtenerse el mail dinamicamente
	htmlContent := buildHtml(summary)
	htmlContent += "<img src=cid:stori-logo></img>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	base64Image, err := loadImageBase64()
	if err != nil {
		panic(err)
	}

	attachment := mail.NewAttachment()
	attachment.SetContent(base64Image)
	attachment.SetType("image/png")
	attachment.SetFilename("logo.png")
	attachment.SetDisposition("inline")
	attachment.SetContentID("stori-logo")

	message.AddAttachment(attachment)

	_, err = client.Send(message)
	if err != nil {
		panic(err)
	}

}

func loadImageBase64() (string, error) {
	imgPath := "resources/assets/stori-logo.png"
	data, err := os.ReadFile(imgPath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func buildHtml(summary *model.Summary) string {
	body := "<ul> Your transactions summary:"
	element := "<li><b>%v: </b>%v</li>"

	body += fmt.Sprintf(element, "Total balance", summary.TotalBalance)
	body += fmt.Sprintf(element, "Average debit amount", summary.AvarageDebit)
	body += fmt.Sprintf(element, "Average credit amount", summary.AvarageCredit)

	for _, total := range summary.MonthlyTransactions {
		body += fmt.Sprintf(element, "Number of transactions in "+total.Month.String(), total.Total)
	}

	body += "</ul> <br> <br> <br>"

	return body
}
