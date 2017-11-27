from subprocess import Popen, PIPE, STDOUT
# import re
# import os

def pyDockerCmd(command):
    # wd = os.chdir("C:\\Program Files\\Docker Toolbox\\")
    # VM = 'default'
    # bash_exec = 'C:\\Program Files\\Git\\bin\\bash.exe'
    bash_exec = 'bash.exe'

    # dockerenv = Popen([bash_exec, '-c', f'docker-machine env {VM}'], cwd = wd, \
    #         stdout = PIPE, stderr = PIPE).communicate()[0].decode('utf-8')
    
    # var = re.compile('([A-Z_]+)?=').findall(dockerenv)
    # val = re.compile('"(.*)?"').findall(dockerenv)

    # env_dict = dict(zip(var, val[0:len(var)]))

    # bashenv = Popen([bash_exec, '-c', 'printenv'], cwd = wd, \
    #         stdout = PIPE, stderr = PIPE).communicate()[0].decode('utf-8').split('\n')
    # for i in bashenv:
    #     try:
    #         name,value = tuple(i.split('='))
    #         env_dict[name] = value
    #     except ValueError:
    #         pass

    # out, err = Popen([bash_exec, '-c', f'{command}'], cwd = wd, \
    #         stdout = PIPE, env = env_dict, stderr = PIPE, encoding='utf-8').communicate()
    out, err = Popen([bash_exec, '-c', f'{command}'], \
            stdout = PIPE, stderr = PIPE, encoding='utf-8').communicate()

    if err == '':
        print(out)
    else:
        print(err)
