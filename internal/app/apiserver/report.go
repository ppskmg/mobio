package apiserver

import (
	"crypto/tls"
	"encoding/json"
	gomail "github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"mobio/internal/app/model/report"
	"mobio/pkg/mwr"
	"net/http"
	"os"
)

func (mr *muxRouter) voteRouter() *httprouter.Router {
	h := httprouter.New()
	h.POST(mr.apiUrl.report.send,
		mwr.Middlewares(
			mr.middleware.cors,
		).Then(
			mr.handler.report.send()))
	return h
}

type reportHandle struct {
	*handleResponse
	//store Client
}

func (rh *reportHandle) send() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		req := &report.Report{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			rh.error(w, r, http.StatusBadRequest, err)
		}
		ee := &report.Report{
			Email: req.Email,
		}

		type emailConfig struct {
			name     string
			username string
			password string
			smtpHost string
			smtpPort int
			to       []string
		}
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
		username, _ := os.LookupEnv("GM_MOBIO_USERNAME")
		password, _ := os.LookupEnv("GM_MOBIO_PASSWORD")
		config := &emailConfig{
			name:     "Mobio",
			username: username,
			password: password,
			smtpHost: "smtp.gmail.com",
			smtpPort: 587,
			to:        []string{ee.Email},
		}

		const message = `
						<h1>Здравствуйте!</h1>
						<p>Благодарим за интерес к отчету.</p>
						<p>Готовы помочь вам в продвижении мобильных приложений<br> на рынках РФ и за рубежом.</p>
						<p>С уважением, команда Mobio.</p>
						
						<b>Как с нами связаться:</b><br>
						*Ответьте на это письмо.<br>
						*Телеграмм: <a href="">telegram</a><br>
						*Заполните форму обратной связи на сайте: <a href="/link">link</a><br>
						*Напишите нам на почту: <a href="mailto:example@example.com">example@example.com</a><br><br>`

		m := gomail.NewMessage()
		m.SetHeader("From", config.username)
		m.SetHeader("To", ee.Email)
		//m.SetAddressHeader("Cc", config.username, config.name)
		m.SetHeader("Subject", "Mobio — Отчет в приложении к письму.")
		m.SetBody("text/html", message)
		m.Attach("./AppMagic.pdf")

		d := gomail.NewDialer(config.smtpHost, config.smtpPort, config.username, config.password)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := d.DialAndSend(m); err != nil {
			rh.error(w, r, http.StatusUnprocessableEntity, err)
		}
		rh.respond(w, r, http.StatusCreated, "report +")
	}
}
