#Getting started

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

![Importing SDK into Eclipse - Step 1](http://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](http://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](http://apidocs.io/illustration/go?step=import2&workspaceFolder=Message360-GoLang&projectName=message360_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](http://apidocs.io/illustration/go?step=import3&projectName=message360_lib)

## How to Use

The following section explains how to use the Message360 library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](http://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](http://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](http://apidocs.io/illustration/go?step=createNewProject2&projectName=message360_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](http://apidocs.io/illustration/go?step=createNewProject3&projectName=message360_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](http://apidocs.io/illustration/go?step=createNewProject4&projectName=message360_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](http://apidocs.io/illustration/go?step=createNewProject5&projectName=message360_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](http://apidocs.io/illustration/go?step=testProject0&projectName=message360_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](http://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](http://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Message360-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](http://apidocs.io/illustration/go?step=buildProject&projectName=message360_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](http://apidocs.io/illustration/go?step=runProject&projectName=message360_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


## Class Reference

### <a name="list_of_controllers"></a>List of Controllers

* [conference_pkg](#conference_pkg)
* [email_pkg](#email_pkg)
* [numberverification_pkg](#numberverification_pkg)
* [carrier_pkg](#carrier_pkg)
* [call_pkg](#call_pkg)
* [sms_pkg](#sms_pkg)
* [account_pkg](#account_pkg)
* [webrtc_pkg](#webrtc_pkg)
* [subaccount_pkg](#subaccount_pkg)
* [address_pkg](#address_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [recording_pkg](#recording_pkg)
* [transcription_pkg](#transcription_pkg)
* [usage_pkg](#usage_pkg)

### <a name="conference_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".conference_pkg") conference_pkg

#### Get instance

Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.

```go
conference := conference_pkg.NewCONFERENCE()
```

#### <a name="create_deaf_mute_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateDeafMuteParticipant") CreateDeafMuteParticipant

> Deaf Mute Participant


```go
func (me *CONFERENCE_IMPL) CreateDeafMuteParticipant(input *CreateDeafMuteParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | TODO: Add a parameter description |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (conference_pkg.CreateDeafMuteParticipantInput)

conferenceSid := "conferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.CreateDeafMuteParticipant(collect)

```


#### <a name="create_list_conference"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListConference") CreateListConference

> List Conference


```go
func (me *CONFERENCE_IMPL) CreateListConference(input *CreateListConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| dateUpdated |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.CreateListConferenceInput)

page,_ := strconv.ParseInt("114", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("114", 10, 8)
collect.PageSize = pageSize

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

status := models_pkg.InterruptedCallStatus_CANCELED
collect.Status = status

dateCreated := "DateCreated"
collect.DateCreated = dateCreated

dateUpdated := "DateUpdated"
collect.DateUpdated = dateUpdated

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.CreateListConference(collect)

```


#### <a name="create_view_conference"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewConference") CreateViewConference

> View Conference


```go
func (me *CONFERENCE_IMPL) CreateViewConference(input *CreateViewConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | The unique identifier of each conference resource |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.CreateViewConferenceInput)

conferencesid := "conferencesid"
collect.Conferencesid = conferencesid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.CreateViewConference(collect)

```


#### <a name="add_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.AddParticipant") AddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) AddParticipant(input *AddParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | Unique Conference Sid |
| participantnumber |  ``` Required ```  | Particiant Number |
| tocountrycode |  ``` Required ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.AddParticipantInput)

conferencesid := "conferencesid"
collect.Conferencesid = conferencesid

participantnumber := "participantnumber"
collect.Participantnumber = participantnumber

tocountrycode,_ := strconv.ParseInt("114", 10, 8)
collect.Tocountrycode = tocountrycode

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.AddParticipant(collect)

```


#### <a name="create_list_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListParticipant") CreateListParticipant

> List Participant


```go
func (me *CONFERENCE_IMPL) CreateListParticipant(input *CreateListParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| page |  ``` Optional ```  | page number |
| pagesize |  ``` Optional ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
collect := new (conference_pkg.CreateListParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

page,_ := strconv.ParseInt("114", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("114", 10, 8)
collect.Pagesize = pagesize

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.CreateListParticipant(collect)

```


#### <a name="create_view_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewParticipant") CreateViewParticipant

> View Participant


```go
func (me *CONFERENCE_IMPL) CreateViewParticipant(input *CreateViewParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (conference_pkg.CreateViewParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = conference.CreateViewParticipant(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="email_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".email_pkg") email_pkg

#### Get instance

Factory for the ``` EMAIL ``` interface can be accessed from the package email_pkg.

```go
email := email_pkg.NewEMAIL()
```

#### <a name="create_delete_invalid"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteInvalid") CreateDeleteInvalid

> This endpoint allows you to delete entries in the Invalid Emails list.


```go
func (me *EMAIL_IMPL) CreateDeleteInvalid(input *CreateDeleteInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |


#### Example Usage

```go
collect := new (email_pkg.CreateDeleteInvalidInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateDeleteInvalid(collect)

```


#### <a name="create_list_blocks"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListBlocks") CreateListBlocks

> Outputs email addresses on your blocklist


```go
func (me *EMAIL_IMPL) CreateListBlocks(input *CreateListBlocksInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | Where to start in the output list |
| limit |  ``` Optional ```  | Maximum number of records to return |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateListBlocksInput)

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateListBlocks(collect)

```


#### <a name="create_list_spam"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListSpam") CreateListSpam

> List out all email addresses marked as spam


```go
func (me *EMAIL_IMPL) CreateListSpam(input *CreateListSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The record number to start the list at |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.CreateListSpamInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.CreateListSpam(collect)

```


#### <a name="create_list_bounces"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListBounces") CreateListBounces

> List out all email addresses that have bounced


```go
func (me *EMAIL_IMPL) CreateListBounces(input *CreateListBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | The record to start the list at |
| limit |  ``` Optional ```  | The maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.CreateListBouncesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.CreateListBounces(collect)

```


#### <a name="create_delete_bounces"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteBounces") CreateDeleteBounces

> Delete an email address from the bounced address list


```go
func (me *EMAIL_IMPL) CreateDeleteBounces(input *CreateDeleteBouncesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to remove from the bounce list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateDeleteBouncesInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateDeleteBounces(collect)

```


#### <a name="create_list_invalid"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListInvalid") CreateListInvalid

> List out all invalid email addresses


```go
func (me *EMAIL_IMPL) CreateListInvalid(input *CreateListInvalidInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |
| offet |  ``` Optional ```  | Starting record for listing out emails |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
collect := new (email_pkg.CreateListInvalidInput)

responseType := "json"
collect.ResponseType = responseType

offet := "offet"
collect.Offet = offet

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.CreateListInvalid(collect)

```


#### <a name="create_list_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListUnsubscribes") CreateListUnsubscribes

> List all unsubscribed email addresses


```go
func (me *EMAIL_IMPL) CreateListUnsubscribes(input *CreateListUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |
| offset |  ``` Optional ```  | Starting record of the list |
| limit |  ``` Optional ```  | Maximum number of records to be returned |


#### Example Usage

```go
collect := new (email_pkg.CreateListUnsubscribesInput)

responseType := "json"
collect.ResponseType = responseType

offset := "offset"
collect.Offset = offset

limit := "limit"
collect.Limit = limit


var result string
result,_ = email.CreateListUnsubscribes(collect)

```


#### <a name="create_delete_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteUnsubscribes") CreateDeleteUnsubscribes

> Delete emails from the unsubscribe list


```go
func (me *EMAIL_IMPL) CreateDeleteUnsubscribes(input *CreateDeleteUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to remove from the unsubscribe list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateDeleteUnsubscribesInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateDeleteUnsubscribes(collect)

```


#### <a name="add_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.AddUnsubscribes") AddUnsubscribes

> Add an email to the unsubscribe list


```go
func (me *EMAIL_IMPL) AddUnsubscribes(input *AddUnsubscribesInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to add to the unsubscribe list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


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


#### <a name="create_delete_block"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteBlock") CreateDeleteBlock

> Deletes a blocked email


```go
func (me *EMAIL_IMPL) CreateDeleteBlock(input *CreateDeleteBlockInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address to remove from block list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateDeleteBlockInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateDeleteBlock(collect)

```


#### <a name="create_delete_spam"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteSpam") CreateDeleteSpam

> Deletes a email address marked as spam from the spam list


```go
func (me *EMAIL_IMPL) CreateDeleteSpam(input *CreateDeleteSpamInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateDeleteSpamInput)

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateDeleteSpam(collect)

```


#### <a name="create_send_email"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateSendEmail") CreateSendEmail

> Send out an email


```go
func (me *EMAIL_IMPL) CreateSendEmail(input *CreateSendEmailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | The to email address |
| from |  ``` Required ```  | The from email address |
| mtype |  ``` Required ```  ``` DefaultValue ```  | email format type, html or text |
| subject |  ``` Required ```  | Email subject |
| message |  ``` Required ```  | The body of the email message |
| cc |  ``` Optional ```  | CC Email address |
| bcc |  ``` Optional ```  | BCC Email address |
| attachment |  ``` Optional ```  | File to be attached.File must be less than 7MB. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (email_pkg.CreateSendEmailInput)

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

cc := "cc"
collect.Cc = cc

bcc := "bcc"
collect.Bcc = bcc

attachment := "attachment"
collect.Attachment = attachment

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = email.CreateSendEmail(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="numberverification_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".numberverification_pkg") numberverification_pkg

#### Get instance

Factory for the ``` NUMBERVERIFICATION ``` interface can be accessed from the package numberverification_pkg.

```go
numberVerification := numberverification_pkg.NewNUMBERVERIFICATION()
```

#### <a name="create_verify_number"></a>![Method: ](http://apidocs.io/img/method.png ".numberverification_pkg.CreateVerifyNumber") CreateVerifyNumber

> Number Verification


```go
func (me *NUMBERVERIFICATION_IMPL) CreateVerifyNumber(input *CreateVerifyNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | TODO: Add a parameter description |
| mtype |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (numberverification_pkg.CreateVerifyNumberInput)

phonenumber := "phonenumber"
collect.Phonenumber = phonenumber

mtype := "type"
collect.Mtype = mtype

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = numberVerification.CreateVerifyNumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="carrier_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg

#### Get instance

Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.

```go
carrier := carrier_pkg.NewCARRIER()
```

#### <a name="create_carrier_lookup"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookup") CreateCarrierLookup

> Get the Carrier Lookup


```go
func (me *CARRIER_IMPL) CreateCarrierLookup(input *CreateCarrierLookupInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | The number to lookup |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (carrier_pkg.CreateCarrierLookupInput)

phonenumber := "phonenumber"
collect.Phonenumber = phonenumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = carrier.CreateCarrierLookup(collect)

```


#### <a name="create_carrier_lookup_list"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupList") CreateCarrierLookupList

> Get the All Purchase Number's Carrier lookup


```go
func (me *CARRIER_IMPL) CreateCarrierLookupList(input *CreateCarrierLookupListInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page Number |
| pagesize |  ``` Optional ```  | Page Size |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (carrier_pkg.CreateCarrierLookupListInput)

page,_ := strconv.ParseInt("114", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("114", 10, 8)
collect.Pagesize = pagesize

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = carrier.CreateCarrierLookupList(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="call_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".call_pkg") call_pkg

#### Get instance

Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.

```go
call := call_pkg.NewCALL()
```

#### <a name="create_view_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateViewCall") CreateViewCall

> View Call Response


```go
func (me *CALL_IMPL) CreateViewCall(input *CreateViewCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | Call Sid id for particular Call |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateViewCallInput)

callsid := "callsid"
collect.Callsid = callsid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateViewCall(collect)

```


#### <a name="create_group_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateGroupCall") CreateGroupCall

> Group Call


```go
func (me *CALL_IMPL) CreateGroupCall(input *CreateGroupCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromCountryCode |  ``` Required ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| from |  ``` Required ```  | TODO: Add a parameter description |
| toCountryCode |  ``` Required ```  ``` DefaultValue ```  | TODO: Add a parameter description |
| to |  ``` Required ```  | TODO: Add a parameter description |
| url |  ``` Required ```  | TODO: Add a parameter description |
| method |  ``` Optional ```  | TODO: Add a parameter description |
| statusCallBackUrl |  ``` Optional ```  | TODO: Add a parameter description |
| statusCallBackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| fallBackUrl |  ``` Optional ```  | TODO: Add a parameter description |
| fallBackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| heartBeatUrl |  ``` Optional ```  | TODO: Add a parameter description |
| heartBeatMethod |  ``` Optional ```  | TODO: Add a parameter description |
| timeout |  ``` Optional ```  | TODO: Add a parameter description |
| playDtmf |  ``` Optional ```  | TODO: Add a parameter description |
| hideCallerId |  ``` Optional ```  | TODO: Add a parameter description |
| record |  ``` Optional ```  | TODO: Add a parameter description |
| recordCallBackUrl |  ``` Optional ```  | TODO: Add a parameter description |
| recordCallBackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| transcribe |  ``` Optional ```  | TODO: Add a parameter description |
| transcribeCallBackUrl |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | TODO: Add a parameter description |


#### Example Usage

```go
collect := new (call_pkg.CreateGroupCallInput)

fromCountryCode := "1"
collect.FromCountryCode = fromCountryCode

from := "From"
collect.From = from

toCountryCode := "1"
collect.ToCountryCode = toCountryCode

to := "To"
collect.To = to

url := "Url"
collect.Url = url

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

timeout,_ := strconv.ParseInt("114", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := "HideCallerId"
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

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateGroupCall(collect)

```


#### <a name="create_voice_effect"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateVoiceEffect") CreateVoiceEffect

> Voice Effect


```go
func (me *CALL_IMPL) CreateVoiceEffect(input *CreateVoiceEffectInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | TODO: Add a parameter description |
| audioDirection |  ``` Optional ```  | TODO: Add a parameter description |
| pitchSemiTones |  ``` Optional ```  | value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | value between -1 and 1 |
| pitch |  ``` Optional ```  | value greater than 0 |
| rate |  ``` Optional ```  | value greater than 0 |
| tempo |  ``` Optional ```  | value greater than 0 |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateVoiceEffectInput)

callSid := "CallSid"
collect.CallSid = callSid

audioDirection := models_pkg.AudioDirection_IN
collect.AudioDirection = audioDirection

pitchSemiTones := 114.525013104791
collect.PitchSemiTones = pitchSemiTones

pitchOctaves := 114.525013104791
collect.PitchOctaves = pitchOctaves

pitch := 114.525013104791
collect.Pitch = pitch

rate := 114.525013104791
collect.Rate = rate

tempo := 114.525013104791
collect.Tempo = tempo

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateVoiceEffect(collect)

```


#### <a name="create_record_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateRecordCall") CreateRecordCall

> Record a Call


```go
func (me *CALL_IMPL) CreateRecordCall(input *CreateRecordCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| record |  ``` Required ```  | Set true to initiate recording or false to terminate recording |
| direction |  ``` Optional ```  | The leg of the call to record |
| timeLimit |  ``` Optional ```  | Time in seconds the recording duration should not exceed |
| callBackUrl |  ``` Optional ```  | URL consulted after the recording completes |
| fileformat |  ``` Optional ```  | Format of the recording file. Can be .mp3 or .wav |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateRecordCallInput)

callSid := "CallSid"
collect.CallSid = callSid

record := false
collect.Record = record

direction := models_pkg.Direction_IN
collect.Direction = direction

timeLimit,_ := strconv.ParseInt("72", 10, 8)
collect.TimeLimit = timeLimit

callBackUrl := "CallBackUrl"
collect.CallBackUrl = callBackUrl

fileformat := models_pkg.AudioFormat_MP3
collect.Fileformat = fileformat

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateRecordCall(collect)

```


#### <a name="create_play_audio"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreatePlayAudio") CreatePlayAudio

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreatePlayAudio(input *CreatePlayAudioInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| audioUrl |  ``` Required ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |
| length |  ``` Optional ```  | Time limit in seconds for audio play back |
| direction |  ``` Optional ```  | The leg of the call audio will be played to |
| loop |  ``` Optional ```  | Repeat audio playback indefinitely |
| mix |  ``` Optional ```  | If false, all other audio will be muted |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreatePlayAudioInput)

callSid := "CallSid"
collect.CallSid = callSid

audioUrl := "AudioUrl"
collect.AudioUrl = audioUrl

length,_ := strconv.ParseInt("72", 10, 8)
collect.Length = length

direction := models_pkg.Direction_IN
collect.Direction = direction

loop := false
collect.Loop = loop

mix := false
collect.Mix = mix

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreatePlayAudio(collect)

```


#### <a name="create_list_calls"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateListCalls") CreateListCalls

> A list of calls associated with your Message360 account


```go
func (me *CALL_IMPL) CreateListCalls(input *CreateListCallsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Only list calls to this number |
| from |  ``` Optional ```  | Only list calls from this number |
| dateCreated |  ``` Optional ```  | Only list calls starting within the specified date range |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateListCallsInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("72", 10, 8)
collect.PageSize = pageSize

to := "To"
collect.To = to

from := "From"
collect.From = from

dateCreated := "DateCreated"
collect.DateCreated = dateCreated

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateListCalls(collect)

```


#### <a name="create_interrupted_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateInterruptedCall") CreateInterruptedCall

> Interrupt the Call by Call Sid


```go
func (me *CALL_IMPL) CreateInterruptedCall(input *CreateInterruptedCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | Call SId |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateInterruptedCallInput)

callSid := "CallSid"
collect.CallSid = callSid

url := "Url"
collect.Url = url

method := models_pkg.HttpAction_GET
collect.Method = method

status := models_pkg.InterruptedCallStatus_CANCELED
collect.Status = status

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateInterruptedCall(collect)

```


#### <a name="create_send_digit"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateSendDigit") CreateSendDigit

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreateSendDigit(input *CreateSendDigitInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateSendDigitInput)

callSid := "CallSid"
collect.CallSid = callSid

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

playDtmfDirection := models_pkg.Direction_IN
collect.PlayDtmfDirection = playDtmfDirection

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateSendDigit(collect)

```


#### <a name="create_make_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateMakeCall") CreateMakeCall

> You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) CreateMakeCall(input *CreateMakeCallInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromCountryCode |  ``` Required ```  | from country code |
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| toCountryCode |  ``` Required ```  | To cuntry code number |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (call_pkg.CreateMakeCallInput)

fromCountryCode := "FromCountryCode"
collect.FromCountryCode = fromCountryCode

from := "From"
collect.From = from

toCountryCode := "ToCountryCode"
collect.ToCountryCode = toCountryCode

to := "To"
collect.To = to

url := "Url"
collect.Url = url

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

heartBeatMethod := false
collect.HeartBeatMethod = heartBeatMethod

timeout,_ := strconv.ParseInt("72", 10, 8)
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

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = call.CreateMakeCall(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="sms_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".sms_pkg") sms_pkg

#### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

#### <a name="create_view_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateViewSMS") CreateViewSMS

> View Particular SMS


```go
func (me *SMS_IMPL) CreateViewSMS(input *CreateViewSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.CreateViewSMSInput)

messagesid := "messagesid"
collect.Messagesid = messagesid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.CreateViewSMS(collect)

```


#### <a name="create_list_inbound_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListInboundSMS") CreateListInboundSMS

> List All Inbound SMS


```go
func (me *SMS_IMPL) CreateListInboundSMS(input *CreateListInboundSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound SMS |
| to |  ``` Optional ```  | To Number to get Inbound SMS |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.CreateListInboundSMSInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pagesize := "pagesize"
collect.Pagesize = pagesize

from := "from"
collect.From = from

to := "to"
collect.To = to

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.CreateListInboundSMS(collect)

```


#### <a name="create_list_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListSMS") CreateListSMS

> List All SMS


```go
func (me *SMS_IMPL) CreateListSMS(input *CreateListSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.CreateListSMSInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("72", 10, 8)
collect.Pagesize = pagesize

from := "from"
collect.From = from

to := "to"
collect.To = to

datesent := "datesent"
collect.Datesent = datesent

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.CreateListSMS(collect)

```


#### <a name="create_send_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateSendSMS") CreateSendSMS

> Send an SMS from a message360 number


```go
func (me *SMS_IMPL) CreateSendSMS(input *CreateSendSMSInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| fromcountrycode |  ``` Required ```  ``` DefaultValue ```  | From Country Code |
| from |  ``` Required ```  | SMS enabled Message360 number to send this message from |
| tocountrycode |  ``` Required ```  ``` DefaultValue ```  | To country code |
| to |  ``` Required ```  | Number to send the SMS to |
| body |  ``` Required ```  | Text Message To Send |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (sms_pkg.CreateSendSMSInput)

fromcountrycode,_ := strconv.ParseInt("1", 10, 8)
collect.Fromcountrycode = fromcountrycode

from := "from"
collect.From = from

tocountrycode,_ := strconv.ParseInt("1", 10, 8)
collect.Tocountrycode = tocountrycode

to := "to"
collect.To = to

body := "body"
collect.Body = body

method := models_pkg.HttpAction_GET
collect.Method = method

messagestatuscallback := "messagestatuscallback"
collect.Messagestatuscallback = messagestatuscallback

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = sMS.CreateSendSMS(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="account_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".account_pkg") account_pkg

#### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

#### <a name="create_view_account"></a>![Method: ](http://apidocs.io/img/method.png ".account_pkg.CreateViewAccount") CreateViewAccount

> Display Account Description


```go
func (me *ACCOUNT_IMPL) CreateViewAccount(input *CreateViewAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (account_pkg.CreateViewAccountInput)

date := "date"
collect.Date = date

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = account.CreateViewAccount(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="webrtc_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".webrtc_pkg") webrtc_pkg

#### Get instance

Factory for the ``` WEBRTC ``` interface can be accessed from the package webrtc_pkg.

```go
webRTC := webrtc_pkg.NewWEBRTC()
```

#### <a name="create_token"></a>![Method: ](http://apidocs.io/img/method.png ".webrtc_pkg.CreateToken") CreateToken

> message360 webrtc


```go
func (me *WEBRTC_IMPL) CreateToken(input *CreateTokenInput)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your message360 Account SID |
| authToken |  ``` Required ```  | Your message360 Token |


#### Example Usage

```go
collect := new (webrtc_pkg.CreateTokenInput)

accountSid := "account_sid"
collect.AccountSid = accountSid

authToken := "auth_token"
collect.AuthToken = authToken


var result 
result,_ = webRTC.CreateToken(collect)

```


#### <a name="create_check_funds"></a>![Method: ](http://apidocs.io/img/method.png ".webrtc_pkg.CreateCheckFunds") CreateCheckFunds

> TODO: Add a method description


```go
func (me *WEBRTC_IMPL) CreateCheckFunds(input *CreateCheckFundsInput)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| accountSid |  ``` Required ```  | Your message360 Account SID |
| authToken |  ``` Required ```  | Your message360 Token |


#### Example Usage

```go
collect := new (webrtc_pkg.CreateCheckFundsInput)

accountSid := "account_sid"
collect.AccountSid = accountSid

authToken := "auth_token"
collect.AuthToken = authToken


var result 
result,_ = webRTC.CreateCheckFunds(collect)

```


#### <a name="create_authenticate_number"></a>![Method: ](http://apidocs.io/img/method.png ".webrtc_pkg.CreateAuthenticateNumber") CreateAuthenticateNumber

> Authenticate a message360 number for use


```go
func (me *WEBRTC_IMPL) CreateAuthenticateNumber(input *CreateAuthenticateNumberInput)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to authenticate for use |
| accountSid |  ``` Required ```  | Your message360 Account SID |
| authToken |  ``` Required ```  | Your message360 token |


#### Example Usage

```go
collect := new (webrtc_pkg.CreateAuthenticateNumberInput)

phoneNumber := "phone_number"
collect.PhoneNumber = phoneNumber

accountSid := "account_sid"
collect.AccountSid = accountSid

authToken := "auth_token"
collect.AuthToken = authToken


var result 
result,_ = webRTC.CreateAuthenticateNumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="subaccount_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".subaccount_pkg") subaccount_pkg

#### Get instance

Factory for the ``` SUBACCOUNT ``` interface can be accessed from the package subaccount_pkg.

```go
subAccount := subaccount_pkg.NewSUBACCOUNT()
```

#### <a name="create_sub_account"></a>![Method: ](http://apidocs.io/img/method.png ".subaccount_pkg.CreateSubAccount") CreateSubAccount

> Create Sub account


```go
func (me *SUBACCOUNT_IMPL) CreateSubAccount(input *CreateSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| firstname |  ``` Required ```  | TODO: Add a parameter description |
| lastname |  ``` Required ```  | TODO: Add a parameter description |
| email |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | ResponseType Format either json or xml |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateSubAccountInput)

firstname := "firstname"
collect.Firstname = firstname

lastname := "lastname"
collect.Lastname = lastname

email := "email"
collect.Email = email

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = subAccount.CreateSubAccount(collect)

```


#### <a name="create_suspend_sub_account"></a>![Method: ](http://apidocs.io/img/method.png ".subaccount_pkg.CreateSuspendSubAccount") CreateSuspendSubAccount

> Suspend or unsuspend


```go
func (me *SUBACCOUNT_IMPL) CreateSuspendSubAccount(input *CreateSuspendSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subaccountsid |  ``` Required ```  | TODO: Add a parameter description |
| activate |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateSuspendSubAccountInput)

subaccountsid := "subaccountsid"
collect.Subaccountsid = subaccountsid

activate := models_pkg.ActivateStatus_ACTIVATE
collect.Activate = activate

responseType := "ResponseType"
collect.ResponseType = responseType


var result string
result,_ = subAccount.CreateSuspendSubAccount(collect)

```


#### <a name="create_delete_merge_sub_account"></a>![Method: ](http://apidocs.io/img/method.png ".subaccount_pkg.CreateDeleteMergeSubAccount") CreateDeleteMergeSubAccount

> Delete or Merge Sub account


```go
func (me *SUBACCOUNT_IMPL) CreateDeleteMergeSubAccount(input *CreateDeleteMergeSubAccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subaccountsid |  ``` Required ```  | TODO: Add a parameter description |
| mergenumber |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format either json or xml |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateDeleteMergeSubAccountInput)

subaccountsid := "subaccountsid"
collect.Subaccountsid = subaccountsid

mergenumber := models_pkg.MergeNumberStatus_DELETE
collect.Mergenumber = mergenumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = subAccount.CreateDeleteMergeSubAccount(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="address_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".address_pkg") address_pkg

#### Get instance

Factory for the ``` ADDRESS ``` interface can be accessed from the package address_pkg.

```go
address := address_pkg.NewADDRESS()
```

#### <a name="create_address"></a>![Method: ](http://apidocs.io/img/method.png ".address_pkg.CreateAddress") CreateAddress

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
| description |  ``` Optional ```  | Description of addresses. |
| email |  ``` Optional ```  | Email Id of user. |
| phone |  ``` Optional ```  | Phone number of user. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response Type Either json or xml |


#### Example Usage

```go
collect := new (address_pkg.CreateAddressInput)

name := "name"
collect.Name = name

address := "address"
collect.Address = address

country := "country"
collect.Country = country

state := "state"
collect.State = state

city := "city"
collect.City = city

zip := "zip"
collect.Zip = zip

description := "description"
collect.Description = description

email := "email"
collect.Email = email

phone := "phone"
collect.Phone = phone

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.CreateAddress(collect)

```


#### <a name="create_delete_address"></a>![Method: ](http://apidocs.io/img/method.png ".address_pkg.CreateDeleteAddress") CreateDeleteAddress

> To delete Address to your address book


```go
func (me *ADDRESS_IMPL) CreateDeleteAddress(input *CreateDeleteAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be deleted. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.CreateDeleteAddressInput)

addressid := "addressid"
collect.Addressid = addressid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.CreateDeleteAddress(collect)

```


#### <a name="create_verify_address"></a>![Method: ](http://apidocs.io/img/method.png ".address_pkg.CreateVerifyAddress") CreateVerifyAddress

> Validates an address given.


```go
func (me *ADDRESS_IMPL) CreateVerifyAddress(input *CreateVerifyAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be verified. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type either JSON or xml |


#### Example Usage

```go
collect := new (address_pkg.CreateVerifyAddressInput)

addressid := "addressid"
collect.Addressid = addressid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.CreateVerifyAddress(collect)

```


#### <a name="create_list_address"></a>![Method: ](http://apidocs.io/img/method.png ".address_pkg.CreateListAddress") CreateListAddress

> List All Address 


```go
func (me *ADDRESS_IMPL) CreateListAddress(input *CreateListAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | Return requested # of items starting the value, default=0, must be an integer |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | How many results to return, default=10, max 100, must be an integer |
| addressId |  ``` Optional ```  | addresses Sid |
| dateCreated |  ``` Optional ```  | date created address. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.CreateListAddressInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

addressId := "addressId"
collect.AddressId = addressId

dateCreated := "dateCreated"
collect.DateCreated = dateCreated

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.CreateListAddress(collect)

```


#### <a name="create_view_address"></a>![Method: ](http://apidocs.io/img/method.png ".address_pkg.CreateViewAddress") CreateViewAddress

> View Address Specific address Book by providing the address id


```go
func (me *ADDRESS_IMPL) CreateViewAddress(input *CreateViewAddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressId |  ``` Required ```  | The identifier of the address to be retrieved. |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response Type either json or xml |


#### Example Usage

```go
collect := new (address_pkg.CreateViewAddressInput)

addressId := "addressId"
collect.AddressId = addressId

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = address.CreateViewAddress(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="phonenumber_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

#### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

#### <a name="update_phone_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.UpdatePhoneNumber") UpdatePhoneNumber

> Update Phone Number Details


```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber(input *UpdatePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |
| voiceUrl |  ``` Optional ```  | URL requested once the call connects |
| voiceMethod |  ``` Optional ```  | TODO: Add a parameter description |
| voiceFallbackUrl |  ``` Optional ```  | URL requested if the voice URL is not available |
| voiceFallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallback |  ``` Optional ```  | TODO: Add a parameter description |
| hangupCallbackMethod |  ``` Optional ```  | TODO: Add a parameter description |
| heartbeatUrl |  ``` Optional ```  | URL requested once the call connects |
| heartbeatMethod |  ``` Optional ```  | URL that can be requested every 60 seconds during the call to notify of elapsed time |
| smsUrl |  ``` Optional ```  | URL requested when an SMS is received |
| smsMethod |  ``` Optional ```  | TODO: Add a parameter description |
| smsFallbackUrl |  ``` Optional ```  | URL requested once the call connects |
| smsFallbackMethod |  ``` Optional ```  | URL requested if the sms URL is not available |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.UpdatePhoneNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

voiceUrl := "VoiceUrl"
collect.VoiceUrl = voiceUrl

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

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.UpdatePhoneNumber(collect)

```


#### <a name="create_buy_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateBuyNumber") CreateBuyNumber

> Buy Phone Number 


```go
func (me *PHONENUMBER_IMPL) CreateBuyNumber(input *CreateBuyNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be purchase |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateBuyNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateBuyNumber(collect)

```


#### <a name="create_release_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateReleaseNumber") CreateReleaseNumber

> Release number from account


```go
func (me *PHONENUMBER_IMPL) CreateReleaseNumber(input *CreateReleaseNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be relase |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateReleaseNumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateReleaseNumber(collect)

```


#### <a name="create_view_number_details"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateViewNumberDetails") CreateViewNumberDetails

> Get Phone Number Details


```go
func (me *PHONENUMBER_IMPL) CreateViewNumberDetails(input *CreateViewNumberDetailsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Get Phone number Detail |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateViewNumberDetailsInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateViewNumberDetails(collect)

```


#### <a name="create_list_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateListNumber") CreateListNumber

> List Account's Phone number details


```go
func (me *PHONENUMBER_IMPL) CreateListNumber(input *CreateListNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| numberType |  ``` Optional ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateListNumberInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("72", 10, 8)
collect.PageSize = pageSize

numberType := models_pkg.Number Type_ALL
collect.NumberType = numberType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateListNumber(collect)

```


#### <a name="create_available_phone_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateAvailablePhoneNumber") CreateAvailablePhoneNumber

> Available Phone Number


```go
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumber(input *CreateAvailablePhoneNumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | Number type either SMS,Voice or all |
| areaCode |  ``` Required ```  | Phone Number Area Code |
| pageSize |  ``` Optional ```  | Page Size |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateAvailablePhoneNumberInput)

numberType := models_pkg.Number Type_ALL
collect.NumberType = numberType

areaCode := "AreaCode"
collect.AreaCode = areaCode

pageSize,_ := strconv.ParseInt("72", 10, 8)
collect.PageSize = pageSize

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateAvailablePhoneNumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="recording_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".recording_pkg") recording_pkg

#### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

#### <a name="create_list_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateListRecording") CreateListRecording

> List out Recordings


```go
func (me *RECORDING_IMPL) CreateListRecording(input *CreateListRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| callSid |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.CreateListRecordingInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("72", 10, 8)
collect.PageSize = pageSize

dateCreated := "DateCreated"
collect.DateCreated = dateCreated

callSid := "CallSid"
collect.CallSid = callSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.CreateListRecording(collect)

```


#### <a name="create_delete_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateDeleteRecording") CreateDeleteRecording

> Delete Recording Record


```go
func (me *RECORDING_IMPL) CreateDeleteRecording(input *CreateDeleteRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording Sid to be delete |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.CreateDeleteRecordingInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.CreateDeleteRecording(collect)

```


#### <a name="create_view_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateViewRecording") CreateViewRecording

> View a specific Recording


```go
func (me *RECORDING_IMPL) CreateViewRecording(input *CreateViewRecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Search Recording sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (recording_pkg.CreateViewRecordingInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = recording.CreateViewRecording(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="transcription_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

#### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

#### <a name="create_audio_url_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateAudioURLTranscription") CreateAudioURLTranscription

> Audio URL Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateAudioURLTranscription(input *CreateAudioURLTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audioUrl |  ``` Required ```  | Audio url |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.CreateAudioURLTranscriptionInput)

audioUrl := "AudioUrl"
collect.AudioUrl = audioUrl

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.CreateAudioURLTranscription(collect)

```


#### <a name="create_recording_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateRecordingTranscription") CreateRecordingTranscription

> Recording Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateRecordingTranscription(input *CreateRecordingTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.CreateRecordingTranscriptionInput)

recordingSid := "RecordingSid"
collect.RecordingSid = recordingSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.CreateRecordingTranscription(collect)

```


#### <a name="create_view_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateViewTranscription") CreateViewTranscription

> View Specific Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateViewTranscription(input *CreateViewTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionSid |  ``` Required ```  | Unique Transcription ID |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.CreateViewTranscriptionInput)

transcriptionSid := "TranscriptionSid"
collect.TranscriptionSid = transcriptionSid

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.CreateViewTranscription(collect)

```


#### <a name="create_list_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateListTranscription") CreateListTranscription

> Get All transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateListTranscription(input *CreateListTranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateTranscribed |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (transcription_pkg.CreateListTranscriptionInput)

page,_ := strconv.ParseInt("72", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("72", 10, 8)
collect.PageSize = pageSize

status := models_pkg.Status_INPROGRESS
collect.Status = status

dateTranscribed := "DateTranscribed"
collect.DateTranscribed = dateTranscribed

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = transcription.CreateListTranscription(collect)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="usage_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".usage_pkg") usage_pkg

#### Get instance

Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.

```go
usage := usage_pkg.NewUSAGE()
```

#### <a name="create_list_usage"></a>![Method: ](http://apidocs.io/img/method.png ".usage_pkg.CreateListUsage") CreateListUsage

> Get all usage 


```go
func (me *USAGE_IMPL) CreateListUsage(input *CreateListUsageInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| productCode |  ``` Required ```  ``` DefaultValue ```  | Product Code |
| startDate |  ``` Required ```  ``` DefaultValue ```  | Start Usage Date |
| endDate |  ``` Required ```  ``` DefaultValue ```  | End Usage Date |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (usage_pkg.CreateListUsageInput)

productCode := models_pkg.Product Code_ALL
collect.ProductCode = productCode

startDate := "2016-09-06"
collect.StartDate = startDate

endDate := "2016-09-06"
collect.EndDate = endDate

responseType := "json"
collect.ResponseType = responseType


var result string
result,_ = usage.CreateListUsage(collect)

```


[Back to List of Controllers](#list_of_controllers)



