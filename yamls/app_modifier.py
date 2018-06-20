"""
Retrieves a truncated version of the latest git commit sha and updates 
the go-bftx container image tag in app.yaml
"""

import sys
import ruamel.yaml
from subprocess import check_output

yaml_path = 'stateful-set.yaml'
for app in ruamel.yaml.round_trip_load_all(stream=open(yaml_path)):
    pass

# parse the most recent git commit sha from command line
image_tag = 'blockfreight/go-bftx:ci-cd-' + check_output('git log -1 --pretty=format:%h'.split()).decode()

# update go-bftx image with most recent git-commit-sha tag in the StatefulSet blocks
app['spec']['template']['spec']['containers'][1]['image'] = image_tag

ruamel.yaml.round_trip_dump(app, sys.stdout)