import requests

# search_text = '410641'.replace(" ", "%20")
land_use = 'NONE'

res = set()

# for i in range(1000000):
for i in range(410641, 410642):
    print(f"{i:06d}")
    search_text = f"{i:06d}"

    try:
        coord_json = None
        coord_url = 'https://developers.onemap.sg/commonapi/search?searchVal={}&returnGeom=Y&getAddrDetails=Y&pageNum=1'\
            .format(search_text)

        coord_req = requests.get(coord_url)

        coord_json = coord_req.json() if coord_req.ok and coord_req.json() != 'The request you have just typed is not allowed!' else None

    except requests.exceptions.RequestException as e:
        print(e)

    try:
        if coord_json and coord_json['results']:
            address = coord_json['results'][0]['ADDRESS']
            lat, lng = coord_json['results'][0]['LATITUDE'], coord_json['results'][0]['LONGITUDE']

            print(lat, lng)

            land_url = 'https://www.ura.gov.sg/arcgis/rest/services/MP19/Updated_Landuse_gaz/MapServer/45/' \
                       'query?returnGeometry=true&where=1%3D1&outSR=4326&outFields=*&inSr=4326&' \
                       'geometry=%7B%22x%22%3A{}%2C%22y%22%3A{}%2C%22spatialReference%22%3A%7B%22wkid%22%3A4326%7D%7D&' \
                       'geometryType=esriGeometryPoint&spatialRel=esriSpatialRelWithin&f=json'.format(lng, lat)

            land_r = requests.get(land_url).json()
            land_use = land_r['features'][0]['attributes']['LU_DESC']
    except requests.exceptions.RequestException as e:
        print(e)

    print(land_use)

    if land_use != "NONE":
        # res[land_use] = res[land_use] + [address] if land_use in res else [address]
        if land_use not in res:
            res.add(land_use)

print(res)

