from sklearn import tree
# import pydotplus
from sklearn.tree import DecisionTreeClassifier
import matplotlib.pyplot as plt
import matplotlib.image as pltimg
import pandas as pd
from mapping import *


def decision(business, zone_no, unitzone, condition='0'):
    # Convert paramters to int type
    business, zone_no = int(business), int(zone_no)

    # Read in CSV and build dataframe
    df = pd.read_csv("Assessment_Criteria.csv")

    df['Unit_Zone'] = df['Unit_Zone'].fillna('0')
    df['Condition_Name'] = df['Condition_Name'].fillna('0')

    df['Business_Use_Type'], business_df = pd.factorize(df.Business_Use_Type)
    business_df = business_df.tolist()
    df['Property_Type'], property_df = pd.factorize(df.Property_Type)
    property_df = property_df.tolist()
    df['Unit_Zone'], unitzone_df = pd.factorize(df.Unit_Zone)
    unitzone_df = unitzone_df.tolist()
    df['Condition_Name'], condition_df = pd.factorize(df.Condition_Name)
    condition_df = condition_df.tolist()
    df['Assessment_Criteria'], assessment_mapping = pd.factorize(df.Assessment_Criteria)
    assessment_mapping = assessment_mapping.tolist()

    # Build Decision Tree
    features = ['Business_Use_Type', 'Property_Type', 'Unit_Zone', 'Condition_Name']
    X = df[features]
    y = df['Assessment_Criteria']
    dtree = DecisionTreeClassifier()
    dtree = dtree.fit(X, y)

    print(business_df)
    print(property_df)
    # print(len(unitzone_df), unitzone_df)
    # print(len(assessment_mapping), assessment_mapping)

    # Map Zone Number to Property Type
    zone = zoneNoToZone[zone_no]
    if zone in zoneToBiz:
        print("zone", zone)
        property_type = property_df.index(zoneToBiz[zone]) + 1
        print("property", property_type, property_df[property_type-1])
        decision = dtree.predict([[business, property_type, unitzone, condition]])[0]
        print("decision", decision, classification_mapping[decision])
        return assessmentToAssessmentNo[assessment_mapping[decision]]
    else:
        return 5


    # data = tree.export_graphviz(dtree, out_file=None, feature_names=features)
    # graph = pydotplus.graph_from_dot_data(data)
    # graph.write_png('mydecisiontree.png')
    #
    # img = pltimg.imread('mydecisiontree.png')
    # imgplot = plt.imshow(img)
    # plt.show()

# print(decision(4, 3, 0))
print(decision(1, 4, 0))

