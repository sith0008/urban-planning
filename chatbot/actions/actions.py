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

from testing import test

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
        postal = int(tracker.get_slot("postal"))
        lotnum = tracker.get_slot("lotnum")
        floor = int(tracker.get_slot("floor"))
        unit = int(tracker.get_slot("unit"))
        


        # slot variables are currently in string format. 
        testresponse = test()
        dispatcher.utter_message(testresponse)
        dispatcher.utter_message(" ")
        
        locSpec = locationSpec(postal, lotnum, floor, unit)

        # result_1, result_2 = query(use_class, use_desc, gfa, locSpec)
        
        # return whatever i want into the chatbot
        
        # dispatcher.utter_message("return whatever ")

        return []
