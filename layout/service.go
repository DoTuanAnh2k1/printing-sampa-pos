package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceService(layout string, serviceList []model.Service) string {
	var serviceOut string
	serviceOut += "[SERVICES]\n"
	serviceSectionLayout := getServiceSection(layout)

	for _, service := range serviceList {
		tmp := serviceSectionLayout
		tmp = strings.ReplaceAll(tmp, "{CALCULATION NAME}", service.Name)
		tmp = strings.ReplaceAll(tmp, "{CALCULATION TOTAL}", service.Total)
		tmp += "\n"
		serviceOut = serviceOut + tmp
	}
	return serviceOut
}

func getServiceSection(layout string) (serviceSectionLayout string) {
	// get payment section
	reService := regexp.MustCompile(`(?s)\[SERVICES\](.*?)\[`)
	matchesService := reService.FindStringSubmatch(layout)
	if len(matchesService) > 1 {
		serviceSectionLayout = matchesService[1]
	}
	return
}
