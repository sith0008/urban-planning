from sklearn import tree
# import pydotplus
from sklearn.tree import DecisionTreeClassifier
import matplotlib.pyplot as plt
import matplotlib.image as pltimg
import pandas as pd
from mapping import *


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

    # print(len(business_df), business_df)
    # print(len(property_df), property_df)
    # print(len(unitzone_df), unitzone_df)
    # print(len(assessment_mapping), assessment_mapping)

    # Map Zone Number to Property Type
    zone = zoneNoToZone[zone_no]
    if zone in zoneToBiz:
        property_type = property_df.index(zoneToBiz[zone])+1
        return assessmentToAssessmentNo[assessment_mapping[dtree.predict([[business, property_type, unitzone]])[0]]]
    else:
        return 4


    # data = tree.export_graphviz(dtree, out_file=None, feature_names=features)
    # graph = pydotplus.graph_from_dot_data(data)
    # graph.write_png('mydecisiontree.png')
    #
    # img = pltimg.imread('mydecisiontree.png')
    # imgplot = plt.imshow(img)
    # plt.show()

