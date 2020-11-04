import requests

def locationSpec(use_class, use_desc, gfa, postal, lotnum, floor, unit):

    API_key = "Your API Key here"
    base_url = "http://localhost:8080/getSimilarCases"
    
    # Final_url = base_url + "appid=" + API_key + "&q=" + city + "&units=metric"
    req = {
        "proposedUseClass":use_class,
		"proposedUseDesc":use_desc,
		"GFA": gfa,
        "postalCode":postal,
        "lotNumber": lotnum,
        "floor":floor,
        "unit":unit
    }
    response = requests.post(base_url, json=req)
    print(response)
    # return data depending on how the json looks like
    # return locSpec_data['main']
    

