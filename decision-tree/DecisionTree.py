from sklearn import tree
# import pydotplus
from sklearn.tree import DecisionTreeClassifier
import matplotlib.pyplot as plt
import matplotlib.image as pltimg
import pandas as pd


def decision(business, property, unitzone):
    df = pd.read_csv("Assessment_Criteria.csv")
    df['Unit_Zone'] = df['Unit_Zone'].fillna('Others')
    df['Business_Use_Type'], business_mapping = pd.factorize(df.Business_Use_Type)
    business_mapping = business_mapping.tolist()
    df['Property_Type'], property_mapping = pd.factorize(df.Property_Type)
    property_mapping = property_mapping.tolist()
    df['Unit_Zone'], unitzone_mapping = pd.factorize(df.Unit_Zone)
    unitzone_mapping = unitzone_mapping.tolist()
    df['Assessment_Criteria'], assessment_mapping = pd.factorize(df.Assessment_Criteria)
    assessment_mapping = assessment_mapping.tolist()

    print(business_mapping)

    features = ['Business_Use_Type', 'Property_Type', 'Unit_Zone']

    X = df[features]
    y = df['Assessment_Criteria']

    dtree = DecisionTreeClassifier()
    dtree = dtree.fit(X, y)
    # data = tree.export_graphviz(dtree, out_file=None, feature_names=features)
    # graph = pydotplus.graph_from_dot_data(data)
    # graph.write_png('mydecisiontree.png')
    #
    # img = pltimg.imread('mydecisiontree.png')
    # imgplot = plt.imshow(img)
    # plt.show()
    #
    # print(dtree.predict([[0, 4, 2]]))

    return dtree.predict([[business, property, unitzone]])

decision(1,1,1)