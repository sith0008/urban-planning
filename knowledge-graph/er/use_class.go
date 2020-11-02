package er

type SpecificUseClass string

const (
	// Food and Beverage
	SpecificUseClass_Restaurant    SpecificUseClass = "Restaurant"
	SpecificUseClass_BarPub        SpecificUseClass = "Bar/Pub"
	SpecificUseClass_RestaurantBar SpecificUseClass = "Restaurant and Bar"
	SpecificUseClass_Nightclub     SpecificUseClass = "Nightclub"

	// Shop, Office and Services
	SpecificUseClass_Shop                 SpecificUseClass = "Shop"
	SpecificUseClass_Laundromat           SpecificUseClass = "Laundromat"
	SpecificUseClass_Office               SpecificUseClass = "Office"
	SpecificUseClass_MassageEstablishment SpecificUseClass = "Massage Establishment"
	SpecificUseClass_MedicalClinic        SpecificUseClass = "Medical Clinic"
	SpecificUseClass_PetShop              SpecificUseClass = "Pet Shop"
	SpecificUseClass_PetBoarding          SpecificUseClass = "Pet Boarding"

	// Education
	SpecificUseClass_CommercialSchool SpecificUseClass = "Commercial School"
	SpecificUseClass_Childcare        SpecificUseClass = "Childcare Centre"

	// Sports and Recreation
	SpecificUseClass_FitnessCentre   SpecificUseClass = "Fitness Centre"
	SpecificUseClass_AmusementCentre SpecificUseClass = "Amusement Centre"

	// Accommodation
	SpecificUseClass_Residential      SpecificUseClass = "Residential"
	SpecificUseClass_BackpackerHostel SpecificUseClass = "Backpackers' Hostel"
	SpecificUseClass_Hotel            SpecificUseClass = "Hotel"
	SpecificUseClass_StudentHostel    SpecificUseClass = "Students' Hostel"
	SpecificUseClass_ServiceApartment SpecificUseClass = "Serviced Apartment"
	SpecificUseClass_WorkerDorm       SpecificUseClass = "Workers' Dormitories"

	// Industrial-related uses
	SpecificUseClass_LightIndUse         SpecificUseClass = "Light Industrial Use"
	SpecificUseClass_GeneralIndUse       SpecificUseClass = "General Industrial Use"
	SpecificUseClass_IndTraining         SpecificUseClass = "Industrial Training"
	SpecificUseClass_Warehouse           SpecificUseClass = "Warehouse"
	SpecificUseClass_IndCanteen          SpecificUseClass = "Industrial Canteen"
	SpecificUseClass_Showroom            SpecificUseClass = "Showroom"
	SpecificUseClass_EBusiness           SpecificUseClass = "E-Business"
	SpecificUseClass_CoreMediaActivities SpecificUseClass = "Core Media Activities"

	// Community-related uses
	SpecificUseClass_CommunityClub SpecificUseClass = "Association/CommunityClub/Family Service Centre"

	// Religious Use
	SpecificUseClass_ReligiousActivities SpecificUseClass = "Religious Activities"
	SpecificUseClass_LimitedReligiousUse SpecificUseClass = "Limited & Non-Exclusive Religious Use"
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
