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

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Ytel%20API%20V3-GoLang&projectName=ytelapiv3_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=ytelapiv3_lib)

## How to Use

The following section explains how to use the Ytelapiv3Lib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=ytelapiv3_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=ytelapiv3_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=ytelapiv3_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=ytelapiv3_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=ytelapiv3_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Ytel%20API%20V3-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=ytelapiv3_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=ytelapiv3_lib)

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

* [shortcode_pkg](#shortcode_pkg)
* [areamail_pkg](#areamail_pkg)
* [postcard_pkg](#postcard_pkg)
* [letter_pkg](#letter_pkg)
* [call_pkg](#call_pkg)
* [phonenumber_pkg](#phonenumber_pkg)
* [sms_pkg](#sms_pkg)
* [sharedshortcode_pkg](#sharedshortcode_pkg)
* [conference_pkg](#conference_pkg)
* [carrier_pkg](#carrier_pkg)
* [email_pkg](#email_pkg)
* [account_pkg](#account_pkg)
* [subaccount_pkg](#subaccount_pkg)
* [address_pkg](#address_pkg)
* [recording_pkg](#recording_pkg)
* [transcription_pkg](#transcription_pkg)
* [usage_pkg](#usage_pkg)

## <a name="shortcode_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".shortcode_pkg") shortcode_pkg

### Get instance

Factory for the ``` SHORTCODE ``` interface can be accessed from the package shortcode_pkg.

```go
shortCode := shortcode_pkg.NewSHORTCODE()
```

### <a name="create_list_shortcodes"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateListShortcodes") CreateListShortcodes

> Retrieve a list of Short Code assignment associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) CreateListShortcodes(
            shortcode *string,
            page *string,
            pagesize *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Optional ```  | Only list Short Code Assignment sent from this Short Code |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |


#### Example Usage

```go
shortcode := "Shortcode"
page := "page"
pagesize := "pagesize"

var result string
result,_ = shortCode.CreateListShortcodes(shortcode, page, pagesize)

```


### <a name="create_view_sms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateViewSMS") CreateViewSMS

> Retrieve a single Short Code object by its shortcode assignment.


```go
func (me *SHORTCODE_IMPL) CreateViewSMS(shortcode string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Dedicated Short Code to your Ytel account |


#### Example Usage

```go
shortcode := "Shortcode"

var result string
result,_ = shortCode.CreateViewSMS(shortcode)

```


### <a name="create_view_sms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateViewSMS") CreateViewSMS

> View a single Sms Short Code message.


```go
func (me *SHORTCODE_IMPL) CreateViewSMS(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for the sms resource |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = shortCode.CreateViewSMS(messageSid)

```


### <a name="create_list_sms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateListSMS") CreateListSMS

> Retrieve a list of Short Code messages.


```go
func (me *SHORTCODE_IMPL) CreateListSMS(
            shortcode *string,
            to *string,
            dateSent *string,
            page *int64,
            pageSize *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Optional ```  | Only list messages sent from this Short Code |
| to |  ``` Optional ```  | Only list messages sent to this number |
| dateSent |  ``` Optional ```  | Only list messages sent with the specified date |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | The count of objects to return per page. |


#### Example Usage

```go
shortcode := "Shortcode"
to := "To"
dateSent := "DateSent"
page,_ := strconv.ParseInt("71", 10, 8)
pageSize,_ := strconv.ParseInt("71", 10, 8)

var result string
result,_ = shortCode.CreateListSMS(shortcode, to, dateSent, page, pageSize)

```


### <a name="create_send_sms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateSendSMS") CreateSendSMS

> Send Dedicated Shortcode


```go
func (me *SHORTCODE_IMPL) CreateSendSMS(
            shortcode int64,
            to float64,
            body string,
            method *string,
            messagestatuscallback *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | Your dedicated shortcode |
| to |  ``` Required ```  | The number to send your SMS to |
| body |  ``` Required ```  | The body of your message |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once the Short Code message is sent.GET or POST |
| messagestatuscallback |  ``` Optional ```  | URL that can be requested to receive notification when Short Code message was sent. |


#### Example Usage

```go
shortcode,_ := strconv.ParseInt("71", 10, 8)
to := 71.9849464120273
body := "body"
method := "method"
messagestatuscallback := "messagestatuscallback"

var result string
result,_ = shortCode.CreateSendSMS(shortcode, to, body, method, messagestatuscallback)

```


### <a name="create_list_inbound_sms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateListInboundSMS") CreateListInboundSMS

> Retrive a list of inbound Sms Short Code messages associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) CreateListInboundSMS(
            page *int64,
            pagesize *int64,
            from *string,
            shortcode *string,
            datecreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Only list SMS messages sent from this number |
| shortcode |  ``` Optional ```  | Only list SMS messages sent to Shortcode |
| datecreated |  ``` Optional ```  | Only list SMS messages sent in the specified date MAKE REQUEST |


#### Example Usage

```go
page,_ := strconv.ParseInt("71", 10, 8)
pagesize,_ := strconv.ParseInt("71", 10, 8)
from := "From"
shortcode := "Shortcode"
datecreated := "Datecreated"

var result string
result,_ = shortCode.CreateListInboundSMS(page, pagesize, from, shortcode, datecreated)

```


### <a name="update_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.UpdateShortcode") UpdateShortcode

> Update a dedicated shortcode.


```go
func (me *SHORTCODE_IMPL) UpdateShortcode(
            shortcode string,
            friendlyName *string,
            callbackMethod *string,
            callbackUrl *string,
            fallbackMethod *string,
            fallbackUrl *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid dedicated shortcode to your Ytel account. |
| friendlyName |  ``` Optional ```  | User generated name of the dedicated shortcode. |
| callbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |
| callbackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| fallbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |
| fallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST. |


#### Example Usage

```go
shortcode := "Shortcode"
friendlyName := "FriendlyName"
callbackMethod := "CallbackMethod"
callbackUrl := "CallbackUrl"
fallbackMethod := "FallbackMethod"
fallbackUrl := "FallbackUrl"

var result string
result,_ = shortCode.UpdateShortcode(shortcode, friendlyName, callbackMethod, callbackUrl, fallbackMethod, fallbackUrl)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="areamail_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".areamail_pkg") areamail_pkg

### Get instance

Factory for the ``` AREAMAIL ``` interface can be accessed from the package areamail_pkg.

```go
areaMail := areamail_pkg.NewAREAMAIL()
```

### <a name="create_delete_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateDeleteAreaMail") CreateDeleteAreaMail

> Remove an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) CreateDeleteAreaMail(areamailid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |


#### Example Usage

```go
areamailid := "areamailid"

var result string
result,_ = areaMail.CreateDeleteAreaMail(areamailid)

```


### <a name="create_view_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateViewAreaMail") CreateViewAreaMail

> Retrieve an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) CreateViewAreaMail(areamailid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |


#### Example Usage

```go
areamailid := "areamailid"

var result string
result,_ = areaMail.CreateViewAreaMail(areamailid)

```


### <a name="create_list_area_mails"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateListAreaMails") CreateListAreaMails

> Retrieve a list of AreaMail objects.


```go
func (me *AREAMAIL_IMPL) CreateListAreaMails(
            page *int64,
            pagesize *int64,
            areamailsid *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| areamailsid |  ``` Optional ```  | The unique identifier for an AreaMail object. |
| dateCreated |  ``` Optional ```  | The date the AreaMail was created. |


#### Example Usage

```go
page,_ := strconv.ParseInt("71", 10, 8)
pagesize,_ := strconv.ParseInt("71", 10, 8)
areamailsid := "areamailsid"
dateCreated := "dateCreated"

var result string
result,_ = areaMail.CreateListAreaMails(page, pagesize, areamailsid, dateCreated)

```


### <a name="create_area_mail"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreaMail") CreateAreaMail

> Create a new AreaMail object.


```go
func (me *AREAMAIL_IMPL) CreateAreaMail(
            routes string,
            attachbyid string,
            front string,
            back string,
            description *string,
            targettype *string,
            htmldata *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| routes |  ``` Required ```  | List of routes that AreaMail should be delivered to.  A single route can be either a zipcode or a carrier route.List of routes that AreaMail should be delivered to.  A single route can be either a zipcode or a carrier route. A carrier route is in the form of 92610-C043 where the carrier route is composed of 5 numbers for zipcode, 1 letter for carrier route type, and 3 numbers for carrier route code. Delivery can be sent to mutliple routes by separating them with a commas. Valid Values: 92656, 92610-C043 |
| attachbyid |  ``` Required ```  | Set an existing areamail by attaching its AreamailId. |
| front |  ``` Required ```  | The front of the AreaMail item to be created. This can be a URL, local file, or an HTML string. Supported file types are PDF, PNG, and JPEG. Back required |
| back |  ``` Required ```  | The back of the AreaMail item to be created. This can be a URL, local file, or an HTML string. Supported file types are PDF, PNG, and JPEG. |
| description |  ``` Optional ```  | A string value to use as a description for this AreaMail item. |
| targettype |  ``` Optional ```  | The delivery location type. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
routes := "routes"
attachbyid := "attachbyid"
front := "front"
back := "back"
description := "description"
targettype := "targettype"
htmldata := "htmldata"

var result string
result,_ = areaMail.CreateAreaMail(routes, attachbyid, front, back, description, targettype, htmldata)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="postcard_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".postcard_pkg") postcard_pkg

### Get instance

Factory for the ``` POSTCARD ``` interface can be accessed from the package postcard_pkg.

```go
postCard := postcard_pkg.NewPOSTCARD()
```

### <a name="create_delete_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.CreateDeletePostcard") CreateDeletePostcard

> Remove a postcard object.


```go
func (me *POSTCARD_IMPL) CreateDeletePostcard(postcardid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier of a postcard object. |


#### Example Usage

```go
postcardid := "postcardid"

var result string
result,_ = postCard.CreateDeletePostcard(postcardid)

```


### <a name="create_view_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.CreateViewPostcard") CreateViewPostcard

> Retrieve a postcard object by its PostcardId.


```go
func (me *POSTCARD_IMPL) CreateViewPostcard(postcardid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier for a postcard object. |


#### Example Usage

```go
postcardid := "postcardid"

var result string
result,_ = postCard.CreateViewPostcard(postcardid)

```


### <a name="create_postcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.CreatePostcard") CreatePostcard

> Create, print, and mail a postcard to an address. The postcard front must be supplied as a PDF, image, or an HTML string. The back can be a PDF, image, HTML string, or it can be automatically generated by supplying a custom message.


```go
func (me *POSTCARD_IMPL) CreatePostcard(
            to string,
            from string,
            attachbyid string,
            front string,
            back string,
            message string,
            setting string,
            description *string,
            htmldata *string)(string,error)
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
| description |  ``` Optional ```  | A description for the postcard. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
to := "to"
from := "from"
attachbyid := "attachbyid"
front := "front"
back := "back"
message := "message"
setting := "setting"
description := "description"
htmldata := "htmldata"

var result string
result,_ = postCard.CreatePostcard(to, from, attachbyid, front, back, message, setting, description, htmldata)

```


### <a name="create_list_postcards"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.CreateListPostcards") CreateListPostcards

> Retrieve a list of postcard objects. The postcards objects are sorted by creation date, with the most recently created postcards appearing first.


```go
func (me *POSTCARD_IMPL) CreateListPostcards(
            page *int64,
            pagesize *int64,
            postcardid *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| postcardid |  ``` Optional ```  | The unique identifier for a postcard object. |
| dateCreated |  ``` Optional ```  | The date the postcard was created. |


#### Example Usage

```go
page,_ := strconv.ParseInt("71", 10, 8)
pagesize,_ := strconv.ParseInt("71", 10, 8)
postcardid := "postcardid"
dateCreated := "dateCreated"

var result string
result,_ = postCard.CreateListPostcards(page, pagesize, postcardid, dateCreated)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="letter_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".letter_pkg") letter_pkg

### Get instance

Factory for the ``` LETTER ``` interface can be accessed from the package letter_pkg.

```go
letter := letter_pkg.NewLETTER()
```

### <a name="create_delete_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateDeleteLetter") CreateDeleteLetter

> Remove a letter object by its LetterId.


```go
func (me *LETTER_IMPL) CreateDeleteLetter(lettersid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |


#### Example Usage

```go
lettersid := "lettersid"

var result string
result,_ = letter.CreateDeleteLetter(lettersid)

```


### <a name="create_view_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateViewLetter") CreateViewLetter

> Retrieve a letter object by its LetterSid.


```go
func (me *LETTER_IMPL) CreateViewLetter(lettersid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |


#### Example Usage

```go
lettersid := "lettersid"

var result string
result,_ = letter.CreateViewLetter(lettersid)

```


### <a name="create_letter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetter") CreateLetter

> Create, print, and mail a letter to an address. The letter file must be supplied as a PDF or an HTML string.


```go
func (me *LETTER_IMPL) CreateLetter(
            to string,
            from string,
            attachbyid string,
            file string,
            color string,
            description *string,
            extraservice *string,
            doublesided *string,
            template *string,
            htmldata *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| from |  ``` Required ```  | The AddressId or an object structured when creating an address by addresses/Create. |
| attachbyid |  ``` Required ```  | Set an existing letter by attaching its LetterId. |
| file |  ``` Required ```  | File can be a 8.5"x11" PDF uploaded file or URL that links to a file. |
| color |  ``` Required ```  | Specify if letter should be printed in color. |
| description |  ``` Optional ```  | A description for the letter. |
| extraservice |  ``` Optional ```  | Add an extra service to your letter. Options are "certified" or "registered". Certified provides tracking and delivery confirmation for domestic destinations and is an additional $5.00. Registered provides tracking and confirmation for international addresses and is an additional $16.50. |
| doublesided |  ``` Optional ```  | Specify if letter should be printed on both sides. |
| template |  ``` Optional ```  | Boolean that defaults to true. When set to false, this specifies that your letter does not follow the m360 address template. In this case, a blank address page will be inserted at the beginning of your file and you will be charged for the extra page. |
| htmldata |  ``` Optional ```  | A string value that contains HTML markup. |


#### Example Usage

```go
to := "to"
from := "from"
attachbyid := "attachbyid"
file := "file"
color := "color"
description := "description"
extraservice := "extraservice"
doublesided := "doublesided"
template := "template"
htmldata := "htmldata"

var result string
result,_ = letter.CreateLetter(to, from, attachbyid, file, color, description, extraservice, doublesided, template, htmldata)

```


### <a name="create_list_letters"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateListLetters") CreateListLetters

> Retrieve a list of letter objects. The letter objects are sorted by creation date, with the most recently created letters appearing first.


```go
func (me *LETTER_IMPL) CreateListLetters(
            page *int64,
            pagesize *int64,
            lettersid *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| lettersid |  ``` Optional ```  | The unique identifier for a letter object. |
| dateCreated |  ``` Optional ```  | The date the letter was created. |


#### Example Usage

```go
page,_ := strconv.ParseInt("113", 10, 8)
pagesize,_ := strconv.ParseInt("113", 10, 8)
lettersid := "lettersid"
dateCreated := "dateCreated"

var result string
result,_ = letter.CreateListLetters(page, pagesize, lettersid, dateCreated)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="call_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".call_pkg") call_pkg

### Get instance

Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.

```go
call := call_pkg.NewCALL()
```

### <a name="create_view_call1"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateViewCall1") CreateViewCall1

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) CreateViewCall1(callSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the voice call. |


#### Example Usage

```go
callSid := "callSid"

var result string
result,_ = call.CreateViewCall1(callSid)

```


### <a name="create_view_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateViewCall") CreateViewCall

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) CreateViewCall(callsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | The unique identifier for the voice call. |


#### Example Usage

```go
callsid := "callsid"

var result string
result,_ = call.CreateViewCall(callsid)

```


### <a name="create_play_dtmf"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreatePlayDTMF") CreatePlayDTMF

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreatePlayDTMF(
            callSid string,
            playDtmf string,
            playDtmfDirection models_pkg.PlayDtmfDirectionEnum)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |


#### Example Usage

```go
callSid := "CallSid"
playDtmf := "PlayDtmf"
playDtmfDirection := models_pkg.PlayDtmfDirection_IN

var result string
result,_ = call.CreatePlayDTMF(callSid, playDtmf, playDtmfDirection)

```


### <a name="create_make_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateMakeCall") CreateMakeCall

> You can experiment with initiating a call through Ytel and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) CreateMakeCall(
            from string,
            to string,
            url string,
            method *string,
            statusCallBackUrl *string,
            statusCallBackMethod *string,
            fallBackUrl *string,
            fallBackMethod *string,
            heartBeatUrl *string,
            heartBeatMethod *string,
            timeout *int64,
            playDtmf *string,
            hideCallerId *bool,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod *string,
            transcribe *bool,
            transcribeCallBackUrl *string,
            ifMachine models_pkg.IfMachineEnum,
            ifMachineUrl *string,
            ifMachineMethod *string,
            feedback *bool,
            surveyId *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call. |
| to |  ``` Required ```  | To number |
| url |  ``` Required ```  | URL requested once the call connects |
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
from := "From"
to := "To"
url := "Url"
method := "Method"
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := "StatusCallBackMethod"
fallBackUrl := "FallBackUrl"
fallBackMethod := "FallBackMethod"
heartBeatUrl := "HeartBeatUrl"
heartBeatMethod := "HeartBeatMethod"
timeout,_ := strconv.ParseInt("113", 10, 8)
playDtmf := "PlayDtmf"
hideCallerId := false
record := false
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := "RecordCallBackMethod"
transcribe := false
transcribeCallBackUrl := "TranscribeCallBackUrl"
ifMachine := models_pkg.IfMachine_CONTINUE
ifMachineUrl := "IfMachineUrl"
ifMachineMethod := "IfMachineMethod"
feedback := false
surveyId := "SurveyId"

var result string
result,_ = call.CreateMakeCall(from, to, url, method, statusCallBackUrl, statusCallBackMethod, fallBackUrl, fallBackMethod, heartBeatUrl, heartBeatMethod, timeout, playDtmf, hideCallerId, record, recordCallBackUrl, recordCallBackMethod, transcribe, transcribeCallBackUrl, ifMachine, ifMachineUrl, ifMachineMethod, feedback, surveyId)

```


### <a name="create_play_audio"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreatePlayAudio") CreatePlayAudio

> Play Audio from a url


```go
func (me *CALL_IMPL) CreatePlayAudio(
            callSid string,
            audioUrl string,
            sayText string,
            length *int64,
            direction models_pkg.DirectionEnum,
            mix *bool)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| audioUrl |  ``` Required ```  | URL to sound that should be played. You also can add more than one audio file using semicolons e.g. http://example.com/audio1.mp3;http://example.com/audio2.wav |
| sayText |  ``` Required ```  | Valid alphanumeric string that should be played to the In-progress call. |
| length |  ``` Optional ```  | Time limit in seconds for audio play back |
| direction |  ``` Optional ```  | The leg of the call audio will be played to |
| mix |  ``` Optional ```  | If false, all other audio will be muted |


#### Example Usage

```go
callSid := "CallSid"
audioUrl := "AudioUrl"
sayText := "SayText"
length,_ := strconv.ParseInt("113", 10, 8)
direction := models_pkg.Direction_IN
mix := false

var result string
result,_ = call.CreatePlayAudio(callSid, audioUrl, sayText, length, direction, mix)

```


### <a name="create_record_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateRecordCall") CreateRecordCall

> Start or stop recording of an in-progress voice call.


```go
func (me *CALL_IMPL) CreateRecordCall(
            callSid string,
            record bool,
            direction models_pkg.Direction4Enum,
            timeLimit *int64,
            callBackUrl *string,
            fileformat models_pkg.FileformatEnum)(string,error)
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


#### Example Usage

```go
callSid := "CallSid"
record := false
direction := models_pkg.Direction4_IN
timeLimit,_ := strconv.ParseInt("113", 10, 8)
callBackUrl := "CallBackUrl"
fileformat := models_pkg.Fileformat_MP3

var result string
result,_ = call.CreateRecordCall(callSid, record, direction, timeLimit, callBackUrl, fileformat)

```


### <a name="create_voice_effect"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateVoiceEffect") CreateVoiceEffect

> Add audio voice effects to the an in-progress voice call.


```go
func (me *CALL_IMPL) CreateVoiceEffect(
            callSid string,
            audioDirection models_pkg.AudioDirectionEnum,
            pitchSemiTones *float64,
            pitchOctaves *float64,
            pitch *float64,
            rate *float64,
            tempo *float64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the in-progress voice call. |
| audioDirection |  ``` Optional ```  | The direction the audio effect should be placed on. If IN, the effects will occur on the incoming audio stream. If OUT, the effects will occur on the outgoing audio stream. |
| pitchSemiTones |  ``` Optional ```  | Set the pitch in semitone (half-step) intervals. Value between -14 and 14 |
| pitchOctaves |  ``` Optional ```  | Set the pitch in octave intervals.. Value between -1 and 1 |
| pitch |  ``` Optional ```  | Set the pitch (lowness/highness) of the audio. The higher the value, the higher the pitch. Value greater than 0 |
| rate |  ``` Optional ```  | Set the rate for audio. The lower the value, the lower the rate. value greater than 0 |
| tempo |  ``` Optional ```  | Set the tempo (speed) of the audio. A higher value denotes a faster tempo. Value greater than 0 |


#### Example Usage

```go
callSid := "CallSid"
audioDirection := models_pkg.AudioDirection_IN
pitchSemiTones := 113.708119638594
pitchOctaves := 113.708119638594
pitch := 113.708119638594
rate := 113.708119638594
tempo := 113.708119638594

var result string
result,_ = call.CreateVoiceEffect(callSid, audioDirection, pitchSemiTones, pitchOctaves, pitch, rate, tempo)

```


### <a name="create_interrupt_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateInterruptCall") CreateInterruptCall

> Interrupt the Call by Call Sid


```go
func (me *CALL_IMPL) CreateInterruptCall(
            callSid string,
            url *string,
            method *string,
            status models_pkg.StatusEnum)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for voice call that is in progress. |
| url |  ``` Optional ```  | URL the in-progress call will be redirected to |
| method |  ``` Optional ```  | The method used to request the above Url parameter |
| status |  ``` Optional ```  | Status to set the in-progress call to |


#### Example Usage

```go
callSid := "CallSid"
url := "Url"
method := "Method"
status := models_pkg.Status_CANCELED

var result string
result,_ = call.CreateInterruptCall(callSid, url, method, status)

```


### <a name="create_list_calls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateListCalls") CreateListCalls

> A list of calls associated with your Ytel account


```go
func (me *CALL_IMPL) CreateListCalls(
            page *int64,
            pageSize *int64,
            to *string,
            from *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Filter calls that were sent to this 10-digit number (E.164 format). |
| from |  ``` Optional ```  | Filter calls that were sent from this 10-digit number (E.164 format). |
| dateCreated |  ``` Optional ```  | Return calls that are from a specified date. |


#### Example Usage

```go
page,_ := strconv.ParseInt("113", 10, 8)
pageSize,_ := strconv.ParseInt("113", 10, 8)
to := "To"
from := "From"
dateCreated := "DateCreated"

var result string
result,_ = call.CreateListCalls(page, pageSize, to, from, dateCreated)

```


### <a name="create_send_rvm"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateSendRVM") CreateSendRVM

> Initiate an outbound Ringless Voicemail through Ytel.


```go
func (me *CALL_IMPL) CreateSendRVM(
            from string,
            rVMCallerId string,
            to string,
            voiceMailURL string,
            method *string,
            statusCallBackUrl *string,
            statsCallBackMethod *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call. |
| rVMCallerId |  ``` Required ```  | A required secondary Caller ID for RVM to work. |
| to |  ``` Required ```  | A valid number (E.164 format) that will receive the phone call. |
| voiceMailURL |  ``` Required ```  | The URL requested once the RVM connects. A set of default parameters will be sent here. |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| statsCallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |


#### Example Usage

```go
from := "From"
rVMCallerId := "RVMCallerId"
to := "To"
voiceMailURL := "VoiceMailURL"
method := "Method"
statusCallBackUrl := "StatusCallBackUrl"
statsCallBackMethod := "StatsCallBackMethod"

var result string
result,_ = call.CreateSendRVM(from, rVMCallerId, to, voiceMailURL, method, statusCallBackUrl, statsCallBackMethod)

```


### <a name="create_group_call"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateGroupCall") CreateGroupCall

> Group Call


```go
func (me *CALL_IMPL) CreateGroupCall(
            from string,
            to string,
            url string,
            groupConfirmKey string,
            groupConfirmFile models_pkg.GroupConfirmFileEnum,
            method *string,
            statusCallBackUrl *string,
            statusCallBackMethod *string,
            fallBackUrl *string,
            fallBackMethod *string,
            heartBeatUrl *string,
            heartBeatMethod *string,
            timeout *int64,
            playDtmf *string,
            hideCallerId *string,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod *string,
            transcribe *bool,
            transcribeCallBackUrl *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | This number to display on Caller ID as calling |
| to |  ``` Required ```  | Please enter multiple E164 number. You can add max 10 numbers. Add numbers separated with comma. e.g : 1111111111,2222222222 |
| url |  ``` Required ```  | URL requested once the call connects |
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
from := "From"
to := "To"
url := "Url"
groupConfirmKey := "GroupConfirmKey"
groupConfirmFile := models_pkg.GroupConfirmFile_MP3
method := "Method"
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := "StatusCallBackMethod"
fallBackUrl := "FallBackUrl"
fallBackMethod := "FallBackMethod"
heartBeatUrl := "HeartBeatUrl"
heartBeatMethod := "HeartBeatMethod"
timeout,_ := strconv.ParseInt("113", 10, 8)
playDtmf := "PlayDtmf"
hideCallerId := "HideCallerId"
record := false
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := "RecordCallBackMethod"
transcribe := false
transcribeCallBackUrl := "TranscribeCallBackUrl"

var result string
result,_ = call.CreateGroupCall(from, to, url, groupConfirmKey, groupConfirmFile, method, statusCallBackUrl, statusCallBackMethod, fallBackUrl, fallBackMethod, heartBeatUrl, heartBeatMethod, timeout, playDtmf, hideCallerId, record, recordCallBackUrl, recordCallBackMethod, transcribe, transcribeCallBackUrl)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="phonenumber_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

### <a name="create_get_did_score"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateGetDIDScore") CreateGetDIDScore

> Get DID Score Number


```go
func (me *PHONENUMBER_IMPL) CreateGetDIDScore(phonenumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | Specifies the multiple phone numbers for check updated spamscore . |


#### Example Usage

```go
phonenumber := "Phonenumber"

var result string
result,_ = phoneNumber.CreateGetDIDScore(phonenumber)

```


### <a name="create_move_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateMoveNumber") CreateMoveNumber

> Transfer phone number that has been purchased for from one account to another account.


```go
func (me *PHONENUMBER_IMPL) CreateMoveNumber(
            phonenumber string,
            fromaccountsid string,
            toaccountsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
| fromaccountsid |  ``` Required ```  | A specific Accountsid from where Number is getting transfer. |
| toaccountsid |  ``` Required ```  | A specific Accountsid to which Number is getting transfer. |


#### Example Usage

```go
phonenumber := "phonenumber"
fromaccountsid := "fromaccountsid"
toaccountsid := "toaccountsid"

var result string
result,_ = phoneNumber.CreateMoveNumber(phonenumber, fromaccountsid, toaccountsid)

```


### <a name="create_purchase_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreatePurchaseNumber") CreatePurchaseNumber

> Purchase a phone number to be used with your Ytel account


```go
func (me *PHONENUMBER_IMPL) CreatePurchaseNumber(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreatePurchaseNumber(phoneNumber)

```


### <a name="create_release_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateReleaseNumber") CreateReleaseNumber

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) CreateReleaseNumber(
            responseType string,
            phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| responseType |  ``` Required ```  | Response type format xml or json |
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |


#### Example Usage

```go
responseType := "ResponseType"
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateReleaseNumber(responseType, phoneNumber)

```


### <a name="create_view_details"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateViewDetails") CreateViewDetails

> Retrieve the details for a phone number by its number.


```go
func (me *PHONENUMBER_IMPL) CreateViewDetails(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel 10-digit phone number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateViewDetails(phoneNumber)

```


### <a name="create_bulk_release"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateBulkRelease") CreateBulkRelease

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) CreateBulkRelease(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel comma separated numbers (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateBulkRelease(phoneNumber)

```


### <a name="create_available_numbers"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateAvailableNumbers") CreateAvailableNumbers

> Retrieve a list of available phone numbers that can be purchased and used for your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateAvailableNumbers(
            numbertype models_pkg.NumbertypeEnum,
            areacode string,
            pagesize *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numbertype |  ``` Required ```  | Number type either SMS,Voice or all |
| areacode |  ``` Required ```  | Specifies the area code for the returned list of available numbers. Only available for North American numbers. |
| pagesize |  ``` Optional ```  | The count of objects to return. |


#### Example Usage

```go
numbertype := models_pkg.numbertype_ALL
areacode := "areacode"
pagesize,_ := strconv.ParseInt("113", 10, 8)

var result string
result,_ = phoneNumber.CreateAvailableNumbers(numbertype, areacode, pagesize)

```


### <a name="update_number"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.UpdateNumber") UpdateNumber

> Update properties for a Ytel number that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) UpdateNumber(
            phoneNumber string,
            voiceUrl string,
            friendlyName *string,
            voiceMethod *string,
            voiceFallbackUrl *string,
            voiceFallbackMethod *string,
            hangupCallback *string,
            hangupCallbackMethod *string,
            heartbeatUrl *string,
            heartbeatMethod *string,
            smsUrl *string,
            smsMethod *string,
            smsFallbackUrl *string,
            smsFallbackMethod *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel number (E.164 format). |
| voiceUrl |  ``` Required ```  | URL requested once the call connects |
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
phoneNumber := "PhoneNumber"
voiceUrl := "VoiceUrl"
friendlyName := "FriendlyName"
voiceMethod := "VoiceMethod"
voiceFallbackUrl := "VoiceFallbackUrl"
voiceFallbackMethod := "VoiceFallbackMethod"
hangupCallback := "HangupCallback"
hangupCallbackMethod := "HangupCallbackMethod"
heartbeatUrl := "HeartbeatUrl"
heartbeatMethod := "HeartbeatMethod"
smsUrl := "SmsUrl"
smsMethod := "SmsMethod"
smsFallbackUrl := "SmsFallbackUrl"
smsFallbackMethod := "SmsFallbackMethod"

var result string
result,_ = phoneNumber.UpdateNumber(phoneNumber, voiceUrl, friendlyName, voiceMethod, voiceFallbackUrl, voiceFallbackMethod, hangupCallback, hangupCallbackMethod, heartbeatUrl, heartbeatMethod, smsUrl, smsMethod, smsFallbackUrl, smsFallbackMethod)

```


### <a name="create_list_numbers"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateListNumbers") CreateListNumbers

> Retrieve a list of purchased phones numbers associated with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateListNumbers(
            page *int64,
            pageSize *int64,
            numberType models_pkg.NumberType14Enum,
            friendlyName *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | Which page of the overall response will be returned. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| numberType |  ``` Optional ```  | The capability supported by the number.Number type either SMS,Voice or all |
| friendlyName |  ``` Optional ```  | A human-readable label added to the number object. |


#### Example Usage

```go
page,_ := strconv.ParseInt("113", 10, 8)
pageSize,_ := strconv.ParseInt("113", 10, 8)
numberType := models_pkg.NumberType14_ALL
friendlyName := "FriendlyName"

var result string
result,_ = phoneNumber.CreateListNumbers(page, pageSize, numberType, friendlyName)

```


### <a name="create_bulk_update_numbers"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateBulkUpdateNumbers") CreateBulkUpdateNumbers

> Update properties for a Ytel numbers that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) CreateBulkUpdateNumbers(
            phoneNumber string,
            voiceUrl string,
            friendlyName *string,
            voiceMethod *string,
            voiceFallbackUrl *string,
            voiceFallbackMethod *string,
            hangupCallback *string,
            hangupCallbackMethod *string,
            heartbeatUrl *string,
            heartbeatMethod *string,
            smsUrl *string,
            smsMethod *string,
            smsFallbackUrl *string,
            smsFallbackMethod *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid comma(,) separated Ytel numbers. (E.164 format). |
| voiceUrl |  ``` Required ```  | The URL returning InboundXML incoming calls should execute when connected. |
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
phoneNumber := "PhoneNumber"
voiceUrl := "VoiceUrl"
friendlyName := "FriendlyName"
voiceMethod := "VoiceMethod"
voiceFallbackUrl := "VoiceFallbackUrl"
voiceFallbackMethod := "VoiceFallbackMethod"
hangupCallback := "HangupCallback"
hangupCallbackMethod := "HangupCallbackMethod"
heartbeatUrl := "HeartbeatUrl"
heartbeatMethod := "HeartbeatMethod"
smsUrl := "SmsUrl"
smsMethod := "SmsMethod"
smsFallbackUrl := "SmsFallbackUrl"
smsFallbackMethod := "SmsFallbackMethod"

var result string
result,_ = phoneNumber.CreateBulkUpdateNumbers(phoneNumber, voiceUrl, friendlyName, voiceMethod, voiceFallbackUrl, voiceFallbackMethod, hangupCallback, hangupCallbackMethod, heartbeatUrl, heartbeatMethod, smsUrl, smsMethod, smsFallbackUrl, smsFallbackMethod)

```


### <a name="create_bulk_buy_numbers"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateBulkBuyNumbers") CreateBulkBuyNumbers

> Purchase a selected number of DID's from specific area codes to be used with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateBulkBuyNumbers(
            numberType models_pkg.NumberType15Enum,
            areaCode string,
            quantity string,
            leftover *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numberType |  ``` Required ```  | The capability the number supports. |
| areaCode |  ``` Required ```  | Specifies the area code for the returned list of available numbers. Only available for North American numbers. |
| quantity |  ``` Required ```  | A positive integer that tells how many number you want to buy at a time. |
| leftover |  ``` Optional ```  | If desired quantity is unavailable purchase what is available . |


#### Example Usage

```go
numberType := models_pkg.NumberType15_ALL
areaCode := "AreaCode"
quantity := "Quantity"
leftover := "Leftover"

var result string
result,_ = phoneNumber.CreateBulkBuyNumbers(numberType, areaCode, quantity, leftover)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sms_pkg") sms_pkg

### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

### <a name="create_view_sms1"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateViewSMS1") CreateViewSMS1

> Retrieve a single SMS message object with details by its SmsSid.


```go
func (me *SMS_IMPL) CreateViewSMS1(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = sMS.CreateViewSMS1(messageSid)

```


### <a name="create_view_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateViewSMS") CreateViewSMS

> Retrieve a single SMS message object by its SmsSid.


```go
func (me *SMS_IMPL) CreateViewSMS(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = sMS.CreateViewSMS(messageSid)

```


### <a name="create_send_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSendSMS") CreateSendSMS

> Send an SMS from a Ytel number


```go
func (me *SMS_IMPL) CreateSendSMS(
            from string,
            to string,
            body string,
            method *string,
            messageStatusCallback *string,
            smartsms *bool,
            deliveryStatus *bool)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | The 10-digit SMS-enabled Ytel number (E.164 format) in which the message is sent. |
| to |  ``` Required ```  | The 10-digit phone number (E.164 format) that will receive the message. |
| body |  ``` Required ```  | The body message that is to be sent in the text. |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |
| smartsms |  ``` Optional ```  | Check's 'to' number can receive sms or not using Carrier API, if wireless = true then text sms is sent, else wireless = false then call is recieved to end user with audible message. |
| deliveryStatus |  ``` Optional ```  | Delivery reports are a method to tell your system if the message has arrived on the destination phone. |


#### Example Usage

```go
from := "From"
to := "To"
body := "Body"
method := "Method"
messageStatusCallback := "MessageStatusCallback"
smartsms := false
deliveryStatus := false

var result string
result,_ = sMS.CreateSendSMS(from, to, body, method, messageStatusCallback, smartsms, deliveryStatus)

```


### <a name="create_list_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateListSMS") CreateListSMS

> Retrieve a list of Outbound SMS message objects.


```go
func (me *SMS_IMPL) CreateListSMS(
            page *int64,
            pageSize *int64,
            from *string,
            to *string,
            dateSent *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |


#### Example Usage

```go
page,_ := strconv.ParseInt("113", 10, 8)
pageSize,_ := strconv.ParseInt("113", 10, 8)
from := "From"
to := "To"
dateSent := "DateSent"

var result string
result,_ = sMS.CreateListSMS(page, pageSize, from, to, dateSent)

```


### <a name="create_list_inbound_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateListInboundSMS") CreateListInboundSMS

> Retrieve a list of Inbound SMS message objects.


```go
func (me *SMS_IMPL) CreateListInboundSMS(
            page *int64,
            pageSize *int64,
            from *string,
            to *string,
            dateSent *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | The count of objects to return per page. |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Filter sms message objects by this date. |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pageSize,_ := strconv.ParseInt("22", 10, 8)
from := "From"
to := "To"
dateSent := "DateSent"

var result string
result,_ = sMS.CreateListInboundSMS(page, pageSize, from, to, dateSent)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sharedshortcode_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sharedshortcode_pkg") sharedshortcode_pkg

### Get instance

Factory for the ``` SHAREDSHORTCODE ``` interface can be accessed from the package sharedshortcode_pkg.

```go
sharedShortCode := sharedshortcode_pkg.NewSHAREDSHORTCODE()
```

### <a name="create_view_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateViewShortcode") CreateViewShortcode

> The response returned here contains all resource properties associated with the given Shortcode.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateViewShortcode(shortcode string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Shortcode to your Ytel account |


#### Example Usage

```go
shortcode := "Shortcode"

var result string
result,_ = sharedShortCode.CreateViewShortcode(shortcode)

```


### <a name="create_view_keyword"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateViewKeyword") CreateViewKeyword

> View a set of properties for a single keyword.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateViewKeyword(keywordid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keywordid |  ``` Required ```  | The unique identifier of each keyword |


#### Example Usage

```go
keywordid := "Keywordid"

var result string
result,_ = sharedShortCode.CreateViewKeyword(keywordid)

```


### <a name="create_view_template"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateViewTemplate") CreateViewTemplate

> View a Shared ShortCode Template


```go
func (me *SHAREDSHORTCODE_IMPL) CreateViewTemplate(templateId uuid.UUID)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| templateId |  ``` Required ```  | The unique identifier for a template object |


#### Example Usage

```go
templateId := uuid.NewV4()

var result string
result,_ = sharedShortCode.CreateViewTemplate(templateId)

```


### <a name="create_list_inbound_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateListInboundSMS") CreateListInboundSMS

> List All Inbound ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) CreateListInboundSMS(
            datecreated *string,
            page *int64,
            pagesize *int64,
            from *string,
            shortcode *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| datecreated |  ``` Optional ```  | Only list messages sent with the specified date |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound ShortCode |
| shortcode |  ``` Optional ```  | Only list messages sent to this Short Code |


#### Example Usage

```go
datecreated := "Datecreated"
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
from := "from"
shortcode := "Shortcode"

var result string
result,_ = sharedShortCode.CreateListInboundSMS(datecreated, page, pagesize, from, shortcode)

```


### <a name="create_send_sms"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateSendSMS") CreateSendSMS

> Send an SMS from a Ytel ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) CreateSendSMS(
            shortcode string,
            to string,
            templateid uuid.UUID,
            data string,
            method *string,
            messageStatusCallback *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | The Short Code number that is the sender of this message |
| to |  ``` Required ```  | A valid 10-digit number that should receive the message |
| templateid |  ``` Required ```  | The unique identifier for the template used for the message |
| data |  ``` Required ```  | format of your data, example: {companyname}:test,{otpcode}:1234 |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once the Short Code message is sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when Short Code message was sent. |


#### Example Usage

```go
shortcode := "shortcode"
to := "to"
templateid := uuid.NewV4()
data := "data"
method := "Method"
messageStatusCallback := "MessageStatusCallback"

var result string
result,_ = sharedShortCode.CreateSendSMS(shortcode, to, templateid, data, method, messageStatusCallback)

```


### <a name="create_list_templates"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateListTemplates") CreateListTemplates

> List Shortcode Templates by Type


```go
func (me *SHAREDSHORTCODE_IMPL) CreateListTemplates(
            mtype *string,
            page *int64,
            pagesize *int64,
            shortcode *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mtype |  ``` Optional ```  | The type (category) of template Valid values: marketing, authorization |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| shortcode |  ``` Optional ```  | Only list templates of type |


#### Example Usage

```go
mtype := "type"
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
shortcode := "Shortcode"

var result string
result,_ = sharedShortCode.CreateListTemplates(mtype, page, pagesize, shortcode)

```


### <a name="create_list_keywords"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateListKeywords") CreateListKeywords

> Retrieve a list of keywords associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateListKeywords(
            page *int64,
            pagesize *int64,
            keyword *string,
            shortcode *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| keyword |  ``` Optional ```  | Only list keywords of keyword |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
keyword := "Keyword"
shortcode,_ := strconv.ParseInt("22", 10, 8)

var result string
result,_ = sharedShortCode.CreateListKeywords(page, pagesize, keyword, shortcode)

```


### <a name="create_list_shortcodes"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateListShortcodes") CreateListShortcodes

> Retrieve a list of shortcode assignment associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateListShortcodes(
            shortcode *string,
            page *int64,
            pagesize *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |


#### Example Usage

```go
shortcode := "Shortcode"
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)

var result string
result,_ = sharedShortCode.CreateListShortcodes(shortcode, page, pagesize)

```


### <a name="update_shortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.UpdateShortcode") UpdateShortcode

> Update Assignment


```go
func (me *SHAREDSHORTCODE_IMPL) UpdateShortcode(
            shortcode string,
            friendlyName *string,
            callbackUrl *string,
            callbackMethod *string,
            fallbackUrl *string,
            fallbackUrlMethod *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid shortcode to your Ytel account |
| friendlyName |  ``` Optional ```  | User generated name of the shortcode |
| callbackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| callbackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |
| fallbackUrl |  ``` Optional ```  | URL used if any errors occur during execution of InboundXML or at initial request of the required Url provided with the POST. |
| fallbackUrlMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required FallbackUrl once call connects. |


#### Example Usage

```go
shortcode := "Shortcode"
friendlyName := "FriendlyName"
callbackUrl := "CallbackUrl"
callbackMethod := "CallbackMethod"
fallbackUrl := "FallbackUrl"
fallbackUrlMethod := "FallbackUrlMethod"

var result string
result,_ = sharedShortCode.UpdateShortcode(shortcode, friendlyName, callbackUrl, callbackMethod, fallbackUrl, fallbackUrlMethod)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="conference_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".conference_pkg") conference_pkg

### Get instance

Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.

```go
conference := conference_pkg.NewCONFERENCE()
```

### <a name="create_play_audio"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreatePlayAudio") CreatePlayAudio

> Play an audio file during a conference.


```go
func (me *CONFERENCE_IMPL) CreatePlayAudio(
            conferenceSid string,
            participantSid string,
            audioUrl models_pkg.AudioUrlEnum)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
| audioUrl |  ``` Required ```  | The URL for the audio file that is to be played during the conference. Multiple audio files can be chained by using a semicolon. |


#### Example Usage

```go
conferenceSid := "ConferenceSid"
participantSid := "ParticipantSid"
audioUrl := models_pkg.AudioUrl_MP3

var result string
result,_ = conference.CreatePlayAudio(conferenceSid, participantSid, audioUrl)

```


### <a name="create_hangup_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateHangupParticipant") CreateHangupParticipant

> Remove a participant from a conference.


```go
func (me *CONFERENCE_IMPL) CreateHangupParticipant(
            participantSid string,
            conferenceSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |


#### Example Usage

```go
participantSid := "ParticipantSid"
conferenceSid := "ConferenceSid"

var result string
result,_ = conference.CreateHangupParticipant(participantSid, conferenceSid)

```


### <a name="create_view_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateViewConference") CreateViewConference

> Retrieve information about a conference by its ConferenceSid.


```go
func (me *CONFERENCE_IMPL) CreateViewConference(conferenceSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier of each conference resource |


#### Example Usage

```go
conferenceSid := "ConferenceSid"

var result string
result,_ = conference.CreateViewConference(conferenceSid)

```


### <a name="create_view_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateViewParticipant") CreateViewParticipant

> Retrieve information about a participant by its ParticipantSid.


```go
func (me *CONFERENCE_IMPL) CreateViewParticipant(
            conferenceSid string,
            participantSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |


#### Example Usage

```go
conferenceSid := "ConferenceSid"
participantSid := "ParticipantSid"

var result string
result,_ = conference.CreateViewParticipant(conferenceSid, participantSid)

```


### <a name="create_silence_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateSilenceParticipant") CreateSilenceParticipant

> Deaf Mute Participant


```go
func (me *CONFERENCE_IMPL) CreateSilenceParticipant(
            conferenceSid string,
            participantSid string,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | ID of the active conference |
| participantSid |  ``` Required ```  | ID of an active participant |
| muted |  ``` Optional ```  | Mute a participant |
| deaf |  ``` Optional ```  | Make it so a participant cant hear |


#### Example Usage

```go
conferenceSid := "conferenceSid"
participantSid := "ParticipantSid"
muted := false
deaf := false

var result string
result,_ = conference.CreateSilenceParticipant(conferenceSid, participantSid, muted, deaf)

```


### <a name="add_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.AddParticipant") AddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) AddParticipant(
            conferenceSid string,
            participantNumber string,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantNumber |  ``` Required ```  | The phone number of the participant to be added. |
| muted |  ``` Optional ```  | Specifies if participant should be muted. |
| deaf |  ``` Optional ```  | Specifies if the participant should hear audio in the conference. |


#### Example Usage

```go
conferenceSid := "ConferenceSid"
participantNumber := "ParticipantNumber"
muted := false
deaf := false

var result string
result,_ = conference.AddParticipant(conferenceSid, participantNumber, muted, deaf)

```


### <a name="create_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConference") CreateConference

> Here you can experiment with initiating a conference call through Ytel and view the request response generated when doing so.


```go
func (me *CONFERENCE_IMPL) CreateConference(
            url string,
            from string,
            to string,
            method *string,
            statusCallBackUrl *string,
            statusCallBackMethod *string,
            fallbackUrl *string,
            fallbackMethod *string,
            record *bool,
            recordCallBackUrl *string,
            recordCallBackMethod *string,
            scheduleTime *string,
            timeout *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| url |  ``` Required ```  | URL requested once the conference connects |
| from |  ``` Required ```  | A valid 10-digit number (E.164 format) that will be initiating the conference call. |
| to |  ``` Required ```  | A valid 10-digit number (E.164 format) that is to receive the conference call. |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once call connects. |
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
url := "Url"
from := "From"
to := "To"
method := "Method"
statusCallBackUrl := "StatusCallBackUrl"
statusCallBackMethod := "StatusCallBackMethod"
fallbackUrl := "FallbackUrl"
fallbackMethod := "FallbackMethod"
record := false
recordCallBackUrl := "RecordCallBackUrl"
recordCallBackMethod := "RecordCallBackMethod"
scheduleTime := "ScheduleTime"
timeout,_ := strconv.ParseInt("22", 10, 8)

var result string
result,_ = conference.CreateConference(url, from, to, method, statusCallBackUrl, statusCallBackMethod, fallbackUrl, fallbackMethod, record, recordCallBackUrl, recordCallBackMethod, scheduleTime, timeout)

```


### <a name="create_list_participants"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateListParticipants") CreateListParticipants

> Retrieve a list of participants for an in-progress conference.


```go
func (me *CONFERENCE_IMPL) CreateListParticipants(
            conferenceSid string,
            page *int64,
            pagesize *int64,
            muted *bool,
            deaf *bool)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference. |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| muted |  ``` Optional ```  | Specifies if participant should be muted. |
| deaf |  ``` Optional ```  | Specifies if the participant should hear audio in the conference. |


#### Example Usage

```go
conferenceSid := "ConferenceSid"
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
muted := false
deaf := false

var result string
result,_ = conference.CreateListParticipants(conferenceSid, page, pagesize, muted, deaf)

```


### <a name="create_list_conferences"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateListConferences") CreateListConferences

> Retrieve a list of conference objects.


```go
func (me *CONFERENCE_IMPL) CreateListConferences(
            page *int64,
            pagesize *int64,
            friendlyName *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| dateCreated |  ``` Optional ```  | Conference created date |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
friendlyName := "FriendlyName"
dateCreated := "DateCreated"

var result string
result,_ = conference.CreateListConferences(page, pagesize, friendlyName, dateCreated)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="carrier_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg

### Get instance

Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.

```go
carrier := carrier_pkg.NewCARRIER()
```

### <a name="create_lookup_carrier"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CreateLookupCarrier") CreateLookupCarrier

> Get the Carrier Lookup


```go
func (me *CARRIER_IMPL) CreateLookupCarrier(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = carrier.CreateLookupCarrier(phoneNumber)

```


### <a name="create_carrier_results"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierResults") CreateCarrierResults

> Retrieve a list of carrier lookup objects.


```go
func (me *CARRIER_IMPL) CreateCarrierResults(
            page *int64,
            pageSize *int64)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  | The count of objects to return per page. |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pageSize,_ := strconv.ParseInt("22", 10, 8)

var result string
result,_ = carrier.CreateCarrierResults(page, pageSize)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="email_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".email_pkg") email_pkg

### Get instance

Factory for the ``` EMAIL ``` interface can be accessed from the package email_pkg.

```go
email := email_pkg.NewEMAIL()
```

### <a name="create_remove_invalid_email"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateRemoveInvalidEmail") CreateRemoveInvalidEmail

> Remove an email from the invalid email list.


```go
func (me *EMAIL_IMPL) CreateRemoveInvalidEmail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the invalid email list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateRemoveInvalidEmail(email)

```


### <a name="create_blocked_emails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateBlockedEmails") CreateBlockedEmails

> Retrieve a list of emails that have been blocked.


```go
func (me *EMAIL_IMPL) CreateBlockedEmails(
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of blocked emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
offset := "Offset"
limit := "Limit"

var result string
result,_ = email.CreateBlockedEmails(offset, limit)

```


### <a name="create_spam_emails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateSpamEmails") CreateSpamEmails

> Retrieve a list of emails that are on the spam list.


```go
func (me *EMAIL_IMPL) CreateSpamEmails(
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of spam emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
offset := "Offset"
limit := "Limit"

var result string
result,_ = email.CreateSpamEmails(offset, limit)

```


### <a name="create_bounced_emails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateBouncedEmails") CreateBouncedEmails

> Retrieve a list of emails that have bounced.


```go
func (me *EMAIL_IMPL) CreateBouncedEmails(
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of bounced emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
offset := "Offset"
limit := "Limit"

var result string
result,_ = email.CreateBouncedEmails(offset, limit)

```


### <a name="create_remove_bounced_email"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateRemoveBouncedEmail") CreateRemoveBouncedEmail

> Remove an email address from the bounced list.


```go
func (me *EMAIL_IMPL) CreateRemoveBouncedEmail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to be remove from the bounced email list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateRemoveBouncedEmail(email)

```


### <a name="create_invalid_emails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateInvalidEmails") CreateInvalidEmails

> Retrieve a list of invalid email addresses.


```go
func (me *EMAIL_IMPL) CreateInvalidEmails(
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of invalid emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
offset := "Offset"
limit := "Limit"

var result string
result,_ = email.CreateInvalidEmails(offset, limit)

```


### <a name="create_list_unsubscribed_emails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateListUnsubscribedEmails") CreateListUnsubscribedEmails

> Retrieve a list of email addresses from the unsubscribe list.


```go
func (me *EMAIL_IMPL) CreateListUnsubscribedEmails(
            offset *string,
            limit *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of unsubscribed emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
offset := "Offset"
limit := "Limit"

var result string
result,_ = email.CreateListUnsubscribedEmails(offset, limit)

```


### <a name="create_remove_unsubscribed_email"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateRemoveUnsubscribedEmail") CreateRemoveUnsubscribedEmail

> Remove an email address from the list of unsubscribed emails.


```go
func (me *EMAIL_IMPL) CreateRemoveUnsubscribedEmail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the unsubscribe list. |


#### Example Usage

```go
email := "email"

var result string
result,_ = email.CreateRemoveUnsubscribedEmail(email)

```


### <a name="add_email_unsubscribe"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.AddEmailUnsubscribe") AddEmailUnsubscribe

> Add an email to the unsubscribe list


```go
func (me *EMAIL_IMPL) AddEmailUnsubscribe(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be added to the unsubscribe list |


#### Example Usage

```go
email := "email"

var result string
result,_ = email.AddEmailUnsubscribe(email)

```


### <a name="create_remove_blocked_address"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateRemoveBlockedAddress") CreateRemoveBlockedAddress

> Remove an email from blocked emails list.


```go
func (me *EMAIL_IMPL) CreateRemoveBlockedAddress(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to be remove from the blocked list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateRemoveBlockedAddress(email)

```


### <a name="create_remove_spam_address"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateRemoveSpamAddress") CreateRemoveSpamAddress

> Remove an email from the spam email list.


```go
func (me *EMAIL_IMPL) CreateRemoveSpamAddress(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the spam list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateRemoveSpamAddress(email)

```


### <a name="create_send_email"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateSendEmail") CreateSendEmail

> Create and submit an email message to one or more email addresses.


```go
func (me *EMAIL_IMPL) CreateSendEmail(
            to string,
            mtype models_pkg.TypeEnum,
            subject string,
            message string,
            from *string,
            cc *string,
            bcc *string,
            attachment *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| to |  ``` Required ```  | A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| mtype |  ``` Required ```  | Specifies the type of email to be sent |
| subject |  ``` Required ```  | The subject of the mail. Must be a valid string. |
| message |  ``` Required ```  | The email message that is to be sent in the text. |
| from |  ``` Optional ```  | A valid address that will send the email. |
| cc |  ``` Optional ```  | Carbon copy. A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| bcc |  ``` Optional ```  | Blind carbon copy. A valid address that will receive the email. Multiple addresses can be separated by using commas. |
| attachment |  ``` Optional ```  | A file that is attached to the email. Must be less than 7 MB in size. |


#### Example Usage

```go
to := "To"
mtype := models_pkg.Type_TEXT
subject := "Subject"
message := "Message"
from := "From"
cc := "Cc"
bcc := "Bcc"
attachment := "Attachment"

var result string
result,_ = email.CreateSendEmail(to, mtype, subject, message, from, cc, bcc, attachment)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="account_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".account_pkg") account_pkg

### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

### <a name="create_view_account"></a>![Method: ](https://apidocs.io/img/method.png ".account_pkg.CreateViewAccount") CreateViewAccount

> Retrieve information regarding your Ytel account by a specific date. The response object will contain data such as account status, balance, and account usage totals.


```go
func (me *ACCOUNT_IMPL) CreateViewAccount(date string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | Filter account information based on date. |


#### Example Usage

```go
date := "Date"

var result string
result,_ = account.CreateViewAccount(date)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="subaccount_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subaccount_pkg") subaccount_pkg

### Get instance

Factory for the ``` SUBACCOUNT ``` interface can be accessed from the package subaccount_pkg.

```go
subAccount := subaccount_pkg.NewSUBACCOUNT()
```

### <a name="create_toggle_subaccount_status"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateToggleSubaccountStatus") CreateToggleSubaccountStatus

> Suspend or unsuspend


```go
func (me *SUBACCOUNT_IMPL) CreateToggleSubaccountStatus(
            subAccountSID string,
            mActivate models_pkg.MActivateEnum)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be activated or suspended |
| mActivate |  ``` Required ```  | 0 to suspend or 1 to activate |


#### Example Usage

```go
subAccountSID := "SubAccountSID"
mActivate := models_pkg.mActivate_ENUM_1

var result string
result,_ = subAccount.CreateToggleSubaccountStatus(subAccountSID, mActivate)

```


### <a name="create_delete_subaccount"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateDeleteSubaccount") CreateDeleteSubaccount

> Delete sub account or merge numbers into parent


```go
func (me *SUBACCOUNT_IMPL) CreateDeleteSubaccount(
            subAccountSID string,
            mergeNumber models_pkg.MergeNumberEnum)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be deleted |
| mergeNumber |  ``` Required ```  | 0 to delete or 1 to merge numbers to parent account. |


#### Example Usage

```go
subAccountSID := "SubAccountSID"
mergeNumber := models_pkg.MergeNumber_ENUM_0

var result string
result,_ = subAccount.CreateDeleteSubaccount(subAccountSID, mergeNumber)

```


### <a name="create_subaccount"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateSubaccount") CreateSubaccount

> Create a sub user account under the parent account


```go
func (me *SUBACCOUNT_IMPL) CreateSubaccount(
            firstName string,
            lastName string,
            email string,
            friendlyName string,
            password string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| firstName |  ``` Required ```  | Sub account user first name |
| lastName |  ``` Required ```  | sub account user last name |
| email |  ``` Required ```  | Sub account email address |
| friendlyName |  ``` Required ```  | Descriptive name of the sub account |
| password |  ``` Required ```  | The password of the sub account.  Please make sure to make as password that is at least 8 characters long, contain a symbol, uppercase and a number. |


#### Example Usage

```go
firstName := "FirstName"
lastName := "LastName"
email := "Email"
friendlyName := "FriendlyName"
password := "Password"

var result string
result,_ = subAccount.CreateSubaccount(firstName, lastName, email, friendlyName, password)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="address_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".address_pkg") address_pkg

### Get instance

Factory for the ``` ADDRESS ``` interface can be accessed from the package address_pkg.

```go
address := address_pkg.NewADDRESS()
```

### <a name="create_delete_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateDeleteAddress") CreateDeleteAddress

> To delete Address to your address book


```go
func (me *ADDRESS_IMPL) CreateDeleteAddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be deleted. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.CreateDeleteAddress(addressid)

```


### <a name="create_verify_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateVerifyAddress") CreateVerifyAddress

> Validates an address given.


```go
func (me *ADDRESS_IMPL) CreateVerifyAddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be verified. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.CreateVerifyAddress(addressid)

```


### <a name="create_view_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateViewAddress") CreateViewAddress

> View Address Specific address Book by providing the address id


```go
func (me *ADDRESS_IMPL) CreateViewAddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be retrieved. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.CreateViewAddress(addressid)

```


### <a name="create_list_addresses"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateListAddresses") CreateListAddresses

> List All Address 


```go
func (me *ADDRESS_IMPL) CreateListAddresses(
            page *int64,
            pagesize *int64,
            addressid *string,
            dateCreated *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | How many results to return, default is 10, max is 100, must be an integer |
| addressid |  ``` Optional ```  | addresses Sid |
| dateCreated |  ``` Optional ```  | date created address. |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
addressid := "addressid"
dateCreated := "dateCreated"

var result string
result,_ = address.CreateListAddresses(page, pagesize, addressid, dateCreated)

```


### <a name="create_address"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.CreateAddress") CreateAddress

> To add an address to your address book, you create a new address object. You can retrieve and delete individual addresses as well as get a list of addresses. Addresses are identified by a unique random ID.


```go
func (me *ADDRESS_IMPL) CreateAddress(
            name string,
            address string,
            country string,
            state string,
            city string,
            zip string,
            description *string,
            email *string,
            phone *string)(string,error)
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


#### Example Usage

```go
name := "Name"
address := "Address"
country := "Country"
state := "State"
city := "City"
zip := "Zip"
description := "Description"
email := "email"
phone := "Phone"

var result string
result,_ = address.CreateAddress(name, address, country, state, city, zip, description, email, phone)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="recording_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recording_pkg") recording_pkg

### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

### <a name="create_delete_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateDeleteRecording") CreateDeleteRecording

> Remove a recording from your Ytel account.


```go
func (me *RECORDING_IMPL) CreateDeleteRecording(recordingsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for a recording. |


#### Example Usage

```go
recordingsid := "recordingsid"

var result string
result,_ = recording.CreateDeleteRecording(recordingsid)

```


### <a name="create_view_recording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateViewRecording") CreateViewRecording

> Retrieve the recording of a call by its RecordingSid. This resource will return information regarding the call such as start time, end time, duration, and so forth.


```go
func (me *RECORDING_IMPL) CreateViewRecording(recordingsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for the recording. |


#### Example Usage

```go
recordingsid := "recordingsid"

var result string
result,_ = recording.CreateViewRecording(recordingsid)

```


### <a name="create_list_recordings"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateListRecordings") CreateListRecordings

> Retrieve a list of recording objects.


```go
func (me *RECORDING_IMPL) CreateListRecordings(
            page *int64,
            pagesize *int64,
            datecreated *string,
            callsid *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| datecreated |  ``` Optional ```  | Filter results by creation date |
| callsid |  ``` Optional ```  | The unique identifier for a call. |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
datecreated := "Datecreated"
callsid := "callsid"

var result string
result,_ = recording.CreateListRecordings(page, pagesize, datecreated, callsid)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transcription_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

### <a name="create_transcribe_audio_url"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscribeAudioURL") CreateTranscribeAudioURL

> Transcribe an audio recording from a file.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscribeAudioURL(audiourl string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audiourl |  ``` Required ```  | URL pointing to the location of the audio file that is to be transcribed. |


#### Example Usage

```go
audiourl := "audiourl"

var result string
result,_ = transcription.CreateTranscribeAudioURL(audiourl)

```


### <a name="create_transcribe_recording"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscribeRecording") CreateTranscribeRecording

> Transcribe a recording by its RecordingSid.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscribeRecording(recordingSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | The unique identifier for a recording object. |


#### Example Usage

```go
recordingSid := "recordingSid"

var result string
result,_ = transcription.CreateTranscribeRecording(recordingSid)

```


### <a name="create_view_transcription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateViewTranscription") CreateViewTranscription

> Retrieve information about a transaction by its TranscriptionSid.


```go
func (me *TRANSCRIPTION_IMPL) CreateViewTranscription(transcriptionsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionsid |  ``` Required ```  | The unique identifier for a transcription object. |


#### Example Usage

```go
transcriptionsid := "transcriptionsid"

var result string
result,_ = transcription.CreateViewTranscription(transcriptionsid)

```


### <a name="create_list_transcriptions"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateListTranscriptions") CreateListTranscriptions

> Retrieve a list of transcription objects for your Ytel account.


```go
func (me *TRANSCRIPTION_IMPL) CreateListTranscriptions(
            page *int64,
            pagesize *int64,
            status models_pkg.Status12Enum,
            dateTranscribed *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |
| status |  ``` Optional ```  | The state of the transcription. |
| dateTranscribed |  ``` Optional ```  | The date the transcription took place. |


#### Example Usage

```go
page,_ := strconv.ParseInt("22", 10, 8)
pagesize,_ := strconv.ParseInt("22", 10, 8)
status := models_pkg.status12_INPROGRESS
dateTranscribed := "dateTranscribed"

var result string
result,_ = transcription.CreateListTranscriptions(page, pagesize, status, dateTranscribed)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="usage_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".usage_pkg") usage_pkg

### Get instance

Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.

```go
usage := usage_pkg.NewUSAGE()
```

### <a name="create_list_usage"></a>![Method: ](https://apidocs.io/img/method.png ".usage_pkg.CreateListUsage") CreateListUsage

> Retrieve usage metrics regarding your Ytel account. The results includes inbound/outbound voice calls and inbound/outbound SMS messages as well as carrier lookup requests.


```go
func (me *USAGE_IMPL) CreateListUsage(
            productCode models_pkg.ProductCodeEnum,
            startDate *string,
            endDate *string,
            includeSubAccounts *string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| productCode |  ``` Optional ```  | Filter usage results by product type. |
| startDate |  ``` Optional ```  | Filter usage objects by start date. |
| endDate |  ``` Optional ```  | Filter usage objects by end date. |
| includeSubAccounts |  ``` Optional ```  | Will include all subaccount usage data |


#### Example Usage

```go
productCode := models_pkg.ProductCode_ENUM_0
startDate := "startDate"
endDate := "endDate"
includeSubAccounts := "IncludeSubAccounts"

var result string
result,_ = usage.CreateListUsage(productCode, startDate, endDate, includeSubAccounts)

```


[Back to List of Controllers](#list_of_controllers)



