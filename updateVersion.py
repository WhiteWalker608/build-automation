import json
import sys

binaryName = sys.argv[1]
versionNumber = sys.argv[2]

with open('versions.json', 'r+') as f:
    data = json.load(f)
    data[binaryName] = versionNumber
    f.seek(0)
    json.dump(data, f, indent=4)
    f.truncate()
