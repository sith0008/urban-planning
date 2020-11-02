package er

type SpecificPropType string

const (
	SpecificPropType_Residential               SpecificPropType = "Residential"
	SpecificPropType_ResComm1st                SpecificPropType = "Residential with Commercial at 1st storey"
	SpecificPropType_ResComm                   SpecificPropType = "Commercial & Residential"
	SpecificPropType_ResInstitution            SpecificPropType = "Residential/Institution"
	SpecificPropType_Commercial                SpecificPropType = "Commercial"
	SpecificPropType_CommInstitution           SpecificPropType = "Commercial/Institution"
	SpecificPropType_Hotel                     SpecificPropType = "Hotel"
	SpecificPropType_White                     SpecificPropType = "White"
	SpecificPropType_BusinessPark              SpecificPropType = "Business Park"
	SpecificPropType_BusinessParkWhite         SpecificPropType = "Business Park - White"
	SpecificPropType_Business1                 SpecificPropType = "Business 1 (B1)"
	SpecificPropType_Business2                 SpecificPropType = "Business 2 (B2)"
	SpecificPropType_Business1White            SpecificPropType = "Business 1 - White"
	SpecificPropType_Business2White            SpecificPropType = "Business 2 - White"
	SpecificPropType_HealthMed                 SpecificPropType = "Health & Medical Care"
	SpecificPropType_EduInstitution            SpecificPropType = "Educational Institution"
	SpecificPropType_PlaceOfWorship            SpecificPropType = "Place of Worship"
	SpecificPropType_CivicCommunityInstitution SpecificPropType = "Civic & Community Institution"
	// SpecificPropType_OpenSpace                 SpecificPropType = "OpenSpace"
	// SpecificPropType_Park                      SpecificPropType = "Park"
	// SpecificPropType_BeachArea                 SpecificPropType = "BeachArea"
	// SpecificPropType_SportsRec                 SpecificPropType = "SportsRec"
	// SpecificPropType_WaterBody                 SpecificPropType = "WaterBody"
	// SpecificPropType_Road                      SpecificPropType = "Road"
	// SpecificPropType_TransportFacil            SpecificPropType = "TransportFacil"
	// SpecificPropType_RapidTransit              SpecificPropType = "RapidTransit"
	// SpecificPropType_Utility                   SpecificPropType = "Utility"
	// SpecificPropType_Cemetery                  SpecificPropType = "Cemetery"
	// SpecificPropType_Agriculture               SpecificPropType = "Agriculture"
	// SpecificPropType_Port                      SpecificPropType = "Port"
	// SpecificPropType_ReserveSite               SpecificPropType = "ReserveSite"
	// SpecificPropType_SpecialUse                SpecificPropType = "SpecialUse"
)

type GenericPropType string

const (
	GenericPropType_Business       GenericPropType = "Business"
	GenericPropType_Commercial     GenericPropType = "Commercial"
	GenericPropType_Residential    GenericPropType = "Residential"
	GenericPropType_MixedDev       GenericPropType = "MixedDev"
	GenericPropType_SpecialisedUse GenericPropType = "SpecialisedUse"
)

var PropTypeMap = map[SpecificPropType]GenericPropType{
	SpecificPropType_Residential:               GenericPropType_Residential,
	SpecificPropType_ResComm1st:                GenericPropType_MixedDev,
	SpecificPropType_ResComm:                   GenericPropType_MixedDev,
	SpecificPropType_ResInstitution:            GenericPropType_Residential,
	SpecificPropType_Commercial:                GenericPropType_Commercial,
	SpecificPropType_CommInstitution:           GenericPropType_Commercial,
	SpecificPropType_Hotel:                     GenericPropType_SpecialisedUse,
	SpecificPropType_White:                     GenericPropType_MixedDev,
	SpecificPropType_BusinessPark:              GenericPropType_Business,
	SpecificPropType_BusinessParkWhite:         GenericPropType_Business,
	SpecificPropType_Business1:                 GenericPropType_Business,
	SpecificPropType_Business2:                 GenericPropType_Business,
	SpecificPropType_Business1White:            GenericPropType_Business,
	SpecificPropType_Business2White:            GenericPropType_Business,
	SpecificPropType_HealthMed:                 GenericPropType_SpecialisedUse,
	SpecificPropType_EduInstitution:            GenericPropType_SpecialisedUse,
	SpecificPropType_PlaceOfWorship:            GenericPropType_SpecialisedUse,
	SpecificPropType_CivicCommunityInstitution: GenericPropType_SpecialisedUse,
}
