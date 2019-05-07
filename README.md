# go-notify
GOlang CLI that allows you to report any type of information to some group chat software like Microsoft Teams.
   
## Table of Contents
1. [Develop](#develop) <br>
2. [Usage](#usage) <br>
    2.1. [Microsoft Teams](#microsoftTeams) <br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.1.1 [messageCard](#messageCard)<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2.1.2 [adaptiveCard](#adaptiveCard)<br>
3. [Examples](#examples) <br>
    3.1. [Microsoft Teams](#mtExamples) <br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3.1.1 [Desktop](#messageCardDesktop)<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;3.1.2 [Mobile](#messageCardMobile)<br>

### 1.Develop <a name="develop"></a>

If your git command is authenticated globally:

```
go get -u github.com/hugo-lorenzo-mato/go-notify.git
```

Else:

```
mkdir -p src/github.com/hugo-lorenzo-mato
cd srcgithub.com/hugo-lorenzo-mato
git clone https://github.com/hugo-lorenzo-mato/go-notify.git
cd go-notify
export GOPATH={Absolute path to the directory immediately before .../src/ created previously}
go build
```

### 2.Usage <a name="usage"></a>

```sh
$go-notify.exe --help

NAME:
   go-notify.exe - Reports any type of incident to the configured messaging services

USAGE:
   go-notify.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     microsoftTeams  Microsoft Teams Channel report options
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d               Activate debug mode
   --config value, -c value  Yaml config file path (default: ".")
   --help, -h                show help
   --version, -v             print the version

```


##### 2.1 microsftTeams <a name="microsoftTeams"></a>

```sh
$go-notify.exe microsoftTeams --help

NAME:
   go-notify.exe microsoftTeams - Microsoft Teams Channel report options

USAGE:
   go-notify.exe microsoftTeams command [command options] [arguments...]

COMMANDS:
     messageCard   Report any type of information to Microsoft Teams Channel using messageCard card type
     adaptiveCard  Report any type of information to Microsoft Teams Channel using adaptiveCard card type

OPTIONS:
   --help, -h  show help
```


##### 2.1.1 messageCard <a name="messageCard"></a>

```
$go-notify.exe microsoftTeams messageCard --help

NAME:
   go-notify.exe microsoftTeams messageCard - Report any type of information to Microsoft Teams Channel using messageCard card type

USAGE:
   go-notify.exe microsoftTeams messageCard [command options] [arguments...]

OPTIONS:
   --hookUrl value    Channel webhook URL
   --cardTitle value  Card title
   --cardText value   Card text
   --uriAction value  URI action
   --planKey value    Plan key
```

##### 2.1.2 adaptiveCard <a name="adaptiveCard"></a>

```
NOT YET IMPLEMENTED
```

### 3.Examples <a name="examples"></a>

##### 3.1 Microsoft Teams <a name="mtExamples"></a>

###### 3.1.1 Desktop <a name="messageCardDesktop"></a>


<img src="https://i.imgur.com/cgZ6jtt.png"
     alt="Microsoft Teams Golang CLI Desktop Message Card Black"
     width="700" />
     
<img src="https://i.imgur.com/OLHcWu4.png"
     alt="Microsoft Teams Golang CLI Desktop Message Card White"
     width="600"/>
     
     
###### 3.1.2 Mobile <a name="messageCardMobile"></a>

    
<img src="https://i.imgur.com/u6kwQqe.jpg"
     alt="Microsoft Teams Golang CLI Desktop Message Card Black Mobile"
     width="300" 
     height="500"/>
     
<img src="https://i.imgur.com/mhAGaz0.jpg"
     alt="Microsoft Teams Golang CLI Desktop Message Card Black Notification"
     width="320" 
     height="230"/>
     

 
