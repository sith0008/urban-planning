import requests

def query(use_class, use_desc, gfa, locSpec):

    API_key = "Your API Key here"
    base_url = "http://api.openweathermap.org/data/2.5/weather?"
    
    Final_url = base_url + "appid=" + API_key + "&q=" + city + "&units=metric"
    query_data = requests.get(Final_url).json()
    

    # return data depending on how the json looks like
    return query_data['main']
    

