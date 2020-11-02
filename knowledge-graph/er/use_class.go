package er

type SpecificUseClass string

const (
	// Food and Beverage
	SpecificUseClass_Restaurant    SpecificUseClass = "Restaurant"
	SpecificUseClass_BarPub        SpecificUseClass = "BarPub"
	SpecificUseClass_RestaurantBar SpecificUseClass = "RestaurantBar"
	SpecificUseClass_Nightclub     SpecificUseClass = "Nightclub"

	// Shop, Office and Services
	SpecificUseClass_Shop                 SpecificUseClass = "Shop"
	SpecificUseClass_Laundromat           SpecificUseClass = "Laundromat"
	SpecificUseClass_Office               SpecificUseClass = "Office"
	SpecificUseClass_MassageEstablishment SpecificUseClass = "MassageEstablishment"
	SpecificUseClass_MedicalClinic        SpecificUseClass = "MedicalClinic"
	SpecificUseClass_PetShop              SpecificUseClass = "PetShop"
	SpecificUseClass_PetBoarding          SpecificUseClass = "PetBoarding"

	// Education
	SpecificUseClass_CommercialSchool SpecificUseClass = "CommercialSchool"
	SpecificUseClass_Childcare        SpecificUseClass = "Childcare"

	// Sports and Recreation
	SpecificUseClass_FitnessCentre   SpecificUseClass = "FitnessCentre"
	SpecificUseClass_AmusementCentre SpecificUseClass = "AmusementCentre"

	// Accommodation
	SpecificUseClass_Residential      SpecificUseClass = "Residential"
	SpecificUseClass_BackpackerHostel SpecificUseClass = "BackpackerHostel"
	SpecificUseClass_Hotel            SpecificUseClass = "Hotel"
	SpecificUseClass_StudentHostel    SpecificUseClass = "StudentHostel"
	SpecificUseClass_ServiceApartment SpecificUseClass = "ServiceApartment"
	SpecificUseClass_WorkerDorm       SpecificUseClass = "WorkerDorm"

	// Industrial-related uses
	SpecificUseClass_LightIndUse         SpecificUseClass = "LightIndUse"
	SpecificUseClass_GeneralIndUse       SpecificUseClass = "GeneralIndUse"
	SpecificUseClass_IndTraining         SpecificUseClass = "IndTraining"
	SpecificUseClass_Warehouse           SpecificUseClass = "Warehouse"
	SpecificUseClass_IndCanteen          SpecificUseClass = "IndCanteen"
	SpecificUseClass_Showroom            SpecificUseClass = "Showroom"
	SpecificUseClass_EBusiness           SpecificUseClass = "EBusiness"
	SpecificUseClass_CoreMediaActivities SpecificUseClass = "CoreMediaActivities"

	// Community-related uses
	SpecificUseClass_CommunityClub SpecificUseClass = "CommunityClub"

	// Religious Use
	SpecificUseClass_ReligiousActivities SpecificUseClass = "ReligiousActivities"
	SpecificUseClass_LimitedReligiousUse SpecificUseClass = "LimitedReligiousUse"
)

type GenericUseClass string

const (
	GenericUseClass_FoodBev              GenericUseClass = "FoodBev"
	GenericUseClass_ShopOfficeSvc        GenericUseClass = "ShopOfficeSvc"
	GenericUseClass_Education            GenericUseClass = "Education"
	GenericUseClass_SportsRec            GenericUseClass = "SportsRec"
	GenericUseClass_Accommodation        GenericUseClass = "Accommodation"
	GenericUseClass_IndRelatedUse        GenericUseClass = "IndRelatedUse"
	GenericUseClass_CoommunityRelatedUse GenericUseClass = "CoommunityRelatedUse"
	GenericUseClass_ReligiousUse         GenericUseClass = "ReligiousUse"
)

var UseClassMap = map[SpecificUseClass]GenericUseClass{
	SpecificUseClass_Restaurant:           GenericUseClass_FoodBev,
	SpecificUseClass_BarPub:               GenericUseClass_FoodBev,
	SpecificUseClass_RestaurantBar:        GenericUseClass_FoodBev,
	SpecificUseClass_Nightclub:            GenericUseClass_FoodBev,
	SpecificUseClass_Shop:                 GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_Laundromat:           GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_Office:               GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_MassageEstablishment: GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_MedicalClinic:        GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_PetShop:              GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_PetBoarding:          GenericUseClass_ShopOfficeSvc,
	SpecificUseClass_CommercialSchool:     GenericUseClass_Education,
	SpecificUseClass_Childcare:            GenericUseClass_Education,
	SpecificUseClass_FitnessCentre:        GenericUseClass_SportsRec,
	SpecificUseClass_AmusementCentre:      GenericUseClass_SportsRec,
	SpecificUseClass_Residential:          GenericUseClass_Accommodation,
	SpecificUseClass_BackpackerHostel:     GenericUseClass_Accommodation,
	SpecificUseClass_Hotel:                GenericUseClass_Accommodation,
	SpecificUseClass_StudentHostel:        GenericUseClass_Accommodation,
	SpecificUseClass_ServiceApartment:     GenericUseClass_Accommodation,
	SpecificUseClass_WorkerDorm:           GenericUseClass_Accommodation,
	SpecificUseClass_LightIndUse:          GenericUseClass_IndRelatedUse,
	SpecificUseClass_GeneralIndUse:        GenericUseClass_IndRelatedUse,
	SpecificUseClass_IndTraining:          GenericUseClass_IndRelatedUse,
	SpecificUseClass_Warehouse:            GenericUseClass_IndRelatedUse,
	SpecificUseClass_IndCanteen:           GenericUseClass_IndRelatedUse,
	SpecificUseClass_Showroom:             GenericUseClass_IndRelatedUse,
	SpecificUseClass_EBusiness:            GenericUseClass_IndRelatedUse,
	SpecificUseClass_CoreMediaActivities:  GenericUseClass_IndRelatedUse,
	SpecificUseClass_CommunityClub:        GenericUseClass_CoommunityRelatedUse,
	SpecificUseClass_ReligiousActivities:  GenericUseClass_ReligiousUse,
	SpecificUseClass_LimitedReligiousUse:  GenericUseClass_ReligiousUse,
}
