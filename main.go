package main

import (
	"email-cleaner/service"
	"email-cleaner/utils"
	"log"
	"time"
)

func main() {
	deleteEmails := []string{"no-reply", "info", "cskh", "newsletter", "hello", "News", "replies", "linkedin", "service@mailer.scalable.capital", 
		"indeed", "marketing", "notifications", "support", "zalando-outlet", "relaxdays", "vietcombank", "news", "jacobs-university",
		"gitguardian", "fontawesome", "WallStreetJournal@em.dowjones.com", "emirates@e.emirates.email", "service@mailer.scalable.capital", 
		"webform@altium.com", "st@event.st.com", "teamzoom@e.zoom.us", "mailrobot@mail.xing.com", "Barrons@em.dowjones.com", "wandb", 
		"MarketWatch@em.dowjones.com","neues@funkelkram.de", "singaporeair@email.singaporeair.com", "ilscan@ef.com", 
		"crm_asia@samsung.com", "admin@bibcitation.com", "germany@bolt-rider.com", "discover@airbnb.com", "quincy@freecodecamp.org",
		"community@email.hays.com", "flashback@gopro.com", "vietnam@seao.crm.samsung.com", "wellfound", "success@gitkraken.com",
		"MITx_OL@mit.edu", "michaelpage@mail.michaelpage.de", "team@m.ngrok.com", "promotion5@amazon.de", "offers@m.myunidays.com",
		"german-personalized-digest@quora.com", "appstore@insideapple.apple.com", "study@constructor.university", "microsoft.start@email2.microsoft.com", 
		"vietpride295@gmail.com", "apply@42heilbronn.de", "promotions@bambooairways.com", "thy@mail.turkishairlines.com", "service@service.scalable.capital", 
		"english-personalized-digest@quora.com", "support@anvil.works", "gopro@e-mail.gopro.com", "Xbox@engage.xbox.com", "IRBCG@bamboocap.com.vn",
		"Info@bambooairways.com", "lavanyashukla@wandb.com", "gdpr@email.hays.com"}

	deletedMess := 0
	for {
		mess := service.ListMessId()
		for _, msg := range mess {
			if deletedMess == 40{
				deletedMess = 0
				break
			} else {
				sender := service.GetSender(msg.Id)
				if utils.ContainString(deleteEmails, sender) == true {
					deletedMess++
					go service.DeleteMess(msg.Id)
					log.Println("Trashed sender: ", sender)
				}
			}
		}
		time.Sleep(12 * time.Hour)
	}
}
