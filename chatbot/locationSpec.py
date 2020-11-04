import requests

def locationSpec(postal, lotnum, floor, unit):

    API_key = "Your API Key here"
    base_url = "http://api.openweathermap.org/data/2.5/weather?"
    
    Final_url = base_url + "appid=" + API_key + "&q=" + city + "&units=metric"
    locSpec_data = requests.post(Final_url).json()
    

    # return data depending on how the json looks like
    return locSpec_data['main']
    

