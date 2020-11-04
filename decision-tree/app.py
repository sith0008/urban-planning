# app.py

from flask import Flask, request
from DecisionTree import decision


app = Flask(__name__)


@app.route('/query')
def query():
    _business = request.args.get('business') if request.args.get('business') else 0
    _property = request.args.get('property') if request.args.get('property') else 0
    _unitzone = request.args.get('unitzone') if request.args.get('unitzone') else 0
    return str(decision(_business, _property, _unitzone))


if __name__ == "__main__":
    app.run(debug=True, port=5000)  # run app in debug mode on port 5000
