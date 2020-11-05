business_mapping = {
    'Restaurant': 1,
    'Bar/Pub': 2,
    'Restaurant and Bar': 3,
    'Nightclub': 4,
    'Shop': 5,
    'Laundromat': 6,
    'Office': 7,
    'Massage Establishment': 8,
    'Medical Clinic': 9,
    'Pet Shop': 10,
    'Pet Boarding': 11,
    'Commercial School': 12,
    'Childcare Centre': 13,
    'Fitness Centre/Gymnasium': 14,
    'Amusement Centre': 15,
    'Residential': 16,
    "Backpackers' Hostel": 17,
    'Hotel': 18,
    "Students' Hostel": 19,
    'Serviced Apartment': 20,
    "Workers' Dormitories": 21,
    'Light Industrial Use': 22,
    'General Industrial Use': 23,
    'Industrial Training': 24,
    'Warehouse': 25,
    'Industrial Canteen': 26,
    'Showroom': 27,
    'E-business': 28,
    'Core Media Activities': 29,
    'Association/Community Club/Family Service Centre': 30,
    'Religious Activities': 31,
    'Limited & Non-Exclusive Religious Use': 32
}

classification_mapping = {
    1: "No Planning Permission Required",
    2: "Instant Approval",
    3: "Submit Change Of Use Application For Evaluation",
    4: "Unlikely",
    5: "Not Allowed"
}

assessmentToAssessmentNo = {
    "No Planning Permission Required": 1,
    "Instant Approval": 2,
    "Submit Change Of Use Application For Evaluation": 3,
    "Unlikely": 4,
    "Not Allowed": 5
}

zoneToZoneNo = {
        'RESIDENTIAL': 1,
        'RESIDENTIAL WITH COMMERCIAL AT 1ST STOREY': 2,
        'COMMERCIAL & RESIDENTIAL': 3,
        'COMMERCIAL': 4,
        'HOTEL': 5,
        'WHITE': 6,
        'BUSINESS PARK': 7,
        'BUSINESS PARK - WHITE': 8,
        'BUSINESS 1': 9,
        'BUSINESS 2': 10,
        'BUSINESS 1 - WHITE': 11,
        'BUSINESS 2 - WHITE': 12,
        'RESIDENTIAL / INSTITUTION': 13,
        'COMMERCIAL / INSTITUTION': 14,
        'HEALTH & MEDICAL CARE': 15,
        'EDUCATIONAL INSTITUTION': 16,
        'PLACE OF WORSHIP': 17,
        'CIVIC & COMMUNITY INSTITUTION': 18,
        'OPEN SPACE': 19,
        'PARK': 20,
        'BEACH AREA': 21,
        'SPORTS & RECREATION': 22,
        'WATERBODY': 23,
        'ROAD': 24,
        'TRANSPORT FACILITIES': 25,
        'RAPID TRANSIT': 26,
        'UTILITY': 27,
        'CEMETERY': 28,
        'AGRICULTURE': 29,
        'PORT / AIRPORT': 30,
        'RESERVE SITE': 31,
        'SPECIAL USE': 32,
    }

zoneNoToZone = {
        1: 'RESIDENTIAL',
        2: 'RESIDENTIAL WITH COMMERCIAL AT 1ST STOREY',
        3: 'COMMERCIAL & RESIDENTIAL',
        4: 'COMMERCIAL',
        5: 'HOTEL',
        6: 'WHITE',
        7: 'BUSINESS PARK',
        8: 'BUSINESS PARK - WHITE',
        9: 'BUSINESS 1',
        10: 'BUSINESS 2',
        11: 'BUSINESS 1 - WHITE',
        12: 'BUSINESS 2 - WHITE',
        13: 'RESIDENTIAL / INSTITUTION',
        14: 'COMMERCIAL / INSTITUTION',
        15: 'HEALTH & MEDICAL CARE',
        16: 'EDUCATIONAL INSTITUTION',
        17: 'PLACE OF WORSHIP',
        18: 'CIVIC & COMMUNITY INSTITUTION',
        19: 'OPEN SPACE',
        20: 'PARK',
        21: 'BEACH AREA',
        22: 'SPORTS & RECREATION',
        23: 'WATERBODY',
        24: 'ROAD',
        25: 'TRANSPORT FACILITIES',
        26: 'RAPID TRANSIT',
        27: 'UTILITY',
        28: 'CEMETERY',
        29: 'AGRICULTURE',
        30: 'PORT / AIRPORT',
        31: 'RESERVE SITE',
        32: 'SPECIAL USE',
    }

zoneToBiz = {
        'RESIDENTIAL': 'HDB Commercial Premises',
        'RESIDENTIAL WITH COMMERCIAL AT 1ST STOREY': 'Mixed Commercial & Residential Developments',
        'COMMERCIAL & RESIDENTIAL': 'Mixed Commercial & Residential Developments',
        'COMMERCIAL': 'Commercial Buildings',
        'HOTEL': 'Hotel',
        'BUSINESS PARK': 'Business Park',
        'BUSINESS 1': 'Industrial Buildings',
        'BUSINESS 2': 'Industrial Buildings',
        'BUSINESS 1 - White': 'Business 1-White Buildings',
        'BUSINESS 2 - White': 'Business 2-White Buildings',
        'HEALTH & MEDICAL CARE': 'Medical and Healthcare',
        'EDUCATIONAL INSTITUTION': 'Educational Institution',
        'PLACE OF WORSHIP': 'Place of Worship',
        'CIVIC & COMMUNITY INSTITUTION': 'Civic and Community Institution'
    }