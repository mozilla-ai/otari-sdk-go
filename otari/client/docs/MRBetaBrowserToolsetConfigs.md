# MRBetaBrowserToolsetConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseTab** | Pointer to [**NullableMRBetaBrowserCloseTabConfig**](MRBetaBrowserCloseTabConfig.md) |  | [optional] 
**DoubleClick** | Pointer to [**NullableMRBetaBrowserDoubleClickConfig**](MRBetaBrowserDoubleClickConfig.md) |  | [optional] 
**FileUpload** | Pointer to [**NullableMRBetaBrowserFileUploadConfig**](MRBetaBrowserFileUploadConfig.md) |  | [optional] 
**Find** | Pointer to [**NullableMRBetaBrowserFindConfig**](MRBetaBrowserFindConfig.md) |  | [optional] 
**FormInput** | Pointer to [**NullableMRBetaBrowserFormInputConfig**](MRBetaBrowserFormInputConfig.md) |  | [optional] 
**GetPageText** | Pointer to [**NullableMRBetaBrowserGetPageTextConfig**](MRBetaBrowserGetPageTextConfig.md) |  | [optional] 
**HoldKey** | Pointer to [**NullableMRBetaBrowserHoldKeyConfig**](MRBetaBrowserHoldKeyConfig.md) |  | [optional] 
**Hover** | Pointer to [**NullableMRBetaBrowserHoverConfig**](MRBetaBrowserHoverConfig.md) |  | [optional] 
**JavascriptExec** | Pointer to [**NullableMRBetaBrowserJavascriptExecConfig**](MRBetaBrowserJavascriptExecConfig.md) |  | [optional] 
**Key** | Pointer to [**NullableMRBetaBrowserKeyConfig**](MRBetaBrowserKeyConfig.md) |  | [optional] 
**LeftClick** | Pointer to [**NullableMRBetaBrowserLeftClickConfig**](MRBetaBrowserLeftClickConfig.md) |  | [optional] 
**LeftClickDrag** | Pointer to [**NullableMRBetaBrowserLeftClickDragConfig**](MRBetaBrowserLeftClickDragConfig.md) |  | [optional] 
**LeftMouseDown** | Pointer to [**NullableMRBetaBrowserLeftMouseDownConfig**](MRBetaBrowserLeftMouseDownConfig.md) |  | [optional] 
**LeftMouseUp** | Pointer to [**NullableMRBetaBrowserLeftMouseUpConfig**](MRBetaBrowserLeftMouseUpConfig.md) |  | [optional] 
**ListTabs** | Pointer to [**NullableMRBetaBrowserListTabsConfig**](MRBetaBrowserListTabsConfig.md) |  | [optional] 
**MiddleClick** | Pointer to [**NullableMRBetaBrowserMiddleClickConfig**](MRBetaBrowserMiddleClickConfig.md) |  | [optional] 
**MouseMove** | Pointer to [**NullableMRBetaBrowserMouseMoveConfig**](MRBetaBrowserMouseMoveConfig.md) |  | [optional] 
**Navigate** | Pointer to [**NullableMRBetaBrowserNavigateConfig**](MRBetaBrowserNavigateConfig.md) |  | [optional] 
**NewTab** | Pointer to [**NullableMRBetaBrowserNewTabConfig**](MRBetaBrowserNewTabConfig.md) |  | [optional] 
**ReadConsole** | Pointer to [**NullableMRBetaBrowserReadConsoleConfig**](MRBetaBrowserReadConsoleConfig.md) |  | [optional] 
**ReadNetwork** | Pointer to [**NullableMRBetaBrowserReadNetworkConfig**](MRBetaBrowserReadNetworkConfig.md) |  | [optional] 
**ReadPage** | Pointer to [**NullableMRBetaBrowserReadPageConfig**](MRBetaBrowserReadPageConfig.md) |  | [optional] 
**RightClick** | Pointer to [**NullableMRBetaBrowserRightClickConfig**](MRBetaBrowserRightClickConfig.md) |  | [optional] 
**Screenshot** | Pointer to [**NullableMRBetaBrowserScreenshotConfig**](MRBetaBrowserScreenshotConfig.md) |  | [optional] 
**Scroll** | Pointer to [**NullableMRBetaBrowserScrollConfig**](MRBetaBrowserScrollConfig.md) |  | [optional] 
**ScrollTo** | Pointer to [**NullableMRBetaBrowserScrollToConfig**](MRBetaBrowserScrollToConfig.md) |  | [optional] 
**SwitchTab** | Pointer to [**NullableMRBetaBrowserSwitchTabConfig**](MRBetaBrowserSwitchTabConfig.md) |  | [optional] 
**TripleClick** | Pointer to [**NullableMRBetaBrowserTripleClickConfig**](MRBetaBrowserTripleClickConfig.md) |  | [optional] 
**Type** | Pointer to [**NullableMRBetaBrowserTypeConfig**](MRBetaBrowserTypeConfig.md) |  | [optional] 
**Wait** | Pointer to [**NullableMRBetaBrowserWaitConfig**](MRBetaBrowserWaitConfig.md) |  | [optional] 
**Zoom** | Pointer to [**NullableMRBetaBrowserZoomConfig**](MRBetaBrowserZoomConfig.md) |  | [optional] 

## Methods

### NewMRBetaBrowserToolsetConfigs

`func NewMRBetaBrowserToolsetConfigs() *MRBetaBrowserToolsetConfigs`

NewMRBetaBrowserToolsetConfigs instantiates a new MRBetaBrowserToolsetConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaBrowserToolsetConfigsWithDefaults

`func NewMRBetaBrowserToolsetConfigsWithDefaults() *MRBetaBrowserToolsetConfigs`

NewMRBetaBrowserToolsetConfigsWithDefaults instantiates a new MRBetaBrowserToolsetConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseTab

`func (o *MRBetaBrowserToolsetConfigs) GetCloseTab() MRBetaBrowserCloseTabConfig`

GetCloseTab returns the CloseTab field if non-nil, zero value otherwise.

### GetCloseTabOk

`func (o *MRBetaBrowserToolsetConfigs) GetCloseTabOk() (*MRBetaBrowserCloseTabConfig, bool)`

GetCloseTabOk returns a tuple with the CloseTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTab

`func (o *MRBetaBrowserToolsetConfigs) SetCloseTab(v MRBetaBrowserCloseTabConfig)`

SetCloseTab sets CloseTab field to given value.

### HasCloseTab

`func (o *MRBetaBrowserToolsetConfigs) HasCloseTab() bool`

HasCloseTab returns a boolean if a field has been set.

### SetCloseTabNil

`func (o *MRBetaBrowserToolsetConfigs) SetCloseTabNil(b bool)`

 SetCloseTabNil sets the value for CloseTab to be an explicit nil

### UnsetCloseTab
`func (o *MRBetaBrowserToolsetConfigs) UnsetCloseTab()`

UnsetCloseTab ensures that no value is present for CloseTab, not even an explicit nil
### GetDoubleClick

`func (o *MRBetaBrowserToolsetConfigs) GetDoubleClick() MRBetaBrowserDoubleClickConfig`

GetDoubleClick returns the DoubleClick field if non-nil, zero value otherwise.

### GetDoubleClickOk

`func (o *MRBetaBrowserToolsetConfigs) GetDoubleClickOk() (*MRBetaBrowserDoubleClickConfig, bool)`

GetDoubleClickOk returns a tuple with the DoubleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleClick

`func (o *MRBetaBrowserToolsetConfigs) SetDoubleClick(v MRBetaBrowserDoubleClickConfig)`

SetDoubleClick sets DoubleClick field to given value.

### HasDoubleClick

`func (o *MRBetaBrowserToolsetConfigs) HasDoubleClick() bool`

HasDoubleClick returns a boolean if a field has been set.

### SetDoubleClickNil

`func (o *MRBetaBrowserToolsetConfigs) SetDoubleClickNil(b bool)`

 SetDoubleClickNil sets the value for DoubleClick to be an explicit nil

### UnsetDoubleClick
`func (o *MRBetaBrowserToolsetConfigs) UnsetDoubleClick()`

UnsetDoubleClick ensures that no value is present for DoubleClick, not even an explicit nil
### GetFileUpload

`func (o *MRBetaBrowserToolsetConfigs) GetFileUpload() MRBetaBrowserFileUploadConfig`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *MRBetaBrowserToolsetConfigs) GetFileUploadOk() (*MRBetaBrowserFileUploadConfig, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *MRBetaBrowserToolsetConfigs) SetFileUpload(v MRBetaBrowserFileUploadConfig)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *MRBetaBrowserToolsetConfigs) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### SetFileUploadNil

`func (o *MRBetaBrowserToolsetConfigs) SetFileUploadNil(b bool)`

 SetFileUploadNil sets the value for FileUpload to be an explicit nil

### UnsetFileUpload
`func (o *MRBetaBrowserToolsetConfigs) UnsetFileUpload()`

UnsetFileUpload ensures that no value is present for FileUpload, not even an explicit nil
### GetFind

`func (o *MRBetaBrowserToolsetConfigs) GetFind() MRBetaBrowserFindConfig`

GetFind returns the Find field if non-nil, zero value otherwise.

### GetFindOk

`func (o *MRBetaBrowserToolsetConfigs) GetFindOk() (*MRBetaBrowserFindConfig, bool)`

GetFindOk returns a tuple with the Find field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFind

`func (o *MRBetaBrowserToolsetConfigs) SetFind(v MRBetaBrowserFindConfig)`

SetFind sets Find field to given value.

### HasFind

`func (o *MRBetaBrowserToolsetConfigs) HasFind() bool`

HasFind returns a boolean if a field has been set.

### SetFindNil

`func (o *MRBetaBrowserToolsetConfigs) SetFindNil(b bool)`

 SetFindNil sets the value for Find to be an explicit nil

### UnsetFind
`func (o *MRBetaBrowserToolsetConfigs) UnsetFind()`

UnsetFind ensures that no value is present for Find, not even an explicit nil
### GetFormInput

`func (o *MRBetaBrowserToolsetConfigs) GetFormInput() MRBetaBrowserFormInputConfig`

GetFormInput returns the FormInput field if non-nil, zero value otherwise.

### GetFormInputOk

`func (o *MRBetaBrowserToolsetConfigs) GetFormInputOk() (*MRBetaBrowserFormInputConfig, bool)`

GetFormInputOk returns a tuple with the FormInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormInput

`func (o *MRBetaBrowserToolsetConfigs) SetFormInput(v MRBetaBrowserFormInputConfig)`

SetFormInput sets FormInput field to given value.

### HasFormInput

`func (o *MRBetaBrowserToolsetConfigs) HasFormInput() bool`

HasFormInput returns a boolean if a field has been set.

### SetFormInputNil

`func (o *MRBetaBrowserToolsetConfigs) SetFormInputNil(b bool)`

 SetFormInputNil sets the value for FormInput to be an explicit nil

### UnsetFormInput
`func (o *MRBetaBrowserToolsetConfigs) UnsetFormInput()`

UnsetFormInput ensures that no value is present for FormInput, not even an explicit nil
### GetGetPageText

`func (o *MRBetaBrowserToolsetConfigs) GetGetPageText() MRBetaBrowserGetPageTextConfig`

GetGetPageText returns the GetPageText field if non-nil, zero value otherwise.

### GetGetPageTextOk

`func (o *MRBetaBrowserToolsetConfigs) GetGetPageTextOk() (*MRBetaBrowserGetPageTextConfig, bool)`

GetGetPageTextOk returns a tuple with the GetPageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetPageText

`func (o *MRBetaBrowserToolsetConfigs) SetGetPageText(v MRBetaBrowserGetPageTextConfig)`

SetGetPageText sets GetPageText field to given value.

### HasGetPageText

`func (o *MRBetaBrowserToolsetConfigs) HasGetPageText() bool`

HasGetPageText returns a boolean if a field has been set.

### SetGetPageTextNil

`func (o *MRBetaBrowserToolsetConfigs) SetGetPageTextNil(b bool)`

 SetGetPageTextNil sets the value for GetPageText to be an explicit nil

### UnsetGetPageText
`func (o *MRBetaBrowserToolsetConfigs) UnsetGetPageText()`

UnsetGetPageText ensures that no value is present for GetPageText, not even an explicit nil
### GetHoldKey

`func (o *MRBetaBrowserToolsetConfigs) GetHoldKey() MRBetaBrowserHoldKeyConfig`

GetHoldKey returns the HoldKey field if non-nil, zero value otherwise.

### GetHoldKeyOk

`func (o *MRBetaBrowserToolsetConfigs) GetHoldKeyOk() (*MRBetaBrowserHoldKeyConfig, bool)`

GetHoldKeyOk returns a tuple with the HoldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldKey

`func (o *MRBetaBrowserToolsetConfigs) SetHoldKey(v MRBetaBrowserHoldKeyConfig)`

SetHoldKey sets HoldKey field to given value.

### HasHoldKey

`func (o *MRBetaBrowserToolsetConfigs) HasHoldKey() bool`

HasHoldKey returns a boolean if a field has been set.

### SetHoldKeyNil

`func (o *MRBetaBrowserToolsetConfigs) SetHoldKeyNil(b bool)`

 SetHoldKeyNil sets the value for HoldKey to be an explicit nil

### UnsetHoldKey
`func (o *MRBetaBrowserToolsetConfigs) UnsetHoldKey()`

UnsetHoldKey ensures that no value is present for HoldKey, not even an explicit nil
### GetHover

`func (o *MRBetaBrowserToolsetConfigs) GetHover() MRBetaBrowserHoverConfig`

GetHover returns the Hover field if non-nil, zero value otherwise.

### GetHoverOk

`func (o *MRBetaBrowserToolsetConfigs) GetHoverOk() (*MRBetaBrowserHoverConfig, bool)`

GetHoverOk returns a tuple with the Hover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHover

`func (o *MRBetaBrowserToolsetConfigs) SetHover(v MRBetaBrowserHoverConfig)`

SetHover sets Hover field to given value.

### HasHover

`func (o *MRBetaBrowserToolsetConfigs) HasHover() bool`

HasHover returns a boolean if a field has been set.

### SetHoverNil

`func (o *MRBetaBrowserToolsetConfigs) SetHoverNil(b bool)`

 SetHoverNil sets the value for Hover to be an explicit nil

### UnsetHover
`func (o *MRBetaBrowserToolsetConfigs) UnsetHover()`

UnsetHover ensures that no value is present for Hover, not even an explicit nil
### GetJavascriptExec

`func (o *MRBetaBrowserToolsetConfigs) GetJavascriptExec() MRBetaBrowserJavascriptExecConfig`

GetJavascriptExec returns the JavascriptExec field if non-nil, zero value otherwise.

### GetJavascriptExecOk

`func (o *MRBetaBrowserToolsetConfigs) GetJavascriptExecOk() (*MRBetaBrowserJavascriptExecConfig, bool)`

GetJavascriptExecOk returns a tuple with the JavascriptExec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptExec

`func (o *MRBetaBrowserToolsetConfigs) SetJavascriptExec(v MRBetaBrowserJavascriptExecConfig)`

SetJavascriptExec sets JavascriptExec field to given value.

### HasJavascriptExec

`func (o *MRBetaBrowserToolsetConfigs) HasJavascriptExec() bool`

HasJavascriptExec returns a boolean if a field has been set.

### SetJavascriptExecNil

`func (o *MRBetaBrowserToolsetConfigs) SetJavascriptExecNil(b bool)`

 SetJavascriptExecNil sets the value for JavascriptExec to be an explicit nil

### UnsetJavascriptExec
`func (o *MRBetaBrowserToolsetConfigs) UnsetJavascriptExec()`

UnsetJavascriptExec ensures that no value is present for JavascriptExec, not even an explicit nil
### GetKey

`func (o *MRBetaBrowserToolsetConfigs) GetKey() MRBetaBrowserKeyConfig`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MRBetaBrowserToolsetConfigs) GetKeyOk() (*MRBetaBrowserKeyConfig, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MRBetaBrowserToolsetConfigs) SetKey(v MRBetaBrowserKeyConfig)`

SetKey sets Key field to given value.

### HasKey

`func (o *MRBetaBrowserToolsetConfigs) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MRBetaBrowserToolsetConfigs) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MRBetaBrowserToolsetConfigs) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetLeftClick

`func (o *MRBetaBrowserToolsetConfigs) GetLeftClick() MRBetaBrowserLeftClickConfig`

GetLeftClick returns the LeftClick field if non-nil, zero value otherwise.

### GetLeftClickOk

`func (o *MRBetaBrowserToolsetConfigs) GetLeftClickOk() (*MRBetaBrowserLeftClickConfig, bool)`

GetLeftClickOk returns a tuple with the LeftClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftClick

`func (o *MRBetaBrowserToolsetConfigs) SetLeftClick(v MRBetaBrowserLeftClickConfig)`

SetLeftClick sets LeftClick field to given value.

### HasLeftClick

`func (o *MRBetaBrowserToolsetConfigs) HasLeftClick() bool`

HasLeftClick returns a boolean if a field has been set.

### SetLeftClickNil

`func (o *MRBetaBrowserToolsetConfigs) SetLeftClickNil(b bool)`

 SetLeftClickNil sets the value for LeftClick to be an explicit nil

### UnsetLeftClick
`func (o *MRBetaBrowserToolsetConfigs) UnsetLeftClick()`

UnsetLeftClick ensures that no value is present for LeftClick, not even an explicit nil
### GetLeftClickDrag

`func (o *MRBetaBrowserToolsetConfigs) GetLeftClickDrag() MRBetaBrowserLeftClickDragConfig`

GetLeftClickDrag returns the LeftClickDrag field if non-nil, zero value otherwise.

### GetLeftClickDragOk

`func (o *MRBetaBrowserToolsetConfigs) GetLeftClickDragOk() (*MRBetaBrowserLeftClickDragConfig, bool)`

GetLeftClickDragOk returns a tuple with the LeftClickDrag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftClickDrag

`func (o *MRBetaBrowserToolsetConfigs) SetLeftClickDrag(v MRBetaBrowserLeftClickDragConfig)`

SetLeftClickDrag sets LeftClickDrag field to given value.

### HasLeftClickDrag

`func (o *MRBetaBrowserToolsetConfigs) HasLeftClickDrag() bool`

HasLeftClickDrag returns a boolean if a field has been set.

### SetLeftClickDragNil

`func (o *MRBetaBrowserToolsetConfigs) SetLeftClickDragNil(b bool)`

 SetLeftClickDragNil sets the value for LeftClickDrag to be an explicit nil

### UnsetLeftClickDrag
`func (o *MRBetaBrowserToolsetConfigs) UnsetLeftClickDrag()`

UnsetLeftClickDrag ensures that no value is present for LeftClickDrag, not even an explicit nil
### GetLeftMouseDown

`func (o *MRBetaBrowserToolsetConfigs) GetLeftMouseDown() MRBetaBrowserLeftMouseDownConfig`

GetLeftMouseDown returns the LeftMouseDown field if non-nil, zero value otherwise.

### GetLeftMouseDownOk

`func (o *MRBetaBrowserToolsetConfigs) GetLeftMouseDownOk() (*MRBetaBrowserLeftMouseDownConfig, bool)`

GetLeftMouseDownOk returns a tuple with the LeftMouseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftMouseDown

`func (o *MRBetaBrowserToolsetConfigs) SetLeftMouseDown(v MRBetaBrowserLeftMouseDownConfig)`

SetLeftMouseDown sets LeftMouseDown field to given value.

### HasLeftMouseDown

`func (o *MRBetaBrowserToolsetConfigs) HasLeftMouseDown() bool`

HasLeftMouseDown returns a boolean if a field has been set.

### SetLeftMouseDownNil

`func (o *MRBetaBrowserToolsetConfigs) SetLeftMouseDownNil(b bool)`

 SetLeftMouseDownNil sets the value for LeftMouseDown to be an explicit nil

### UnsetLeftMouseDown
`func (o *MRBetaBrowserToolsetConfigs) UnsetLeftMouseDown()`

UnsetLeftMouseDown ensures that no value is present for LeftMouseDown, not even an explicit nil
### GetLeftMouseUp

`func (o *MRBetaBrowserToolsetConfigs) GetLeftMouseUp() MRBetaBrowserLeftMouseUpConfig`

GetLeftMouseUp returns the LeftMouseUp field if non-nil, zero value otherwise.

### GetLeftMouseUpOk

`func (o *MRBetaBrowserToolsetConfigs) GetLeftMouseUpOk() (*MRBetaBrowserLeftMouseUpConfig, bool)`

GetLeftMouseUpOk returns a tuple with the LeftMouseUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftMouseUp

`func (o *MRBetaBrowserToolsetConfigs) SetLeftMouseUp(v MRBetaBrowserLeftMouseUpConfig)`

SetLeftMouseUp sets LeftMouseUp field to given value.

### HasLeftMouseUp

`func (o *MRBetaBrowserToolsetConfigs) HasLeftMouseUp() bool`

HasLeftMouseUp returns a boolean if a field has been set.

### SetLeftMouseUpNil

`func (o *MRBetaBrowserToolsetConfigs) SetLeftMouseUpNil(b bool)`

 SetLeftMouseUpNil sets the value for LeftMouseUp to be an explicit nil

### UnsetLeftMouseUp
`func (o *MRBetaBrowserToolsetConfigs) UnsetLeftMouseUp()`

UnsetLeftMouseUp ensures that no value is present for LeftMouseUp, not even an explicit nil
### GetListTabs

`func (o *MRBetaBrowserToolsetConfigs) GetListTabs() MRBetaBrowserListTabsConfig`

GetListTabs returns the ListTabs field if non-nil, zero value otherwise.

### GetListTabsOk

`func (o *MRBetaBrowserToolsetConfigs) GetListTabsOk() (*MRBetaBrowserListTabsConfig, bool)`

GetListTabsOk returns a tuple with the ListTabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListTabs

`func (o *MRBetaBrowserToolsetConfigs) SetListTabs(v MRBetaBrowserListTabsConfig)`

SetListTabs sets ListTabs field to given value.

### HasListTabs

`func (o *MRBetaBrowserToolsetConfigs) HasListTabs() bool`

HasListTabs returns a boolean if a field has been set.

### SetListTabsNil

`func (o *MRBetaBrowserToolsetConfigs) SetListTabsNil(b bool)`

 SetListTabsNil sets the value for ListTabs to be an explicit nil

### UnsetListTabs
`func (o *MRBetaBrowserToolsetConfigs) UnsetListTabs()`

UnsetListTabs ensures that no value is present for ListTabs, not even an explicit nil
### GetMiddleClick

`func (o *MRBetaBrowserToolsetConfigs) GetMiddleClick() MRBetaBrowserMiddleClickConfig`

GetMiddleClick returns the MiddleClick field if non-nil, zero value otherwise.

### GetMiddleClickOk

`func (o *MRBetaBrowserToolsetConfigs) GetMiddleClickOk() (*MRBetaBrowserMiddleClickConfig, bool)`

GetMiddleClickOk returns a tuple with the MiddleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleClick

`func (o *MRBetaBrowserToolsetConfigs) SetMiddleClick(v MRBetaBrowserMiddleClickConfig)`

SetMiddleClick sets MiddleClick field to given value.

### HasMiddleClick

`func (o *MRBetaBrowserToolsetConfigs) HasMiddleClick() bool`

HasMiddleClick returns a boolean if a field has been set.

### SetMiddleClickNil

`func (o *MRBetaBrowserToolsetConfigs) SetMiddleClickNil(b bool)`

 SetMiddleClickNil sets the value for MiddleClick to be an explicit nil

### UnsetMiddleClick
`func (o *MRBetaBrowserToolsetConfigs) UnsetMiddleClick()`

UnsetMiddleClick ensures that no value is present for MiddleClick, not even an explicit nil
### GetMouseMove

`func (o *MRBetaBrowserToolsetConfigs) GetMouseMove() MRBetaBrowserMouseMoveConfig`

GetMouseMove returns the MouseMove field if non-nil, zero value otherwise.

### GetMouseMoveOk

`func (o *MRBetaBrowserToolsetConfigs) GetMouseMoveOk() (*MRBetaBrowserMouseMoveConfig, bool)`

GetMouseMoveOk returns a tuple with the MouseMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseMove

`func (o *MRBetaBrowserToolsetConfigs) SetMouseMove(v MRBetaBrowserMouseMoveConfig)`

SetMouseMove sets MouseMove field to given value.

### HasMouseMove

`func (o *MRBetaBrowserToolsetConfigs) HasMouseMove() bool`

HasMouseMove returns a boolean if a field has been set.

### SetMouseMoveNil

`func (o *MRBetaBrowserToolsetConfigs) SetMouseMoveNil(b bool)`

 SetMouseMoveNil sets the value for MouseMove to be an explicit nil

### UnsetMouseMove
`func (o *MRBetaBrowserToolsetConfigs) UnsetMouseMove()`

UnsetMouseMove ensures that no value is present for MouseMove, not even an explicit nil
### GetNavigate

`func (o *MRBetaBrowserToolsetConfigs) GetNavigate() MRBetaBrowserNavigateConfig`

GetNavigate returns the Navigate field if non-nil, zero value otherwise.

### GetNavigateOk

`func (o *MRBetaBrowserToolsetConfigs) GetNavigateOk() (*MRBetaBrowserNavigateConfig, bool)`

GetNavigateOk returns a tuple with the Navigate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigate

`func (o *MRBetaBrowserToolsetConfigs) SetNavigate(v MRBetaBrowserNavigateConfig)`

SetNavigate sets Navigate field to given value.

### HasNavigate

`func (o *MRBetaBrowserToolsetConfigs) HasNavigate() bool`

HasNavigate returns a boolean if a field has been set.

### SetNavigateNil

`func (o *MRBetaBrowserToolsetConfigs) SetNavigateNil(b bool)`

 SetNavigateNil sets the value for Navigate to be an explicit nil

### UnsetNavigate
`func (o *MRBetaBrowserToolsetConfigs) UnsetNavigate()`

UnsetNavigate ensures that no value is present for Navigate, not even an explicit nil
### GetNewTab

`func (o *MRBetaBrowserToolsetConfigs) GetNewTab() MRBetaBrowserNewTabConfig`

GetNewTab returns the NewTab field if non-nil, zero value otherwise.

### GetNewTabOk

`func (o *MRBetaBrowserToolsetConfigs) GetNewTabOk() (*MRBetaBrowserNewTabConfig, bool)`

GetNewTabOk returns a tuple with the NewTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTab

`func (o *MRBetaBrowserToolsetConfigs) SetNewTab(v MRBetaBrowserNewTabConfig)`

SetNewTab sets NewTab field to given value.

### HasNewTab

`func (o *MRBetaBrowserToolsetConfigs) HasNewTab() bool`

HasNewTab returns a boolean if a field has been set.

### SetNewTabNil

`func (o *MRBetaBrowserToolsetConfigs) SetNewTabNil(b bool)`

 SetNewTabNil sets the value for NewTab to be an explicit nil

### UnsetNewTab
`func (o *MRBetaBrowserToolsetConfigs) UnsetNewTab()`

UnsetNewTab ensures that no value is present for NewTab, not even an explicit nil
### GetReadConsole

`func (o *MRBetaBrowserToolsetConfigs) GetReadConsole() MRBetaBrowserReadConsoleConfig`

GetReadConsole returns the ReadConsole field if non-nil, zero value otherwise.

### GetReadConsoleOk

`func (o *MRBetaBrowserToolsetConfigs) GetReadConsoleOk() (*MRBetaBrowserReadConsoleConfig, bool)`

GetReadConsoleOk returns a tuple with the ReadConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadConsole

`func (o *MRBetaBrowserToolsetConfigs) SetReadConsole(v MRBetaBrowserReadConsoleConfig)`

SetReadConsole sets ReadConsole field to given value.

### HasReadConsole

`func (o *MRBetaBrowserToolsetConfigs) HasReadConsole() bool`

HasReadConsole returns a boolean if a field has been set.

### SetReadConsoleNil

`func (o *MRBetaBrowserToolsetConfigs) SetReadConsoleNil(b bool)`

 SetReadConsoleNil sets the value for ReadConsole to be an explicit nil

### UnsetReadConsole
`func (o *MRBetaBrowserToolsetConfigs) UnsetReadConsole()`

UnsetReadConsole ensures that no value is present for ReadConsole, not even an explicit nil
### GetReadNetwork

`func (o *MRBetaBrowserToolsetConfigs) GetReadNetwork() MRBetaBrowserReadNetworkConfig`

GetReadNetwork returns the ReadNetwork field if non-nil, zero value otherwise.

### GetReadNetworkOk

`func (o *MRBetaBrowserToolsetConfigs) GetReadNetworkOk() (*MRBetaBrowserReadNetworkConfig, bool)`

GetReadNetworkOk returns a tuple with the ReadNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadNetwork

`func (o *MRBetaBrowserToolsetConfigs) SetReadNetwork(v MRBetaBrowserReadNetworkConfig)`

SetReadNetwork sets ReadNetwork field to given value.

### HasReadNetwork

`func (o *MRBetaBrowserToolsetConfigs) HasReadNetwork() bool`

HasReadNetwork returns a boolean if a field has been set.

### SetReadNetworkNil

`func (o *MRBetaBrowserToolsetConfigs) SetReadNetworkNil(b bool)`

 SetReadNetworkNil sets the value for ReadNetwork to be an explicit nil

### UnsetReadNetwork
`func (o *MRBetaBrowserToolsetConfigs) UnsetReadNetwork()`

UnsetReadNetwork ensures that no value is present for ReadNetwork, not even an explicit nil
### GetReadPage

`func (o *MRBetaBrowserToolsetConfigs) GetReadPage() MRBetaBrowserReadPageConfig`

GetReadPage returns the ReadPage field if non-nil, zero value otherwise.

### GetReadPageOk

`func (o *MRBetaBrowserToolsetConfigs) GetReadPageOk() (*MRBetaBrowserReadPageConfig, bool)`

GetReadPageOk returns a tuple with the ReadPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPage

`func (o *MRBetaBrowserToolsetConfigs) SetReadPage(v MRBetaBrowserReadPageConfig)`

SetReadPage sets ReadPage field to given value.

### HasReadPage

`func (o *MRBetaBrowserToolsetConfigs) HasReadPage() bool`

HasReadPage returns a boolean if a field has been set.

### SetReadPageNil

`func (o *MRBetaBrowserToolsetConfigs) SetReadPageNil(b bool)`

 SetReadPageNil sets the value for ReadPage to be an explicit nil

### UnsetReadPage
`func (o *MRBetaBrowserToolsetConfigs) UnsetReadPage()`

UnsetReadPage ensures that no value is present for ReadPage, not even an explicit nil
### GetRightClick

`func (o *MRBetaBrowserToolsetConfigs) GetRightClick() MRBetaBrowserRightClickConfig`

GetRightClick returns the RightClick field if non-nil, zero value otherwise.

### GetRightClickOk

`func (o *MRBetaBrowserToolsetConfigs) GetRightClickOk() (*MRBetaBrowserRightClickConfig, bool)`

GetRightClickOk returns a tuple with the RightClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightClick

`func (o *MRBetaBrowserToolsetConfigs) SetRightClick(v MRBetaBrowserRightClickConfig)`

SetRightClick sets RightClick field to given value.

### HasRightClick

`func (o *MRBetaBrowserToolsetConfigs) HasRightClick() bool`

HasRightClick returns a boolean if a field has been set.

### SetRightClickNil

`func (o *MRBetaBrowserToolsetConfigs) SetRightClickNil(b bool)`

 SetRightClickNil sets the value for RightClick to be an explicit nil

### UnsetRightClick
`func (o *MRBetaBrowserToolsetConfigs) UnsetRightClick()`

UnsetRightClick ensures that no value is present for RightClick, not even an explicit nil
### GetScreenshot

`func (o *MRBetaBrowserToolsetConfigs) GetScreenshot() MRBetaBrowserScreenshotConfig`

GetScreenshot returns the Screenshot field if non-nil, zero value otherwise.

### GetScreenshotOk

`func (o *MRBetaBrowserToolsetConfigs) GetScreenshotOk() (*MRBetaBrowserScreenshotConfig, bool)`

GetScreenshotOk returns a tuple with the Screenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshot

`func (o *MRBetaBrowserToolsetConfigs) SetScreenshot(v MRBetaBrowserScreenshotConfig)`

SetScreenshot sets Screenshot field to given value.

### HasScreenshot

`func (o *MRBetaBrowserToolsetConfigs) HasScreenshot() bool`

HasScreenshot returns a boolean if a field has been set.

### SetScreenshotNil

`func (o *MRBetaBrowserToolsetConfigs) SetScreenshotNil(b bool)`

 SetScreenshotNil sets the value for Screenshot to be an explicit nil

### UnsetScreenshot
`func (o *MRBetaBrowserToolsetConfigs) UnsetScreenshot()`

UnsetScreenshot ensures that no value is present for Screenshot, not even an explicit nil
### GetScroll

`func (o *MRBetaBrowserToolsetConfigs) GetScroll() MRBetaBrowserScrollConfig`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *MRBetaBrowserToolsetConfigs) GetScrollOk() (*MRBetaBrowserScrollConfig, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *MRBetaBrowserToolsetConfigs) SetScroll(v MRBetaBrowserScrollConfig)`

SetScroll sets Scroll field to given value.

### HasScroll

`func (o *MRBetaBrowserToolsetConfigs) HasScroll() bool`

HasScroll returns a boolean if a field has been set.

### SetScrollNil

`func (o *MRBetaBrowserToolsetConfigs) SetScrollNil(b bool)`

 SetScrollNil sets the value for Scroll to be an explicit nil

### UnsetScroll
`func (o *MRBetaBrowserToolsetConfigs) UnsetScroll()`

UnsetScroll ensures that no value is present for Scroll, not even an explicit nil
### GetScrollTo

`func (o *MRBetaBrowserToolsetConfigs) GetScrollTo() MRBetaBrowserScrollToConfig`

GetScrollTo returns the ScrollTo field if non-nil, zero value otherwise.

### GetScrollToOk

`func (o *MRBetaBrowserToolsetConfigs) GetScrollToOk() (*MRBetaBrowserScrollToConfig, bool)`

GetScrollToOk returns a tuple with the ScrollTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollTo

`func (o *MRBetaBrowserToolsetConfigs) SetScrollTo(v MRBetaBrowserScrollToConfig)`

SetScrollTo sets ScrollTo field to given value.

### HasScrollTo

`func (o *MRBetaBrowserToolsetConfigs) HasScrollTo() bool`

HasScrollTo returns a boolean if a field has been set.

### SetScrollToNil

`func (o *MRBetaBrowserToolsetConfigs) SetScrollToNil(b bool)`

 SetScrollToNil sets the value for ScrollTo to be an explicit nil

### UnsetScrollTo
`func (o *MRBetaBrowserToolsetConfigs) UnsetScrollTo()`

UnsetScrollTo ensures that no value is present for ScrollTo, not even an explicit nil
### GetSwitchTab

`func (o *MRBetaBrowserToolsetConfigs) GetSwitchTab() MRBetaBrowserSwitchTabConfig`

GetSwitchTab returns the SwitchTab field if non-nil, zero value otherwise.

### GetSwitchTabOk

`func (o *MRBetaBrowserToolsetConfigs) GetSwitchTabOk() (*MRBetaBrowserSwitchTabConfig, bool)`

GetSwitchTabOk returns a tuple with the SwitchTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchTab

`func (o *MRBetaBrowserToolsetConfigs) SetSwitchTab(v MRBetaBrowserSwitchTabConfig)`

SetSwitchTab sets SwitchTab field to given value.

### HasSwitchTab

`func (o *MRBetaBrowserToolsetConfigs) HasSwitchTab() bool`

HasSwitchTab returns a boolean if a field has been set.

### SetSwitchTabNil

`func (o *MRBetaBrowserToolsetConfigs) SetSwitchTabNil(b bool)`

 SetSwitchTabNil sets the value for SwitchTab to be an explicit nil

### UnsetSwitchTab
`func (o *MRBetaBrowserToolsetConfigs) UnsetSwitchTab()`

UnsetSwitchTab ensures that no value is present for SwitchTab, not even an explicit nil
### GetTripleClick

`func (o *MRBetaBrowserToolsetConfigs) GetTripleClick() MRBetaBrowserTripleClickConfig`

GetTripleClick returns the TripleClick field if non-nil, zero value otherwise.

### GetTripleClickOk

`func (o *MRBetaBrowserToolsetConfigs) GetTripleClickOk() (*MRBetaBrowserTripleClickConfig, bool)`

GetTripleClickOk returns a tuple with the TripleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripleClick

`func (o *MRBetaBrowserToolsetConfigs) SetTripleClick(v MRBetaBrowserTripleClickConfig)`

SetTripleClick sets TripleClick field to given value.

### HasTripleClick

`func (o *MRBetaBrowserToolsetConfigs) HasTripleClick() bool`

HasTripleClick returns a boolean if a field has been set.

### SetTripleClickNil

`func (o *MRBetaBrowserToolsetConfigs) SetTripleClickNil(b bool)`

 SetTripleClickNil sets the value for TripleClick to be an explicit nil

### UnsetTripleClick
`func (o *MRBetaBrowserToolsetConfigs) UnsetTripleClick()`

UnsetTripleClick ensures that no value is present for TripleClick, not even an explicit nil
### GetType

`func (o *MRBetaBrowserToolsetConfigs) GetType() MRBetaBrowserTypeConfig`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaBrowserToolsetConfigs) GetTypeOk() (*MRBetaBrowserTypeConfig, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaBrowserToolsetConfigs) SetType(v MRBetaBrowserTypeConfig)`

SetType sets Type field to given value.

### HasType

`func (o *MRBetaBrowserToolsetConfigs) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MRBetaBrowserToolsetConfigs) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MRBetaBrowserToolsetConfigs) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWait

`func (o *MRBetaBrowserToolsetConfigs) GetWait() MRBetaBrowserWaitConfig`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *MRBetaBrowserToolsetConfigs) GetWaitOk() (*MRBetaBrowserWaitConfig, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *MRBetaBrowserToolsetConfigs) SetWait(v MRBetaBrowserWaitConfig)`

SetWait sets Wait field to given value.

### HasWait

`func (o *MRBetaBrowserToolsetConfigs) HasWait() bool`

HasWait returns a boolean if a field has been set.

### SetWaitNil

`func (o *MRBetaBrowserToolsetConfigs) SetWaitNil(b bool)`

 SetWaitNil sets the value for Wait to be an explicit nil

### UnsetWait
`func (o *MRBetaBrowserToolsetConfigs) UnsetWait()`

UnsetWait ensures that no value is present for Wait, not even an explicit nil
### GetZoom

`func (o *MRBetaBrowserToolsetConfigs) GetZoom() MRBetaBrowserZoomConfig`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *MRBetaBrowserToolsetConfigs) GetZoomOk() (*MRBetaBrowserZoomConfig, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *MRBetaBrowserToolsetConfigs) SetZoom(v MRBetaBrowserZoomConfig)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *MRBetaBrowserToolsetConfigs) HasZoom() bool`

HasZoom returns a boolean if a field has been set.

### SetZoomNil

`func (o *MRBetaBrowserToolsetConfigs) SetZoomNil(b bool)`

 SetZoomNil sets the value for Zoom to be an explicit nil

### UnsetZoom
`func (o *MRBetaBrowserToolsetConfigs) UnsetZoom()`

UnsetZoom ensures that no value is present for Zoom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


