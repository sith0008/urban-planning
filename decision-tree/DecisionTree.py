from sklearn import tree
# import pydotplus
from sklearn.tree import DecisionTreeClassifier
import matplotlib.pyplot as plt
import matplotlib.image as pltimg
import pandas as pd


def decision(business, zone_no, unitzone):
    # Convert paramters to int type
    business, zone_no = int(business), int(zone_no)

    # Read in CSV and build dataframe
    df = pd.read_csv("Assessment_Criteria.csv")
    df['Unit_Zone'] = df['Unit_Zone'].fillna('Others')
    df['Business_Use_Type'], business_df = pd.factorize(df.Business_Use_Type)
    business_df = business_df.tolist()
    df['Property_Type'], property_df = pd.factorize(df.Property_Type)
    property_df = property_df.tolist()
    df['Unit_Zone'], unitzone_df = pd.factorize(df.Unit_Zone)
    unitzone_df = unitzone_df.tolist()
    df['Assessment_Criteria'], assessment_mapping = pd.factorize(df.Assessment_Criteria)
    assessment_mapping = assessment_mapping.tolist()

    # Build Decision Tree
    features = ['Business_Use_Type', 'Property_Type', 'Unit_Zone']
    X = df[features]
    y = df['Assessment_Criteria']
    dtree = DecisionTreeClassifier()
    dtree = dtree.fit(X, y)

    print(len(business_df), business_df)
    print(len(property_df), property_df)
    print(len(unitzone_df), unitzone_df)
    print(len(assessment_mapping), assessment_mapping)

    # Init mappings
    zoneNoToZone = {
        1: 'RESIDENTIAL',
        2: 'RESIDENTIAL WITH COMMERCIAL AT 1ST STOREY',
        3: 'COMMERCIAL & RESIDENTIAL',
        4: 'COMMERCIAL',
        5: 'HOTEL',
        6: 'WHITE',
        7: 'BUSINESS PARK',
        8: 'BUSINESS PARK - WHITE',
        9: 'BUSINESS 1 (B1)',
        10: 'BUSINESS 2 (B2)',
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
        'BUSINESS 1 (B1)': 'Industrial Buildings',
        'BUSINESS 2 (B2)': 'Industrial Buildings',
        'BUSINESS 1 - White': 'Business 1-White Buildings',
        'BUSINESS 2 - White': 'Business 2-White Buildings',
        'HEALTH & MEDICAL CARE': 'Medical and Healthcare',
        'EDUCATIONAL INSTITUTION': 'Educational Institution',
        'PLACE OF WORSHIP': 'Place of Worship',
        'CIVIC & COMMUNITY INSTITUTION': 'Civic & Community Institution'
    }

    # Map Zone Number to Property Type
    print("zone_no", zone_no, type(zone_no))
    zone = zoneNoToZone[zone_no]
    print(zone)
    if zone in zoneToBiz:
        property_type = property_df.index(zoneToBiz[zone])+1
        print(business, business_df[business-1])
        print(assessment_mapping[dtree.predict([[business, property_type, unitzone]])[0]])
        return assessment_mapping[dtree.predict([[business, property_type, unitzone]])[0]]
    else:
        return 4


    # data = tree.export_graphviz(dtree, out_file=None, feature_names=features)
    # graph = pydotplus.graph_from_dot_data(data)
    # graph.write_png('mydecisiontree.png')
    #
    # img = pltimg.imread('mydecisiontree.png')
    # imgplot = plt.imshow(img)
    # plt.show()

