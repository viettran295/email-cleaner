package main

import (
	"email-cleaner/service"
	"email-cleaner/utils"
	"log"
	"time"
)

func main() {
	deleteEmails := []string{"WallStreetJournal@em.dowjones.com", "newsletter@email.blockchain.com", "jobs-noreply@linkedin.com", "emirates@e.emirates.email",
		"service@mailer.scalable.capital", "service@mailer.scalable.capital", "messages-noreply@linkedin.com", "linkedin@e.linkedin.com", "no-reply@khunpon.de",
		"updates-noreply@linkedin.com", "notifications-noreply@linkedin.com", "invitations@linkedin.com", "no-reply@mail.goodreads.com", "donotreply@amazon.de",
		"webform@altium.com", "noreply@medium.com", "info@info.vietcombank.com.vn", "st@newsletter.st.com", "st@event.st.com", "teamzoom@e.zoom.us",
		"hello@info.zalando-outlet.de", "newsletter@c.rituals.com", "hello@medium.com", "mailrobot@mail.xing.com", "no-reply@m.mail.coursera.org",
		"Barrons@em.dowjones.com", "newsletter@khunpon.de", "noreply-sg@email.decathlon.com", "academy@mail.wandb.ai", "jory@m.fontawesome.com",
		"noreply@glassdoor.com", "no_reply@bahn.de", "MarketWatch@em.dowjones.com", "noreplay@newsletter.saalesparkasse.de", "service@info.kleinanzeigen.de",
		"neues@funkelkram.de", "singaporeair@email.singaporeair.com", "ilscan@ef.com", "crm_asia@samsung.com", "News@em.acb.com.vn", "opportunities@careeralerts.hellofresh.com",
		"admin@bibcitation.com", "info@bjjfanatics.com", "germany@bolt-rider.com", "discover@airbnb.com", "noreply@e.klarna.com", "quincy@freecodecamp.org",
		"matt@m.fontawesome.com", "community@email.hays.com", "flashback@gopro.com", "mackenzie.jackson@gitguardian.com", "newsletter@membership.cgv.vn",
		"vietnam@seao.crm.samsung.com", "service@info.misterspex.de", "replies@oracle-mail.com", "MITx_OL@mit.edu", "michaelpage@mail.michaelpage.de",
		"team@m.ngrok.com", "hello@knog.com.au", "hello@mail.blinkist.com", "promotion5@amazon.de", "offers@m.myunidays.com", "success@gitkraken.com",
		"german-personalized-digest@quora.com", "appstore@insideapple.apple.com", "do_not_reply@mailersp1.binance.com", "study@constructor.university",
		"noreply@huobi.ug", "no-reply@leetcode.com", "microsoft.start@email2.microsoft.com", "info@simplize.vn", "cskh@abs.vn", "vietpride295@gmail.com",
		"apply@42heilbronn.de", "noreply@immobilienscout24.de", "do-not-reply@wohnungshelden.de", "noreply@meinverein.de", "promotions@bambooairways.com",
		"st@information.st.com", "hello@taptapsend.os.tc", "info@glassdoor.com", "newsletter@email.n26.com", "thy@mail.turkishairlines.com", "no-reply@grab.com",
		"hello@udacity.com", "no-reply@mail.wandb.ai", "account-security-noreply@accountprotection.microsoft.com", "aws-marketing-email-replies@amazon.com",
		"service@service.scalable.capital", "cskh@abs.vn"}

	for {
		mess := service.ListMessId()
		for _, msg := range mess {
			sender := service.GetSender(msg.Id)
			if utils.ContainString(deleteEmails, sender) == true {
				go service.DeleteMess(msg.Id)
				log.Println("Deleted sender: ", sender)
			}
		}
		time.Sleep(6 * time.Hour)
	}
}
