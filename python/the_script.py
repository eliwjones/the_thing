# Process JSON from the_server.
import json
import os
import requests
import sys

THE_THING_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

key = sys.argv[1]
url = "http://127.0.0.1:9999"
try:
    r = requests.get(url + "/api/" + key)
    message = json.dumps(r.json())
except:
    message = "[Error] key: %s, url: %s" % (key, url)

with open(THE_THING_DIR + "/log/the_log", "a") as f:
    f.write(message + '\n')
