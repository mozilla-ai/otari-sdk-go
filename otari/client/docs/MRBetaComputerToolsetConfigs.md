# MRBetaComputerToolsetConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CursorPosition** | Pointer to [**NullableMRBetaComputerCursorPositionConfig**](MRBetaComputerCursorPositionConfig.md) |  | [optional] 
**DoubleClick** | Pointer to [**NullableMRBetaComputerDoubleClickConfig**](MRBetaComputerDoubleClickConfig.md) |  | [optional] 
**HoldKey** | Pointer to [**NullableMRBetaComputerHoldKeyConfig**](MRBetaComputerHoldKeyConfig.md) |  | [optional] 
**Key** | Pointer to [**NullableMRBetaComputerKeyConfig**](MRBetaComputerKeyConfig.md) |  | [optional] 
**LeftClick** | Pointer to [**NullableMRBetaComputerLeftClickConfig**](MRBetaComputerLeftClickConfig.md) |  | [optional] 
**LeftClickDrag** | Pointer to [**NullableMRBetaComputerLeftClickDragConfig**](MRBetaComputerLeftClickDragConfig.md) |  | [optional] 
**LeftMouseDown** | Pointer to [**NullableMRBetaComputerLeftMouseDownConfig**](MRBetaComputerLeftMouseDownConfig.md) |  | [optional] 
**LeftMouseUp** | Pointer to [**NullableMRBetaComputerLeftMouseUpConfig**](MRBetaComputerLeftMouseUpConfig.md) |  | [optional] 
**MiddleClick** | Pointer to [**NullableMRBetaComputerMiddleClickConfig**](MRBetaComputerMiddleClickConfig.md) |  | [optional] 
**MouseMove** | Pointer to [**NullableMRBetaComputerMouseMoveConfig**](MRBetaComputerMouseMoveConfig.md) |  | [optional] 
**RightClick** | Pointer to [**NullableMRBetaComputerRightClickConfig**](MRBetaComputerRightClickConfig.md) |  | [optional] 
**Screenshot** | Pointer to [**NullableMRBetaComputerScreenshotConfig**](MRBetaComputerScreenshotConfig.md) |  | [optional] 
**Scroll** | Pointer to [**NullableMRBetaComputerScrollConfig**](MRBetaComputerScrollConfig.md) |  | [optional] 
**TripleClick** | Pointer to [**NullableMRBetaComputerTripleClickConfig**](MRBetaComputerTripleClickConfig.md) |  | [optional] 
**Type** | Pointer to [**NullableMRBetaComputerTypeConfig**](MRBetaComputerTypeConfig.md) |  | [optional] 
**Wait** | Pointer to [**NullableMRBetaComputerWaitConfig**](MRBetaComputerWaitConfig.md) |  | [optional] 
**Zoom** | Pointer to [**NullableMRBetaComputerZoomConfig**](MRBetaComputerZoomConfig.md) |  | [optional] 

## Methods

### NewMRBetaComputerToolsetConfigs

`func NewMRBetaComputerToolsetConfigs() *MRBetaComputerToolsetConfigs`

NewMRBetaComputerToolsetConfigs instantiates a new MRBetaComputerToolsetConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaComputerToolsetConfigsWithDefaults

`func NewMRBetaComputerToolsetConfigsWithDefaults() *MRBetaComputerToolsetConfigs`

NewMRBetaComputerToolsetConfigsWithDefaults instantiates a new MRBetaComputerToolsetConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursorPosition

`func (o *MRBetaComputerToolsetConfigs) GetCursorPosition() MRBetaComputerCursorPositionConfig`

GetCursorPosition returns the CursorPosition field if non-nil, zero value otherwise.

### GetCursorPositionOk

`func (o *MRBetaComputerToolsetConfigs) GetCursorPositionOk() (*MRBetaComputerCursorPositionConfig, bool)`

GetCursorPositionOk returns a tuple with the CursorPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorPosition

`func (o *MRBetaComputerToolsetConfigs) SetCursorPosition(v MRBetaComputerCursorPositionConfig)`

SetCursorPosition sets CursorPosition field to given value.

### HasCursorPosition

`func (o *MRBetaComputerToolsetConfigs) HasCursorPosition() bool`

HasCursorPosition returns a boolean if a field has been set.

### SetCursorPositionNil

`func (o *MRBetaComputerToolsetConfigs) SetCursorPositionNil(b bool)`

 SetCursorPositionNil sets the value for CursorPosition to be an explicit nil

### UnsetCursorPosition
`func (o *MRBetaComputerToolsetConfigs) UnsetCursorPosition()`

UnsetCursorPosition ensures that no value is present for CursorPosition, not even an explicit nil
### GetDoubleClick

`func (o *MRBetaComputerToolsetConfigs) GetDoubleClick() MRBetaComputerDoubleClickConfig`

GetDoubleClick returns the DoubleClick field if non-nil, zero value otherwise.

### GetDoubleClickOk

`func (o *MRBetaComputerToolsetConfigs) GetDoubleClickOk() (*MRBetaComputerDoubleClickConfig, bool)`

GetDoubleClickOk returns a tuple with the DoubleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleClick

`func (o *MRBetaComputerToolsetConfigs) SetDoubleClick(v MRBetaComputerDoubleClickConfig)`

SetDoubleClick sets DoubleClick field to given value.

### HasDoubleClick

`func (o *MRBetaComputerToolsetConfigs) HasDoubleClick() bool`

HasDoubleClick returns a boolean if a field has been set.

### SetDoubleClickNil

`func (o *MRBetaComputerToolsetConfigs) SetDoubleClickNil(b bool)`

 SetDoubleClickNil sets the value for DoubleClick to be an explicit nil

### UnsetDoubleClick
`func (o *MRBetaComputerToolsetConfigs) UnsetDoubleClick()`

UnsetDoubleClick ensures that no value is present for DoubleClick, not even an explicit nil
### GetHoldKey

`func (o *MRBetaComputerToolsetConfigs) GetHoldKey() MRBetaComputerHoldKeyConfig`

GetHoldKey returns the HoldKey field if non-nil, zero value otherwise.

### GetHoldKeyOk

`func (o *MRBetaComputerToolsetConfigs) GetHoldKeyOk() (*MRBetaComputerHoldKeyConfig, bool)`

GetHoldKeyOk returns a tuple with the HoldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldKey

`func (o *MRBetaComputerToolsetConfigs) SetHoldKey(v MRBetaComputerHoldKeyConfig)`

SetHoldKey sets HoldKey field to given value.

### HasHoldKey

`func (o *MRBetaComputerToolsetConfigs) HasHoldKey() bool`

HasHoldKey returns a boolean if a field has been set.

### SetHoldKeyNil

`func (o *MRBetaComputerToolsetConfigs) SetHoldKeyNil(b bool)`

 SetHoldKeyNil sets the value for HoldKey to be an explicit nil

### UnsetHoldKey
`func (o *MRBetaComputerToolsetConfigs) UnsetHoldKey()`

UnsetHoldKey ensures that no value is present for HoldKey, not even an explicit nil
### GetKey

`func (o *MRBetaComputerToolsetConfigs) GetKey() MRBetaComputerKeyConfig`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MRBetaComputerToolsetConfigs) GetKeyOk() (*MRBetaComputerKeyConfig, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MRBetaComputerToolsetConfigs) SetKey(v MRBetaComputerKeyConfig)`

SetKey sets Key field to given value.

### HasKey

`func (o *MRBetaComputerToolsetConfigs) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MRBetaComputerToolsetConfigs) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MRBetaComputerToolsetConfigs) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetLeftClick

`func (o *MRBetaComputerToolsetConfigs) GetLeftClick() MRBetaComputerLeftClickConfig`

GetLeftClick returns the LeftClick field if non-nil, zero value otherwise.

### GetLeftClickOk

`func (o *MRBetaComputerToolsetConfigs) GetLeftClickOk() (*MRBetaComputerLeftClickConfig, bool)`

GetLeftClickOk returns a tuple with the LeftClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftClick

`func (o *MRBetaComputerToolsetConfigs) SetLeftClick(v MRBetaComputerLeftClickConfig)`

SetLeftClick sets LeftClick field to given value.

### HasLeftClick

`func (o *MRBetaComputerToolsetConfigs) HasLeftClick() bool`

HasLeftClick returns a boolean if a field has been set.

### SetLeftClickNil

`func (o *MRBetaComputerToolsetConfigs) SetLeftClickNil(b bool)`

 SetLeftClickNil sets the value for LeftClick to be an explicit nil

### UnsetLeftClick
`func (o *MRBetaComputerToolsetConfigs) UnsetLeftClick()`

UnsetLeftClick ensures that no value is present for LeftClick, not even an explicit nil
### GetLeftClickDrag

`func (o *MRBetaComputerToolsetConfigs) GetLeftClickDrag() MRBetaComputerLeftClickDragConfig`

GetLeftClickDrag returns the LeftClickDrag field if non-nil, zero value otherwise.

### GetLeftClickDragOk

`func (o *MRBetaComputerToolsetConfigs) GetLeftClickDragOk() (*MRBetaComputerLeftClickDragConfig, bool)`

GetLeftClickDragOk returns a tuple with the LeftClickDrag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftClickDrag

`func (o *MRBetaComputerToolsetConfigs) SetLeftClickDrag(v MRBetaComputerLeftClickDragConfig)`

SetLeftClickDrag sets LeftClickDrag field to given value.

### HasLeftClickDrag

`func (o *MRBetaComputerToolsetConfigs) HasLeftClickDrag() bool`

HasLeftClickDrag returns a boolean if a field has been set.

### SetLeftClickDragNil

`func (o *MRBetaComputerToolsetConfigs) SetLeftClickDragNil(b bool)`

 SetLeftClickDragNil sets the value for LeftClickDrag to be an explicit nil

### UnsetLeftClickDrag
`func (o *MRBetaComputerToolsetConfigs) UnsetLeftClickDrag()`

UnsetLeftClickDrag ensures that no value is present for LeftClickDrag, not even an explicit nil
### GetLeftMouseDown

`func (o *MRBetaComputerToolsetConfigs) GetLeftMouseDown() MRBetaComputerLeftMouseDownConfig`

GetLeftMouseDown returns the LeftMouseDown field if non-nil, zero value otherwise.

### GetLeftMouseDownOk

`func (o *MRBetaComputerToolsetConfigs) GetLeftMouseDownOk() (*MRBetaComputerLeftMouseDownConfig, bool)`

GetLeftMouseDownOk returns a tuple with the LeftMouseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftMouseDown

`func (o *MRBetaComputerToolsetConfigs) SetLeftMouseDown(v MRBetaComputerLeftMouseDownConfig)`

SetLeftMouseDown sets LeftMouseDown field to given value.

### HasLeftMouseDown

`func (o *MRBetaComputerToolsetConfigs) HasLeftMouseDown() bool`

HasLeftMouseDown returns a boolean if a field has been set.

### SetLeftMouseDownNil

`func (o *MRBetaComputerToolsetConfigs) SetLeftMouseDownNil(b bool)`

 SetLeftMouseDownNil sets the value for LeftMouseDown to be an explicit nil

### UnsetLeftMouseDown
`func (o *MRBetaComputerToolsetConfigs) UnsetLeftMouseDown()`

UnsetLeftMouseDown ensures that no value is present for LeftMouseDown, not even an explicit nil
### GetLeftMouseUp

`func (o *MRBetaComputerToolsetConfigs) GetLeftMouseUp() MRBetaComputerLeftMouseUpConfig`

GetLeftMouseUp returns the LeftMouseUp field if non-nil, zero value otherwise.

### GetLeftMouseUpOk

`func (o *MRBetaComputerToolsetConfigs) GetLeftMouseUpOk() (*MRBetaComputerLeftMouseUpConfig, bool)`

GetLeftMouseUpOk returns a tuple with the LeftMouseUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftMouseUp

`func (o *MRBetaComputerToolsetConfigs) SetLeftMouseUp(v MRBetaComputerLeftMouseUpConfig)`

SetLeftMouseUp sets LeftMouseUp field to given value.

### HasLeftMouseUp

`func (o *MRBetaComputerToolsetConfigs) HasLeftMouseUp() bool`

HasLeftMouseUp returns a boolean if a field has been set.

### SetLeftMouseUpNil

`func (o *MRBetaComputerToolsetConfigs) SetLeftMouseUpNil(b bool)`

 SetLeftMouseUpNil sets the value for LeftMouseUp to be an explicit nil

### UnsetLeftMouseUp
`func (o *MRBetaComputerToolsetConfigs) UnsetLeftMouseUp()`

UnsetLeftMouseUp ensures that no value is present for LeftMouseUp, not even an explicit nil
### GetMiddleClick

`func (o *MRBetaComputerToolsetConfigs) GetMiddleClick() MRBetaComputerMiddleClickConfig`

GetMiddleClick returns the MiddleClick field if non-nil, zero value otherwise.

### GetMiddleClickOk

`func (o *MRBetaComputerToolsetConfigs) GetMiddleClickOk() (*MRBetaComputerMiddleClickConfig, bool)`

GetMiddleClickOk returns a tuple with the MiddleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleClick

`func (o *MRBetaComputerToolsetConfigs) SetMiddleClick(v MRBetaComputerMiddleClickConfig)`

SetMiddleClick sets MiddleClick field to given value.

### HasMiddleClick

`func (o *MRBetaComputerToolsetConfigs) HasMiddleClick() bool`

HasMiddleClick returns a boolean if a field has been set.

### SetMiddleClickNil

`func (o *MRBetaComputerToolsetConfigs) SetMiddleClickNil(b bool)`

 SetMiddleClickNil sets the value for MiddleClick to be an explicit nil

### UnsetMiddleClick
`func (o *MRBetaComputerToolsetConfigs) UnsetMiddleClick()`

UnsetMiddleClick ensures that no value is present for MiddleClick, not even an explicit nil
### GetMouseMove

`func (o *MRBetaComputerToolsetConfigs) GetMouseMove() MRBetaComputerMouseMoveConfig`

GetMouseMove returns the MouseMove field if non-nil, zero value otherwise.

### GetMouseMoveOk

`func (o *MRBetaComputerToolsetConfigs) GetMouseMoveOk() (*MRBetaComputerMouseMoveConfig, bool)`

GetMouseMoveOk returns a tuple with the MouseMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseMove

`func (o *MRBetaComputerToolsetConfigs) SetMouseMove(v MRBetaComputerMouseMoveConfig)`

SetMouseMove sets MouseMove field to given value.

### HasMouseMove

`func (o *MRBetaComputerToolsetConfigs) HasMouseMove() bool`

HasMouseMove returns a boolean if a field has been set.

### SetMouseMoveNil

`func (o *MRBetaComputerToolsetConfigs) SetMouseMoveNil(b bool)`

 SetMouseMoveNil sets the value for MouseMove to be an explicit nil

### UnsetMouseMove
`func (o *MRBetaComputerToolsetConfigs) UnsetMouseMove()`

UnsetMouseMove ensures that no value is present for MouseMove, not even an explicit nil
### GetRightClick

`func (o *MRBetaComputerToolsetConfigs) GetRightClick() MRBetaComputerRightClickConfig`

GetRightClick returns the RightClick field if non-nil, zero value otherwise.

### GetRightClickOk

`func (o *MRBetaComputerToolsetConfigs) GetRightClickOk() (*MRBetaComputerRightClickConfig, bool)`

GetRightClickOk returns a tuple with the RightClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightClick

`func (o *MRBetaComputerToolsetConfigs) SetRightClick(v MRBetaComputerRightClickConfig)`

SetRightClick sets RightClick field to given value.

### HasRightClick

`func (o *MRBetaComputerToolsetConfigs) HasRightClick() bool`

HasRightClick returns a boolean if a field has been set.

### SetRightClickNil

`func (o *MRBetaComputerToolsetConfigs) SetRightClickNil(b bool)`

 SetRightClickNil sets the value for RightClick to be an explicit nil

### UnsetRightClick
`func (o *MRBetaComputerToolsetConfigs) UnsetRightClick()`

UnsetRightClick ensures that no value is present for RightClick, not even an explicit nil
### GetScreenshot

`func (o *MRBetaComputerToolsetConfigs) GetScreenshot() MRBetaComputerScreenshotConfig`

GetScreenshot returns the Screenshot field if non-nil, zero value otherwise.

### GetScreenshotOk

`func (o *MRBetaComputerToolsetConfigs) GetScreenshotOk() (*MRBetaComputerScreenshotConfig, bool)`

GetScreenshotOk returns a tuple with the Screenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshot

`func (o *MRBetaComputerToolsetConfigs) SetScreenshot(v MRBetaComputerScreenshotConfig)`

SetScreenshot sets Screenshot field to given value.

### HasScreenshot

`func (o *MRBetaComputerToolsetConfigs) HasScreenshot() bool`

HasScreenshot returns a boolean if a field has been set.

### SetScreenshotNil

`func (o *MRBetaComputerToolsetConfigs) SetScreenshotNil(b bool)`

 SetScreenshotNil sets the value for Screenshot to be an explicit nil

### UnsetScreenshot
`func (o *MRBetaComputerToolsetConfigs) UnsetScreenshot()`

UnsetScreenshot ensures that no value is present for Screenshot, not even an explicit nil
### GetScroll

`func (o *MRBetaComputerToolsetConfigs) GetScroll() MRBetaComputerScrollConfig`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *MRBetaComputerToolsetConfigs) GetScrollOk() (*MRBetaComputerScrollConfig, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *MRBetaComputerToolsetConfigs) SetScroll(v MRBetaComputerScrollConfig)`

SetScroll sets Scroll field to given value.

### HasScroll

`func (o *MRBetaComputerToolsetConfigs) HasScroll() bool`

HasScroll returns a boolean if a field has been set.

### SetScrollNil

`func (o *MRBetaComputerToolsetConfigs) SetScrollNil(b bool)`

 SetScrollNil sets the value for Scroll to be an explicit nil

### UnsetScroll
`func (o *MRBetaComputerToolsetConfigs) UnsetScroll()`

UnsetScroll ensures that no value is present for Scroll, not even an explicit nil
### GetTripleClick

`func (o *MRBetaComputerToolsetConfigs) GetTripleClick() MRBetaComputerTripleClickConfig`

GetTripleClick returns the TripleClick field if non-nil, zero value otherwise.

### GetTripleClickOk

`func (o *MRBetaComputerToolsetConfigs) GetTripleClickOk() (*MRBetaComputerTripleClickConfig, bool)`

GetTripleClickOk returns a tuple with the TripleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripleClick

`func (o *MRBetaComputerToolsetConfigs) SetTripleClick(v MRBetaComputerTripleClickConfig)`

SetTripleClick sets TripleClick field to given value.

### HasTripleClick

`func (o *MRBetaComputerToolsetConfigs) HasTripleClick() bool`

HasTripleClick returns a boolean if a field has been set.

### SetTripleClickNil

`func (o *MRBetaComputerToolsetConfigs) SetTripleClickNil(b bool)`

 SetTripleClickNil sets the value for TripleClick to be an explicit nil

### UnsetTripleClick
`func (o *MRBetaComputerToolsetConfigs) UnsetTripleClick()`

UnsetTripleClick ensures that no value is present for TripleClick, not even an explicit nil
### GetType

`func (o *MRBetaComputerToolsetConfigs) GetType() MRBetaComputerTypeConfig`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaComputerToolsetConfigs) GetTypeOk() (*MRBetaComputerTypeConfig, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaComputerToolsetConfigs) SetType(v MRBetaComputerTypeConfig)`

SetType sets Type field to given value.

### HasType

`func (o *MRBetaComputerToolsetConfigs) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MRBetaComputerToolsetConfigs) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MRBetaComputerToolsetConfigs) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWait

`func (o *MRBetaComputerToolsetConfigs) GetWait() MRBetaComputerWaitConfig`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *MRBetaComputerToolsetConfigs) GetWaitOk() (*MRBetaComputerWaitConfig, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *MRBetaComputerToolsetConfigs) SetWait(v MRBetaComputerWaitConfig)`

SetWait sets Wait field to given value.

### HasWait

`func (o *MRBetaComputerToolsetConfigs) HasWait() bool`

HasWait returns a boolean if a field has been set.

### SetWaitNil

`func (o *MRBetaComputerToolsetConfigs) SetWaitNil(b bool)`

 SetWaitNil sets the value for Wait to be an explicit nil

### UnsetWait
`func (o *MRBetaComputerToolsetConfigs) UnsetWait()`

UnsetWait ensures that no value is present for Wait, not even an explicit nil
### GetZoom

`func (o *MRBetaComputerToolsetConfigs) GetZoom() MRBetaComputerZoomConfig`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *MRBetaComputerToolsetConfigs) GetZoomOk() (*MRBetaComputerZoomConfig, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *MRBetaComputerToolsetConfigs) SetZoom(v MRBetaComputerZoomConfig)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *MRBetaComputerToolsetConfigs) HasZoom() bool`

HasZoom returns a boolean if a field has been set.

### SetZoomNil

`func (o *MRBetaComputerToolsetConfigs) SetZoomNil(b bool)`

 SetZoomNil sets the value for Zoom to be an explicit nil

### UnsetZoom
`func (o *MRBetaComputerToolsetConfigs) UnsetZoom()`

UnsetZoom ensures that no value is present for Zoom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


