"""
Retrieves a truncated version of the latest git commit sha and updates 
the go-bftx container image tag in app.yaml
"""

import ruamel.yaml
import sys
from subprocess import check_output

yaml_path = 'blockfreight-statefulset.yaml'
for i, yaml_block in enumerate(ruamel.yaml.round_trip_load_all(stream=open(yaml_path))):
    pass

# parse the most recent git commit sha from command line
docker_image = 'blockfreight/go-bftx:ci-cd-' + check_output('git log -1 --pretty=format:%h'.split()).decode()

# update go-bftx image with most recent git-commit-sha tag in the StatefulSet block
yaml_block['spec']['template']['spec']['containers'][1]['image'] = docker_image

ruamel.yaml.round_trip_dump(yaml_block, sys.stdout)