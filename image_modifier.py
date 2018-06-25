"""
Retrieves a truncated version of the latest git commit sha and updates 
the go-bftx container image tag in app.yaml
"""

import ruamel.yaml
import sys
from subprocess import check_output

yaml = ruamel.yaml.YAML()
yaml.preserve_quotes = True
yaml_file = list(yaml.load_all(open('statefulset.yaml')))

docker_image = (sys.argv[1] + ':' + sys.argv[2]) # pass in travis environment variable to the program

yaml_file[0]['spec']['template']['spec']['containers'][1]['image'] = docker_image
with open('statefulset.yaml', 'w') as f:
    yaml.dump_all(yaml_file, f)