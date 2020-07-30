package services

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"time"
	"truora-server/models"
	"truora-server/resources"
	"truora-server/utils"
)

type ServiceDomain struct {
	Database resources.DataBase
}

//Get the information of the server
func (d *ServiceDomain) GetInfo(ctx *fasthttp.RequestCtx) {
	domain, ok := ctx.UserValue("domain").(string)
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	enc := json.NewEncoder(ctx)

	//Obtener los datos de los servidores de SSLLabs
	var response models.Domain
	response.Name = domain
	dataFromAPI := resources.GetInfoDomainFromAPI(domain)
	response.IsDown = dataFromAPI.Status != "READY"
	utils.SetInfoServers(dataFromAPI.Servers, &response)

	dataDomain, err := d.Database.GetDomainByName(domain)
	exists := !(err != nil && dataDomain.SSLGrade == "")

	if exists {
		currentTime := time.Now()
		difference := currentTime.Sub(dataDomain.DateChecked)

		response.Logo = dataDomain.Logo
		response.Title = dataDomain.Title

		if difference.Hours() >= 1 {
			response.PreviousSSLGrade = dataDomain.SSLGrade
			oldServers, err := d.Database.GetDomainServers(domain)
			if !(err != nil && len(oldServers) == 0) {
				if haveChanged(oldServers, response.Servers) {
					response.ServersChanged = true
					d.Database.UpdateDomainServersChanged(domain)
				}
			}
		}
	} else {
		bodyReader := utils.GetHtmlReader(domain)
		icon := utils.GetIcon(bodyReader)
		titleText := utils.GetTitle(bodyReader)
		response.Logo = icon
		response.Title = titleText

		//Save Domain in database
		err = d.Database.CreateDomain(response, domain)
		if err != nil {
			log.Fatal(err)
		}

		//Save Servers in database
		for _, server := range response.Servers {
			err = d.Database.CreateServer(server, domain)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	enc.Encode(&response)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

//Get the history of requests
func (d *ServiceDomain) GetHistory(ctx *fasthttp.RequestCtx) {
	enc := json.NewEncoder(ctx)
	var history models.History
	items, err := d.Database.GetDomains()
	history.Items = items
	if err != nil {
		log.Fatal(err)
	}
	enc.Encode(&history)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func haveChanged(oldServers []models.Server, newServers []models.Server) bool {
	if len(oldServers) != len(newServers) {
		return false
	}
	isDifferent := true
	for _, server := range oldServers {
		isDifferent = true
		for _, serverFromApi := range newServers {
			if server == serverFromApi {
				isDifferent = false
				break
			}
		}

		if isDifferent {
			return true
		}
	}

	return isDifferent
}