package util

import (
	"os"

	"github.com/google/uuid"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udm/internal/context"
	"github.com/free5gc/udm/internal/logger"
	"github.com/free5gc/udm/pkg/factory"
)

func InitUDMContext(udmContext *context.UDMContext) {
	config := factory.UdmConfig
	logger.UtilLog.Info("udmconfig Info: Version[", config.Info.Version, "] Description[", config.Info.Description, "]")
	configuration := config.Configuration
	udmContext.NfId = uuid.New().String()
	sbi := configuration.Sbi
	udmContext.UriScheme = ""
	udmContext.SBIPort = factory.UdmSbiDefaultPort
	udmContext.RegisterIP = factory.UdmSbiDefaultIP
	if sbi != nil {
		if sbi.Scheme != "" {
			udmContext.UriScheme = models.UriScheme(sbi.Scheme)
		}
		if sbi.RegisterIP != "" {
			udmContext.RegisterIP = sbi.RegisterIP
		}
		if sbi.Port != 0 {
			udmContext.SBIPort = sbi.Port
		}

		udmContext.BindingIP = os.Getenv(sbi.BindingIP)
		if udmContext.BindingIP != "" {
			logger.UtilLog.Info("Parsing ServerIP address from ENV Variable.")
		} else {
			udmContext.BindingIP = sbi.BindingIP
			if udmContext.BindingIP == "" {
				logger.UtilLog.Warn("Error parsing ServerIPv4 address as string. Using the 0.0.0.0 address as default.")
				udmContext.BindingIP = "0.0.0.0"
			}
		}
	}
	udmContext.NrfUri = configuration.NrfUri
	servingNameList := configuration.ServiceNameList

	udmContext.SuciProfiles = configuration.SuciProfiles

	udmContext.InitNFService(servingNameList, config.Info.Version)
}
