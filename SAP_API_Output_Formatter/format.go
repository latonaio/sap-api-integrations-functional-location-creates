package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-functional-location-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		FunctionalLocation:             data.FunctionalLocation,
		FunctionalLocationLabelName:    data.FunctionalLocationLabelName,
		FunctionalLocationName:         data.FunctionalLocationName,
		FuncLocationStructure:          data.FuncLocationStructure,
		FunctionalLocationCategory:     data.FunctionalLocationCategory,
		SuperiorFunctionalLocation:     data.SuperiorFunctionalLocation,
		SuperiorFuncnlLocLabelName:     data.SuperiorFuncnlLocLabelName,
		TechnicalObjectType:            data.TechnicalObjectType,
		GrossWeight:                    data.GrossWeight,
		GrossWeightUnit:                data.GrossWeightUnit,
		OperationStartDate:             data.OperationStartDate,
		InventoryNumber:                data.InventoryNumber,
		AcquisitionValue:               data.AcquisitionValue,
		Currency:                       data.Currency,
		AcquisitionDate:                data.AcquisitionDate,
		AssetManufacturerName:          data.AssetManufacturerName,
		ManufacturerPartNmbr:           data.ManufacturerPartNmbr,
		ManufacturerCountry:            data.ManufacturerCountry,
		ManufacturerPartTypeName:       data.ManufacturerPartTypeName,
		ConstructionMonth:              data.ConstructionMonth,
		ConstructionYear:               data.ConstructionYear,
		ManufacturerSerialNumber:       data.ManufacturerSerialNumber,
		MaintenancePlant:               data.MaintenancePlant,
		AssetLocation:                  data.AssetLocation,
		AssetRoom:                      data.AssetRoom,
		PlantSection:                   data.PlantSection,
		WorkCenter:                     data.WorkCenter,
		WorkCenterInternalID:           data.WorkCenterInternalID,
		WorkCenterPlant:                data.WorkCenterPlant,
		ABCIndicator:                   data.ABCIndicator,
		MaintObjectFreeDefinedAttrib:   data.MaintObjectFreeDefinedAttrib,
		FormOfAddress:                  data.FormOfAddress,
		BusinessPartnerName1:           data.BusinessPartnerName1,
		BusinessPartnerName2:           data.BusinessPartnerName2,
		CityName:                       data.CityName,
		PostalCode:                     data.PostalCode,
		StreetName:                     data.StreetName,
		Region:                         data.Region,
		Country:                        data.Country,
		PhoneNumber:                    data.PhoneNumber,
		FaxNumber:                      data.FaxNumber,
		CompanyCode:                    data.CompanyCode,
		BusinessArea:                   data.BusinessArea,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		CostCenter:                     data.CostCenter,
		ControllingArea:                data.ControllingArea,
		WBSElementExternalID:           data.WBSElementExternalID,
		SettlementOrder:                data.SettlementOrder,
		ConstructionMaterial:           data.ConstructionMaterial,
		MaintenancePlannerGroup:        data.MaintenancePlannerGroup,
		MaintenancePlanningPlant:       data.MaintenancePlanningPlant,
		MainWorkCenterPlant:            data.MainWorkCenterPlant,
		MainWorkCenter:                 data.MainWorkCenter,
		MainWorkCenterInternalID:       data.MainWorkCenterInternalID,
		CatalogProfile:                 data.CatalogProfile,
		EquipmentInstallationIsAllowed: data.EquipmentInstallationIsAllowed,
		OnePieceOfEquipmentIsAllowed:   data.OnePieceOfEquipmentIsAllowed,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		SalesOffice:                    data.SalesOffice,
		OrganizationDivision:           data.OrganizationDivision,
		SalesGroup:                     data.SalesGroup,
		FunctionalLocationHasEquipment: data.FunctionalLocationHasEquipment,
		FuncnlLocHasSubOrdinateFuncLoc: data.FuncnlLocHasSubOrdinateFuncLoc,
		LastChangeDateTime:             data.LastChangeDateTime,
		FuncnlLocIsMarkedForDeletion:   data.FuncnlLocIsMarkedForDeletion,
		FuncnlLocIsDeleted:             data.FuncnlLocIsDeleted,
		FunctionalLocationIsActive:     data.FunctionalLocationIsActive,
		FuncnlLocIsDeactivated:         data.FuncnlLocIsDeactivated,
		CreationDate:                   data.CreationDate,
	}
	return header, nil
}
