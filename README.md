# Getting started

Ytel API version 3

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

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Ytel-GoLang&projectName=ytel_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=ytel_lib)

## How to Use

The following section explains how to use the Ytel library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=ytel_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=ytel_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=ytel_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=ytel_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=ytel_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Ytel-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=ytel_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=ytel_lib)

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

* [webrtc_pkg](#webrtc_pkg)
* [sharedshortcode_pkg](#sharedshortcode_pkg)
* [conference_pkg](#conference_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [transcription_pkg](#transcription_pkg)
* [recording_pkg](#recording_pkg)
* [email_pkg](#email_pkg)
* [sms_pkg](#sms_pkg)
* [call_pkg](#call_pkg)
* [carrier_pkg](#carrier_pkg)
* [address_pkg](#address_pkg)
* [subaccount_pkg](#subaccount_pkg)
* [account_pkg](#account_pkg)
* [usage_pkg](#usage_pkg)
* [shortcode_pkg](#shortcode_pkg)
* [postcard_pkg](#postcard_pkg)
* [letter_pkg](#letter_pkg)
* [areamail_pkg](#areamail_pkg)

## <a name="webrtc_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".webrtc_pkg") webrtc_pkg

### Get instance

Factory for the ``` WEBRTC ``` interface can be accessed from the package webrtc_pkg.

```go
webRTC := webrtc_pkg.NewWEBRTC()
```

### <a name="create_token"></a>![Method: ](https://apidocs.io/img/method.png ".webrtc_pkg.CreateToken") CreateToken

> Ytel webrtc


```go
func (me *WEBRTC_IMPL) CreateToken(input *CreateTokenInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your Ytel Account SID |
| authToken |  ``` Required ```  | Your Ytel Token |
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


### <a name="check_funds"></a>![Method: ](https://apidocs.io/img/method.png ".webrtc_pkg.CheckFunds") CheckFunds

> TODO: Add a method description


```go
func (me *WEBRTC_IMPL) CheckFunds(input *CheckFundsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your Ytel Account SID |
| authToken |  ``` Required ```  | Your Ytel Token |


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


[Back to List of Controllers](#list_of_controllers)

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
| templateId |  ``` Required ```  | The unique identifier for a template object |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ViewTemplateInput)

templateId := uuid.NewV4()
collect.TemplateId = templateId

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
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| shortcode |  ``` Optional ```  | Only list messages sent from this Short Code |
| to |  ``` Optional ```  | Only list messages sent to this number |
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

shortcode := "Shortcode"
collect.Shortcode = shortcode

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
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound ShortCode |
| shortcode |  ``` Optional ```  | Only list messages sent to this Short Code |
| datecreated |  ``` Optional ```  | Only list messages sent with the specified date |


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

datecreated := "Datecreated"
collect.Datecreated = datecreated


var result string
result,_ = sharedShortCode.ListInboundSharedShortcodes(collect)

```


### <a name="send_shared_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.SendSharedShortcode") SendSharedShortcode

> Send an SMS from a Ytel ShortCode


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

> Retrieve a list of keywords associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) ListKeyword(input *ListKeywordInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| keyword |  ``` Optional ```  | Only list keywords of keyword |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListKeywordInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

keyword := "Keyword"
collect.Keyword = keyword

shortcode,_ := strconv.ParseInt("105", 10, 8)
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
| shortcode |  ``` Required ```  | List of valid Shortcode to your Ytel account |
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

> Retrieve a list of shortcode assignment associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) ListAssignment(input *ListAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.ListAssignmentInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

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
| shortcode |  ``` Required ```  | List of valid shortcode to your Ytel account |
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

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf


var result string
result,_ = conference.DeafMuteParticipant(collect)

```


### <a name="view_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ViewParticipant") ViewParticipant

> Retrieve information about a participant by its ParticipantSid.


```go
func (me *CONFERENCE_IMPL) ViewParticipant(input *ViewParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
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


### <a name="view_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.ViewConference") ViewConference

> Retrieve information about a conference by its ConferenceSid.


```go
func (me *CONFERENCE_IMPL) ViewConference(input *ViewConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier of each conference resource |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.ViewConferenceInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.ViewConference(collect)

```


### <a name="add_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.AddParticipant") AddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) AddParticipant(input *AddParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantNumber |  ``` Required ```  | The phone number of the participant to be added. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| muted |  ``` Optional ```  | Specifies if participant should be muted. |
| deaf |  ``` Optional ```  | Specifies if the participant should hear audio in the conference. |


#### Example Usage

```go
collect := new (conference_pkg.AddParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantNumber := "ParticipantNumber"
collect.ParticipantNumber = participantNumber

responseType := "json"
collect.ResponseType = responseType

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf


var result string
result,_ = conference.AddParticipant(collect)

```


### <a name="create_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConference") CreateConference

> Here you can experiment with initiating a conference call through Ytel and view the request response generated when doing so.


```go
func (me *CONFERENCE_IMPL) CreateConference(input *CreateConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid 10-digit number (E.164 format) that will be initiating the conference call. |
| to |  ``` Required ```  | A valid 10-digit number (E.164 format) that is to receive the conference call. |
| url |  ``` Required ```  | URL requested once the conference connects |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the conference is finished. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallbackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| record |  ``` Optional ```  | Specifies if the conference should be recorded. |
| recordCallBackUrl |  ``` Optional ```  | Recording parameters will be sent here upon completion. |
| recordCallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once conference connects. |
| scheduleTime |  ``` Optional ```  | Schedule conference in future. Schedule time must be greater than current time |
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

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_POST
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := models_pkg.HttpAction_GET
collect.StatusCallBackMethod = statusCallBackMethod

fallbackUrl := "FallbackUrl"
collect.FallbackUrl = fallbackUrl

fallbackMethod := models_pkg.HttpAction_GET
collect.FallbackMethod = fallbackMethod

record := false
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := models_pkg.HttpAction_GET
collect.RecordCallBackMethod = recordCallBackMethod

scheduleTime := "ScheduleTime"
collect.ScheduleTime = scheduleTime

timeout,_ := strconv.ParseInt("105", 10, 8)
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

> Retrieve a list of participants for an in-progress conference.


```go
func (me *CONFERENCE_IMPL) ListParticipant(input *ListParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response format, xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| muted |  ``` Optional ```  | Specifies if participant should be muted. |
| deaf |  ``` Optional ```  | Specifies if the participant should hear audio in the conference. |


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

> Retrieve a list of conference objects.


```go
func (me *CONFERENCE_IMPL) ListConference(input *ListConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| dateCreated |  ``` Optional ```  | Conference created date |


#### Example Usage

```go
collect := new (conference_pkg.ListConferenceInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

dateCreated := "DateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = conference.ListConference(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="phonenumber_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

### <a name="available_phone_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.AvailablePhoneNumber") AvailablePhoneNumber

> Retrieve a list of available phone numbers that can be purchased and used for your Ytel account.


```go
func (me *PHONENUMBER_IMPL) AvailablePhoneNumber(input *AvailablePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numbertype |  ``` Required ```  | Number type either SMS,Voice or all |
| areacode |  ``` Required ```  | Specifies the area code for the returned list of available numbers. Only available for North American numbers. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return. |


#### Example Usage

```go
collect := new (phonenumber_pkg.AvailablePhoneNumberInput)

numbertype := models_pkg.Number Type_ALL
collect.Numbertype = numbertype

areacode := "areacode"
collect.Areacode = areacode

responseType := "json"
collect.ResponseType = responseType

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize


var result string
result,_ = phoneNumber.AvailablePhoneNumber(collect)

```


### <a name="mass_release_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.MassReleaseNumber") MassReleaseNumber

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) MassReleaseNumber(input *MassReleaseNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel comma separated numbers (E.164 format). |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.MassReleaseNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.MassReleaseNumber(collect)

```


### <a name="view_number_details"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.ViewNumberDetails") ViewNumberDetails

> Retrieve the details for a phone number by its number.


```go
func (me *PHONENUMBER_IMPL) ViewNumberDetails(input *ViewNumberDetailsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel 10-digit phone number (E.164 format). |
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

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) ReleaseNumber(input *ReleaseNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
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

> Purchase a phone number to be used with your Ytel account


```go
func (me *PHONENUMBER_IMPL) BuyNumber(input *BuyNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
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

> Update properties for a Ytel number that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber(input *UpdatePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel number (E.164 format). |
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
| smsFallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML from an SMS or at initial request of the SmsUrl. |
| smsFallbackMethod |  ``` Optional ```  | The HTTP method Ytel will use when URL requested if the SmsUrl is not available. |


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


### <a name="transfer_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.TransferNumber") TransferNumber

> Transfer phone number that has been purchased for from one account to another account.


```go
func (me *PHONENUMBER_IMPL) TransferNumber(input *TransferNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
| fromaccountsid |  ``` Required ```  | A specific Accountsid from where Number is getting transfer. |
| toaccountsid |  ``` Required ```  | A specific Accountsid to which Number is getting transfer. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.TransferNumberInput)

phonenumber := "phonenumber"
collect.Phonenumber = phonenumber

fromaccountsid := "fromaccountsid"
collect.Fromaccountsid = fromaccountsid

toaccountsid := "toaccountsid"
collect.Toaccountsid = toaccountsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.TransferNumber(collect)

```


### <a name="list_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.ListNumber") ListNumber

> Retrieve a list of purchased phones numbers associated with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) ListNumber(input *ListNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| numberType |  ``` Optional ```  | The capability supported by the number.Number type either SMS,Voice or all |
| friendlyName |  ``` Optional ```  | A human-readable label added to the number object. |


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


### <a name="mass_update_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.MassUpdateNumber") MassUpdateNumber

> Update properties for a Ytel numbers that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) MassUpdateNumber(input *MassUpdateNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid comma(,) separated Ytel numbers. (E.164 format). |
| voiceUrl |  ``` Required ```  | The URL returning InboundXML incoming calls should execute when connected. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| friendlyName |  ``` Optional ```  | A human-readable value for labeling the number. |
| voiceMethod |  ``` Optional ```  | Specifies the HTTP method used to request the VoiceUrl once incoming call connects. |
| voiceFallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML on a call or at initial request of the voice url |
| voiceFallbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the VoiceFallbackUrl once incoming call connects. |
| hangupCallback |  ``` Optional ```  | URL that can be requested to receive notification when and how incoming call has ended. |
| hangupCallbackMethod |  ``` Optional ```  | The HTTP method Ytel will use when requesting the HangupCallback URL. |
| heartbeatUrl |  ``` Optional ```  | URL that can be used to monitor the phone number. |
| heartbeatMethod |  ``` Optional ```  | The HTTP method Ytel will use when requesting the HeartbeatUrl. |
| smsUrl |  ``` Optional ```  | URL requested when an SMS is received. |
| smsMethod |  ``` Optional ```  | The HTTP method Ytel will use when requesting the SmsUrl. |
| smsFallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML from an SMS or at initial request of the SmsUrl. |
| smsFallbackMethod |  ``` Optional ```  | The HTTP method Ytel will use when URL requested if the SmsUrl is not available. |


#### Example Usage

```go
collect := new (phonenumber_pkg.MassUpdateNumberInput)

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
result,_ = phoneNumber.MassUpdateNumber(collect)

```


### <a name="get_did_score_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.GetDIDScoreNumber") GetDIDScoreNumber

> Get DID Score Number


```go
func (me *PHONENUMBER_IMPL) GetDIDScoreNumber(input *GetDIDScoreNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | Specifies the multiple phone numbers for check updated spamscore . |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.GetDIDScoreNumberInput)

phonenumber := "Phonenumber"
collect.Phonenumber = phonenumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.GetDIDScoreNumber(collect)

```


### <a name="bulk_buy_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.BulkBuyNumber") BulkBuyNumber

> Purchase a selected number of DID's from specific area codes to be used with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) BulkBuyNumber(input *BulkBuyNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | The capability the number supports. |
| areaCode |  ``` Required ```  | Specifies the area code for the returned list of available numbers. Only available for North American numbers. |
| quantity |  ``` Required ```  | A positive integer that tells how many number you want to buy at a time. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| leftover |  ``` Optional ```  | If desired quantity is unavailable purchase what is available . |


#### Example Usage

```go
collect := new (phonenumber_pkg.BulkBuyNumberInput)

numberType := models_pkg.Number Type_ALL
collect.NumberType = numberType

areaCode := "AreaCode"
collect.AreaCode = areaCode

quantity := "Quantity"
collect.Quantity = quantity

responseType := "json"
collect.ResponseType = responseType

leftover := "Leftover"
collect.Leftover = leftover


var result string
result,_ = phoneNumber.BulkBuyNumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transcription_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

### <a name="view_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.ViewTranscription") ViewTranscription

> Retrieve information about a transaction by its TranscriptionSid.


```go
func (me *TRANSCRIPTION_IMPL) ViewTranscription(input *ViewTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionsid |  ``` Required ```  | The unique identifier for a transcription object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.ViewTranscriptionInput)

transcriptionsid := "transcriptionsid"
collect.Transcriptionsid = transcriptionsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.ViewTranscription(collect)

```


### <a name="recording_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.RecordingTranscription") RecordingTranscription

> Transcribe a message360 recording by its RecordingSid.


```go
func (me *TRANSCRIPTION_IMPL) RecordingTranscription(input *RecordingTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | The unique identifier for a recording object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.RecordingTranscriptionInput)

recordingSid := "recordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.RecordingTranscription(collect)

```


### <a name="audio_url_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.AudioURLTranscription") AudioURLTranscription

> Transcribe an audio recording from a file.


```go
func (me *TRANSCRIPTION_IMPL) AudioURLTranscription(input *AudioURLTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audiourl |  ``` Required ```  | URL pointing to the location of the audio file that is to be transcribed. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.AudioURLTranscriptionInput)

audiourl := "audiourl"
collect.Audiourl = audiourl

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.AudioURLTranscription(collect)

```


### <a name="list_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.ListTranscription") ListTranscription

> Retrieve a list of transcription objects for your Ytel account.


```go
func (me *TRANSCRIPTION_IMPL) ListTranscription(input *ListTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| status |  ``` Optional ```  | The state of the transcription. |
| dateTranscribed |  ``` Optional ```  | The date the transcription took place. |


#### Example Usage

```go
collect := new (transcription_pkg.ListTranscriptionInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

status := models_pkg.Status_INPROGRESS
collect.Status = status

dateTranscribed := "dateTranscribed"
collect.DateTranscribed = dateTranscribed


var result string
result,_ = transcription.ListTranscription(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="recording_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recording_pkg") recording_pkg

### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

### <a name="view_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.ViewRecording") ViewRecording

> Retrieve the recording of a call by its RecordingSid. This resource will return information regarding the call such as start time, end time, duration, and so forth.


```go
func (me *RECORDING_IMPL) ViewRecording(input *ViewRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for the recording. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.ViewRecordingInput)

recordingsid := "recordingsid"
collect.Recordingsid = recordingsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.ViewRecording(collect)

```


### <a name="delete_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.DeleteRecording") DeleteRecording

> Remove a recording from your Ytel account.


```go
func (me *RECORDING_IMPL) DeleteRecording(input *DeleteRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for a recording. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.DeleteRecordingInput)

recordingsid := "recordingsid"
collect.Recordingsid = recordingsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.DeleteRecording(collect)

```


### <a name="list_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.ListRecording") ListRecording

> Retrieve a list of recording objects.


```go
func (me *RECORDING_IMPL) ListRecording(input *ListRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| datecreated |  ``` Optional ```  | Filter results by creation date |
| callsid |  ``` Optional ```  | The unique identifier for a call. |


#### Example Usage

```go
collect := new (recording_pkg.ListRecordingInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

datecreated := "Datecreated"
collect.Datecreated = datecreated

callsid := "callsid"
collect.Callsid = callsid


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

> Remove an email from the spam email list.


```go
func (me *EMAIL_IMPL) DeleteSpam(input *DeleteSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| email |  ``` Required ```  | A valid email address that is to be remove from the spam list. |


#### Example Usage

```go
collect := new (email_pkg.DeleteSpamInput)

responseType := "json"
collect.ResponseType = responseType

email := "Email"
collect.Email = email


var result string
result,_ = email.DeleteSpam(collect)

```


### <a name="delete_block"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteBlock") DeleteBlock

> Remove an email from blocked emails list.


```go
func (me *EMAIL_IMPL) DeleteBlock(input *DeleteBlockInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to be remove from the blocked list. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.DeleteBlockInput)

email := "Email"
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
| email |  ``` Required ```  | A valid email address that is to be added to the unsubscribe list |
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

> Create and submit an email message to one or more email addresses.


```go
func (me *EMAIL_IMPL) SendEmail(input *SendEmailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| mtype |  ``` Required ```  ``` DefaultValue ```  | Specifies the type of email to be sent |
| subject |  ``` Required ```  | The subject of the mail. Must be a valid string. |
| message |  ``` Required ```  | The email message that is to be sent in the text. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| from |  ``` Optional ```  | A valid address that will send the email. |
| cc |  ``` Optional ```  | Carbon copy. A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| bcc |  ``` Optional ```  | Blind carbon copy. A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| attachment |  ``` Optional ```  | A file that is attached to the email. Must be less than 7 MB in size. |


#### Example Usage

```go
collect := new (email_pkg.SendEmailInput)

to := "To"
collect.To = to

mtype := models_pkg.SendEmailAs_HTML
collect.Mtype = mtype

subject := "Subject"
collect.Subject = subject

message := "Message"
collect.Message = message

responseType := "json"
collect.ResponseType = responseType

from := "From"
collect.From = from

cc := "Cc"
collect.Cc = cc

bcc := "Bcc"
collect.Bcc = bcc

attachment := "Attachment"
collect.Attachment = attachment


var result string
result,_ = email.SendEmail(collect)

```


### <a name="delete_unsubscribes"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteUnsubscribes") DeleteUnsubscribes

> Remove an email address from the list of unsubscribed emails.


```go
func (me *EMAIL_IMPL) DeleteUnsubscribes(input *DeleteUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the unsubscribe list. |
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

> Retrieve a list of email addresses from the unsubscribe list.


```go
func (me *EMAIL_IMPL) ListUnsubscribes(input *ListUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The starting point of the list of unsubscribed emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.ListUnsubscribesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.ListUnsubscribes(collect)

```


### <a name="list_invalid"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListInvalid") ListInvalid

> Retrieve a list of invalid email addresses.


```go
func (me *EMAIL_IMPL) ListInvalid(input *ListInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The starting point of the list of invalid emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.ListInvalidInput)

responseType := "json"
collect.ResponseType = responseType

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.ListInvalid(collect)

```


### <a name="delete_bounces"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteBounces") DeleteBounces

> Remove an email address from the bounced list.


```go
func (me *EMAIL_IMPL) DeleteBounces(input *DeleteBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| email |  ``` Required ```  | The email address to be remove from the bounced email list. |


#### Example Usage

```go
collect := new (email_pkg.DeleteBouncesInput)

responseType := "json"
collect.ResponseType = responseType

email := "Email"
collect.Email = email


var result string
result,_ = email.DeleteBounces(collect)

```


### <a name="list_bounces"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListBounces") ListBounces

> Retrieve a list of emails that have bounced.


```go
func (me *EMAIL_IMPL) ListBounces(input *ListBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The starting point of the list of bounced emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.ListBouncesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.ListBounces(collect)

```


### <a name="list_spam"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListSpam") ListSpam

> Retrieve a list of emails that are on the spam list.


```go
func (me *EMAIL_IMPL) ListSpam(input *ListSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The starting point of the list of spam emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.ListSpamInput)

responseType := "json"
collect.ResponseType = responseType

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.ListSpam(collect)

```


### <a name="list_blocks"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.ListBlocks") ListBlocks

> Retrieve a list of emails that have been blocked.


```go
func (me *EMAIL_IMPL) ListBlocks(input *ListBlocksInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The starting point of the list of blocked emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.ListBlocksInput)

responseType := "json"
collect.ResponseType = responseType

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.ListBlocks(collect)

```


### <a name="delete_invalid"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.DeleteInvalid") DeleteInvalid

> Remove an email from the invalid email list.


```go
func (me *EMAIL_IMPL) DeleteInvalid(input *DeleteInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the invalid email list. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (email_pkg.DeleteInvalidInput)

email := "Email"
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

> Send an SMS from a Ytel number


```go
func (me *SMS_IMPL) SendSMS(input *SendSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | The 10-digit SMS-enabled Ytel number (E.164 format) in which the message is sent. |
| to |  ``` Required ```  | The 10-digit phone number (E.164 format) that will receive the message. |
| body |  ``` Required ```  | The body message that is to be sent in the text. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |
| smartsms |  ``` Optional ```  ``` DefaultValue ```  | Check's 'to' number can receive sms or not using Carrier API, if wireless = true then text sms is sent, else wireless = false then call is recieved to end user with audible message. |
| deliveryStatus |  ``` Optional ```  ``` DefaultValue ```  | Delivery reports are a method to tell your system if the message has arrived on the destination phone. |


#### Example Usage

```go
collect := new (sms_pkg.SendSMSInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

body := "Body"
collect.Body = body

responseType := "json"
collect.ResponseType = responseType

method := models_pkg.HttpAction_GET
collect.Method = method

messageStatusCallback := "MessageStatusCallback"
collect.MessageStatusCallback = messageStatusCallback

smartsms := false
collect.Smartsms = smartsms

deliveryStatus := false
collect.DeliveryStatus = deliveryStatus


var result string
result,_ = sMS.SendSMS(collect)

```


### <a name="view_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ViewSMS") ViewSMS

> Retrieve a single SMS message object by its SmsSid.


```go
func (me *SMS_IMPL) ViewSMS(input *ViewSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.ViewSMSInput)

messageSid := "MessageSid"
collect.MessageSid = messageSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.ViewSMS(collect)

```


### <a name="list_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ListSMS") ListSMS

> Retrieve a list of Outbound SMS message objects.


```go
func (me *SMS_IMPL) ListSMS(input *ListSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |


#### Example Usage

```go
collect := new (sms_pkg.ListSMSInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

from := "From"
collect.From = from

to := "To"
collect.To = to

dateSent := "DateSent"
collect.DateSent = dateSent


var result string
result,_ = sMS.ListSMS(collect)

```


### <a name="list_inbound_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ListInboundSMS") ListInboundSMS

> Retrieve a list of Inbound SMS message objects.


```go
func (me *SMS_IMPL) ListInboundSMS(input *ListInboundSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Filter sms message objects by this date. |


#### Example Usage

```go
collect := new (sms_pkg.ListInboundSMSInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

from := "From"
collect.From = from

to := "To"
collect.To = to

dateSent := "DateSent"
collect.DateSent = dateSent


var result string
result,_ = sMS.ListInboundSMS(collect)

```


### <a name="view_detail_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.ViewDetailSMS") ViewDetailSMS

> Retrieve a single SMS message object with details by its SmsSid.


```go
func (me *SMS_IMPL) ViewDetailSMS(input *ViewDetailSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.ViewDetailSMSInput)

messageSid := "MessageSid"
collect.MessageSid = messageSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.ViewDetailSMS(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="call_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".call_pkg") call_pkg

### Get instance

Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.

```go
call := call_pkg.NewCALL()
```

### <a name="make_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.MakeCall") MakeCall

> You can experiment with initiating a call through Ytel and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) MakeCall(input *MakeCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call. |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed tim |
| heartBeatMethod |  ``` Optional ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Optional ```  | Time (in seconds) Ytel should wait while the call is ringing before canceling the call |
| playDtmf |  ``` Optional ```  | DTMF Digits to play to the call once it connects. 0-9, #, or * |
| hideCallerId |  ``` Optional ```  | Specifies if the caller id will be hidden |
| record |  ``` Optional ```  | Specifies if the call should be recorded |
| recordCallBackUrl |  ``` Optional ```  | Recording parameters will be sent here upon completion |
| recordCallBackMethod |  ``` Optional ```  | Method used to request the RecordCallback URL. |
| transcribe |  ``` Optional ```  | Specifies if the call recording should be transcribed |
| transcribeCallBackUrl |  ``` Optional ```  | Transcription parameters will be sent here upon completion |
| ifMachine |  ``` Optional ```  | How Ytel should handle the receiving numbers voicemail machine |
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

timeout,_ := strconv.ParseInt("196", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := true
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

ifMachine := models_pkg.ifMachine_CONTINUE
collect.IfMachine = ifMachine

ifMachineUrl := "IfMachineUrl"
collect.IfMachineUrl = ifMachineUrl

ifMachineMethod := models_pkg.HttpAction_GET
collect.IfMachineMethod = ifMachineMethod

feedback := true
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

length,_ := strconv.ParseInt("196", 10, 8)
collect.Length = length

direction := models_pkg.Direction_IN
collect.Direction = direction

mix := true
collect.Mix = mix


var result string
result,_ = call.PlayAudio(collect)

```


### <a name="record_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.RecordCall") RecordCall

> Start or stop recording of an in-progress voice call.


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

record := true
collect.Record = record

responseType := "json"
collect.ResponseType = responseType

direction := models_pkg.Direction_IN
collect.Direction = direction

timeLimit,_ := strconv.ParseInt("196", 10, 8)
collect.TimeLimit = timeLimit

callBackUrl := "CallBackUrl"
collect.CallBackUrl = callBackUrl

fileformat := models_pkg.AudioFormat_MP3
collect.Fileformat = fileformat


var result string
result,_ = call.RecordCall(collect)

```


### <a name="voice_effect"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.VoiceEffect") VoiceEffect

> Add audio voice effects to the an in-progress voice call.


```go
func (me *CALL_IMPL) VoiceEffect(input *VoiceEffectInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the in-progress voice call. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| audioDirection |  ``` Optional ```  | The direction the audio effect should be placed on. If IN, the effects will occur on the incoming audio stream. If OUT, the effects will occur on the outgoing audio stream. |
| pitchSemiTones |  ``` Optional ```  | Set the pitch in semitone (half-step) intervals. Value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | Set the pitch in octave intervals.. Value between -1 and 1 |
| pitch |  ``` Optional ```  | Set the pitch (lowness/highness) of the audio. The higher the value, the higher the pitch. Value greater than 0 |
| rate |  ``` Optional ```  | Set the rate for audio. The lower the value, the lower the rate. value greater than 0 |
| tempo |  ``` Optional ```  | Set the tempo (speed) of the audio. A higher value denotes a faster tempo. Value greater than 0 |


#### Example Usage

```go
collect := new (call_pkg.VoiceEffectInput)

callSid := "CallSid"
collect.CallSid = callSid

responseType := "json"
collect.ResponseType = responseType

audioDirection := models_pkg.AudioDirection_IN
collect.AudioDirection = audioDirection

pitchSemiTones := 196.738698457712
collect.PitchSemiTones = pitchSemiTones

pitchOctaves := 196.738698457712
collect.PitchOctaves = pitchOctaves

pitch := 196.738698457712
collect.Pitch = pitch

rate := 196.738698457712
collect.Rate = rate

tempo := 196.738698457712
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
| callSid |  ``` Required ```  | The unique identifier for voice call that is in progress. |
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


### <a name="list_calls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.ListCalls") ListCalls

> A list of calls associated with your Ytel account


```go
func (me *CALL_IMPL) ListCalls(input *ListCallsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Filter calls that were sent to this 10-digit number (E.164 format). |
| from |  ``` Optional ```  | Filter calls that were sent from this 10-digit number (E.164 format). |
| dateCreated |  ``` Optional ```  | Return calls that are from a specified date. |


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

> Initiate an outbound Ringless Voicemail through Ytel.


```go
func (me *CALL_IMPL) SendRinglessVM(input *SendRinglessVMInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call. |
| rVMCallerId |  ``` Required ```  | A required secondary Caller ID for RVM to work. |
| to |  ``` Required ```  | A valid number (E.164 format) that will receive the phone call. |
| voiceMailURL |  ``` Required ```  | The URL requested once the RVM connects. A set of default parameters will be sent here. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| method |  ``` Optional ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| statsCallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |


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

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) ViewCall(input *ViewCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | The unique identifier for the voice call. |
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


### <a name="view_call_detail"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.ViewCallDetail") ViewCallDetail

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) ViewCallDetail(
            callSid string,
            responseType string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the voice call. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
callSid := "callSid"
responseType := "json"

var result string
result,_ = call.ViewCallDetail(callSid, responseType)

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
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| statusCallBackMethod |  ``` Optional ```  | Specifies the HTTP methodlinkclass used to request StatusCallbackUrl. |
| fallBackUrl |  ``` Optional ```  | URL requested if the initial Url parameter fails or encounters an error |
| fallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| heartBeatUrl |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time and pass other general information. |
| heartBeatMethod |  ``` Optional ```  | Specifies the HTTP method used to request HeartbeatUrl. |
| timeout |  ``` Optional ```  | Time (in seconds) we should wait while the call is ringing before canceling the call |
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

timeout,_ := strconv.ParseInt("196", 10, 8)
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


[Back to List of Controllers](#list_of_controllers)

## <a name="carrier_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg

### Get instance

Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.

```go
carrier := carrier_pkg.NewCARRIER()
```

### <a name="carrier_lookup_list"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CarrierLookupList") CarrierLookupList

> Retrieve a list of carrier lookup objects.


```go
func (me *CARRIER_IMPL) CarrierLookupList(input *CarrierLookupListInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (carrier_pkg.CarrierLookupListInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize


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
| phoneNumber |  ``` Required ```  | A valid 10-digit number (E.164 format). |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (carrier_pkg.CarrierLookupInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

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
| addressid |  ``` Required ```  | The identifier of the address to be retrieved. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.ViewAddressInput)

addressid := "addressid"
collect.Addressid = addressid

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
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | How many results to return, default is 10, max is 100, must be an integer |
| addressid |  ``` Optional ```  | addresses Sid |
| dateCreated |  ``` Optional ```  | date created address. |


#### Example Usage

```go
collect := new (address_pkg.ListAddressInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

addressid := "addressid"
collect.Addressid = addressid

dateCreated := "dateCreated"
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
| addressid |  ``` Required ```  | The identifier of the address to be verified. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.VerifyAddressInput)

addressid := "addressid"
collect.Addressid = addressid

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
| addressid |  ``` Required ```  | The identifier of the address to be deleted. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.DeleteAddressInput)

addressid := "addressid"
collect.Addressid = addressid

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

> Retrieve information regarding your Ytel account by a specific date. The response object will contain data such as account status, balance, and account usage totals.


```go
func (me *ACCOUNT_IMPL) ViewAccount(input *ViewAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | Filter account information based on date. |
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

## <a name="usage_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".usage_pkg") usage_pkg

### Get instance

Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.

```go
usage := usage_pkg.NewUSAGE()
```

### <a name="list_usage"></a>![Method: ](https://apidocs.io/img/method.png ".usage_pkg.ListUsage") ListUsage

> Retrieve usage metrics regarding your Ytel account. The results includes inbound/outbound voice calls and inbound/outbound SMS messages as well as carrier lookup requests.


```go
func (me *USAGE_IMPL) ListUsage(input *ListUsageInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| productCode |  ``` Optional ```  ``` DefaultValue ```  | Filter usage results by product type. |
| startDate |  ``` Optional ```  ``` DefaultValue ```  | Filter usage objects by start date. |
| endDate |  ``` Optional ```  ``` DefaultValue ```  | Filter usage objects by end date. |
| includeSubAccounts |  ``` Optional ```  | Will include all subaccount usage data |


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

includeSubAccounts := "IncludeSubAccounts"
collect.IncludeSubAccounts = includeSubAccounts


var result string
result,_ = usage.ListUsage(collect)

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
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once the Short Code message is sent.GET or POST |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when Short Code message was sent. |


#### Example Usage

```go
collect := new (shortcode_pkg.SendDedicatedShortcodeInput)

shortcode,_ := strconv.ParseInt("155", 10, 8)
collect.Shortcode = shortcode

to := 155.015525231145
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

> Retrive a list of inbound Sms Short Code messages associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) ListInboundShortcode(input *ListInboundShortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Only list SMS messages sent from this number |
| shortcode |  ``` Optional ```  | Only list SMS messages sent to Shortcode |
| datecreated |  ``` Optional ```  | Only list SMS messages sent in the specified date MAKE REQUEST |


#### Example Usage

```go
collect := new (shortcode_pkg.ListInboundShortcodeInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

from := "From"
collect.From = from

shortcode := "Shortcode"
collect.Shortcode = shortcode

datecreated := "Datecreated"
collect.Datecreated = datecreated


var result string
result,_ = shortCode.ListInboundShortcode(collect)

```


### <a name="view_dedicated_shortcode_assignment"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.ViewDedicatedShortcodeAssignment") ViewDedicatedShortcodeAssignment

> Retrieve a single Short Code object by its shortcode assignment.


```go
func (me *SHORTCODE_IMPL) ViewDedicatedShortcodeAssignment(input *ViewDedicatedShortcodeAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Dedicated Short Code to your Ytel account |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (shortcode_pkg.ViewDedicatedShortcodeAssignmentInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = shortCode.ViewDedicatedShortcodeAssignment(collect)

```


### <a name="update_dedicated_short_code_assignment"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.UpdateDedicatedShortCodeAssignment") UpdateDedicatedShortCodeAssignment

> Update a dedicated shortcode.


```go
func (me *SHORTCODE_IMPL) UpdateDedicatedShortCodeAssignment(input *UpdateDedicatedShortCodeAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid dedicated shortcode to your Ytel account. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| friendlyName |  ``` Optional ```  | User generated name of the dedicated shortcode. |
| callbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |
| callbackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| fallbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| fallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST. |


#### Example Usage

```go
collect := new (shortcode_pkg.UpdateDedicatedShortCodeAssignmentInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

responseType := "json"
collect.ResponseType = responseType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

callbackMethod := "CallbackMethod"
collect.CallbackMethod = callbackMethod

callbackUrl := "CallbackUrl"
collect.CallbackUrl = callbackUrl

fallbackMethod := "FallbackMethod"
collect.FallbackMethod = fallbackMethod

fallbackUrl := "FallbackUrl"
collect.FallbackUrl = fallbackUrl


var result string
result,_ = shortCode.UpdateDedicatedShortCodeAssignment(collect)

```


### <a name="list_short_code_assignment"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.ListShortCodeAssignment") ListShortCodeAssignment

> Retrieve a list of Short Code assignment associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) ListShortCodeAssignment(input *ListShortCodeAssignmentInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response type format xml or json |
| shortcode |  ``` Optional ```  | Only list Short Code Assignment sent from this Short Code |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (shortcode_pkg.ListShortCodeAssignmentInput)

responseType := "json"
collect.ResponseType = responseType

shortcode := "Shortcode"
collect.Shortcode = shortcode

page := "page"
collect.Page = page

pagesize := "pagesize"
collect.Pagesize = pagesize


var result string
result,_ = shortCode.ListShortCodeAssignment(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="postcard_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".postcard_pkg") postcard_pkg

### Get instance

Factory for the ``` POSTCARD ``` interface can be accessed from the package postcard_pkg.

```go
postCard := postcard_pkg.NewPOSTCARD()
```

### <a name="view_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.ViewPostcard") ViewPostcard

> Retrieve a postcard object by its PostcardId.


```go
func (me *POSTCARD_IMPL) ViewPostcard(input *ViewPostcardInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier for a postcard object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (postcard_pkg.ViewPostcardInput)

postcardid := "postcardid"
collect.Postcardid = postcardid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = postCard.ViewPostcard(collect)

```


### <a name="create_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.CreatePostcard") CreatePostcard

> Create, print, and mail a postcard to an address. The postcard front must be supplied as a PDF, image, or an HTML string. The back can be a PDF, image, HTML string, or it can be automatically generated by supplying a custom message.


```go
func (me *POSTCARD_IMPL) CreatePostcard(input *CreatePostcardInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| from |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| attachbyid |  ``` Required ```  | Set an existing postcard by attaching its PostcardId. |
| front |  ``` Required ```  | A 4.25"x6.25" or 6.25"x11.25" image to use as the front of the postcard.  This can be a URL, local file, or an HTML string. Supported file types are PDF, PNG, and JPEG. |
| back |  ``` Required ```  | A 4.25"x6.25" or 6.25"x11.25" image to use as the back of the postcard, supplied as a URL, local file, or HTML string.  This allows you to customize your back design, but we will still insert the recipient address for you. |
| message |  ``` Required ```  | The message for the back of the postcard with a maximum of 350 characters. |
| setting |  ``` Required ```  | Code for the dimensions of the media to be printed. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| description |  ``` Optional ```  | A description for the postcard. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
collect := new (postcard_pkg.CreatePostcardInput)

to := "to"
collect.To = to

from := "from"
collect.From = from

attachbyid := "attachbyid"
collect.Attachbyid = attachbyid

front := "front"
collect.Front = front

back := "back"
collect.Back = back

message := "message"
collect.Message = message

setting := "setting"
collect.Setting = setting

responseType := "json"
collect.ResponseType = responseType

description := "description"
collect.Description = description

htmldata := "htmldata"
collect.Htmldata = htmldata


var result string
result,_ = postCard.CreatePostcard(collect)

```


### <a name="list_postcards"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.ListPostcards") ListPostcards

> Retrieve a list of postcard objects. The postcards objects are sorted by creation date, with the most recently created postcards appearing first.


```go
func (me *POSTCARD_IMPL) ListPostcards(input *ListPostcardsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| postcardid |  ``` Optional ```  | The unique identifier for a postcard object. |
| dateCreated |  ``` Optional ```  | The date the postcard was created. |


#### Example Usage

```go
collect := new (postcard_pkg.ListPostcardsInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

postcardid := "postcardid"
collect.Postcardid = postcardid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = postCard.ListPostcards(collect)

```


### <a name="delete_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.DeletePostcard") DeletePostcard

> Remove a postcard object.


```go
func (me *POSTCARD_IMPL) DeletePostcard(input *DeletePostcardInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier of a postcard object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (postcard_pkg.DeletePostcardInput)

postcardid := "postcardid"
collect.Postcardid = postcardid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = postCard.DeletePostcard(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="letter_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".letter_pkg") letter_pkg

### Get instance

Factory for the ``` LETTER ``` interface can be accessed from the package letter_pkg.

```go
letter := letter_pkg.NewLETTER()
```

### <a name="view_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.ViewLetter") ViewLetter

> Retrieve a letter object by its LetterSid.


```go
func (me *LETTER_IMPL) ViewLetter(input *ViewLetterInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (letter_pkg.ViewLetterInput)

lettersid := "lettersid"
collect.Lettersid = lettersid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = letter.ViewLetter(collect)

```


### <a name="create_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetter") CreateLetter

> Create, print, and mail a letter to an address. The letter file must be supplied as a PDF or an HTML string.


```go
func (me *LETTER_IMPL) CreateLetter(input *CreateLetterInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| from |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| attachbyid |  ``` Required ```  | Set an existing letter by attaching its LetterId. |
| file |  ``` Required ```  | File can be a 8.5"x11" PDF uploaded file or URL that links to a file. |
| color |  ``` Required ```  | Specify if letter should be printed in color. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| description |  ``` Optional ```  | A description for the letter. |
| extraservice |  ``` Optional ```  | Add an extra service to your letter. Options are "certified" or "registered". Certified provides tracking and delivery confirmation for domestic destinations and is an additional $5.00. Registered provides tracking and confirmation for international addresses and is an additional $16.50. |
| doublesided |  ``` Optional ```  | Specify if letter should be printed on both sides. |
| template |  ``` Optional ```  | Boolean that defaults to true. When set to false, this specifies that your letter does not follow the m360 address template. In this case, a blank address page will be inserted at the beginning of your file and you will be charged for the extra page. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
collect := new (letter_pkg.CreateLetterInput)

to := "to"
collect.To = to

from := "from"
collect.From = from

attachbyid := "attachbyid"
collect.Attachbyid = attachbyid

file := "file"
collect.File = file

color := "color"
collect.Color = color

responseType := "json"
collect.ResponseType = responseType

description := "description"
collect.Description = description

extraservice := "extraservice"
collect.Extraservice = extraservice

doublesided := "doublesided"
collect.Doublesided = doublesided

template := "template"
collect.Template = template

htmldata := "htmldata"
collect.Htmldata = htmldata


var result string
result,_ = letter.CreateLetter(collect)

```


### <a name="list_letters"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.ListLetters") ListLetters

> Retrieve a list of letter objects. The letter objects are sorted by creation date, with the most recently created letters appearing first.


```go
func (me *LETTER_IMPL) ListLetters(input *ListLettersInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| lettersid |  ``` Optional ```  | The unique identifier for a letter object. |
| dateCreated |  ``` Optional ```  | The date the letter was created. |


#### Example Usage

```go
collect := new (letter_pkg.ListLettersInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

lettersid := "lettersid"
collect.Lettersid = lettersid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = letter.ListLetters(collect)

```


### <a name="delete_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.DeleteLetter") DeleteLetter

> Remove a letter object by its LetterId.


```go
func (me *LETTER_IMPL) DeleteLetter(input *DeleteLetterInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (letter_pkg.DeleteLetterInput)

lettersid := "lettersid"
collect.Lettersid = lettersid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = letter.DeleteLetter(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="areamail_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".areamail_pkg") areamail_pkg

### Get instance

Factory for the ``` AREAMAIL ``` interface can be accessed from the package areamail_pkg.

```go
areaMail := areamail_pkg.NewAREAMAIL()
```

### <a name="create_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreaMail") CreateAreaMail

> Create a new AreaMail object.


```go
func (me *AREAMAIL_IMPL) CreateAreaMail(input *CreateAreaMailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| routes |  ``` Required ```  | List of routes that AreaMail should be delivered to.  A single route can be either a zipcode or a carrier route.List of routes that AreaMail should be delivered to.  A single route can be either a zipcode or a carrier route. A carrier route is in the form of 92610-C043 where the carrier route is composed of 5 numbers for zipcode, 1 letter for carrier route type, and 3 numbers for carrier route code. Delivery can be sent to mutliple routes by separating them with a commas. Valid Values: 92656, 92610-C043 |
| attachbyid |  ``` Required ```  | Set an existing areamail by attaching its AreamailId. |
| front |  ``` Required ```  | The front of the AreaMail item to be created. This can be a URL, local file, or an HTML string. Supported file types are PDF, PNG, and JPEG. Back required |
| back |  ``` Required ```  | The back of the AreaMail item to be created. This can be a URL, local file, or an HTML string. Supported file types are PDF, PNG, and JPEG. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| description |  ``` Optional ```  | A string value to use as a description for this AreaMail item. |
| targettype |  ``` Optional ```  | The delivery location type. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
collect := new (areamail_pkg.CreateAreaMailInput)

routes := "routes"
collect.Routes = routes

attachbyid := "attachbyid"
collect.Attachbyid = attachbyid

front := "front"
collect.Front = front

back := "back"
collect.Back = back

responseType := "json"
collect.ResponseType = responseType

description := "description"
collect.Description = description

targettype := "targettype"
collect.Targettype = targettype

htmldata := "htmldata"
collect.Htmldata = htmldata


var result string
result,_ = areaMail.CreateAreaMail(collect)

```


### <a name="view_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.ViewAreaMail") ViewAreaMail

> Retrieve an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) ViewAreaMail(input *ViewAreaMailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (areamail_pkg.ViewAreaMailInput)

areamailid := "areamailid"
collect.Areamailid = areamailid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = areaMail.ViewAreaMail(collect)

```


### <a name="list_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.ListAreaMail") ListAreaMail

> Retrieve a list of AreaMail objects.


```go
func (me *AREAMAIL_IMPL) ListAreaMail(input *ListAreaMailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| areamailsid |  ``` Optional ```  | The unique identifier for an AreaMail object. |
| dateCreated |  ``` Optional ```  | The date the AreaMail was created. |


#### Example Usage

```go
collect := new (areamail_pkg.ListAreaMailInput)

responseType := "json"
collect.ResponseType = responseType

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

areamailsid := "areamailsid"
collect.Areamailsid = areamailsid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = areaMail.ListAreaMail(collect)

```


### <a name="delete_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.DeleteAreaMail") DeleteAreaMail

> Remove an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) DeleteAreaMail(input *DeleteAreaMailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (areamail_pkg.DeleteAreaMailInput)

areamailid := "areamailid"
collect.Areamailid = areamailid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = areaMail.DeleteAreaMail(collect)

```


[Back to List of Controllers](#list_of_controllers)



