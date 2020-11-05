# app.py

from flask import Flask, request
from DecisionTree import decision
from LandUse import getLandUse

app = Flask(__name__)


@app.route('/landuse')
def landuse():
    _postal = request.args.get('postal') if request.args.get('postal') else 0
    _property = getLandUse(_postal)
    return _property


@app.route('/query')
def query():
    _business = request.args.get('business') if request.args.get('business') else 0
    _zone = request.args.get('zone') if request.args.get('zone') else 0
    _unitzone = request.args.get('unit') if request.args.get('unit') else 0
    return str(decision(_business, _zone, _unitzone))


if __name__ == "__main__":
    app.run(debug=True, port=5000)  # run app in debug mode on port 5000
