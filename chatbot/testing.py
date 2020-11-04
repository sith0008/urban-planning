import requests

def test():

    response = requests.get('http://dummy.restapiexample.com/api/v1/employees')
    print(response.status_code)
    print(response)

    
    

    # return data depending on how the json looks like
    return response
    