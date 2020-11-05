# This files contains your custom actions which can be used to run
# custom Python code.
#
# See this guide on how to implement these action:
# https://rasa.com/docs/rasa/core/actions/#custom-actions/


# This is a simple example for a custom action which utters "Hello World!"

# from typing import Any, Text, Dict, List
#
# from rasa_sdk import Action, Tracker
# from rasa_sdk.executor import CollectingDispatcher
#
#
# class ActionHelloWorld(Action):
#
#     def name(self) -> Text:
#         return "action_hello_world"
#
#     def run(self, dispatcher: CollectingDispatcher,
#             tracker: Tracker,
#             domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
#
#         dispatcher.utter_message(text="Hello World!")
#
#         return []


from typing import Any, Text, Dict, List, Union

from rasa_sdk import Action, Tracker
from rasa_sdk.executor import CollectingDispatcher
from rasa_sdk.forms import FormAction
from rasa_sdk.events import SlotSet

from testing import test
import requests

classification_mapping = {
    1:"No planning permission required",
    2:"Instant approval",
    3:"Submit change of use application for evaluation",
    4:"The use is not within the planning intentions of the site"
}
# class HealthForm(FormAction):

#     def name(self):
#         return "health_form"

#     @staticmethod
#     def required_slots(tracker):

#         if tracker.get_slot('confirm_exercise') == True:
#             return ["confirm_exercise", "exercise", "sleep",
#              "diet", "stress", "goal"]
#         else:
#             return ["confirm_exercise", "sleep",
#              "diet", "stress", "goal"]

#     def slot_mappings(self) -> Dict[Text, Union[Dict, List[Dict]]]:
#         """A dictionary to map required slots to
#             - an extracted entity
#             - intent: value pairs
#             - a whole message
#             or a list of them, where a first match will be picked"""

#         return {
#             "confirm_exercise": [
#                 self.from_intent(intent="affirm", value=True),
#                 self.from_intent(intent="deny", value=False),
#                 self.from_intent(intent="inform", value=True),
#             ],
#             "sleep": [
#                 self.from_entity(entity="sleep"),
#                 self.from_intent(intent="deny", value="None"),
#             ],
#             "diet": [
#                 self.from_text(intent="inform"),
#                 self.from_text(intent="affirm"),
#                 self.from_text(intent="deny"),
#             ],
#             "goal": [
#                 self.from_text(intent="inform"),
#             ],
#         }

#     def submit(
#         self,
#         dispatcher: CollectingDispatcher,
#         tracker: Tracker,
#         domain: Dict[Text, Any],
#     ) -> List[Dict]:

#         dispatcher.utter_message("Thanks, great job!")
#         return []
class COUForm(FormAction):

    def name(self):
        return "cou_form"

    @staticmethod
    def required_slots(tracker):
        return ["use_class", "use_desc", "gfa", "postal", "lotnum", "floor", "unit"]
    



        # if tracker.get_slot('confirm_exercise') == True:
        #     return ["confirm_exercise", "exercise", "sleep",
        #      "diet", "stress", "goal"]
        # else:
        #     return ["confirm_exercise", "sleep",
        #      "diet", "stress", "goal"]

    def slot_mappings(self) -> Dict[Text, Union[Dict, List[Dict]]]:
        """A dictionary to map required slots to
            - an extracted entity
            - intent: value pairs
            - a whole message
            or a list of them, where a first match will be picked"""
            
        # TO_DO validation class#
        return {
            "use_desc":[
                self.from_text()
            ], 
            "gfa": [
                self.from_text()
            ],
            "postal": [
                self.from_text()
            ],
            "lotnum": [
                self.from_text()
            ],
            "floor": [
                self.from_text()
            ],
            "unit": [
                self.from_text()
            ]
        }
            
            # "confirm_building": [
            #     self.from_intent(intent="affirm", value=True),
            #     self.from_intent(intent="deny", value=False),
            #     self.from_intent(intent="inform", value=True),
            # ],
            # # "street": [
            # #     #form backtrack should be added here
            # #     self.from_text(intent="inform"),
            # # ],
            # # "building": [
            # #     # self.from_text(intent="deny"),
            # #     self.from_intent(intent="deny", value="None"),
            # #     self.from_text(intent="inform"),
            # # ],
            # # "unit": [
            # #     self.from_text(intent="inform"),
            # # ],
            # "postal": [
            #     self.from_text(intent="inform"),
            # ],
        # }

    def submit(
        self,
        dispatcher: CollectingDispatcher,
        tracker: Tracker,
        domain: Dict[Text, Any],
    ) -> List[Dict]:

        dispatcher.utter_message("Thanks, great job!")

        use_class = tracker.get_slot("use_class")
        use_desc = tracker.get_slot("use_desc")
        gfa = float(tracker.get_slot("gfa"))
        postal = str(tracker.get_slot("postal"))
        lotnum = tracker.get_slot("lotnum")
        floor = int(tracker.get_slot("floor"))
        unit = int(tracker.get_slot("unit"))
        
        propType = self.getPropertyType(postal)
        subClassification = self.getSubmissionClassification(use_class, propType)
        if subClassification == 1 or subClassification == 2 or subClassification == 4:
            return [SlotSet("classifcation", classification_mapping[subClassification])]
        else:
            similarCases = self.getSimilarCases(use_class, use_desc, gfa, postal, lotnum, floor, unit)
            responses = self.constructResponse(similarCases)
            return [SlotSet("classifcation", classification_mapping[subClassification]),SlotSet("responses", responses if responses is not None else [])]

        # dispatcher.utter_message(f"Found these cases to be similar to your application: {responses}")
        # return []

    def getPropertyType(self, postal):
        url = "http://localhost:5000/landuse"
        req = {
            "postal": postal
        }
        # response is int from 1 - 32
        response = requests.get(url, params=req).json()
        return response

    def getSubmissionClassification(self, useClass, propType):
        url = "http://localhost:5000/query"
        req = {
            "business": useClass,
            "property": propType
        }
        # response is int from 1 - 4
        response = requests.get(url, params=req).json()
        return response

    def getSimilarCases(self, use_class, use_desc, gfa, postal, lotnum, floor, unit):
        url = "http://localhost:8080/getSimilarCases"
        req = {
            "proposedUseClass":use_class,
            "proposedUseDesc":use_desc,
            "GFA": gfa,
            "postalCode":postal,
            "lotNumber": lotnum,
            "floor":floor,
            "unit":unit
        }
        response = requests.post(url, json=req).json()
        return response

    def constructResponse(self, similarCases):
        # TODO add case details
        responses = []
        for c in similarCases:
            responses.append(c["CaseSpec"]["Id"])
        return responses
