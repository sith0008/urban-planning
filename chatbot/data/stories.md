## say goodbye
* goodbye
  - utter_goodbye

## bot challenge
* bot_challenge
  - utter_iamabot

## survey happy path 
<!-- to start form, instead of affirm, it should be start_form -->

<!-- * greet
    - utter_greet
* affirm
    - cou_form
    - form{"name": "cou_form"}
    - form{"name": null}
    - utter_slots_values
* thankyou
    - utter_no_worries
    - utter_goodbye -->

* greet
    - utter_greet
* start_form
    - cou_form
    - form{"name": "cou_form"}
    - form{"name": null}
    - utter_slots_values
    - utter_api_response
* thankyou
    - utter_no_worries
    - utter_goodbye


## survey stop
* greet
    - utter_greet
* start_form
    - cou_form
    - form{"name": "cou_form"}
* out_of_scope
    - utter_ask_continue
* deny
    - action_deactivate_form
    - form{"name": null}
    - utter_goodbye

## survey continue
* greet
    - utter_greet
* start_form
    - cou_form
    - form{"name": "cou_form"}
* out_of_scope
    - utter_ask_continue
* affirm
    - cou_form
    - form{"name": null}
    - utter_slots_values

## no survey
* greet
    - utter_greet
* deny
    - utter_goodbye



<!-- ## survey backtrack
* greet
    - utter_greet
* affirm
    -cou_form
    -form{"name": "cou_form"}
* mistake
    -cou_ -->