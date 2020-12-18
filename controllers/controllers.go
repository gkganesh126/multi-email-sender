package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gkganesh126/multi-email-sender/auth"
	"github.com/gkganesh126/multi-email-sender/models"
	"github.com/golang/glog"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	glog.Info("At SendEmail...\n")

	if !auth.ValidateToken(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var email models.Email
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		glog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	glog.Info("email: ", email)

	for _, provider := range email.Provider {
		switch provider {
		case "mailjet":
			sendMailjetEmail(email)
			break
		case "sendgrid":
			//sendSendgridEmail(email)
			break
		case "amazon-ses":
			//sendAmazonSesEmail(email)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}

func sendMailjetEmail(email models.Email) {
	mailjetClient := mailjet.NewMailjetClient("14641d86d254bf8aaa74a71887475e24", "e46fe3841c7b114aeea940411a30611e")
	var recipients []mailjet.RecipientV31
	for i, _ := range email.To {
		recipient := mailjet.RecipientV31{
			Email: email.To[i].EmailID,
			Name:  email.To[i].Name,
		}
		recipients = append(recipients, recipient)
	}
	toRecipients := mailjet.RecipientsV31(recipients)

	recipients = recipients[:0]
	for i, _ := range email.Cc {
		recipient := mailjet.RecipientV31{
			Email: email.Cc[i].EmailID,
			Name:  email.Cc[i].Name,
		}
		recipients = append(recipients, recipient)
	}
	ccRecipients := mailjet.RecipientsV31(recipients)

	recipients = recipients[:0]
	for i, _ := range email.Bcc {
		recipient := mailjet.RecipientV31{
			Email: email.Bcc[i].EmailID,
			Name:  email.Bcc[i].Name,
		}
		recipients = append(recipients, recipient)
	}
	bccRecipients := mailjet.RecipientsV31(recipients)

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: email.From.EmailID,
				Name:  email.From.Name,
			},
			To:      &toRecipients,
			Cc:      &ccRecipients,
			Bcc:     &bccRecipients,
			Subject: email.Subject,
			//TextPart: "My first Mailjet email",
			HTMLPart: email.Message,
			CustomID: "AppGettingStartedTest",
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		glog.Error(err.Error())
		return
	}
	glog.Info("Data: ", res)
}
