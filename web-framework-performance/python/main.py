from flask import Flask
import logging

log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

app = Flask(__name__)

@app.route('/api/hello')
def hello():
    return 'Hello World'

app.run(debug=False, port=5001)
