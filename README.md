# Getting started

message360 API version 3

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Message360-GoLang&projectName=message360_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=message360_lib)

## How to Use

The following section explains how to use the Message360 library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=message360_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=message360_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=message360_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=message360_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=message360_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Message360-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=message360_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=message360_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [sharedshortcode_pkg](#sharedshortcode_pkg)
* [conference_pkg](#conference_pkg)
* [transcription_pkg](#transcription_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [usage_pkg](#usage_pkg)
* [webrtc_pkg](#webrtc_pkg)
* [recording_pkg](#recording_pkg)
* [email_pkg](#email_pkg)
* [sms_pkg](#sms_pkg)
* [call_pkg](#call_pkg)
* [carrier_pkg](#carrier_pkg)
* [address_pkg](#address_pkg)
* [subaccount_pkg](#subaccount_pkg)
* [account_pkg](#account_pkg)
* [shortcode_pkg](#shortcode_pkg)

## <a name="sharedshortcode_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sharedshortcode_pkg") sharedshortcode_pkg

### Get instance

Factory for the ``` SHAREDSHORTCODE ``` interface can be accessed from the package sharedshortcode_pkg.

```go
sharedShortCode := sharedshortcode_pkg.NewSHAREDSHORTCODE()
```

### <a name="view_template"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ViewTemplate") ViewTemplate

> View a Shared ShortCode Template


```go
func (me *SHAREDSHORTCODE_IMPL) ViewTemplate(input *ViewTemplateInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| templateid |  ``` Required ```  | The unique identifier for a template object |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ViewTemplateInput)

templateid := uuid.NewV4()
collect.Templateid = templateid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sharedShortCode.ViewTemplate(collect)

```


### <a name="view_shared_shortcodes"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ViewSharedShortcodes") ViewSharedShortcodes

> View a ShortCode Message


```go
func (me *SHAREDSHORTCODE_IMPL) ViewSharedShortcodes(input *ViewSharedShortcodesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message sid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ViewSharedShortcodesInput)

messagesid := "messagesid"
collect.Messagesid = messagesid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sharedShortCode.ViewSharedShortcodes(collect)

```


### <a name="list_outbound_shared_shortcodes"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ListOutboundSharedShortcodes") ListOutboundSharedShortcodes

> List ShortCode Messages


```go
func (me *SHAREDSHORTCODE_IMPL) ListOutboundSharedShortcodes(input *ListOutboundSharedShortcodesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListOutboundSharedShortcodesInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

from := "from"
collect.From = from

to := "to"
collect.To = to

datesent := "datesent"
collect.Datesent = datesent


var result string
result,_ = sharedShortCode.ListOutboundSharedShortcodes(collect)

```


### <a name="list_inbound_shared_shortcodes"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ListInboundSharedShortcodes") ListInboundSharedShortcodes

> List All Inbound ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) ListInboundSharedShortcodes(input *ListInboundSharedShortcodesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound ShortCode |
| shortcode |  ``` Optional ```  | Only list messages sent to this Short Code |
| dateReceived |  ``` Optional ```  | Only list messages sent with the specified date |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListInboundSharedShortcodesInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

from := "from"
collect.From = from

shortcode := "Shortcode"
collect.Shortcode = shortcode

dateReceived := "DateReceived"
collect.DateReceived = dateReceived


var result string
result,_ = sharedShortCode.ListInboundSharedShortcodes(collect)

```


### <a name="send_shared_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.SendSharedShortcode") SendSharedShortcode

> Send an SMS from a message360 ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) SendSharedShortcode(input *SendSharedShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | The Short Code number that is the sender of this message |
| to |  ``` Required ```  | A valid 10-digit number that should receive the message |
| templateid |  ``` Required ```  | The unique identifier for the template used for the message |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| data |  ``` Required ```  | format of your data, example: {companyname}:test,{otpcode}:1234 |
| method |  ``` Optional ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once the Short Code message is sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when Short Code message was sent. |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.SendSharedShortcodeInput)

shortcode := "shortcode"
collect.Shortcode = shortcode

to := "to"
collect.To = to

templateid := uuid.NewV4()
collect.Templateid = templateid

responseType := "json"
collect.ResponseType = responseType

data := "data"
collect.Data = data

method := models_pkg.HttpAction_GET
collect.Method = method

messageStatusCallback := "MessageStatusCallback"
collect.MessageStatusCallback = messageStatusCallback


var result string
result,_ = sharedShortCode.SendSharedShortcode(collect)

```


### <a name="list_templates"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ListTemplates") ListTemplates

> List Shortcode Templates by Type


```go
func (me *SHAREDSHORTCODE_IMPL) ListTemplates(input *ListTemplatesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| mtype |  ``` Optional ```  ``` DefaultValue ```  | The type (category) of template Valid values: marketing, authorization |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| shortcode |  ``` Optional ```  | Only list templates of type |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListTemplatesInput)

responseType := "json"
collect.ResponseType = responseType

mtype := "authorization"
collect.Mtype = mtype

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

shortcode := "Shortcode"
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.ListTemplates(collect)

```


### <a name="view_keyword"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ViewKeyword") ViewKeyword

> View a set of properties for a single keyword.


```go
func (me *SHAREDSHORTCODE_IMPL) ViewKeyword(input *ViewKeywordInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keywordid |  ``` Required ```  | The unique identifier of each keyword |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ViewKeywordInput)

keywordid := "Keywordid"
collect.Keywordid = keywordid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sharedShortCode.ViewKeyword(collect)

```


### <a name="list_keyword"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ListKeyword") ListKeyword

> Retrieve a list of keywords associated with your message360 account.


```go
func (me *SHAREDSHORTCODE_IMPL) ListKeyword(input *ListKeywordInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| keyword |  ``` Optional ```  | Only list keywords of keyword |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListKeywordInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

keyword := "Keyword"
collect.Keyword = keyword

shortcode,_ := strconv.ParseInt("172", 10, 8)
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.ListKeyword(collect)

```


### <a name="view_assignement"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ViewAssignement") ViewAssignement

> The response returned here contains all resource properties associated with the given Shortcode.


```go
func (me *SHAREDSHORTCODE_IMPL) ViewAssignement(input *ViewAssignementInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Shortcode to your message360 account |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ViewAssignementInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sharedShortCode.ViewAssignement(collect)

```


### <a name="list_assignment"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.ListAssignment") ListAssignment

> Retrieve a list of shortcode assignment associated with your message360 account.


```go
func (me *SHAREDSHORTCODE_IMPL) ListAssignment(input *ListAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListAssignmentInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

shortcode := "Shortcode"
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.ListAssignment(collect)

```


### <a name="update_assignment"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.UpdateAssignment") UpdateAssignment

> TODO: Add a method description


```go
func (me *SHAREDSHORTCODE_IMPL) UpdateAssignment(input *UpdateAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid shortcode to your message360 account |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| friendlyName |  ``` Optional ```  | User generated name of the shortcode |
| callbackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| callbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |
| fallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST. |
| fallbackUrlMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.UpdateAssignmentInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

responseType := "json"
collect.ResponseType = responseType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

callbackUrl := "CallbackUrl"
collect.CallbackUrl = callbackUrl

callbackMethod := models_pkg.HttpAction_GET
collect.CallbackMethod = callbackMethod

fallbackUrl := "FallbackUrl"
collect.FallbackUrl = fallbackUrl

fallbackUrlMethod := models_pkg.HttpAction_GET
collect.FallbackUrlMethod = fallbackUrlMethod


var result string
result,_ = sharedShortCode.UpdateAssignment(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="conference_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".conference_pkg") conference_pkg

### Get instance

Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.

```go
conference := conference_pkg.NewCONFERENCE()
```

### <a name="deaf_mute_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.DeafMuteParticipant") DeafMuteParticipant

> Deaf Mute Participant


```go
func (me *CONFERENCE_IMPL) DeafMuteParticipant(input *DeafMuteParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | ID of the active conference |
| participantSid |  ``` Required ```  | ID of an active participant |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| muted |  ``` Optional ```  | Mute a participant |
| deaf |  ``` Optional ```  | Make it so a participant cant hear |


#### Example Usage

```go
collect := new (conference_pkg.DeafMuteParticipantInput)

conferenceSid := "conferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

responseType := "json"
collect.ResponseType = responseType

muted := true
collect.Muted = muted

deaf := true
collect.Deaf = deaf


var result string
result,_ = conference.DeafMuteParticipant(collect)

```


### <a name="view_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ViewParticipant") ViewParticipant

> View Participant


```go
func (me *CONFERENCE_IMPL) ViewParticipant(input *ViewParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.ViewParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.ViewParticipant(collect)

```


### <a name="add_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.AddParticipant") AddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) AddParticipant(input *AddParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | Unique Conference Sid |
| participantnumber |  ``` Required ```  | Particiant Number |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| muted |  ``` Optional ```  | add muted |
| deaf |  ``` Optional ```  | add without volume |


#### Example Usage

```go
collect := new (conference_pkg.AddParticipantInput)

conferencesid := "conferencesid"
collect.Conferencesid = conferencesid

participantnumber := "participantnumber"
collect.Participantnumber = participantnumber

responseType := "json"
collect.ResponseType = responseType

muted := true
collect.Muted = muted

deaf := true
collect.Deaf = deaf


var result string
result,_ = conference.AddParticipant(collect)

```


### <a name="view_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ViewConference") ViewConference

> View Conference


```go
func (me *CONFERENCE_IMPL) ViewConference(input *ViewConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | The unique identifier of each conference resource |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.ViewConferenceInput)

conferencesid := "conferencesid"
collect.Conferencesid = conferencesid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.ViewConference(collect)

```


### <a name="create_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConference") CreateConference

> Here you can experiment with initiating a conference call through message360 and view the request response generated when doing so.


```go
func (me *CONFERENCE_IMPL) CreateConference(input *CreateConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
| method |  ``` Required ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once call connects. |
| recordCallbackUrl |  ``` Required ```  | Recording parameters will be sent here upon completion. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the conference is finished. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| record |  ``` Optional ```  | Specifies if the conference should be recorded. |
| recordCallbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once conference connects. |
| schdeuleTime |  ``` Optional ```  | Schedule conference in future. Schedule time must be greater than current time |
| timeout |  ``` Optional ```  | The number of seconds the call stays on the line while waiting for an answer. The max time limit is 999 and the default limit is 60 seconds but lower times can be set. |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferenceInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

method := models_pkg.HttpAction_POST
collect.Method = method

recordCallbackUrl := "RecordCallbackUrl"
collect.RecordCallbackUrl = recordCallbackUrl

responseType := "json"
collect.ResponseType = responseType

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := models_pkg.HttpAction_GET
collect.StatusCallBackMethod = statusCallBackMethod

fallBackUrl := "FallBackUrl"
collect.FallBackUrl = fallBackUrl

fallBackMethod := models_pkg.HttpAction_GET
collect.FallBackMethod = fallBackMethod

record := true
collect.Record = record

recordCallbackMethod := models_pkg.HttpAction_GET
collect.RecordCallbackMethod = recordCallbackMethod

schdeuleTime := "SchdeuleTime"
collect.SchdeuleTime = schdeuleTime

timeout,_ := strconv.ParseInt("172", 10, 8)
collect.Timeout = timeout


var result string
result,_ = conference.CreateConference(collect)

```


### <a name="hangup_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.HangupParticipant") HangupParticipant

> Remove a participant from a conference.


```go
func (me *CONFERENCE_IMPL) HangupParticipant(input *HangupParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.HangupParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.HangupParticipant(collect)

```


### <a name="play_conference_audio"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.PlayConferenceAudio") PlayConferenceAudio

> Play an audio file during a conference.


```go
func (me *CONFERENCE_IMPL) PlayConferenceAudio(input *PlayConferenceAudioInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
| audioUrl |  ``` Required ```  | The URL for the audio file that is to be played during the conference. Multiple audio files can be chained by using a semicolon. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.PlayConferenceAudioInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

audioUrl := models_pkg.AudioFormat_MP3
collect.AudioUrl = audioUrl

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.PlayConferenceAudio(collect)

```


### <a name="list_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ListParticipant") ListParticipant

> List Participant


```go
func (me *CONFERENCE_IMPL) ListParticipant(input *ListParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response format, xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | page number |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Amount of records to return per page |
| muted |  ``` Optional ```  | Participants that are muted |
| deaf |  ``` Optional ```  | Participants cant hear |


#### Example Usage

```go
collect := new (conference_pkg.ListParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

muted := true
collect.Muted = muted

deaf := true
collect.Deaf = deaf


var result string
result,_ = conference.ListParticipant(collect)

```


### <a name="list_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ListConference") ListConference

> List Conference


```go
func (me *CONFERENCE_IMPL) ListConference(input *ListConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| dateCreated |  ``` Optional ```  | Conference created date |


#### Example Usage

```go
collect := new (conference_pkg.ListConferenceInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

dateCreated := "DateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = conference.ListConference(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transcription_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

### <a name="list_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.ListTranscription") ListTranscription

> Get All transcriptions


```go
func (me *TRANSCRIPTION_IMPL) ListTranscription(input *ListTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | page number |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Amount of data per page |
| status |  ``` Optional ```  | Transcription status |
| dateTranscribed |  ``` Optional ```  | Transcription date |


#### Example Usage

```go
collect := new (transcription_pkg.ListTranscriptionInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

status := models_pkg.Status_INPROGRESS
collect.Status = status

dateTranscribed := "DateTranscribed"
collect.DateTranscribed = dateTranscribed


var result string
result,_ = transcription.ListTranscription(collect)

```


### <a name="view_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.ViewTranscription") ViewTranscription

> View Specific Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) ViewTranscription(input *ViewTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionSid |  ``` Required ```  | Unique Transcription ID |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.ViewTranscriptionInput)

transcriptionSid := "TranscriptionSid"
collect.TranscriptionSid = transcriptionSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.ViewTranscription(collect)

```


### <a name="recording_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.RecordingTranscription") RecordingTranscription

> Recording Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) RecordingTranscription(input *RecordingTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording sid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.RecordingTranscriptionInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.RecordingTranscription(collect)

```


### <a name="audio_url_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.AudioURLTranscription") AudioURLTranscription

> Audio URL Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) AudioURLTranscription(input *AudioURLTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audioUrl |  ``` Required ```  | Audio url |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.AudioURLTranscriptionInput)

audioUrl := "AudioUrl"
collect.AudioUrl = audioUrl

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.AudioURLTranscription(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="phonenumber_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

### <a name="available_phone_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.AvailablePhoneNumber") AvailablePhoneNumber

> Available Phone Number


```go
func (me *PHONENUMBER_IMPL) AvailablePhoneNumber(input *AvailablePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | Number type either SMS,Voice or all |
| areaCode |  ``` Required ```  | Phone Number Area Code |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Page Size |


#### Example Usage

```go
collect := new (phonenumber_pkg.AvailablePhoneNumberInput)

numberType := models_pkg.Number Type_ALL
collect.NumberType = numberType

areaCode := "AreaCode"
collect.AreaCode = areaCode

responseType := "json"
collect.ResponseType = responseType

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize


var result string
result,_ = phoneNumber.AvailablePhoneNumber(collect)

```


### <a name="list_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.ListNumber") ListNumber

> List Account's Phone number details


```go
func (me *PHONENUMBER_IMPL) ListNumber(input *ListNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| numberType |  ``` Optional ```  | SMS or Voice |
| friendlyName |  ``` Optional ```  | Friendly name of the number |


#### Example Usage

```go
collect := new (phonenumber_pkg.ListNumberInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

numberType := models_pkg.Number Type_ALL
collect.NumberType = numberType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName


var result string
result,_ = phoneNumber.ListNumber(collect)

```


### <a name="view_number_details"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.ViewNumberDetails") ViewNumberDetails

> Get Phone Number Details


```go
func (me *PHONENUMBER_IMPL) ViewNumberDetails(input *ViewNumberDetailsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Get Phone number Detail |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.ViewNumberDetailsInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.ViewNumberDetails(collect)

```


### <a name="release_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.ReleaseNumber") ReleaseNumber

> Release number from account


```go
func (me *PHONENUMBER_IMPL) ReleaseNumber(input *ReleaseNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be relase |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.ReleaseNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.ReleaseNumber(collect)

```


### <a name="buy_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.BuyNumber") BuyNumber

> Buy Phone Number 


```go
func (me *PHONENUMBER_IMPL) BuyNumber(input *BuyNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be purchase |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.BuyNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.BuyNumber(collect)

```


### <a name="update_phone_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.UpdatePhoneNumber") UpdatePhoneNumber

> Update Phone Number Details


```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber(input *UpdatePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | The phone number to update |
| voiceUrl |  ``` Required ```  | URL requested once the call connects |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| friendlyName |  ``` Optional ```  | Phone number friendly name description |
| voiceMethod |  ``` Optional ```  | Post or Get |
| voiceFallbackUrl |  ``` Optional ```  | URL requested if the voice URL is not available |
| voiceFallbackMethod |  ``` Optional ```  | Post or Get |
| hangupCallback |  ``` Optional ```  | callback url after a hangup occurs |
| hangupCallbackMethod |  ``` Optional ```  | Post or Get |
| heartbeatUrl |  ``` Optional ```  | URL requested once the call connects |
| heartbeatMethod |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time |
| smsUrl |  ``` Optional ```  | URL requested when an SMS is received |
| smsMethod |  ``` Optional ```  | Post or Get |
| smsFallbackUrl |  ``` Optional ```  | URL requested once the call connects |
| smsFallbackMethod |  ``` Optional ```  | URL requested if the sms URL is not available |


#### Example Usage

```go
collect := new (phonenumber_pkg.UpdatePhoneNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

voiceUrl := "VoiceUrl"
collect.VoiceUrl = voiceUrl

responseType := "json"
collect.ResponseType = responseType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

voiceMethod := models_pkg.HttpAction_GET
collect.VoiceMethod = voiceMethod

voiceFallbackUrl := "VoiceFallbackUrl"
collect.VoiceFallbackUrl = voiceFallbackUrl

voiceFallbackMethod := models_pkg.HttpAction_GET
collect.VoiceFallbackMethod = voiceFallbackMethod

hangupCallback := "HangupCallback"
collect.HangupCallback = hangupCallback

hangupCallbackMethod := models_pkg.HttpAction_GET
collect.HangupCallbackMethod = hangupCallbackMethod

heartbeatUrl := "HeartbeatUrl"
collect.HeartbeatUrl = heartbeatUrl

heartbeatMethod := models_pkg.HttpAction_GET
collect.HeartbeatMethod = heartbeatMethod

smsUrl := "SmsUrl"
collect.SmsUrl = smsUrl

smsMethod := models_pkg.HttpAction_GET
collect.SmsMethod = smsMethod

smsFallbackUrl := "SmsFallbackUrl"
collect.SmsFallbackUrl = smsFallbackUrl

smsFallbackMethod := models_pkg.HttpAction_GET
collect.SmsFallbackMethod = smsFallbackMethod


var result string
result,_ = phoneNumber.UpdatePhoneNumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="usage_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".usage_pkg") usage_pkg

### Get instance

Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.

```go
usage := usage_pkg.NewUSAGE()
```

### <a name="list_usage"></a>![Method: ](https://apidocs.io/img/method.png ".usage_pkg.ListUsage") ListUsage

> Get all usage 


```go
func (me *USAGE_IMPL) ListUsage(input *ListUsageInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| productCode |  ``` Optional ```  ``` DefaultValue ```  | Product Code |
| startDate |  ``` Optional ```  ``` DefaultValue ```  | Start Usage Date |
| endDate |  ``` Optional ```  ``` DefaultValue ```  | End Usage Date |


#### Example Usage

```go
collect := new (usage_pkg.ListUsageInput)

responseType := "json"
collect.ResponseType = responseType

productCode := models_pkg.Product Code_ALL
collect.ProductCode = productCode

startDate := "2016-09-06"
collect.StartDate = startDate

endDate := "2016-09-06"
collect.EndDate = endDate


var result string
result,_ = usage.ListUsage(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="webrtc_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".webrtc_pkg") webrtc_pkg

### Get instance

Factory for the ``` WEBRTC ``` interface can be accessed from the package webrtc_pkg.

```go
webRTC := webrtc_pkg.NewWEBRTC()
```

### <a name="check_funds"></a>![Method: ](https://apidocs.io/img/method.png ".webrtc_pkg.CheckFunds") CheckFunds

> TODO: Add a method description


```go
func (me *WEBRTC_IMPL) CheckFunds(input *CheckFundsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your message360 Account SID |
| authToken |  ``` Required ```  | Your message360 Token |


#### Example Usage

```go
collect := new (webrtc_pkg.CheckFundsInput)

accountSid := "account_sid"
collect.AccountSid = accountSid

authToken := "auth_token"
collect.AuthToken = authToken


var result string
result,_ = webRTC.CheckFunds(collect)

```


### <a name="create_token"></a>![Method: ](https://apidocs.io/img/method.png ".webrtc_pkg.CreateToken") CreateToken

> message360 webrtc


```go
func (me *WEBRTC_IMPL) CreateToken(input *CreateTokenInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your message360 Account SID |
| authToken |  ``` Required ```  | Your message360 Token |
| username |  ``` Required ```  | WebRTC username authentication |
| password |  ``` Required ```  | WebRTC password authentication |


#### Example Usage

```go
collect := new (webrtc_pkg.CreateTokenInput)

accountSid := "account_sid"
collect.AccountSid = accountSid

authToken := "auth_token"
collect.AuthToken = authToken

username := "username"
collect.Username = username

password := "password"
collect.Password = password


var result string
result,_ = webRTC.CreateToken(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="recording_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recording_pkg") recording_pkg

### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

### <a name="view_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.ViewRecording") ViewRecording

> View a specific Recording


```go
func (me *RECORDING_IMPL) ViewRecording(input *ViewRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Search Recording sid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.ViewRecordingInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.ViewRecording(collect)

```


### <a name="delete_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.DeleteRecording") DeleteRecording

> Delete Recording Record


```go
func (me *RECORDING_IMPL) DeleteRecording(input *DeleteRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording Sid to be delete |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.DeleteRecordingInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.DeleteRecording(collect)

```


### <a name="list_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.ListRecording") ListRecording

> List out Recordings


```go
func (me *RECORDING_IMPL) ListRecording(input *ListRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| dateCreated |  ``` Optional ```  | Recording date |
| callSid |  ``` Optional ```  | Call ID |


#### Example Usage

```go
collect := new (recording_pkg.ListRecordingInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

dateCreated := "DateCreated"
collect.DateCreated = dateCreated

callSid := "CallSid"
collect.CallSid = callSid


var result string
result,_ = recording.ListRecording(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="email_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".email_pkg") email_pkg

### Get instance

Factory for the ``` EMAIL ``` interface can be accessed from the package email_pkg.

```go
email := email_pkg.NewEMAIL()
```

### <a name="delete_spam"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteSpam") DeleteSpam

> Deletes a email address marked as spam from the spam list


```go
func (me *EMAIL_IMPL) DeleteSpam(input *DeleteSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| email |  ``` Required ```  | Email address |


#### Example Usage

```go
collect := new (email_pkg.DeleteSpamInput)

responseType := "json"
collect.ResponseType = responseType

email := "email"
collect.Email = email


var result string
result,_ = email.DeleteSpam(collect)

```


### <a name="delete_block"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteBlock") DeleteBlock

> Deletes a blocked email


```go
func (me *EMAIL_IMPL) DeleteBlock(input *DeleteBlockInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address to remove from block list |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.DeleteBlockInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.DeleteBlock(collect)

```


### <a name="add_unsubscribes"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.AddUnsubscribes") AddUnsubscribes

> Add an email to the unsubscribe list


```go
func (me *EMAIL_IMPL) AddUnsubscribes(input *AddUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to add to the unsubscribe list |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.AddUnsubscribesInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.AddUnsubscribes(collect)

```


### <a name="send_email"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.SendEmail") SendEmail

> Send out an email


```go
func (me *EMAIL_IMPL) SendEmail(input *SendEmailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | The to email address |
| from |  ``` Required ```  | The from email address |
| mtype |  ``` Required ```  ``` DefaultValue ```  | email format type, html or text |
| subject |  ``` Required ```  | Email subject |
| message |  ``` Required ```  | The body of the email message |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| cc |  ``` Optional ```  | CC Email address |
| bcc |  ``` Optional ```  | BCC Email address |
| attachment |  ``` Optional ```  | File to be attached.File must be less than 7MB. |


#### Example Usage

```go
collect := new (email_pkg.SendEmailInput)

to := "to"
collect.To = to

from := "from"
collect.From = from

mtype := models_pkg.SendEmailAs_HTML
collect.Mtype = mtype

subject := "subject"
collect.Subject = subject

message := "message"
collect.Message = message

responseType := "json"
collect.ResponseType = responseType

cc := "cc"
collect.Cc = cc

bcc := "bcc"
collect.Bcc = bcc

attachment := "attachment"
collect.Attachment = attachment


var result string
result,_ = email.SendEmail(collect)

```


### <a name="delete_unsubscribes"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteUnsubscribes") DeleteUnsubscribes

> Delete emails from the unsubscribe list


```go
func (me *EMAIL_IMPL) DeleteUnsubscribes(input *DeleteUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to remove from the unsubscribe list |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.DeleteUnsubscribesInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.DeleteUnsubscribes(collect)

```


### <a name="list_unsubscribes"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListUnsubscribes") ListUnsubscribes

> List all unsubscribed email addresses


```go
func (me *EMAIL_IMPL) ListUnsubscribes(input *ListUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | Starting record of the list |
| limit |  ``` Optional ```  | Maximum number of records to be returned |


#### Example Usage

```go
collect := new (email_pkg.ListUnsubscribesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.ListUnsubscribes(collect)

```


### <a name="list_invalid"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListInvalid") ListInvalid

> List out all invalid email addresses


```go
func (me *EMAIL_IMPL) ListInvalid(input *ListInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offet |  ``` Optional ```  | Starting record for listing out emails |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.ListInvalidInput)

responseType := "json"
collect.ResponseType = responseType

offet := "offet"
collect.Offet = offet

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.ListInvalid(collect)

```


### <a name="delete_bounces"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteBounces") DeleteBounces

> Delete an email address from the bounced address list


```go
func (me *EMAIL_IMPL) DeleteBounces(input *DeleteBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| email |  ``` Required ```  | The email address to remove from the bounce list |


#### Example Usage

```go
collect := new (email_pkg.DeleteBouncesInput)

responseType := "json"
collect.ResponseType = responseType

email := "email"
collect.Email = email


var result string
result,_ = email.DeleteBounces(collect)

```


### <a name="list_bounces"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListBounces") ListBounces

> List out all email addresses that have bounced


```go
func (me *EMAIL_IMPL) ListBounces(input *ListBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The record to start the list at |
| limit |  ``` Optional ```  | The maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.ListBouncesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.ListBounces(collect)

```


### <a name="list_spam"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListSpam") ListSpam

> List out all email addresses marked as spam


```go
func (me *EMAIL_IMPL) ListSpam(input *ListSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The record number to start the list at |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.ListSpamInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.ListSpam(collect)

```


### <a name="list_blocks"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListBlocks") ListBlocks

> Outputs email addresses on your blocklist


```go
func (me *EMAIL_IMPL) ListBlocks(input *ListBlocksInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | Where to start in the output list |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.ListBlocksInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.ListBlocks(collect)

```


### <a name="delete_invalid"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteInvalid") DeleteInvalid

> This endpoint allows you to delete entries in the Invalid Emails list.


```go
func (me *EMAIL_IMPL) DeleteInvalid(input *DeleteInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email that was marked invalid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Json or xml |


#### Example Usage

```go
collect := new (email_pkg.DeleteInvalidInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.DeleteInvalid(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sms_pkg") sms_pkg

### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

### <a name="send_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.SendSMS") SendSMS

> Send an SMS from a message360 number


```go
func (me *SMS_IMPL) SendSMS(input *SendSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | SMS enabled Message360 number to send this message from |
| to |  ``` Required ```  | Number to send the SMS to |
| body |  ``` Required ```  | Text Message To Send |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |
| smartsms |  ``` Optional ```  ``` DefaultValue ```  | Check's 'to' number can receive sms or not using Carrier API, if wireless = true then text sms is sent, else wireless = false then call is recieved to end user with audible message. |


#### Example Usage

```go
collect := new (sms_pkg.SendSMSInput)

from := "from"
collect.From = from

to := "to"
collect.To = to

body := "body"
collect.Body = body

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_GET
collect.Method = method

messagestatuscallback := "messagestatuscallback"
collect.Messagestatuscallback = messagestatuscallback

smartsms := false
collect.Smartsms = smartsms


var result string
result,_ = sMS.SendSMS(collect)

```


### <a name="view_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ViewSMS") ViewSMS

> View a Particular SMS


```go
func (me *SMS_IMPL) ViewSMS(input *ViewSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message sid |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.ViewSMSInput)

messagesid := "messagesid"
collect.Messagesid = messagesid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.ViewSMS(collect)

```


### <a name="list_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ListSMS") ListSMS

> List All SMS


```go
func (me *SMS_IMPL) ListSMS(input *ListSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |


#### Example Usage

```go
collect := new (sms_pkg.ListSMSInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

from := "from"
collect.From = from

to := "to"
collect.To = to

datesent := "datesent"
collect.Datesent = datesent


var result string
result,_ = sMS.ListSMS(collect)

```


### <a name="list_inbound_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ListInboundSMS") ListInboundSMS

> List All Inbound SMS


```go
func (me *SMS_IMPL) ListInboundSMS(input *ListInboundSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound SMS |
| to |  ``` Optional ```  | To Number to get Inbound SMS |
| dateSent |  ``` Optional ```  | Filter sms message objects by this date. |


#### Example Usage

```go
collect := new (sms_pkg.ListInboundSMSInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

from := "from"
collect.From = from

to := "to"
collect.To = to

dateSent := "DateSent"
collect.DateSent = dateSent


var result string
result,_ = sMS.ListInboundSMS(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="call_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".call_pkg") call_pkg

### Get instance

Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.

```go
call := call_pkg.NewCALL()
```

### <a name="make_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.MakeCall") MakeCall

> You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) MakeCall(input *MakeCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed tim |
| heartBeatMethod |  ``` Optional ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Optional ```  | Time (in seconds) Message360 should wait while the call is ringing before canceling the call |
| playDtmf |  ``` Optional ```  | DTMF Digits to play to the call once it connects. 0-9, #, or * |
| hideCallerId |  ``` Optional ```  | Specifies if the caller id will be hidden |
| record |  ``` Optional ```  | Specifies if the call should be recorded |
| recordCallBackUrl |  ``` Optional ```  | Recording parameters will be sent here upon completion |
| recordCallBackMethod |  ``` Optional ```  | Method used to request the RecordCallback URL. |
| transcribe |  ``` Optional ```  | Specifies if the call recording should be transcribed |
| transcribeCallBackUrl |  ``` Optional ```  | Transcription parameters will be sent here upon completion |
| ifMachine |  ``` Optional ```  | How Message360 should handle the receiving numbers voicemail machine |
| ifMachineUrl |  ``` Optional ```  | URL requested when IfMachine=continue |
| ifMachineMethod |  ``` Optional ```  | Method used to request the IfMachineUrl. |
| feedback |  ``` Optional ```  | Specify if survey should be enable or not |
| surveyId |  ``` Optional ```  | The unique identifier for the survey. |


#### Example Usage

```go
collect := new (call_pkg.MakeCallInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_GET
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := models_pkg.HttpAction_GET
collect.StatusCallBackMethod = statusCallBackMethod

fallBackUrl := "FallBackUrl"
collect.FallBackUrl = fallBackUrl

fallBackMethod := models_pkg.HttpAction_GET
collect.FallBackMethod = fallBackMethod

heartBeatUrl := "HeartBeatUrl"
collect.HeartBeatUrl = heartBeatUrl

heartBeatMethod := models_pkg.HttpAction_GET
collect.HeartBeatMethod = heartBeatMethod

timeout,_ := strconv.ParseInt("9", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := false
collect.HideCallerId = hideCallerId

record := false
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := models_pkg.HttpAction_GET
collect.RecordCallBackMethod = recordCallBackMethod

transcribe := false
collect.Transcribe = transcribe

transcribeCallBackUrl := "TranscribeCallBackUrl"
collect.TranscribeCallBackUrl = transcribeCallBackUrl

ifMachine := models_pkg.ifMachine_CONTINUE
collect.IfMachine = ifMachine

ifMachineUrl := "IfMachineUrl"
collect.IfMachineUrl = ifMachineUrl

ifMachineMethod := models_pkg.HttpAction_GET
collect.IfMachineMethod = ifMachineMethod

feedback := false
collect.Feedback = feedback

surveyId := "SurveyId"
collect.SurveyId = surveyId


var result string
result,_ = call.MakeCall(collect)

```


### <a name="play_audio"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.PlayAudio") PlayAudio

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) PlayAudio(input *PlayAudioInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| audioUrl |  ``` Required ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |
| sayText |  ``` Required ```  | Valid alphanumeric string that should be played to the In-progress call. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| length |  ``` Optional ```  | Time limit in seconds for audio play back |
| direction |  ``` Optional ```  | The leg of the call audio will be played to |
| mix |  ``` Optional ```  | If false, all other audio will be muted |


#### Example Usage

```go
collect := new (call_pkg.PlayAudioInput)

callSid := "CallSid"
collect.CallSid = callSid

audioUrl := "AudioUrl"
collect.AudioUrl = audioUrl

sayText := "SayText"
collect.SayText = sayText

responseType := "json"
collect.ResponseType = responseType

length,_ := strconv.ParseInt("9", 10, 8)
collect.Length = length

direction := models_pkg.Direction_IN
collect.Direction = direction

mix := false
collect.Mix = mix


var result string
result,_ = call.PlayAudio(collect)

```


### <a name="record_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.RecordCall") RecordCall

> Record a Call


```go
func (me *CALL_IMPL) RecordCall(input *RecordCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| record |  ``` Required ```  | Set true to initiate recording or false to terminate recording |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response format, xml or json |
| direction |  ``` Optional ```  | The leg of the call to record |
| timeLimit |  ``` Optional ```  | Time in seconds the recording duration should not exceed |
| callBackUrl |  ``` Optional ```  | URL consulted after the recording completes |
| fileformat |  ``` Optional ```  | Format of the recording file. Can be .mp3 or .wav |


#### Example Usage

```go
collect := new (call_pkg.RecordCallInput)

callSid := "CallSid"
collect.CallSid = callSid

record := false
collect.Record = record

responseType := "json"
collect.ResponseType = responseType

direction := models_pkg.Direction_IN
collect.Direction = direction

timeLimit,_ := strconv.ParseInt("9", 10, 8)
collect.TimeLimit = timeLimit

callBackUrl := "CallBackUrl"
collect.CallBackUrl = callBackUrl

fileformat := models_pkg.AudioFormat_MP3
collect.Fileformat = fileformat


var result string
result,_ = call.RecordCall(collect)

```


### <a name="voice_effect"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.VoiceEffect") VoiceEffect

> Voice Effect


```go
func (me *CALL_IMPL) VoiceEffect(input *VoiceEffectInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the in-progress voice call. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| audioDirection |  ``` Optional ```  | The direction the audio effect should be placed on. If IN, the effects will occur on the incoming audio stream. If OUT, the effects will occur on the outgoing audio stream. |
| pitchSemiTones |  ``` Optional ```  | value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | value between -1 and 1 |
| pitch |  ``` Optional ```  | value greater than 0 |
| rate |  ``` Optional ```  | value greater than 0 |
| tempo |  ``` Optional ```  | value greater than 0 |


#### Example Usage

```go
collect := new (call_pkg.VoiceEffectInput)

callSid := "CallSid"
collect.CallSid = callSid

responseType := "json"
collect.ResponseType = responseType

audioDirection := models_pkg.AudioDirection_IN
collect.AudioDirection = audioDirection

pitchSemiTones := 9.41756666145174
collect.PitchSemiTones = pitchSemiTones

pitchOctaves := 9.41756666145174
collect.PitchOctaves = pitchOctaves

pitch := 9.41756666145174
collect.Pitch = pitch

rate := 9.41756666145174
collect.Rate = rate

tempo := 9.41756666145174
collect.Tempo = tempo


var result string
result,_ = call.VoiceEffect(collect)

```


### <a name="send_digit"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.SendDigit") SendDigit

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) SendDigit(input *SendDigitInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |


#### Example Usage

```go
collect := new (call_pkg.SendDigitInput)

callSid := "CallSid"
collect.CallSid = callSid

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

responseType := "json"
collect.ResponseType = responseType

playDtmfDirection := models_pkg.Direction_IN
collect.PlayDtmfDirection = playDtmfDirection


var result string
result,_ = call.SendDigit(collect)

```


### <a name="interrupted_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.InterruptedCall") InterruptedCall

> Interrupt the Call by Call Sid


```go
func (me *CALL_IMPL) InterruptedCall(input *InterruptedCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | Call SId |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |


#### Example Usage

```go
collect := new (call_pkg.InterruptedCallInput)

callSid := "CallSid"
collect.CallSid = callSid

responseType := "json"
collect.ResponseType = responseType

url := "Url"
collect.Url = url

method := models_pkg.HttpAction_GET
collect.Method = method

status := models_pkg.InterruptedCallStatus_CANCELED
collect.Status = status


var result string
result,_ = call.InterruptedCall(collect)

```


### <a name="group_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.GroupCall") GroupCall

> Group Call


```go
func (me *CALL_IMPL) GroupCall(input *GroupCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| to |  ``` Required ```  | Please enter multiple E164 number. You can add max 10 numbers. Add numbers separated with comma. e.g : 1111111111,2222222222 |
| url |  ``` Required ```  | URL requested once the call connects |
| responseType |  ``` Required ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| groupConfirmKey |  ``` Required ```  | Define the DTMF that the called party should send to bridge the call. Allowed Values : 0-9, #, * |
| groupConfirmFile |  ``` Required ```  | Specify the audio file you want to play when the called party picks up the call |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed tim |
| heartBeatMethod |  ``` Optional ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Optional ```  | Time (in seconds) Message360 should wait while the call is ringing before canceling the call |
| playDtmf |  ``` Optional ```  | DTMF Digits to play to the call once it connects. 0-9, #, or * |
| hideCallerId |  ``` Optional ```  | Specifies if the caller id will be hidden |
| record |  ``` Optional ```  | Specifies if the call should be recorded |
| recordCallBackUrl |  ``` Optional ```  | Recording parameters will be sent here upon completion |
| recordCallBackMethod |  ``` Optional ```  | Method used to request the RecordCallback URL. |
| transcribe |  ``` Optional ```  | Specifies if the call recording should be transcribed |
| transcribeCallBackUrl |  ``` Optional ```  | Transcription parameters will be sent here upon completion |


#### Example Usage

```go
collect := new (call_pkg.GroupCallInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

responseType := "json"
collect.ResponseType = responseType

groupConfirmKey := "GroupConfirmKey"
collect.GroupConfirmKey = groupConfirmKey

groupConfirmFile := models_pkg.AudioFormat_MP3
collect.GroupConfirmFile = groupConfirmFile

method := models_pkg.HttpAction_GET
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := models_pkg.HttpAction_GET
collect.StatusCallBackMethod = statusCallBackMethod

fallBackUrl := "FallBackUrl"
collect.FallBackUrl = fallBackUrl

fallBackMethod := models_pkg.HttpAction_GET
collect.FallBackMethod = fallBackMethod

heartBeatUrl := "HeartBeatUrl"
collect.HeartBeatUrl = heartBeatUrl

heartBeatMethod := models_pkg.HttpAction_GET
collect.HeartBeatMethod = heartBeatMethod

timeout,_ := strconv.ParseInt("222", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := "HideCallerId"
collect.HideCallerId = hideCallerId

record := true
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := models_pkg.HttpAction_GET
collect.RecordCallBackMethod = recordCallBackMethod

transcribe := true
collect.Transcribe = transcribe

transcribeCallBackUrl := "TranscribeCallBackUrl"
collect.TranscribeCallBackUrl = transcribeCallBackUrl


var result string
result,_ = call.GroupCall(collect)

```


### <a name="list_calls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.ListCalls") ListCalls

> A list of calls associated with your Message360 account


```go
func (me *CALL_IMPL) ListCalls(input *ListCallsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Only list calls to this number |
| from |  ``` Optional ```  | Only list calls from this number |
| dateCreated |  ``` Optional ```  | Only list calls starting within the specified date range |


#### Example Usage

```go
collect := new (call_pkg.ListCallsInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

to := "To"
collect.To = to

from := "From"
collect.From = from

dateCreated := "DateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = call.ListCalls(collect)

```


### <a name="send_ringless_vm"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.SendRinglessVM") SendRinglessVM

> API endpoint used to send a Ringless Voicemail


```go
func (me *CALL_IMPL) SendRinglessVM(input *SendRinglessVMInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| rVMCallerId |  ``` Required ```  | Alternate caller ID required for RVM |
| to |  ``` Required ```  | To number |
| voiceMailURL |  ``` Required ```  | URL to an audio file |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  ``` DefaultValue ```  | Not currently used in this version |
| statusCallBackUrl |  ``` Optional ```  | URL to post the status of the Ringless request |
| statsCallBackMethod |  ``` Optional ```  | POST or GET |


#### Example Usage

```go
collect := new (call_pkg.SendRinglessVMInput)

from := "From"
collect.From = from

rVMCallerId := "RVMCallerId"
collect.RVMCallerId = rVMCallerId

to := "To"
collect.To = to

voiceMailURL := "VoiceMailURL"
collect.VoiceMailURL = voiceMailURL

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_GET
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statsCallBackMethod := models_pkg.HttpAction_GET
collect.StatsCallBackMethod = statsCallBackMethod


var result string
result,_ = call.SendRinglessVM(collect)

```


### <a name="view_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.ViewCall") ViewCall

> View Call Response


```go
func (me *CALL_IMPL) ViewCall(input *ViewCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | Call Sid id for particular Call |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.ViewCallInput)

callsid := "callsid"
collect.Callsid = callsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.ViewCall(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="carrier_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg

### Get instance

Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.

```go
carrier := carrier_pkg.NewCARRIER()
```

### <a name="carrier_lookup_list"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CarrierLookupList") CarrierLookupList

> Get the All Purchase Number's Carrier lookup


```go
func (me *CARRIER_IMPL) CarrierLookupList(input *CarrierLookupListInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Page Number |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Page Size |


#### Example Usage

```go
collect := new (carrier_pkg.CarrierLookupListInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize


var result string
result,_ = carrier.CarrierLookupList(collect)

```


### <a name="carrier_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CarrierLookup") CarrierLookup

> Get the Carrier Lookup


```go
func (me *CARRIER_IMPL) CarrierLookup(input *CarrierLookupInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | The number to lookup |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (carrier_pkg.CarrierLookupInput)

phonenumber := "phonenumber"
collect.Phonenumber = phonenumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = carrier.CarrierLookup(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="address_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".address_pkg") address_pkg

### Get instance

Factory for the ``` ADDRESS ``` interface can be accessed from the package address_pkg.

```go
address := address_pkg.NewADDRESS()
```

### <a name="create_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateAddress") CreateAddress

> To add an address to your address book, you create a new address object. You can retrieve and delete individual addresses as well as get a list of addresses. Addresses are identified by a unique random ID.


```go
func (me *ADDRESS_IMPL) CreateAddress(input *CreateAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| name |  ``` Required ```  | Name of user |
| address |  ``` Required ```  | Address of user. |
| country |  ``` Required ```  | Must be a 2 letter country short-name code (ISO 3166) |
| state |  ``` Required ```  | Must be a 2 letter State eg. CA for US. For Some Countries it can be greater than 2 letters. |
| city |  ``` Required ```  | City Name. |
| zip |  ``` Required ```  | Zip code of city. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type either json or xml |
| description |  ``` Optional ```  | Description of addresses. |
| email |  ``` Optional ```  | Email Id of user. |
| phone |  ``` Optional ```  | Phone number of user. |


#### Example Usage

```go
collect := new (address_pkg.CreateAddressInput)

name := "Name"
collect.Name = name

address := "Address"
collect.Address = address

country := "Country"
collect.Country = country

state := "State"
collect.State = state

city := "City"
collect.City = city

zip := "Zip"
collect.Zip = zip

responseType := "json"
collect.ResponseType = responseType

description := "Description"
collect.Description = description

email := "email"
collect.Email = email

phone := "Phone"
collect.Phone = phone


var result string
result,_ = address.CreateAddress(collect)

```


### <a name="view_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.ViewAddress") ViewAddress

> View Address Specific address Book by providing the address id


```go
func (me *ADDRESS_IMPL) ViewAddress(input *ViewAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressSID |  ``` Required ```  | The identifier of the address to be retrieved. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.ViewAddressInput)

addressSID := "AddressSID"
collect.AddressSID = addressSID

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.ViewAddress(collect)

```


### <a name="list_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.ListAddress") ListAddress

> List All Address 


```go
func (me *ADDRESS_IMPL) ListAddress(input *ListAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| page |  ``` Optional ```  ``` DefaultValue ```  | Return requested # of items starting the value, default=0, must be an integer |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | How many results to return, default is 10, max is 100, must be an integer |
| addressSID |  ``` Optional ```  | addresses Sid |
| dateCreated |  ``` Optional ```  | date created address. |


#### Example Usage

```go
collect := new (address_pkg.ListAddressInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

addressSID := "AddressSID"
collect.AddressSID = addressSID

dateCreated := "DateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = address.ListAddress(collect)

```


### <a name="verify_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.VerifyAddress") VerifyAddress

> Validates an address given.


```go
func (me *ADDRESS_IMPL) VerifyAddress(input *VerifyAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressSID |  ``` Required ```  | The identifier of the address to be verified. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.VerifyAddressInput)

addressSID := "AddressSID"
collect.AddressSID = addressSID

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.VerifyAddress(collect)

```


### <a name="delete_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.DeleteAddress") DeleteAddress

> To delete Address to your address book


```go
func (me *ADDRESS_IMPL) DeleteAddress(input *DeleteAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressSID |  ``` Required ```  | The identifier of the address to be deleted. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.DeleteAddressInput)

addressSID := "AddressSID"
collect.AddressSID = addressSID

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.DeleteAddress(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="subaccount_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subaccount_pkg") subaccount_pkg

### Get instance

Factory for the ``` SUBACCOUNT ``` interface can be accessed from the package subaccount_pkg.

```go
subAccount := subaccount_pkg.NewSUBACCOUNT()
```

### <a name="delete_sub_account"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.DeleteSubAccount") DeleteSubAccount

> Delete sub account or merge numbers into parent


```go
func (me *SUBACCOUNT_IMPL) DeleteSubAccount(input *DeleteSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be deleted |
| mergeNumber |  ``` Required ```  ``` DefaultValue ```  | 0 to delete or 1 to merge numbers to parent account. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (subaccount_pkg.DeleteSubAccountInput)

subAccountSID := "SubAccountSID"
collect.SubAccountSID = subAccountSID

mergeNumber := models_pkg.MergeNumberStatus_DELETE
collect.MergeNumber = mergeNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = subAccount.DeleteSubAccount(collect)

```


### <a name="suspend_sub_account"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.SuspendSubAccount") SuspendSubAccount

> Suspend or unsuspend


```go
func (me *SUBACCOUNT_IMPL) SuspendSubAccount(input *SuspendSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be activated or suspended |
| activate |  ``` Required ```  ``` DefaultValue ```  | 0 to suspend or 1 to activate |
| responseType |  ``` Required ```  ``` DefaultValue ```  | TODO: Add a parameter description |


#### Example Usage

```go
collect := new (subaccount_pkg.SuspendSubAccountInput)

subAccountSID := "SubAccountSID"
collect.SubAccountSID = subAccountSID

activate := models_pkg.ActivateStatus_DEACTIVATE
collect.Activate = activate

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = subAccount.SuspendSubAccount(collect)

```


### <a name="create_sub_account"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateSubAccount") CreateSubAccount

> Create a sub user account under the parent account


```go
func (me *SUBACCOUNT_IMPL) CreateSubAccount(input *CreateSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| firstName |  ``` Required ```  | Sub account user first name |
| lastName |  ``` Required ```  | sub account user last name |
| email |  ``` Required ```  | Sub account email address |
| friendlyName |  ``` Required ```  | Descriptive name of the sub account |
| password |  ``` Required ```  | The password of the sub account.  Please make sure to make as password that is at least 8 characters long, contain a symbol, uppercase and a number. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateSubAccountInput)

firstName := "FirstName"
collect.FirstName = firstName

lastName := "LastName"
collect.LastName = lastName

email := "Email"
collect.Email = email

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

password := "Password"
collect.Password = password

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = subAccount.CreateSubAccount(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="account_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".account_pkg") account_pkg

### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

### <a name="view_account"></a>![Method: ](https://apidocs.io/img/method.png ".account_pkg.ViewAccount") ViewAccount

> Display Account Description


```go
func (me *ACCOUNT_IMPL) ViewAccount(input *ViewAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (account_pkg.ViewAccountInput)

date := "Date"
collect.Date = date

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = account.ViewAccount(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="shortcode_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".shortcode_pkg") shortcode_pkg

### Get instance

Factory for the ``` SHORTCODE ``` interface can be accessed from the package shortcode_pkg.

```go
shortCode := shortcode_pkg.NewSHORTCODE()
```

### <a name="send_dedicated_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.SendDedicatedShortcode") SendDedicatedShortcode

> TODO: Add a method description


```go
func (me *SHORTCODE_IMPL) SendDedicatedShortcode(input *SendDedicatedShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | Your dedicated shortcode |
| to |  ``` Required ```  | The number to send your SMS to |
| body |  ``` Required ```  | The body of your message |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  | Callback status method, POST or GET |
| messagestatuscallback |  ``` Optional ```  | Callback url for SMS status |


#### Example Usage

```go
collect := new (shortcode_pkg.SendDedicatedShortcodeInput)

shortcode,_ := strconv.ParseInt("222", 10, 8)
collect.Shortcode = shortcode

to := 222.694393434885
collect.To = to

body := "body"
collect.Body = body

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_GET
collect.Method = method

messagestatuscallback := "messagestatuscallback"
collect.Messagestatuscallback = messagestatuscallback


var result string
result,_ = shortCode.SendDedicatedShortcode(collect)

```


### <a name="view_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.ViewShortcode") ViewShortcode

> View a single Sms Short Code message.


```go
func (me *SHORTCODE_IMPL) ViewShortcode(input *ViewShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for the sms resource |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (shortcode_pkg.ViewShortcodeInput)

messageSid := "MessageSid"
collect.MessageSid = messageSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = shortCode.ViewShortcode(collect)

```


### <a name="list_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.ListShortcode") ListShortcode

> Retrieve a list of Short Code message objects.


```go
func (me *SHORTCODE_IMPL) ListShortcode(input *ListShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| shortcode |  ``` Optional ```  | Only list messages sent from this Short Code |
| to |  ``` Optional ```  | Only list messages sent to this number |
| dateSent |  ``` Optional ```  | Only list messages sent with the specified date |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (shortcode_pkg.ListShortcodeInput)

responseType := "json"
collect.ResponseType = responseType

shortcode := "Shortcode"
collect.Shortcode = shortcode

to := "To"
collect.To = to

dateSent := "DateSent"
collect.DateSent = dateSent

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize


var result string
result,_ = shortCode.ListShortcode(collect)

```


### <a name="list_inbound_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.ListInboundShortcode") ListInboundShortcode

> Retrive a list of inbound Sms Short Code messages associated with your message360 account.


```go
func (me *SHORTCODE_IMPL) ListInboundShortcode(input *ListInboundShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Only list SMS messages sent from this number |
| shortcode |  ``` Optional ```  | Only list SMS messages sent to Shortcode |
| dateReceived |  ``` Optional ```  | Only list SMS messages sent in the specified date MAKE REQUEST |


#### Example Usage

```go
collect := new (shortcode_pkg.ListInboundShortcodeInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

from := "From"
collect.From = from

shortcode := "Shortcode"
collect.Shortcode = shortcode

dateReceived := "DateReceived"
collect.DateReceived = dateReceived


var result string
result,_ = shortCode.ListInboundShortcode(collect)

```


[Back to List of Controllers](#list_of_controllers)



