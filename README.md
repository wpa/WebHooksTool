WebHooksTool
============
That tool is written in __golang__. To build it you need to have golang in version at least go1.4.2.
To forward request __ngrok__ app is in use.

Steps to build
-------------
- Install required libraries

        go get github.com/yosida95/golang-jenkins
        go get code.google.com/p/gcfg
- Build app
        go build

- Download ngrok from http://ngrok.io

Steps to run it
---------------
1. Scp WebHooks, ngrok and webhooks.cfg to jenkins machine.
2. Configure app


    [Main]
    port = 8999  #listen port


    [CServer]   #Jenkins server
    url = http://jenkins:8080
    jobName = Build Job Name


    [CServerAuth]  #Credentials to use with jenkins
    username = joe_monster
    authToken =  API_TOKEN_HERE


    [Job]    #Job on jenkins to start
    targetBranch = develop
    label = To be tested
    paramName = sha1

3. Start ngrok
        ./ngrok http -log=stdout -log-level debug -inspect -log-format json 8999 > ngrok.log &

4. Start WebHooks
        ./WebHooks &
5. OPTIONAL Create tunnel to jenkins ssh jenkins_host -L 4040:localhost:4040
6. Go to page http://localhost:4040/status and check external address of tunnels
7. Use that url in github webhooks config
