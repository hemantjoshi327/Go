from flask import Flask
from flask import request

app = Flask(__name__)

@app.route("/", methods=['POST','GET'])
def hello():
    print (request.headers.get('x-amz-sns-message-type'))
    print("Header: ", request.headers)
    if request.method == 'POST':
        print(request.data)
    return "Hello AWS!\n"

if __name__ == "__main__":
    app.debug = True
    app.run(host='0.0.0.0',port=81)