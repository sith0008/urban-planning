import requests


def getLandUse(search_text):
    search_text = str(search_text).replace(" ", "%20")
    land_use = None

    try:
        coord_json = None
        coord_url = f'https://developers.onemap.sg/commonapi/search?searchVal={search_text}' \
                    f'&returnGeom=Y&getAddrDetails=Y&pageNum=1 '

        print(coord_url)

        coord_req = requests.get(coord_url)

        if coord_req.ok and coord_req.json() != 'The request you have just typed is not allowed!':
            coord_json = coord_req.json()
        else:
            print("OneMap request not ok or request not allowed")
            coord_json = None

    except requests.exceptions.RequestException as e:
        print(e)

    try:
        if coord_json and coord_json['results']:
            print(coord_json)
            blk, road, building_name = coord_json['results'][0]['BLK_NO'], \
                                       coord_json['results'][0]['ROAD_NAME'], \
                                       coord_json['results'][0]['BUILDING']
            print(blk, road, building_name)
            lat, lng = coord_json['results'][0]['LATITUDE'], coord_json['results'][0]['LONGITUDE']

            print(lat, lng)

            land_url = f'https://www.ura.gov.sg/arcgis/rest/services/MP19/Updated_Landuse_gaz/MapServer/45/' \
                       f'query?returnGeometry=true&where=1%3D1&outSR=4326&outFields=*&inSr=4326&' \
                       f'geometry=%7B%22x%22%3A{lng}%2C%22y%22%3A{lat}' \
                       f'%2C%22spatialReference%22%3A%7B%22wkid%22%3A4326%7D%7D&' \
                       f'geometryType=esriGeometryPoint&spatialRel=esriSpatialRelWithin&f=json'

            land_r = requests.get(land_url).json()
            print(land_r)
            land_use = land_r['features'][0]['attributes']['LU_DESC']
    except requests.exceptions.RequestException as e:
        print(e)

    print(land_use)
    return land_use


getLandUse("152 Beach Road")


