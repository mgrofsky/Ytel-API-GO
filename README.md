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

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=Ytel%20API-GoLang&projectName=ytelapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=ytelapi_lib)

## How to Use

The following section explains how to use the YtelapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=ytelapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=ytelapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=ytelapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=ytelapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=ytelapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Ytel%20API-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=ytelapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=ytelapi_lib)

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

### <a name="create_dedicatedshortcode_listshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateDedicatedshortcodeListshortcode") CreateDedicatedshortcodeListshortcode

> Retrieve a list of Short Code assignment associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) CreateDedicatedshortcodeListshortcode(input *CreateDedicatedshortcodeListshortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Optional ```  | Only list Short Code Assignment sent from this Short Code |
| page |  ``` Optional ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (shortcode_pkg.CreateDedicatedshortcodeListshortcodeInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

page := "page"
collect.Page = page

pagesize := "pagesize"
collect.Pagesize = pagesize


var result string
result,_ = shortCode.CreateDedicatedshortcodeListshortcode(collect)

```


### <a name="create_dedicatedshortcode_updateshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateDedicatedshortcodeUpdateshortcode") CreateDedicatedshortcodeUpdateshortcode

> Update a dedicated shortcode.


```go
func (me *SHORTCODE_IMPL) CreateDedicatedshortcodeUpdateshortcode(input *CreateDedicatedshortcodeUpdateshortcodeInput)(string,error)
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
collect := new (shortcode_pkg.CreateDedicatedshortcodeUpdateshortcodeInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

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
result,_ = shortCode.CreateDedicatedshortcodeUpdateshortcode(collect)

```


### <a name="create_dedicatedshortcode_viewshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateDedicatedshortcodeViewshortcode") CreateDedicatedshortcodeViewshortcode

> Retrieve a single Short Code object by its shortcode assignment.


```go
func (me *SHORTCODE_IMPL) CreateDedicatedshortcodeViewshortcode(shortcode string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Dedicated Short Code to your Ytel account |


#### Example Usage

```go
shortcode := "Shortcode"

var result string
result,_ = shortCode.CreateDedicatedshortcodeViewshortcode(shortcode)

```


### <a name="create_shortcode_viewsms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateShortcodeViewsms") CreateShortcodeViewsms

> View a single Sms Short Code message.


```go
func (me *SHORTCODE_IMPL) CreateShortcodeViewsms(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for the sms resource |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = shortCode.CreateShortcodeViewsms(messageSid)

```


### <a name="create_dedicatedshortcode_getinboundsms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateDedicatedshortcodeGetinboundsms") CreateDedicatedshortcodeGetinboundsms

> Retrive a list of inbound Sms Short Code messages associated with your Ytel account.


```go
func (me *SHORTCODE_IMPL) CreateDedicatedshortcodeGetinboundsms(input *CreateDedicatedshortcodeGetinboundsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Only list SMS messages sent from this number |
| shortcode |  ``` Optional ```  | Only list SMS messages sent to Shortcode |
| datecreated |  ``` Optional ```  | Only list SMS messages sent in the specified date MAKE REQUEST |


#### Example Usage

```go
collect := new (shortcode_pkg.CreateDedicatedshortcodeGetinboundsmsInput)

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
result,_ = shortCode.CreateDedicatedshortcodeGetinboundsms(collect)

```


### <a name="create_dedicatedshortcode_sendsms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateDedicatedshortcodeSendsms") CreateDedicatedshortcodeSendsms

> Send Dedicated Shortcode


```go
func (me *SHORTCODE_IMPL) CreateDedicatedshortcodeSendsms(input *CreateDedicatedshortcodeSendsmsInput)(string,error)
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
collect := new (shortcode_pkg.CreateDedicatedshortcodeSendsmsInput)

shortcode,_ := strconv.ParseInt("206", 10, 8)
collect.Shortcode = shortcode

to := 206.488525940333
collect.To = to

body := "body"
collect.Body = body

method := "method"
collect.Method = method

messagestatuscallback := "messagestatuscallback"
collect.Messagestatuscallback = messagestatuscallback


var result string
result,_ = shortCode.CreateDedicatedshortcodeSendsms(collect)

```


### <a name="create_shortcode_listsms"></a>![Method: ](https://apidocs.io/img/method.png ".shortcode_pkg.CreateShortcodeListsms") CreateShortcodeListsms

> Retrieve a list of Short Code messages.


```go
func (me *SHORTCODE_IMPL) CreateShortcodeListsms(input *CreateShortcodeListsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Optional ```  | Only list messages sent from this Short Code |
| to |  ``` Optional ```  | Only list messages sent to this number |
| dateSent |  ``` Optional ```  | Only list messages sent with the specified date |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (shortcode_pkg.CreateShortcodeListsmsInput)

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
result,_ = shortCode.CreateShortcodeListsms(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="areamail_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".areamail_pkg") areamail_pkg

### Get instance

Factory for the ``` AREAMAIL ``` interface can be accessed from the package areamail_pkg.

```go
areaMail := areamail_pkg.NewAREAMAIL()
```

### <a name="create_areamail_delete"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreamailDelete") CreateAreamailDelete

> Remove an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) CreateAreamailDelete(areamailid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |


#### Example Usage

```go
areamailid := "areamailid"

var result string
result,_ = areaMail.CreateAreamailDelete(areamailid)

```


### <a name="create_areamail_create"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreamailCreate") CreateAreamailCreate

> Create a new AreaMail object.


```go
func (me *AREAMAIL_IMPL) CreateAreamailCreate(input *CreateAreamailCreateInput)(string,error)
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
collect := new (areamail_pkg.CreateAreamailCreateInput)

routes := "routes"
collect.Routes = routes

attachbyid := "attachbyid"
collect.Attachbyid = attachbyid

front := "front"
collect.Front = front

back := "back"
collect.Back = back

description := "description"
collect.Description = description

targettype := "targettype"
collect.Targettype = targettype

htmldata := "htmldata"
collect.Htmldata = htmldata


var result string
result,_ = areaMail.CreateAreamailCreate(collect)

```


### <a name="create_areamail_view"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreamailView") CreateAreamailView

> Retrieve an AreaMail object by its AreaMailId.


```go
func (me *AREAMAIL_IMPL) CreateAreamailView(areamailid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| areamailid |  ``` Required ```  | The unique identifier for an AreaMail object. |


#### Example Usage

```go
areamailid := "areamailid"

var result string
result,_ = areaMail.CreateAreamailView(areamailid)

```


### <a name="create_areamail_lists"></a>![Method: ](https://apidocs.io/img/method.png ".areamail_pkg.CreateAreamailLists") CreateAreamailLists

> Retrieve a list of AreaMail objects.


```go
func (me *AREAMAIL_IMPL) CreateAreamailLists(input *CreateAreamailListsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| areamailsid |  ``` Optional ```  | The unique identifier for an AreaMail object. |
| dateCreated |  ``` Optional ```  | The date the AreaMail was created. |


#### Example Usage

```go
collect := new (areamail_pkg.CreateAreamailListsInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

areamailsid := "areamailsid"
collect.Areamailsid = areamailsid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = areaMail.CreateAreamailLists(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="postcard_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".postcard_pkg") postcard_pkg

### Get instance

Factory for the ``` POSTCARD ``` interface can be accessed from the package postcard_pkg.

```go
postCard := postcard_pkg.NewPOSTCARD()
```

### <a name="postcard_deletepostcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.PostcardDeletepostcard") PostcardDeletepostcard

> Remove a postcard object.


```go
func (me *POSTCARD_IMPL) PostcardDeletepostcard(postcardid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier of a postcard object. |


#### Example Usage

```go
postcardid := "postcardid"

var result string
result,_ = postCard.PostcardDeletepostcard(postcardid)

```


### <a name="postcard_viewpostcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.PostcardViewpostcard") PostcardViewpostcard

> Retrieve a postcard object by its PostcardId.


```go
func (me *POSTCARD_IMPL) PostcardViewpostcard(postcardid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| postcardid |  ``` Required ```  | The unique identifier for a postcard object. |


#### Example Usage

```go
postcardid := "postcardid"

var result string
result,_ = postCard.PostcardViewpostcard(postcardid)

```


### <a name="postcard_listpostcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.PostcardListpostcard") PostcardListpostcard

> Retrieve a list of postcard objects. The postcards objects are sorted by creation date, with the most recently created postcards appearing first.


```go
func (me *POSTCARD_IMPL) PostcardListpostcard(input *PostcardListpostcardInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| postcardid |  ``` Optional ```  | The unique identifier for a postcard object. |
| dateCreated |  ``` Optional ```  | The date the postcard was created. |


#### Example Usage

```go
collect := new (postcard_pkg.PostcardListpostcardInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

postcardid := "postcardid"
collect.Postcardid = postcardid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = postCard.PostcardListpostcard(collect)

```


### <a name="postcard_createpostcard"></a>![Method: ](https://apidocs.io/img/method.png ".postcard_pkg.PostcardCreatepostcard") PostcardCreatepostcard

> Create, print, and mail a postcard to an address. The postcard front must be supplied as a PDF, image, or an HTML string. The back can be a PDF, image, HTML string, or it can be automatically generated by supplying a custom message.


```go
func (me *POSTCARD_IMPL) PostcardCreatepostcard(input *PostcardCreatepostcardInput)(string,error)
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
collect := new (postcard_pkg.PostcardCreatepostcardInput)

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

description := "description"
collect.Description = description

htmldata := "htmldata"
collect.Htmldata = htmldata


var result string
result,_ = postCard.PostcardCreatepostcard(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="letter_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".letter_pkg") letter_pkg

### Get instance

Factory for the ``` LETTER ``` interface can be accessed from the package letter_pkg.

```go
letter := letter_pkg.NewLETTER()
```

### <a name="create_letter_delete"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetterDelete") CreateLetterDelete

> Remove a letter object by its LetterId.


```go
func (me *LETTER_IMPL) CreateLetterDelete(lettersid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |


#### Example Usage

```go
lettersid := "lettersid"

var result string
result,_ = letter.CreateLetterDelete(lettersid)

```


### <a name="create_letter_viewletter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetterViewletter") CreateLetterViewletter

> Retrieve a letter object by its LetterSid.


```go
func (me *LETTER_IMPL) CreateLetterViewletter(lettersid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| lettersid |  ``` Required ```  | The unique identifier for a letter object. |


#### Example Usage

```go
lettersid := "lettersid"

var result string
result,_ = letter.CreateLetterViewletter(lettersid)

```


### <a name="create_letter_listsletter"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetterListsletter") CreateLetterListsletter

> Retrieve a list of letter objects. The letter objects are sorted by creation date, with the most recently created letters appearing first.


```go
func (me *LETTER_IMPL) CreateLetterListsletter(input *CreateLetterListsletterInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| lettersid |  ``` Optional ```  | The unique identifier for a letter object. |
| dateCreated |  ``` Optional ```  | The date the letter was created. |


#### Example Usage

```go
collect := new (letter_pkg.CreateLetterListsletterInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

lettersid := "lettersid"
collect.Lettersid = lettersid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = letter.CreateLetterListsletter(collect)

```


### <a name="create_letter_create"></a>![Method: ](https://apidocs.io/img/method.png ".letter_pkg.CreateLetterCreate") CreateLetterCreate

> Create, print, and mail a letter to an address. The letter file must be supplied as a PDF or an HTML string.


```go
func (me *LETTER_IMPL) CreateLetterCreate(input *CreateLetterCreateInput)(string,error)
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
collect := new (letter_pkg.CreateLetterCreateInput)

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
result,_ = letter.CreateLetterCreate(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="call_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".call_pkg") call_pkg

### Get instance

Factory for the ``` CALL ``` interface can be accessed from the package call_pkg.

```go
call := call_pkg.NewCALL()
```

### <a name="create_calls_viewcalldetail"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsViewcalldetail") CreateCallsViewcalldetail

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) CreateCallsViewcalldetail(callSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier for the voice call. |


#### Example Usage

```go
callSid := "callSid"

var result string
result,_ = call.CreateCallsViewcalldetail(callSid)

```


### <a name="create_calls_viewcalls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsViewcalls") CreateCallsViewcalls

> Retrieve a single voice call’s information by its CallSid.


```go
func (me *CALL_IMPL) CreateCallsViewcalls(callsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callsid |  ``` Required ```  | The unique identifier for the voice call. |


#### Example Usage

```go
callsid := "callsid"

var result string
result,_ = call.CreateCallsViewcalls(callsid)

```


### <a name="create_calls_senddigits"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsSenddigits") CreateCallsSenddigits

> Play Dtmf and send the Digit


```go
func (me *CALL_IMPL) CreateCallsSenddigits(input *CreateCallsSenddigitsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| callSid |  ``` Required ```  | The unique identifier of each call resource |
| playDtmf |  ``` Required ```  | DTMF digits to play to the call. 0-9, #, *, W, or w |
| playDtmfDirection |  ``` Optional ```  | The leg of the call DTMF digits should be sent to |


#### Example Usage

```go
collect := new (call_pkg.CreateCallsSenddigitsInput)

callSid := "CallSid"
collect.CallSid = callSid

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

playDtmfDirection := models_pkg.PlayDtmfDirection_IN
collect.PlayDtmfDirection = playDtmfDirection


var result string
result,_ = call.CreateCallsSenddigits(collect)

```


### <a name="create_calls_makervm"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsMakervm") CreateCallsMakervm

> Initiate an outbound Ringless Voicemail through Ytel.


```go
func (me *CALL_IMPL) CreateCallsMakervm(input *CreateCallsMakervmInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid Ytel Voice enabled number (E.164 format) that will be initiating the phone call. |
| rVMCallerId |  ``` Required ```  | A required secondary Caller ID for RVM to work. |
| to |  ``` Required ```  | A valid number (E.164 format) that will receive the phone call. |
| voiceMailURL |  ``` Required ```  | The URL requested once the RVM connects. A set of default parameters will be sent here. |
| method |  ``` Optional ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once call connects. |
| statusCallBackUrl |  ``` Optional ```  | URL that can be requested to receive notification when call has ended. A set of default parameters will be sent here once the call is finished. |
| statsCallBackMethod |  ``` Optional ```  | Specifies the HTTP method used to request the required StatusCallBackUrl once call connects. |


#### Example Usage

```go
collect := new (call_pkg.CreateCallsMakervmInput)

from := "From"
collect.From = from

rVMCallerId := "RVMCallerId"
collect.RVMCallerId = rVMCallerId

to := "To"
collect.To = to

voiceMailURL := "VoiceMailURL"
collect.VoiceMailURL = voiceMailURL

method := "GET"
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statsCallBackMethod := "StatsCallBackMethod"
collect.StatsCallBackMethod = statsCallBackMethod


var result string
result,_ = call.CreateCallsMakervm(collect)

```


### <a name="create_calls_listcalls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsListcalls") CreateCallsListcalls

> A list of calls associated with your Ytel account


```go
func (me *CALL_IMPL) CreateCallsListcalls(input *CreateCallsListcallsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| to |  ``` Optional ```  | Filter calls that were sent to this 10-digit number (E.164 format). |
| from |  ``` Optional ```  | Filter calls that were sent from this 10-digit number (E.164 format). |
| dateCreated |  ``` Optional ```  | Return calls that are from a specified date. |


#### Example Usage

```go
collect := new (call_pkg.CreateCallsListcallsInput)

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
result,_ = call.CreateCallsListcalls(collect)

```


### <a name="create_calls_interruptcalls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsInterruptcalls") CreateCallsInterruptcalls

> Interrupt the Call by Call Sid


```go
func (me *CALL_IMPL) CreateCallsInterruptcalls(input *CreateCallsInterruptcallsInput)(string,error)
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
collect := new (call_pkg.CreateCallsInterruptcallsInput)

callSid := "CallSid"
collect.CallSid = callSid

url := "Url"
collect.Url = url

method := "Method"
collect.Method = method

status := models_pkg.Status24_CANCELED
collect.Status = status


var result string
result,_ = call.CreateCallsInterruptcalls(collect)

```


### <a name="create_calls_recordcalls"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsRecordcalls") CreateCallsRecordcalls

> Start or stop recording of an in-progress voice call.


```go
func (me *CALL_IMPL) CreateCallsRecordcalls(input *CreateCallsRecordcallsInput)(string,error)
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
collect := new (call_pkg.CreateCallsRecordcallsInput)

callSid := "CallSid"
collect.CallSid = callSid

record := true
collect.Record = record

direction := models_pkg.Direction_IN
collect.Direction = direction

timeLimit,_ := strconv.ParseInt("206", 10, 8)
collect.TimeLimit = timeLimit

callBackUrl := "CallBackUrl"
collect.CallBackUrl = callBackUrl

fileformat := models_pkg.Fileformat_MP3
collect.Fileformat = fileformat


var result string
result,_ = call.CreateCallsRecordcalls(collect)

```


### <a name="create_calls_playaudios"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsPlayaudios") CreateCallsPlayaudios

> Play Audio from a url


```go
func (me *CALL_IMPL) CreateCallsPlayaudios(input *CreateCallsPlayaudiosInput)(string,error)
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
collect := new (call_pkg.CreateCallsPlayaudiosInput)

callSid := "CallSid"
collect.CallSid = callSid

audioUrl := "AudioUrl"
collect.AudioUrl = audioUrl

sayText := "SayText"
collect.SayText = sayText

length,_ := strconv.ParseInt("206", 10, 8)
collect.Length = length

direction := models_pkg.Direction_IN
collect.Direction = direction

mix := true
collect.Mix = mix


var result string
result,_ = call.CreateCallsPlayaudios(collect)

```


### <a name="create_calls_voiceeffect"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsVoiceeffect") CreateCallsVoiceeffect

> Add audio voice effects to the an in-progress voice call.


```go
func (me *CALL_IMPL) CreateCallsVoiceeffect(input *CreateCallsVoiceeffectInput)(string,error)
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
collect := new (call_pkg.CreateCallsVoiceeffectInput)

callSid := "CallSid"
collect.CallSid = callSid

audioDirection := models_pkg.AudioDirection_IN
collect.AudioDirection = audioDirection

pitchSemiTones := 206.488525940333
collect.PitchSemiTones = pitchSemiTones

pitchOctaves := 206.488525940333
collect.PitchOctaves = pitchOctaves

pitch := 206.488525940333
collect.Pitch = pitch

rate := 206.488525940333
collect.Rate = rate

tempo := 206.488525940333
collect.Tempo = tempo


var result string
result,_ = call.CreateCallsVoiceeffect(collect)

```


### <a name="create_calls_groupcall"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsGroupcall") CreateCallsGroupcall

> Group Call


```go
func (me *CALL_IMPL) CreateCallsGroupcall(input *CreateCallsGroupcallInput)(string,error)
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
collect := new (call_pkg.CreateCallsGroupcallInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

groupConfirmKey := "GroupConfirmKey"
collect.GroupConfirmKey = groupConfirmKey

groupConfirmFile := models_pkg.GroupConfirmFile_MP3
collect.GroupConfirmFile = groupConfirmFile

method := "Method"
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := "StatusCallBackMethod"
collect.StatusCallBackMethod = statusCallBackMethod

fallBackUrl := "FallBackUrl"
collect.FallBackUrl = fallBackUrl

fallBackMethod := "FallBackMethod"
collect.FallBackMethod = fallBackMethod

heartBeatUrl := "HeartBeatUrl"
collect.HeartBeatUrl = heartBeatUrl

heartBeatMethod := "HeartBeatMethod"
collect.HeartBeatMethod = heartBeatMethod

timeout,_ := strconv.ParseInt("206", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := "HideCallerId"
collect.HideCallerId = hideCallerId

record := true
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := "RecordCallBackMethod"
collect.RecordCallBackMethod = recordCallBackMethod

transcribe := true
collect.Transcribe = transcribe

transcribeCallBackUrl := "TranscribeCallBackUrl"
collect.TranscribeCallBackUrl = transcribeCallBackUrl


var result string
result,_ = call.CreateCallsGroupcall(collect)

```


### <a name="create_calls_makecall"></a>![Method: ](https://apidocs.io/img/method.png ".call_pkg.CreateCallsMakecall") CreateCallsMakecall

> You can experiment with initiating a call through Ytel and view the request response generated when doing so and get the response in json


```go
func (me *CALL_IMPL) CreateCallsMakecall(input *CreateCallsMakecallInput)(string,error)
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
collect := new (call_pkg.CreateCallsMakecallInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

method := "Method"
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := "StatusCallBackMethod"
collect.StatusCallBackMethod = statusCallBackMethod

fallBackUrl := "FallBackUrl"
collect.FallBackUrl = fallBackUrl

fallBackMethod := "FallBackMethod"
collect.FallBackMethod = fallBackMethod

heartBeatUrl := "HeartBeatUrl"
collect.HeartBeatUrl = heartBeatUrl

heartBeatMethod := "HeartBeatMethod"
collect.HeartBeatMethod = heartBeatMethod

timeout,_ := strconv.ParseInt("206", 10, 8)
collect.Timeout = timeout

playDtmf := "PlayDtmf"
collect.PlayDtmf = playDtmf

hideCallerId := true
collect.HideCallerId = hideCallerId

record := true
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := "RecordCallBackMethod"
collect.RecordCallBackMethod = recordCallBackMethod

transcribe := true
collect.Transcribe = transcribe

transcribeCallBackUrl := "TranscribeCallBackUrl"
collect.TranscribeCallBackUrl = transcribeCallBackUrl

ifMachine := models_pkg.ifMachine_CONTINUE
collect.IfMachine = ifMachine

ifMachineUrl := "IfMachineUrl"
collect.IfMachineUrl = ifMachineUrl

ifMachineMethod := "IfMachineMethod"
collect.IfMachineMethod = ifMachineMethod

feedback := true
collect.Feedback = feedback

surveyId := "SurveyId"
collect.SurveyId = surveyId


var result string
result,_ = call.CreateCallsMakecall(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="phonenumber_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".phonenumber_pkg") phonenumber_pkg

### Get instance

Factory for the ``` PHONENUMBER ``` interface can be accessed from the package phonenumber_pkg.

```go
phoneNumber := phonenumber_pkg.NewPHONENUMBER()
```

### <a name="create_incomingphone_getdidscore"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneGetdidscore") CreateIncomingphoneGetdidscore

> Get DID Score Number


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneGetdidscore(phonenumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | Specifies the multiple phone numbers for check updated spamscore . |


#### Example Usage

```go
phonenumber := "Phonenumber"

var result string
result,_ = phoneNumber.CreateIncomingphoneGetdidscore(phonenumber)

```


### <a name="create_incomingphone_bulkbuy"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneBulkbuy") CreateIncomingphoneBulkbuy

> Purchase a selected number of DID's from specific area codes to be used with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneBulkbuy(input *CreateIncomingphoneBulkbuyInput)(string,error)
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
collect := new (phonenumber_pkg.CreateIncomingphoneBulkbuyInput)

numberType := models_pkg.numbertype16_ALL
collect.NumberType = numberType

areaCode := "AreaCode"
collect.AreaCode = areaCode

quantity := "Quantity"
collect.Quantity = quantity

leftover := "Leftover"
collect.Leftover = leftover


var result string
result,_ = phoneNumber.CreateIncomingphoneBulkbuy(collect)

```


### <a name="create_incomingphone_listnumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneListnumber") CreateIncomingphoneListnumber

> Retrieve a list of purchased phones numbers associated with your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneListnumber(input *CreateIncomingphoneListnumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | Which page of the overall response will be returned. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| numberType |  ``` Optional ```  | The capability supported by the number.Number type either SMS,Voice or all |
| friendlyName |  ``` Optional ```  | A human-readable label added to the number object. |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateIncomingphoneListnumberInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize

numberType := models_pkg.numbertype16_ALL
collect.NumberType = numberType

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName


var result string
result,_ = phoneNumber.CreateIncomingphoneListnumber(collect)

```


### <a name="create_incomingphone_buynumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneBuynumber") CreateIncomingphoneBuynumber

> Purchase a phone number to be used with your Ytel account


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneBuynumber(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateIncomingphoneBuynumber(phoneNumber)

```


### <a name="create_incomingphone_releasenumber_by_response_type_post"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneReleasenumberByResponseTypePost") CreateIncomingphoneReleasenumberByResponseTypePost

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneReleasenumberByResponseTypePost(input *CreateIncomingphoneReleasenumberByResponseTypePostInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
| responseType |  ``` Required ```  | Response type format xml or json |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateIncomingphoneReleasenumberByResponseTypePostInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

responseType := "ResponseType"
collect.ResponseType = responseType


var result string
result,_ = phoneNumber.CreateIncomingphoneReleasenumberByResponseTypePost(collect)

```


### <a name="create_incomingphone_viewnumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneViewnumber") CreateIncomingphoneViewnumber

> Retrieve the details for a phone number by its number.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneViewnumber(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel 10-digit phone number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateIncomingphoneViewnumber(phoneNumber)

```


### <a name="create_incomingphone_transferphonenumbers"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneTransferphonenumbers") CreateIncomingphoneTransferphonenumbers

> Transfer phone number that has been purchased for from one account to another account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneTransferphonenumbers(input *CreateIncomingphoneTransferphonenumbersInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phonenumber |  ``` Required ```  | A valid 10-digit Ytel number (E.164 format). |
| fromaccountsid |  ``` Required ```  | A specific Accountsid from where Number is getting transfer. |
| toaccountsid |  ``` Required ```  | A specific Accountsid to which Number is getting transfer. |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateIncomingphoneTransferphonenumbersInput)

phonenumber := "phonenumber"
collect.Phonenumber = phonenumber

fromaccountsid := "fromaccountsid"
collect.Fromaccountsid = fromaccountsid

toaccountsid := "toaccountsid"
collect.Toaccountsid = toaccountsid


var result string
result,_ = phoneNumber.CreateIncomingphoneTransferphonenumbers(collect)

```


### <a name="create_incomingphone_massreleasenumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneMassreleasenumber") CreateIncomingphoneMassreleasenumber

> Remove a purchased Ytel number from your account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneMassreleasenumber(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid Ytel comma separated numbers (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = phoneNumber.CreateIncomingphoneMassreleasenumber(phoneNumber)

```


### <a name="create_incomingphone_massupdatenumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneMassupdatenumber") CreateIncomingphoneMassupdatenumber

> Update properties for a Ytel numbers that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneMassupdatenumber(input *CreateIncomingphoneMassupdatenumberInput)(string,error)
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
collect := new (phonenumber_pkg.CreateIncomingphoneMassupdatenumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

voiceUrl := "VoiceUrl"
collect.VoiceUrl = voiceUrl

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

voiceMethod := "VoiceMethod"
collect.VoiceMethod = voiceMethod

voiceFallbackUrl := "VoiceFallbackUrl"
collect.VoiceFallbackUrl = voiceFallbackUrl

voiceFallbackMethod := "VoiceFallbackMethod"
collect.VoiceFallbackMethod = voiceFallbackMethod

hangupCallback := "HangupCallback"
collect.HangupCallback = hangupCallback

hangupCallbackMethod := "HangupCallbackMethod"
collect.HangupCallbackMethod = hangupCallbackMethod

heartbeatUrl := "HeartbeatUrl"
collect.HeartbeatUrl = heartbeatUrl

heartbeatMethod := "HeartbeatMethod"
collect.HeartbeatMethod = heartbeatMethod

smsUrl := "SmsUrl"
collect.SmsUrl = smsUrl

smsMethod := "SmsMethod"
collect.SmsMethod = smsMethod

smsFallbackUrl := "SmsFallbackUrl"
collect.SmsFallbackUrl = smsFallbackUrl

smsFallbackMethod := "SmsFallbackMethod"
collect.SmsFallbackMethod = smsFallbackMethod


var result string
result,_ = phoneNumber.CreateIncomingphoneMassupdatenumber(collect)

```


### <a name="create_incomingphone_updatenumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneUpdatenumber") CreateIncomingphoneUpdatenumber

> Update properties for a Ytel number that has been purchased for your account. Refer to the parameters list for the list of properties that can be updated.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneUpdatenumber(input *CreateIncomingphoneUpdatenumberInput)(string,error)
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
collect := new (phonenumber_pkg.CreateIncomingphoneUpdatenumberInput)

phoneNumber := "PhoneNumber"
collect.PhoneNumber = phoneNumber

voiceUrl := "VoiceUrl"
collect.VoiceUrl = voiceUrl

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

voiceMethod := "VoiceMethod"
collect.VoiceMethod = voiceMethod

voiceFallbackUrl := "VoiceFallbackUrl"
collect.VoiceFallbackUrl = voiceFallbackUrl

voiceFallbackMethod := "VoiceFallbackMethod"
collect.VoiceFallbackMethod = voiceFallbackMethod

hangupCallback := "HangupCallback"
collect.HangupCallback = hangupCallback

hangupCallbackMethod := "HangupCallbackMethod"
collect.HangupCallbackMethod = hangupCallbackMethod

heartbeatUrl := "HeartbeatUrl"
collect.HeartbeatUrl = heartbeatUrl

heartbeatMethod := "HeartbeatMethod"
collect.HeartbeatMethod = heartbeatMethod

smsUrl := "SmsUrl"
collect.SmsUrl = smsUrl

smsMethod := "SmsMethod"
collect.SmsMethod = smsMethod

smsFallbackUrl := "SmsFallbackUrl"
collect.SmsFallbackUrl = smsFallbackUrl

smsFallbackMethod := "SmsFallbackMethod"
collect.SmsFallbackMethod = smsFallbackMethod


var result string
result,_ = phoneNumber.CreateIncomingphoneUpdatenumber(collect)

```


### <a name="create_incomingphone_availablenumber"></a>![Method: ](https://apidocs.io/img/method.png ".phonenumber_pkg.CreateIncomingphoneAvailablenumber") CreateIncomingphoneAvailablenumber

> Retrieve a list of available phone numbers that can be purchased and used for your Ytel account.


```go
func (me *PHONENUMBER_IMPL) CreateIncomingphoneAvailablenumber(input *CreateIncomingphoneAvailablenumberInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| numbertype |  ``` Required ```  | Number type either SMS,Voice or all |
| areacode |  ``` Required ```  | Specifies the area code for the returned list of available numbers. Only available for North American numbers. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return. |


#### Example Usage

```go
collect := new (phonenumber_pkg.CreateIncomingphoneAvailablenumberInput)

numbertype := models_pkg.numbertype16_ALL
collect.Numbertype = numbertype

areacode := "areacode"
collect.Areacode = areacode

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize


var result string
result,_ = phoneNumber.CreateIncomingphoneAvailablenumber(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sms_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sms_pkg") sms_pkg

### Get instance

Factory for the ``` SMS ``` interface can be accessed from the package sms_pkg.

```go
sMS := sms_pkg.NewSMS()
```

### <a name="create_sms_viewdetailsms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSmsViewdetailsms") CreateSmsViewdetailsms

> Retrieve a single SMS message object with details by its SmsSid.


```go
func (me *SMS_IMPL) CreateSmsViewdetailsms(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = sMS.CreateSmsViewdetailsms(messageSid)

```


### <a name="create_sms_viewsms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSmsViewsms") CreateSmsViewsms

> Retrieve a single SMS message object by its SmsSid.


```go
func (me *SMS_IMPL) CreateSmsViewsms(messageSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| messageSid |  ``` Required ```  | The unique identifier for a sms message. |


#### Example Usage

```go
messageSid := "MessageSid"

var result string
result,_ = sMS.CreateSmsViewsms(messageSid)

```


### <a name="create_sms_getinboundsms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSmsGetinboundsms") CreateSmsGetinboundsms

> Retrieve a list of Inbound SMS message objects.


```go
func (me *SMS_IMPL) CreateSmsGetinboundsms(input *CreateSmsGetinboundsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Filter sms message objects by this date. |


#### Example Usage

```go
collect := new (sms_pkg.CreateSmsGetinboundsmsInput)

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
result,_ = sMS.CreateSmsGetinboundsms(collect)

```


### <a name="create_sms_listsms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSmsListsms") CreateSmsListsms

> Retrieve a list of Outbound SMS message objects.


```go
func (me *SMS_IMPL) CreateSmsListsms(input *CreateSmsListsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | Filter SMS message objects from this valid 10-digit phone number (E.164 format). |
| to |  ``` Optional ```  | Filter SMS message objects to this valid 10-digit phone number (E.164 format). |
| dateSent |  ``` Optional ```  | Only list SMS messages sent in the specified date range |


#### Example Usage

```go
collect := new (sms_pkg.CreateSmsListsmsInput)

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
result,_ = sMS.CreateSmsListsms(collect)

```


### <a name="create_sms_sendsms"></a>![Method: ](https://apidocs.io/img/method.png ".sms_pkg.CreateSmsSendsms") CreateSmsSendsms

> Send an SMS from a Ytel number


```go
func (me *SMS_IMPL) CreateSmsSendsms(input *CreateSmsSendsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | The 10-digit SMS-enabled Ytel number (E.164 format) in which the message is sent. |
| to |  ``` Required ```  | The 10-digit phone number (E.164 format) that will receive the message. |
| body |  ``` Required ```  | The body message that is to be sent in the text. |
| method |  ``` Optional ```  | Specifies the HTTP method used to request the required URL once SMS sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when SMS has Sent. A set of default parameters will be sent here once the SMS is finished. |
| smartsms |  ``` Optional ```  ``` DefaultValue ```  | Check's 'to' number can receive sms or not using Carrier API, if wireless = true then text sms is sent, else wireless = false then call is recieved to end user with audible message. |
| deliveryStatus |  ``` Optional ```  ``` DefaultValue ```  | Delivery reports are a method to tell your system if the message has arrived on the destination phone. |


#### Example Usage

```go
collect := new (sms_pkg.CreateSmsSendsmsInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

body := "Body"
collect.Body = body

method := "Method"
collect.Method = method

messageStatusCallback := "MessageStatusCallback"
collect.MessageStatusCallback = messageStatusCallback

smartsms := false
collect.Smartsms = smartsms

deliveryStatus := false
collect.DeliveryStatus = deliveryStatus


var result string
result,_ = sMS.CreateSmsSendsms(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="sharedshortcode_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".sharedshortcode_pkg") sharedshortcode_pkg

### Get instance

Factory for the ``` SHAREDSHORTCODE ``` interface can be accessed from the package sharedshortcode_pkg.

```go
sharedShortCode := sharedshortcode_pkg.NewSHAREDSHORTCODE()
```

### <a name="create_shortcode_viewshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateShortcodeViewshortcode") CreateShortcodeViewshortcode

> The response returned here contains all resource properties associated with the given Shortcode.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeViewshortcode(shortcode string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | List of valid Shortcode to your Ytel account |


#### Example Usage

```go
shortcode := "Shortcode"

var result string
result,_ = sharedShortCode.CreateShortcodeViewshortcode(shortcode)

```


### <a name="create_keyword_view"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateKeywordView") CreateKeywordView

> View a set of properties for a single keyword.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateKeywordView(keywordid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keywordid |  ``` Required ```  | The unique identifier of each keyword |


#### Example Usage

```go
keywordid := "Keywordid"

var result string
result,_ = sharedShortCode.CreateKeywordView(keywordid)

```


### <a name="create_shortcode_updateshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateShortcodeUpdateshortcode") CreateShortcodeUpdateshortcode

> Update Assignment


```go
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeUpdateshortcode(input *CreateShortcodeUpdateshortcodeInput)(string,error)
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
collect := new (sharedshortcode_pkg.CreateShortcodeUpdateshortcodeInput)

shortcode := "Shortcode"
collect.Shortcode = shortcode

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

callbackUrl := "CallbackUrl"
collect.CallbackUrl = callbackUrl

callbackMethod := "CallbackMethod"
collect.CallbackMethod = callbackMethod

fallbackUrl := "FallbackUrl"
collect.FallbackUrl = fallbackUrl

fallbackUrlMethod := "FallbackUrlMethod"
collect.FallbackUrlMethod = fallbackUrlMethod


var result string
result,_ = sharedShortCode.CreateShortcodeUpdateshortcode(collect)

```


### <a name="create_template_view"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateTemplateView") CreateTemplateView

> View a Shared ShortCode Template


```go
func (me *SHAREDSHORTCODE_IMPL) CreateTemplateView(templateId uuid.UUID)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| templateId |  ``` Required ```  | The unique identifier for a template object |


#### Example Usage

```go
templateId := uuid.NewV4()

var result string
result,_ = sharedShortCode.CreateTemplateView(templateId)

```


### <a name="create_shortcode_listshortcode"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateShortcodeListshortcode") CreateShortcodeListshortcode

> Retrieve a list of shortcode assignment associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeListshortcode(input *CreateShortcodeListshortcodeInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.CreateShortcodeListshortcodeInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

shortcode := "Shortcode"
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.CreateShortcodeListshortcode(collect)

```


### <a name="create_keyword_lists"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateKeywordLists") CreateKeywordLists

> Retrieve a list of keywords associated with your Ytel account.


```go
func (me *SHAREDSHORTCODE_IMPL) CreateKeywordLists(input *CreateKeywordListsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| keyword |  ``` Optional ```  | Only list keywords of keyword |
| shortcode |  ``` Optional ```  | Only list keywords of shortcode |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.CreateKeywordListsInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

keyword := "Keyword"
collect.Keyword = keyword

shortcode,_ := strconv.ParseInt("42", 10, 8)
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.CreateKeywordLists(collect)

```


### <a name="create_template_lists"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateTemplateLists") CreateTemplateLists

> List Shortcode Templates by Type


```go
func (me *SHAREDSHORTCODE_IMPL) CreateTemplateLists(input *CreateTemplateListsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| mtype |  ``` Optional ```  ``` DefaultValue ```  | The type (category) of template Valid values: marketing, authorization |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| shortcode |  ``` Optional ```  | Only list templates of type |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.CreateTemplateListsInput)

mtype := "authorization"
collect.Mtype = mtype

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

shortcode := "Shortcode"
collect.Shortcode = shortcode


var result string
result,_ = sharedShortCode.CreateTemplateLists(collect)

```


### <a name="create_shortcode_sendsms"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateShortcodeSendsms") CreateShortcodeSendsms

> Send an SMS from a Ytel ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeSendsms(input *CreateShortcodeSendsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shortcode |  ``` Required ```  | The Short Code number that is the sender of this message |
| to |  ``` Required ```  | A valid 10-digit number that should receive the message |
| templateid |  ``` Required ```  | The unique identifier for the template used for the message |
| data |  ``` Required ```  | format of your data, example: {companyname}:test,{otpcode}:1234 |
| method |  ``` Optional ```  ``` DefaultValue ```  | Specifies the HTTP method used to request the required URL once the Short Code message is sent. |
| messageStatusCallback |  ``` Optional ```  | URL that can be requested to receive notification when Short Code message was sent. |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.CreateShortcodeSendsmsInput)

shortcode := "shortcode"
collect.Shortcode = shortcode

to := "to"
collect.To = to

templateid := uuid.NewV4()
collect.Templateid = templateid

data := "data"
collect.Data = data

method := "GET"
collect.Method = method

messageStatusCallback := "MessageStatusCallback"
collect.MessageStatusCallback = messageStatusCallback


var result string
result,_ = sharedShortCode.CreateShortcodeSendsms(collect)

```


### <a name="create_shortcode_getinboundsms"></a>![Method: ](https://apidocs.io/img/method.png ".sharedshortcode_pkg.CreateShortcodeGetinboundsms") CreateShortcodeGetinboundsms

> List All Inbound ShortCode


```go
func (me *SHAREDSHORTCODE_IMPL) CreateShortcodeGetinboundsms(input *CreateShortcodeGetinboundsmsInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| from |  ``` Optional ```  | From Number to Inbound ShortCode |
| shortcode |  ``` Optional ```  | Only list messages sent to this Short Code |
| datecreated |  ``` Optional ```  | Only list messages sent with the specified date |


#### Example Usage

```go
collect := new (sharedshortcode_pkg.CreateShortcodeGetinboundsmsInput)

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
result,_ = sharedShortCode.CreateShortcodeGetinboundsms(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="conference_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".conference_pkg") conference_pkg

### Get instance

Factory for the ``` CONFERENCE ``` interface can be accessed from the package conference_pkg.

```go
conference := conference_pkg.NewCONFERENCE()
```

### <a name="create_conferences_play_audio"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesPlayAudio") CreateConferencesPlayAudio

> Play an audio file during a conference.


```go
func (me *CONFERENCE_IMPL) CreateConferencesPlayAudio(input *CreateConferencesPlayAudioInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |
| audioUrl |  ``` Required ```  | The URL for the audio file that is to be played during the conference. Multiple audio files can be chained by using a semicolon. |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferencesPlayAudioInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

audioUrl := models_pkg.AudioUrl_MP3
collect.AudioUrl = audioUrl


var result string
result,_ = conference.CreateConferencesPlayAudio(collect)

```


### <a name="create_conferences_hangup_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesHangupParticipant") CreateConferencesHangupParticipant

> Remove a participant from a conference.


```go
func (me *CONFERENCE_IMPL) CreateConferencesHangupParticipant(input *CreateConferencesHangupParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferencesHangupParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid


var result string
result,_ = conference.CreateConferencesHangupParticipant(collect)

```


### <a name="create_conferences_viewconference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesViewconference") CreateConferencesViewconference

> Retrieve information about a conference by its ConferenceSid.


```go
func (me *CONFERENCE_IMPL) CreateConferencesViewconference(conferenceSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier of each conference resource |


#### Example Usage

```go
conferenceSid := "ConferenceSid"

var result string
result,_ = conference.CreateConferencesViewconference(conferenceSid)

```


### <a name="create_conferences_listconference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesListconference") CreateConferencesListconference

> Retrieve a list of conference objects.


```go
func (me *CONFERENCE_IMPL) CreateConferencesListconference(input *CreateConferencesListconferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | Number of individual resources listed in the response per page |
| friendlyName |  ``` Optional ```  | Only return conferences with the specified FriendlyName |
| dateCreated |  ``` Optional ```  | Conference created date |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferencesListconferenceInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

friendlyName := "FriendlyName"
collect.FriendlyName = friendlyName

dateCreated := "DateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = conference.CreateConferencesListconference(collect)

```


### <a name="create_conferences_list_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesListParticipant") CreateConferencesListParticipant

> Retrieve a list of participants for an in-progress conference.


```go
func (me *CONFERENCE_IMPL) CreateConferencesListParticipant(input *CreateConferencesListParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference. |
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| muted |  ``` Optional ```  | Specifies if participant should be muted. |
| deaf |  ``` Optional ```  | Specifies if the participant should hear audio in the conference. |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferencesListParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf


var result string
result,_ = conference.CreateConferencesListParticipant(collect)

```


### <a name="create_conferences_view_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesViewParticipant") CreateConferencesViewParticipant

> Retrieve information about a participant by its ParticipantSid.


```go
func (me *CONFERENCE_IMPL) CreateConferencesViewParticipant(input *CreateConferencesViewParticipantInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| conferenceSid |  ``` Required ```  | The unique identifier for a conference object. |
| participantSid |  ``` Required ```  | The unique identifier for a participant object. |


#### Example Usage

```go
collect := new (conference_pkg.CreateConferencesViewParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid


var result string
result,_ = conference.CreateConferencesViewParticipant(collect)

```


### <a name="create_conferences_add_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesAddParticipant") CreateConferencesAddParticipant

> Add Participant in conference 


```go
func (me *CONFERENCE_IMPL) CreateConferencesAddParticipant(input *CreateConferencesAddParticipantInput)(string,error)
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
collect := new (conference_pkg.CreateConferencesAddParticipantInput)

conferenceSid := "ConferenceSid"
collect.ConferenceSid = conferenceSid

participantNumber := "ParticipantNumber"
collect.ParticipantNumber = participantNumber

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf


var result string
result,_ = conference.CreateConferencesAddParticipant(collect)

```


### <a name="create_conferences_create_conference"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesCreateConference") CreateConferencesCreateConference

> Here you can experiment with initiating a conference call through Ytel and view the request response generated when doing so.


```go
func (me *CONFERENCE_IMPL) CreateConferencesCreateConference(input *CreateConferencesCreateConferenceInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| from |  ``` Required ```  | A valid 10-digit number (E.164 format) that will be initiating the conference call. |
| to |  ``` Required ```  | A valid 10-digit number (E.164 format) that is to receive the conference call. |
| url |  ``` Required ```  | URL requested once the conference connects |
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
collect := new (conference_pkg.CreateConferencesCreateConferenceInput)

from := "From"
collect.From = from

to := "To"
collect.To = to

url := "Url"
collect.Url = url

method := "POST"
collect.Method = method

statusCallBackUrl := "StatusCallBackUrl"
collect.StatusCallBackUrl = statusCallBackUrl

statusCallBackMethod := "StatusCallBackMethod"
collect.StatusCallBackMethod = statusCallBackMethod

fallbackUrl := "FallbackUrl"
collect.FallbackUrl = fallbackUrl

fallbackMethod := "FallbackMethod"
collect.FallbackMethod = fallbackMethod

record := false
collect.Record = record

recordCallBackUrl := "RecordCallBackUrl"
collect.RecordCallBackUrl = recordCallBackUrl

recordCallBackMethod := "RecordCallBackMethod"
collect.RecordCallBackMethod = recordCallBackMethod

scheduleTime := "ScheduleTime"
collect.ScheduleTime = scheduleTime

timeout,_ := strconv.ParseInt("42", 10, 8)
collect.Timeout = timeout


var result string
result,_ = conference.CreateConferencesCreateConference(collect)

```


### <a name="create_conferences_deaf_mute_participant"></a>![Method: ](https://apidocs.io/img/method.png ".conference_pkg.CreateConferencesDeafMuteParticipant") CreateConferencesDeafMuteParticipant

> Deaf Mute Participant


```go
func (me *CONFERENCE_IMPL) CreateConferencesDeafMuteParticipant(input *CreateConferencesDeafMuteParticipantInput)(string,error)
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
collect := new (conference_pkg.CreateConferencesDeafMuteParticipantInput)

conferenceSid := "conferenceSid"
collect.ConferenceSid = conferenceSid

participantSid := "ParticipantSid"
collect.ParticipantSid = participantSid

muted := false
collect.Muted = muted

deaf := false
collect.Deaf = deaf


var result string
result,_ = conference.CreateConferencesDeafMuteParticipant(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="carrier_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".carrier_pkg") carrier_pkg

### Get instance

Factory for the ``` CARRIER ``` interface can be accessed from the package carrier_pkg.

```go
carrier := carrier_pkg.NewCARRIER()
```

### <a name="create_carrier_lookup"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookup") CreateCarrierLookup

> Get the Carrier Lookup


```go
func (me *CARRIER_IMPL) CreateCarrierLookup(phoneNumber string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| phoneNumber |  ``` Required ```  | A valid 10-digit number (E.164 format). |


#### Example Usage

```go
phoneNumber := "PhoneNumber"

var result string
result,_ = carrier.CreateCarrierLookup(phoneNumber)

```


### <a name="create_carrier_lookuplist"></a>![Method: ](https://apidocs.io/img/method.png ".carrier_pkg.CreateCarrierLookuplist") CreateCarrierLookuplist

> Retrieve a list of carrier lookup objects.


```go
func (me *CARRIER_IMPL) CreateCarrierLookuplist(input *CreateCarrierLookuplistInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pageSize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |


#### Example Usage

```go
collect := new (carrier_pkg.CreateCarrierLookuplistInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pageSize,_ := strconv.ParseInt("10", 10, 8)
collect.PageSize = pageSize


var result string
result,_ = carrier.CreateCarrierLookuplist(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="email_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".email_pkg") email_pkg

### Get instance

Factory for the ``` EMAIL ``` interface can be accessed from the package email_pkg.

```go
email := email_pkg.NewEMAIL()
```

### <a name="create_email_deleteinvalidemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailDeleteinvalidemail") CreateEmailDeleteinvalidemail

> Remove an email from the invalid email list.


```go
func (me *EMAIL_IMPL) CreateEmailDeleteinvalidemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the invalid email list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateEmailDeleteinvalidemail(email)

```


### <a name="create_email_listblockemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailListblockemail") CreateEmailListblockemail

> Retrieve a list of emails that have been blocked.


```go
func (me *EMAIL_IMPL) CreateEmailListblockemail(input *CreateEmailListblockemailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of blocked emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.CreateEmailListblockemailInput)

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.CreateEmailListblockemail(collect)

```


### <a name="create_email_listspamemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailListspamemail") CreateEmailListspamemail

> Retrieve a list of emails that are on the spam list.


```go
func (me *EMAIL_IMPL) CreateEmailListspamemail(input *CreateEmailListspamemailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of spam emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.CreateEmailListspamemailInput)

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.CreateEmailListspamemail(collect)

```


### <a name="create_email_listbounceemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailListbounceemail") CreateEmailListbounceemail

> Retrieve a list of emails that have bounced.


```go
func (me *EMAIL_IMPL) CreateEmailListbounceemail(input *CreateEmailListbounceemailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of bounced emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.CreateEmailListbounceemailInput)

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.CreateEmailListbounceemail(collect)

```


### <a name="create_email_deletebouncesemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailDeletebouncesemail") CreateEmailDeletebouncesemail

> Remove an email address from the bounced list.


```go
func (me *EMAIL_IMPL) CreateEmailDeletebouncesemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to be remove from the bounced email list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateEmailDeletebouncesemail(email)

```


### <a name="create_email_listinvalidemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailListinvalidemail") CreateEmailListinvalidemail

> Retrieve a list of invalid email addresses.


```go
func (me *EMAIL_IMPL) CreateEmailListinvalidemail(input *CreateEmailListinvalidemailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of invalid emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.CreateEmailListinvalidemailInput)

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.CreateEmailListinvalidemail(collect)

```


### <a name="create_email_listunsubscribedemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailListunsubscribedemail") CreateEmailListunsubscribedemail

> Retrieve a list of email addresses from the unsubscribe list.


```go
func (me *EMAIL_IMPL) CreateEmailListunsubscribedemail(input *CreateEmailListunsubscribedemailInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| offset |  ``` Optional ```  | The starting point of the list of unsubscribed emails that should be returned. |
| limit |  ``` Optional ```  | The count of results that should be returned. |


#### Example Usage

```go
collect := new (email_pkg.CreateEmailListunsubscribedemailInput)

offset := "Offset"
collect.Offset = offset

limit := "Limit"
collect.Limit = limit


var result string
result,_ = email.CreateEmailListunsubscribedemail(collect)

```


### <a name="create_email_deleteunsubscribedemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailDeleteunsubscribedemail") CreateEmailDeleteunsubscribedemail

> Remove an email address from the list of unsubscribed emails.


```go
func (me *EMAIL_IMPL) CreateEmailDeleteunsubscribedemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the unsubscribe list. |


#### Example Usage

```go
email := "email"

var result string
result,_ = email.CreateEmailDeleteunsubscribedemail(email)

```


### <a name="create_email_addunsubscribesemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailAddunsubscribesemail") CreateEmailAddunsubscribesemail

> Add an email to the unsubscribe list


```go
func (me *EMAIL_IMPL) CreateEmailAddunsubscribesemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be added to the unsubscribe list |


#### Example Usage

```go
email := "email"

var result string
result,_ = email.CreateEmailAddunsubscribesemail(email)

```


### <a name="create_email_deleteblocksemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailDeleteblocksemail") CreateEmailDeleteblocksemail

> Remove an email from blocked emails list.


```go
func (me *EMAIL_IMPL) CreateEmailDeleteblocksemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | The email address to be remove from the blocked list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateEmailDeleteblocksemail(email)

```


### <a name="create_email_deletespamemail"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailDeletespamemail") CreateEmailDeletespamemail

> Remove an email from the spam email list.


```go
func (me *EMAIL_IMPL) CreateEmailDeletespamemail(email string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| email |  ``` Required ```  | A valid email address that is to be remove from the spam list. |


#### Example Usage

```go
email := "Email"

var result string
result,_ = email.CreateEmailDeletespamemail(email)

```


### <a name="create_email_sendemails"></a>![Method: ](https://apidocs.io/img/method.png ".email_pkg.CreateEmailSendemails") CreateEmailSendemails

> Create and submit an email message to one or more email addresses.


```go
func (me *EMAIL_IMPL) CreateEmailSendemails(input *CreateEmailSendemailsInput)(string,error)
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
collect := new (email_pkg.CreateEmailSendemailsInput)

to := "To"
collect.To = to

mtype := models_pkg.Type_TEXT
collect.Mtype = mtype

subject := "Subject"
collect.Subject = subject

message := "Message"
collect.Message = message

from := "From"
collect.From = from

cc := "Cc"
collect.Cc = cc

bcc := "Bcc"
collect.Bcc = bcc

attachment := "Attachment"
collect.Attachment = attachment


var result string
result,_ = email.CreateEmailSendemails(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="account_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".account_pkg") account_pkg

### Get instance

Factory for the ``` ACCOUNT ``` interface can be accessed from the package account_pkg.

```go
account := account_pkg.NewACCOUNT()
```

### <a name="create_accounts_viewaccount"></a>![Method: ](https://apidocs.io/img/method.png ".account_pkg.CreateAccountsViewaccount") CreateAccountsViewaccount

> Retrieve information regarding your Ytel account by a specific date. The response object will contain data such as account status, balance, and account usage totals.


```go
func (me *ACCOUNT_IMPL) CreateAccountsViewaccount(date string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| date |  ``` Required ```  | Filter account information based on date. |


#### Example Usage

```go
date := "Date"

var result string
result,_ = account.CreateAccountsViewaccount(date)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="subaccount_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".subaccount_pkg") subaccount_pkg

### Get instance

Factory for the ``` SUBACCOUNT ``` interface can be accessed from the package subaccount_pkg.

```go
subAccount := subaccount_pkg.NewSUBACCOUNT()
```

### <a name="create_user_subaccountactivation"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateUserSubaccountactivation") CreateUserSubaccountactivation

> Suspend or unsuspend


```go
func (me *SUBACCOUNT_IMPL) CreateUserSubaccountactivation(input *CreateUserSubaccountactivationInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be activated or suspended |
| mActivate |  ``` Required ```  | 0 to suspend or 1 to activate |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateUserSubaccountactivationInput)

subAccountSID := "SubAccountSID"
collect.SubAccountSID = subAccountSID

mActivate := models_pkg.mActivate_ACTIVATE
collect.MActivate = mActivate


var result string
result,_ = subAccount.CreateUserSubaccountactivation(collect)

```


### <a name="create_user_deletesubaccount"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateUserDeletesubaccount") CreateUserDeletesubaccount

> Delete sub account or merge numbers into parent


```go
func (me *SUBACCOUNT_IMPL) CreateUserDeletesubaccount(input *CreateUserDeletesubaccountInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subAccountSID |  ``` Required ```  | The SubaccountSid to be deleted |
| mergeNumber |  ``` Required ```  | 0 to delete or 1 to merge numbers to parent account. |


#### Example Usage

```go
collect := new (subaccount_pkg.CreateUserDeletesubaccountInput)

subAccountSID := "SubAccountSID"
collect.SubAccountSID = subAccountSID

mergeNumber := models_pkg.MergeNumber_DELETE
collect.MergeNumber = mergeNumber


var result string
result,_ = subAccount.CreateUserDeletesubaccount(collect)

```


### <a name="create_user_createsubaccount"></a>![Method: ](https://apidocs.io/img/method.png ".subaccount_pkg.CreateUserCreatesubaccount") CreateUserCreatesubaccount

> Create a sub user account under the parent account


```go
func (me *SUBACCOUNT_IMPL) CreateUserCreatesubaccount(input *CreateUserCreatesubaccountInput)(string,error)
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
collect := new (subaccount_pkg.CreateUserCreatesubaccountInput)

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


var result string
result,_ = subAccount.CreateUserCreatesubaccount(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="address_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".address_pkg") address_pkg

### Get instance

Factory for the ``` ADDRESS ``` interface can be accessed from the package address_pkg.

```go
address := address_pkg.NewADDRESS()
```

### <a name="address_deleteaddress"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.AddressDeleteaddress") AddressDeleteaddress

> To delete Address to your address book


```go
func (me *ADDRESS_IMPL) AddressDeleteaddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be deleted. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.AddressDeleteaddress(addressid)

```


### <a name="address_verifyaddress"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.AddressVerifyaddress") AddressVerifyaddress

> Validates an address given.


```go
func (me *ADDRESS_IMPL) AddressVerifyaddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be verified. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.AddressVerifyaddress(addressid)

```


### <a name="address_viewaddress"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.AddressViewaddress") AddressViewaddress

> View Address Specific address Book by providing the address id


```go
func (me *ADDRESS_IMPL) AddressViewaddress(addressid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| addressid |  ``` Required ```  | The identifier of the address to be retrieved. |


#### Example Usage

```go
addressid := "addressid"

var result string
result,_ = address.AddressViewaddress(addressid)

```


### <a name="address_createaddress"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.AddressCreateaddress") AddressCreateaddress

> To add an address to your address book, you create a new address object. You can retrieve and delete individual addresses as well as get a list of addresses. Addresses are identified by a unique random ID.


```go
func (me *ADDRESS_IMPL) AddressCreateaddress(input *AddressCreateaddressInput)(string,error)
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
collect := new (address_pkg.AddressCreateaddressInput)

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

description := "Description"
collect.Description = description

email := "email"
collect.Email = email

phone := "Phone"
collect.Phone = phone


var result string
result,_ = address.AddressCreateaddress(collect)

```


### <a name="address_listaddress"></a>![Method: ](https://apidocs.io/img/method.png ".address_pkg.AddressListaddress") AddressListaddress

> List All Address 


```go
func (me *ADDRESS_IMPL) AddressListaddress(input *AddressListaddressInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | How many results to return, default is 10, max is 100, must be an integer |
| addressid |  ``` Optional ```  | addresses Sid |
| dateCreated |  ``` Optional ```  | date created address. |


#### Example Usage

```go
collect := new (address_pkg.AddressListaddressInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

addressid := "addressid"
collect.Addressid = addressid

dateCreated := "dateCreated"
collect.DateCreated = dateCreated


var result string
result,_ = address.AddressListaddress(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="recording_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".recording_pkg") recording_pkg

### Get instance

Factory for the ``` RECORDING ``` interface can be accessed from the package recording_pkg.

```go
recording := recording_pkg.NewRECORDING()
```

### <a name="create_recording_deleterecording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateRecordingDeleterecording") CreateRecordingDeleterecording

> Remove a recording from your Ytel account.


```go
func (me *RECORDING_IMPL) CreateRecordingDeleterecording(recordingsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for a recording. |


#### Example Usage

```go
recordingsid := "recordingsid"

var result string
result,_ = recording.CreateRecordingDeleterecording(recordingsid)

```


### <a name="create_recording_viewrecording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateRecordingViewrecording") CreateRecordingViewrecording

> Retrieve the recording of a call by its RecordingSid. This resource will return information regarding the call such as start time, end time, duration, and so forth.


```go
func (me *RECORDING_IMPL) CreateRecordingViewrecording(recordingsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingsid |  ``` Required ```  | The unique identifier for the recording. |


#### Example Usage

```go
recordingsid := "recordingsid"

var result string
result,_ = recording.CreateRecordingViewrecording(recordingsid)

```


### <a name="create_recording_listrecording"></a>![Method: ](https://apidocs.io/img/method.png ".recording_pkg.CreateRecordingListrecording") CreateRecordingListrecording

> Retrieve a list of recording objects.


```go
func (me *RECORDING_IMPL) CreateRecordingListrecording(input *CreateRecordingListrecordingInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| datecreated |  ``` Optional ```  | Filter results by creation date |
| callsid |  ``` Optional ```  | The unique identifier for a call. |


#### Example Usage

```go
collect := new (recording_pkg.CreateRecordingListrecordingInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

datecreated := "Datecreated"
collect.Datecreated = datecreated

callsid := "callsid"
collect.Callsid = callsid


var result string
result,_ = recording.CreateRecordingListrecording(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="transcription_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".transcription_pkg") transcription_pkg

### Get instance

Factory for the ``` TRANSCRIPTION ``` interface can be accessed from the package transcription_pkg.

```go
transcription := transcription_pkg.NewTRANSCRIPTION()
```

### <a name="create_transcriptions_audiourltranscription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscriptionsAudiourltranscription") CreateTranscriptionsAudiourltranscription

> Transcribe an audio recording from a file.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscriptionsAudiourltranscription(audiourl string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| audiourl |  ``` Required ```  | URL pointing to the location of the audio file that is to be transcribed. |


#### Example Usage

```go
audiourl := "audiourl"

var result string
result,_ = transcription.CreateTranscriptionsAudiourltranscription(audiourl)

```


### <a name="create_transcriptions_recordingtranscription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscriptionsRecordingtranscription") CreateTranscriptionsRecordingtranscription

> Transcribe a recording by its RecordingSid.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscriptionsRecordingtranscription(recordingSid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| recordingSid |  ``` Required ```  | The unique identifier for a recording object. |


#### Example Usage

```go
recordingSid := "recordingSid"

var result string
result,_ = transcription.CreateTranscriptionsRecordingtranscription(recordingSid)

```


### <a name="create_transcriptions_viewtranscription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscriptionsViewtranscription") CreateTranscriptionsViewtranscription

> Retrieve information about a transaction by its TranscriptionSid.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscriptionsViewtranscription(transcriptionsid string)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| transcriptionsid |  ``` Required ```  | The unique identifier for a transcription object. |


#### Example Usage

```go
transcriptionsid := "transcriptionsid"

var result string
result,_ = transcription.CreateTranscriptionsViewtranscription(transcriptionsid)

```


### <a name="create_transcriptions_listtranscription"></a>![Method: ](https://apidocs.io/img/method.png ".transcription_pkg.CreateTranscriptionsListtranscription") CreateTranscriptionsListtranscription

> Retrieve a list of transcription objects for your Ytel account.


```go
func (me *TRANSCRIPTION_IMPL) CreateTranscriptionsListtranscription(input *CreateTranscriptionsListtranscriptionInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| page |  ``` Optional ```  ``` DefaultValue ```  | The page count to retrieve from the total results in the collection. Page indexing starts at 1. |
| pagesize |  ``` Optional ```  ``` DefaultValue ```  | The count of objects to return per page. |
| status |  ``` Optional ```  | The state of the transcription. |
| dateTranscribed |  ``` Optional ```  | The date the transcription took place. |


#### Example Usage

```go
collect := new (transcription_pkg.CreateTranscriptionsListtranscriptionInput)

page,_ := strconv.ParseInt("1", 10, 8)
collect.Page = page

pagesize,_ := strconv.ParseInt("10", 10, 8)
collect.Pagesize = pagesize

status := models_pkg.Status_INPROGRESS
collect.Status = status

dateTranscribed := "dateTranscribed"
collect.DateTranscribed = dateTranscribed


var result string
result,_ = transcription.CreateTranscriptionsListtranscription(collect)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="usage_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".usage_pkg") usage_pkg

### Get instance

Factory for the ``` USAGE ``` interface can be accessed from the package usage_pkg.

```go
usage := usage_pkg.NewUSAGE()
```

### <a name="create_usage_listusage"></a>![Method: ](https://apidocs.io/img/method.png ".usage_pkg.CreateUsageListusage") CreateUsageListusage

> Retrieve usage metrics regarding your Ytel account. The results includes inbound/outbound voice calls and inbound/outbound SMS messages as well as carrier lookup requests.


```go
func (me *USAGE_IMPL) CreateUsageListusage(input *CreateUsageListusageInput)(string,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| productCode |  ``` Optional ```  ``` DefaultValue ```  | Filter usage results by product type. |
| startDate |  ``` Optional ```  ``` DefaultValue ```  | Filter usage objects by start date. |
| endDate |  ``` Optional ```  ``` DefaultValue ```  | Filter usage objects by end date. |
| includeSubAccounts |  ``` Optional ```  | Will include all subaccount usage data |


#### Example Usage

```go
collect := new (usage_pkg.CreateUsageListusageInput)

productCode := models_pkg.ProductCode27_ALL
collect.ProductCode = productCode

startDate := "2016-09-06"
collect.StartDate = startDate

endDate := "2016-09-06"
collect.EndDate = endDate

includeSubAccounts := "IncludeSubAccounts"
collect.IncludeSubAccounts = includeSubAccounts


var result string
result,_ = usage.CreateUsageListusage(collect)

```


[Back to List of Controllers](#list_of_controllers)



