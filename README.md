# Getting Started

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

Now follow the below mentioned steps to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.
2. Select ```General -> Existing Projects into Workspace``` option from the tree list.
3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```
4. The Go library will be imported and its files will be visible in the Project Explorer

## How to Use

The following section explains how to use the Message360 library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

Name the Project as ```Test``` and click ```Finish```

Create a new directory in the ```src``` directory of this project

Name it ```test.com```

Now create a new file inside ```src/test.com```

Name it ```testsdk.go```

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

Append the library path to this GOPATH

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

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
* [transcription_pkg](#transcription_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [usage_pkg](#usage_pkg)
* [email_pkg](#email_pkg)
* [sms_pkg](#sms_pkg)
* [account_pkg](#account_pkg)
* [recording_pkg](#recording_pkg)
* [call_pkg](#call_pkg)
* [carrier_pkg](#carrier_pkg)

### <a name="conference_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".conference_pkg") conference_pkg

#### Get instance

Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.

```go
conference := conference_pkg.NewCONFERENCE()
```

#### <a name="create_view_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewParticipant") CreateViewParticipant

> View Participant


```go
func (me *CONFERENCE_IMPL) CreateViewParticipant(
            conferenceSid string,
            participantSid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | unique conference sid |
| participantSid |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
conferenceSid := "ConferenceSid"
participantSid := "ParticipantSid"
responseType := "json"

var result string
result,_ = conference.CreateViewParticipant(conferenceSid, participantSid, responseType)

```


#### <a name="create_list_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListParticipant") CreateListParticipant

> List Participant


```go
func (me *CONFERENCE_IMPL) CreateListParticipant(
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool,
            responseType *string)(string,error)
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
conferenceSid := "ConferenceSid"
page,_ := strconv.ParseInt("112", 10, 8)
pagesize,_ := strconv.ParseInt("112", 10, 8)
muted := false
deaf := false
responseType := "json"

var result string
result,_ = conference.CreateListParticipant(conferenceSid, page, pagesize, muted, deaf, responseType)

```


#### <a name="add_participant"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.AddParticipant") AddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) AddParticipant(
            conferencesid string,
            participantnumber string,
            tocountrycode int64,
            muted *bool,
            deaf *bool,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | Unique Conference Sid |
| participantnumber |  ``` Required ```  | Particiant Number |
| tocountrycode |  ``` Required ```  | TODO: Add a parameter description |
| muted |  ``` Optional ```  | TODO: Add a parameter description |
| deaf |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
conferencesid := "conferencesid"
participantnumber := "participantnumber"
tocountrycode,_ := strconv.ParseInt("112", 10, 8)
muted := false
deaf := false
responseType := "json"

var result string
result,_ = conference.AddParticipant(conferencesid, participantnumber, tocountrycode, muted, deaf, responseType)

```


#### <a name="create_view_conference"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateViewConference") CreateViewConference

> View Conference


```go
func (me *CONFERENCE_IMPL) CreateViewConference(
            conferencesid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferencesid |  ``` Required ```  | The unique identifier of each conference resource |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
conferencesid := "conferencesid"
responseType := "json"

var result string
result,_ = conference.CreateViewConference(conferencesid, responseType)

```


#### <a name="create_list_conference"></a>![Method: ](http://apidocs.io/img/method.png ".conference_pkg.CreateListConference") CreateListConference

> List Conference


```go
func (me *CONFERENCE_IMPL) CreateListConference(
            page *int64,
            pageSize *int64,
            friendlyName *string,
            status models_pkg.InterruptedCallStatusEnum,
            dateCreated *string,
            dateUpdated *string,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("112", 10, 8)
pageSize,_ := strconv.ParseInt("112", 10, 8)
friendlyName := "FriendlyName"
status := models_pkg.InterruptedCallStatus_CANCELED
dateCreated := "DateCreated"
dateUpdated := "DateUpdated"
responseType := "json"

var result string
result,_ = conference.CreateListConference(page, pageSize, friendlyName, status, dateCreated, dateUpdated, responseType)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="transcription_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

#### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

#### <a name="create_list_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateListTranscription") CreateListTranscription

> Get All transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateListTranscription(
            page *int64,
            pageSize *int64,
            status models_pkg.StatusEnum,
            dateTranscribed *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | TODO: Add a parameter description |
| pageSize |  ``` Optional ```  | TODO: Add a parameter description |
| status |  ``` Optional ```  | TODO: Add a parameter description |
| dateTranscribed |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("112", 10, 8)
pageSize,_ := strconv.ParseInt("112", 10, 8)
status := models_pkg.Status_INPROGRESS
dateTranscribed := "DateTranscribed"
responseType := "json"

var result string
result,_ = transcription.CreateListTranscription(page, pageSize, status, dateTranscribed, responseType)

```


#### <a name="create_recording_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateRecordingTranscription") CreateRecordingTranscription

> Recording Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateRecordingTranscription(
            recordingSid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
recordingSid := "RecordingSid"
responseType := "json"

var result string
result,_ = transcription.CreateRecordingTranscription(recordingSid, responseType)

```


#### <a name="create_view_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateViewTranscription") CreateViewTranscription

> View Specific Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateViewTranscription(
            transcriptionSid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionSid |  ``` Required ```  | Unique Transcription ID |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
transcriptionSid := "TranscriptionSid"
responseType := "json"

var result string
result,_ = transcription.CreateViewTranscription(transcriptionSid, responseType)

```


#### <a name="create_audio_url_transcription"></a>![Method: ](http://apidocs.io/img/method.png ".transcription_pkg.CreateAudioURLTranscription") CreateAudioURLTranscription

> Audio URL Transcriptions


```go
func (me *TRANSCRIPTION_IMPL) CreateAudioURLTranscription(
            audioUrl string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audioUrl |  ``` Required ```  | Audio url |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
audioUrl := "AudioUrl"
responseType := "json"

var result string
result,_ = transcription.CreateAudioURLTranscription(audioUrl, responseType)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="phonenumber_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

#### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

#### <a name="create_available_phone_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateAvailablePhoneNumber") CreateAvailablePhoneNumber

> Available Phone Number


```go
func (me *PHONENUMBER_IMPL) CreateAvailablePhoneNumber(
            numberType string,
            areaCode string,
            pageSize *int64,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | Number type either SMS,Voice or all |
| areaCode |  ``` Required ```  | Phone Number Area Code |
| pageSize |  ``` Optional ```  | Page Size |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
numberType := "NumberType"
areaCode := "AreaCode"
pageSize,_ := strconv.ParseInt("112", 10, 8)
responseType := "json"

var result string
result,_ = phoneNumber.CreateAvailablePhoneNumber(numberType, areaCode, pageSize, responseType)

```


#### <a name="create_list_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateListNumber") CreateListNumber

> List Account's Phone number details


```go
func (me *PHONENUMBER_IMPL) CreateListNumber(
            page *int64,
            pageSize *int64,
            numberType *string,
            friendlyName *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| numberType |  ``` Optional ```  | TODO: Add a parameter description |
| friendlyName |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("112", 10, 8)
pageSize,_ := strconv.ParseInt("112", 10, 8)
numberType := "NumberType"
friendlyName := "FriendlyName"
responseType := "json"

var result string
result,_ = phoneNumber.CreateListNumber(page, pageSize, numberType, friendlyName, responseType)

```


#### <a name="create_release_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateReleaseNumber") CreateReleaseNumber

> Release number from account


```go
func (me *PHONENUMBER_IMPL) CreateReleaseNumber(
            phoneNumber string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be relase |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
phoneNumber := "PhoneNumber"
responseType := "json"

var result string
result,_ = phoneNumber.CreateReleaseNumber(phoneNumber, responseType)

```


#### <a name="create_buy_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateBuyNumber") CreateBuyNumber

> Buy Phone Number 


```go
func (me *PHONENUMBER_IMPL) CreateBuyNumber(
            phoneNumber string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Phone number to be purchase |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
phoneNumber := "PhoneNumber"
responseType := "json"

var result string
result,_ = phoneNumber.CreateBuyNumber(phoneNumber, responseType)

```


#### <a name="create_view_number_details"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.CreateViewNumberDetails") CreateViewNumberDetails

> Get Phone Number Details


```go
func (me *PHONENUMBER_IMPL) CreateViewNumberDetails(
            phoneNumber string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | Get Phone number Detail |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
phoneNumber := "PhoneNumber"
responseType := "json"

var result string
result,_ = phoneNumber.CreateViewNumberDetails(phoneNumber, responseType)

```


#### <a name="update_phone_number"></a>![Method: ](http://apidocs.io/img/method.png ".phonenumber_pkg.UpdatePhoneNumber") UpdatePhoneNumber

> Update Phone Number Details


```go
func (me *PHONENUMBER_IMPL) UpdatePhoneNumber(
            phoneNumber string,
            friendlyName *string,
            voiceUrl *string,
            voiceMethod models_pkg.HttpMethodEnum,
            voiceFallbackUrl *string,
            voiceFallbackMethod models_pkg.HttpMethodEnum,
            hangupCallback *string,
            hangupCallbackMethod models_pkg.HttpMethodEnum,
            heartbeatUrl *string,
            heartbeatMethod models_pkg.HttpMethodEnum,
            smsUrl *string,
            smsMethod models_pkg.HttpMethodEnum,
            smsFallbackUrl *string,
            smsFallbackMethod models_pkg.HttpMethodEnum,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
phoneNumber := "PhoneNumber"
friendlyName := "FriendlyName"
voiceUrl := "VoiceUrl"
voiceMethod := models_pkg.HttpMethod_GET
voiceFallbackUrl := "VoiceFallbackUrl"
voiceFallbackMethod := models_pkg.HttpMethod_GET
hangupCallback := "HangupCallback"
hangupCallbackMethod := models_pkg.HttpMethod_GET
heartbeatUrl := "HeartbeatUrl"
heartbeatMethod := models_pkg.HttpMethod_GET
smsUrl := "SmsUrl"
smsMethod := models_pkg.HttpMethod_GET
smsFallbackUrl := "SmsFallbackUrl"
smsFallbackMethod := models_pkg.HttpMethod_GET
responseType := "json"

var result string
result,_ = phoneNumber.UpdatePhoneNumber(phoneNumber, friendlyName, voiceUrl, voiceMethod, voiceFallbackUrl, voiceFallbackMethod, hangupCallback, hangupCallbackMethod, heartbeatUrl, heartbeatMethod, smsUrl, smsMethod, smsFallbackUrl, smsFallbackMethod, responseType)

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
func (me *USAGE_IMPL) CreateListUsage(
            productCode string,
            startDate string,
            endDate string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| productCode |  ``` Required ```  ``` DefaultValue ```  | Product Code |
| startDate |  ``` Required ```  ``` DefaultValue ```  | Start Usage Date |
| endDate |  ``` Required ```  ``` DefaultValue ```  | End Usage Date |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
productCode := "0"
startDate := "2016-09-06"
endDate := "2016-09-06"
responseType := "json"

var result string
result,_ = usage.CreateListUsage(productCode, startDate, endDate, responseType)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="email_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".email_pkg") email_pkg

#### Get instance

Factory for the ``` EMAIL ``` interface can be accessed from the package email_pkg.

```go
email := email_pkg.NewEMAIL()
```

#### <a name="create_send_email"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateSendEmail") CreateSendEmail

> Send out an email


```go
func (me *EMAIL_IMPL) CreateSendEmail(
            to string,
            from string,
            mtype string,
            subject string,
            message string,
            cc *string,
            bcc *string,
            attachment []byte,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
to := "to"
from := "from"
mtype := "html"
subject := "subject"
message := "message"
cc := "cc"
bcc := "bcc"
attachment :=  []byte("")
responseType := "json"

var result string
result,_ = email.CreateSendEmail(to, from, mtype, subject, message, cc, bcc, attachment, responseType)

```


#### <a name="create_delete_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteUnsubscribes") CreateDeleteUnsubscribes

> Delete emails from the unsubscribe list


```go
func (me *EMAIL_IMPL) CreateDeleteUnsubscribes(
            email string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to remove from the unsubscribe list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
email := "email"
responseType := "json"

var result string
result,_ = email.CreateDeleteUnsubscribes(email, responseType)

```


#### <a name="create_list_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListUnsubscribes") CreateListUnsubscribes

> List all unsubscribed email addresses


```go
func (me *EMAIL_IMPL) CreateListUnsubscribes(
            responseType *string,
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |
| offset |  ``` Optional ```  | Starting record of the list |
| limit |  ``` Optional ```  | Maximum number of records to be returned |


#### Example Usage

```go
responseType := "json"
offset := "offset"
limit := "limit"

var result string
result,_ = email.CreateListUnsubscribes(responseType, offset, limit)

```


#### <a name="add_unsubscribes"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.AddUnsubscribes") AddUnsubscribes

> Add an email to the unsubscribe list


```go
func (me *EMAIL_IMPL) AddUnsubscribes(
            email string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email to add to the unsubscribe list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
email := "email"
responseType := "json"

var result string
result,_ = email.AddUnsubscribes(email, responseType)

```


#### <a name="create_delete_spam"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteSpam") CreateDeleteSpam

> Deletes a email address marked as spam from the spam list


```go
func (me *EMAIL_IMPL) CreateDeleteSpam(
            email string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
email := "email"
responseType := "json"

var result string
result,_ = email.CreateDeleteSpam(email, responseType)

```


#### <a name="create_delete_block"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteBlock") CreateDeleteBlock

> Deletes a blocked email


```go
func (me *EMAIL_IMPL) CreateDeleteBlock(
            email string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | Email address to remove from block list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
email := "email"
responseType := "json"

var result string
result,_ = email.CreateDeleteBlock(email, responseType)

```


#### <a name="create_list_invalid"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListInvalid") CreateListInvalid

> List out all invalid email addresses


```go
func (me *EMAIL_IMPL) CreateListInvalid(
            responseType *string,
            offet *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |
| offet |  ``` Optional ```  | Starting record for listing out emails |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
responseType := "json"
offet := "offet"
limit := "limit"

var result string
result,_ = email.CreateListInvalid(responseType, offet, limit)

```


#### <a name="create_delete_bounces"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateDeleteBounces") CreateDeleteBounces

> Delete an email address from the bounced address list


```go
func (me *EMAIL_IMPL) CreateDeleteBounces(
            email string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to remove from the bounce list |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
email := "email"
responseType := "json"

var result string
result,_ = email.CreateDeleteBounces(email, responseType)

```


#### <a name="create_list_bounces"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListBounces") CreateListBounces

> List out all email addresses that have bounced


```go
func (me *EMAIL_IMPL) CreateListBounces(
            responseType *string,
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |
| offset |  ``` Optional ```  | The record to start the list at |
| limit |  ``` Optional ```  | The maximum number of records to return |


#### Example Usage

```go
responseType := "json"
offset := "offset"
limit := "limit"

var result string
result,_ = email.CreateListBounces(responseType, offset, limit)

```


#### <a name="create_list_spam"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListSpam") CreateListSpam

> List out all email addresses marked as spam


```go
func (me *EMAIL_IMPL) CreateListSpam(
            responseType string,
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  ``` DefaultValue ```  | Response format, xml or json |
| offset |  ``` Optional ```  | The record number to start the list at |
| limit |  ``` Optional ```  | Maximum number of records to return |


#### Example Usage

```go
responseType := "json"
offset := "offset"
limit := "limit"

var result string
result,_ = email.CreateListSpam(responseType, offset, limit)

```


#### <a name="create_list_blocks"></a>![Method: ](http://apidocs.io/img/method.png ".email_pkg.CreateListBlocks") CreateListBlocks

> Outputs email addresses on your blocklist


```go
func (me *EMAIL_IMPL) CreateListBlocks(
            offset *string,
            limit *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | Where to start in the output list |
| limit |  ``` Optional ```  | Maximum number of records to return |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
offset := "offset"
limit := "limit"
responseType := "json"

var result string
result,_ = email.CreateListBlocks(offset, limit, responseType)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="sms_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".sms_pkg") sms_pkg

#### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

#### <a name="create_send_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateSendSMS") CreateSendSMS

> Send an SMS from a message360 number


```go
func (me *SMS_IMPL) CreateSendSMS(
            fromcountrycode int64,
            from string,
            tocountrycode int64,
            to string,
            body string,
            method models_pkg.HttpMethodEnum,
            messagestatuscallback *string,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
fromcountrycode,_ := strconv.ParseInt("1", 10, 8)
from := "from"
tocountrycode,_ := strconv.ParseInt("1", 10, 8)
to := "to"
body := "body"
method := models_pkg.HttpMethod_GET
messagestatuscallback := "messagestatuscallback"
responseType := "json"

var result string
result,_ = sMS.CreateSendSMS(fromcountrycode, from, tocountrycode, to, body, method, messagestatuscallback, responseType)

```


#### <a name="create_view_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateViewSMS") CreateViewSMS

> View Particular SMS


```go
func (me *SMS_IMPL) CreateViewSMS(
            messagesid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messagesid |  ``` Required ```  | Message sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
messagesid := "messagesid"
responseType := "json"

var result string
result,_ = sMS.CreateViewSMS(messagesid, responseType)

```


#### <a name="create_list_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListSMS") CreateListSMS

> List All SMS


```go
func (me *SMS_IMPL) CreateListSMS(
            page *int64,
            pagesize *int64,
            from *string,
            to *string,
            datesent *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Messages sent from this number |
| to |  ``` Optional ```  | Messages sent to this number |
| datesent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("20", 10, 8)
pagesize,_ := strconv.ParseInt("20", 10, 8)
from := "from"
to := "to"
datesent := "datesent"
responseType := "json"

var result string
result,_ = sMS.CreateListSMS(page, pagesize, from, to, datesent, responseType)

```


#### <a name="create_list_inbound_sms"></a>![Method: ](http://apidocs.io/img/method.png ".sms_pkg.CreateListInboundSMS") CreateListInboundSMS

> List All Inbound SMS


```go
func (me *SMS_IMPL) CreateListInboundSMS(
            page *int64,
            pagesize *string,
            from *string,
            to *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound SMS |
| to |  ``` Optional ```  | To Number to get Inbound SMS |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("20", 10, 8)
pagesize := "pagesize"
from := "from"
to := "to"
responseType := "json"

var result string
result,_ = sMS.CreateListInboundSMS(page, pagesize, from, to, responseType)

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
func (me *ACCOUNT_IMPL) CreateViewAccount(
            date string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response type format xml or json |


#### Example Usage

```go
date := "date"
responseType := "json"

var result string
result,_ = account.CreateViewAccount(date, responseType)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="recording_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".recording_pkg") recording_pkg

#### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

#### <a name="create_delete_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateDeleteRecording") CreateDeleteRecording

> Delete Recording Record


```go
func (me *RECORDING_IMPL) CreateDeleteRecording(
            recordingSid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Unique Recording Sid to be delete |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
recordingSid := "RecordingSid"
responseType := "json"

var result string
result,_ = recording.CreateDeleteRecording(recordingSid, responseType)

```


#### <a name="create_view_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateViewRecording") CreateViewRecording

> View a specific Recording


```go
func (me *RECORDING_IMPL) CreateViewRecording(
            recordingSid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | Search Recording sid |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
recordingSid := "RecordingSid"
responseType := "json"

var result string
result,_ = recording.CreateViewRecording(recordingSid, responseType)

```


#### <a name="create_list_recording"></a>![Method: ](http://apidocs.io/img/method.png ".recording_pkg.CreateListRecording") CreateListRecording

> List out Recordings


```go
func (me *RECORDING_IMPL) CreateListRecording(
            page *int64,
            pageSize *int64,
            dateCreated *string,
            callSid *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| dateCreated |  ``` Optional ```  | TODO: Add a parameter description |
| callSid |  ``` Optional ```  | TODO: Add a parameter description |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page,_ := strconv.ParseInt("20", 10, 8)
pageSize,_ := strconv.ParseInt("20", 10, 8)
dateCreated := "DateCreated"
callSid := "CallSid"
responseType := "json"

var result string
result,_ = recording.CreateListRecording(page, pageSize, dateCreated, callSid, responseType)

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
func (me *CALL_IMPL) CreateViewCall(
            callsid string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | Call Sid id for particular Call |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
callsid := "callsid"
responseType := "json"

var result string
result,_ = call.CreateViewCall(callsid, responseType)

```


#### <a name="create_make_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateMakeCall") CreateMakeCall

> You can experiment with initiating a call through Message360 and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) CreateMakeCall(
            fromCountryCode string,
            from string,
            toCountryCode string,
            to string,
            url string,
            method models_pkg.HttpMethodEnum,
            statusCallBackUrl *string,
            statusCallBackMethod models_pkg.HttpMethodEnum,
            fallBackUrl *string,
            fallBackMethod models_pkg.HttpMethodEnum,
            heartBeatUrl *string,
            heartBeatMethod *bool,
            timeout *int64,
            playDtmf *string,
            hideCallerId *bool,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod models_pkg.HttpMethodEnum,
            transcribe *bool,
            transcribeCallBackUrl *string,
            ifMachine models_pkg.IfMachineEnum,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
fromCountryCode := "FromCountryCode"
from := "From"
toCountryCode := "ToCountryCode"
to := "To"
url := "Url"
method := models_pkg.HttpMethod_GET
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := models_pkg.HttpMethod_GET
fallBackUrl := "FallBackUrl"
fallBackMethod := models_pkg.HttpMethod_GET
heartBeatUrl := "HeartBeatUrl"
heartBeatMethod := false
timeout,_ := strconv.ParseInt("20", 10, 8)
playDtmf := "PlayDtmf"
hideCallerId := false
record := false
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := models_pkg.HttpMethod_GET
transcribe := false
transcribeCallBackUrl := "TranscribeCallBackUrl"
ifMachine := models_pkg.ifMachine_CONTINUE
responseType := "json"

var result string
result,_ = call.CreateMakeCall(fromCountryCode, from, toCountryCode, to, url, method, statusCallBackUrl, statusCallBackMethod, fallBackUrl, fallBackMethod, heartBeatUrl, heartBeatMethod, timeout, playDtmf, hideCallerId, record, recordCallBackUrl, recordCallBackMethod, transcribe, transcribeCallBackUrl, ifMachine, responseType)

```


#### <a name="create_play_audio"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreatePlayAudio") CreatePlayAudio

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreatePlayAudio(
            length int64,
            direction models_pkg.DirectionEnum,
            loop bool,
            mix bool,
            callSid *string,
            audioUrl *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| length |  ``` Required ```  | Time limit in seconds for audio play back |
| direction |  ``` Required ```  | The leg of the call audio will be played to |
| loop |  ``` Required ```  | Repeat audio playback indefinitely |
| mix |  ``` Required ```  | If false, all other audio will be muted |
| callSid |  ``` Optional ```  | The unique identifier of each call resource |
| audioUrl |  ``` Optional ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
length,_ := strconv.ParseInt("20", 10, 8)
direction := models_pkg.Direction_IN
loop := false
mix := false
callSid := "CallSid"
audioUrl := "AudioUrl"
responseType := "json"

var result string
result,_ = call.CreatePlayAudio(length, direction, loop, mix, callSid, audioUrl, responseType)

```


#### <a name="create_record_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateRecordCall") CreateRecordCall

> Record a Call


```go
func (me *CALL_IMPL) CreateRecordCall(
            callSid string,
            record bool,
            direction models_pkg.DirectionEnum,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.AudioFormatEnum,
            responseType *string)(string,error)
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
callSid := "CallSid"
record := false
direction := models_pkg.Direction_IN
timeLimit,_ := strconv.ParseInt("20", 10, 8)
callBackUrl := "CallBackUrl"
fileformat := models_pkg.AudioFormat_MP3
responseType := "json"

var result string
result,_ = call.CreateRecordCall(callSid, record, direction, timeLimit, callBackUrl, fileformat, responseType)

```


#### <a name="create_voice_effect"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateVoiceEffect") CreateVoiceEffect

> Voice Effect


```go
func (me *CALL_IMPL) CreateVoiceEffect(
            callSid string,
            audioDirection models_pkg.AudioDirectionEnum,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64,
            responseType *string)(string,error)
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
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
callSid := "CallSid"
audioDirection := models_pkg.AudioDirection_IN
pitchSemiTones := 20.6740192373628
pitchOctaves := 20.6740192373628
pitch := 20.6740192373628
rate := 20.6740192373628
tempo := 20.6740192373628
responseType := "json"

var result string
result,_ = call.CreateVoiceEffect(callSid, audioDirection, pitchSemiTones, pitchOctaves, pitch, rate, tempo, responseType)

```


#### <a name="create_send_digit"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateSendDigit") CreateSendDigit

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreateSendDigit(
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.DirectionEnum,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
callSid := "CallSid"
playDtmf := "PlayDtmf"
playDtmfDirection := models_pkg.Direction_IN
responseType := "json"

var result string
result,_ = call.CreateSendDigit(callSid, playDtmf, playDtmfDirection, responseType)

```


#### <a name="create_interrupted_call"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateInterruptedCall") CreateInterruptedCall

> Interrupt the Call by Call Sid


```go
func (me *CALL_IMPL) CreateInterruptedCall(
            callSid string,
            url *string,
            method models_pkg.HttpMethodEnum,
            status models_pkg.InterruptedCallStatusEnum,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | Call SId |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
callSid := "CallSid"
url := "Url"
method := models_pkg.HttpMethod_GET
status := models_pkg.InterruptedCallStatus_CANCELED
responseType := "json"

var result string
result,_ = call.CreateInterruptedCall(callSid, url, method, status, responseType)

```


#### <a name="create_list_calls"></a>![Method: ](http://apidocs.io/img/method.png ".call_pkg.CreateListCalls") CreateListCalls

> A list of calls associated with your Message360 account


```go
func (me *CALL_IMPL) CreateListCalls(
            page *string,
            pageSize *string,
            to *string,
            from *string,
            dateCreated *string,
            responseType *string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Zero indexed |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Only list calls to this number |
| from |  ``` Optional ```  | Only list calls from this number |
| dateCreated |  ``` Optional ```  | Only list calls starting within the specified date range |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page := "Page"
pageSize := "PageSize"
to := "To"
from := "From"
dateCreated := "DateCreated"
responseType := "json"

var result 
result,_ = call.CreateListCalls(page, pageSize, to, from, dateCreated, responseType)

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
func (me *CARRIER_IMPL) CreateCarrierLookup(
            phonenumber string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | The number to lookup |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
phonenumber := "phonenumber"
responseType := "json"

var result string
result,_ = carrier.CreateCarrierLookup(phonenumber, responseType)

```


#### <a name="create_carrier_lookup_list"></a>![Method: ](http://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookupList") CreateCarrierLookupList

> Get the All Purchase Number's Carrier lookup


```go
func (me *CARRIER_IMPL) CreateCarrierLookupList(
            page *string,
            pagesize *string,
            responseType *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Page Number |
| pagesize |  ``` Optional ```  | Page Size |
| responseType |  ``` Optional ```  ``` DefaultValue ```  | Response format, xml or json |


#### Example Usage

```go
page := "page"
pagesize := "pagesize"
responseType := "json"

var result string
result,_ = carrier.CreateCarrierLookupList(page, pagesize, responseType)

```


[Back to List of Controllers](#list_of_controllers)



