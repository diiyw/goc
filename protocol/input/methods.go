package input

const (
	
	// Dispatches a key event to the page.
	DispatchKeyEvent = "Input.dispatchKeyEvent"
	
	// This method emulates inserting text that doesn't come from a key press,
	// for example an emoji keyboard or an IME.
	InsertText = "Input.insertText"
	
	// Dispatches a mouse event to the page.
	DispatchMouseEvent = "Input.dispatchMouseEvent"
	
	// Dispatches a touch event to the page.
	DispatchTouchEvent = "Input.dispatchTouchEvent"
	
	// Emulates touch event from the mouse event parameters.
	EmulateTouchFromMouseEvent = "Input.emulateTouchFromMouseEvent"
	
	// Ignores input events (useful while auditing page).
	SetIgnoreInputEvents = "Input.setIgnoreInputEvents"
	
	// Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
	SynthesizePinchGesture = "Input.synthesizePinchGesture"
	
	// Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
	SynthesizeScrollGesture = "Input.synthesizeScrollGesture"
	
	// Synthesizes a tap gesture over a time period by issuing appropriate touch events.
	SynthesizeTapGesture = "Input.synthesizeTapGesture"
	
)

// DispatchKeyEvent parameters
type DispatchKeyEventParams struct {
	
	// Type of the key event.
	Type	string	`json:"type"`
	
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8
	// (default: 0).
	Modifiers	int	`json:"modifiers"`
	
	// Time at which the event occurred.
	Timestamp	TimeSinceEpoch	`json:"timestamp"`
	
	// Text as generated by processing a virtual key code with a keyboard layout. Not needed for
	// for `keyUp` and `rawKeyDown` events (default: "")
	Text	string	`json:"text"`
	
	// Text that would have been generated by the keyboard if no modifiers were pressed (except for
	// shift). Useful for shortcut (accelerator) key handling (default: "").
	UnmodifiedText	string	`json:"unmodifiedText"`
	
	// Unique key identifier (e.g., 'U+0041') (default: "").
	KeyIdentifier	string	`json:"keyIdentifier"`
	
	// Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Code	string	`json:"code"`
	
	// Unique DOM defined string value describing the meaning of the key in the context of active
	// modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	Key	string	`json:"key"`
	
	// Windows virtual key code (default: 0).
	WindowsVirtualKeyCode	int	`json:"windowsVirtualKeyCode"`
	
	// Native virtual key code (default: 0).
	NativeVirtualKeyCode	int	`json:"nativeVirtualKeyCode"`
	
	// Whether the event was generated from auto repeat (default: false).
	AutoRepeat	bool	`json:"autoRepeat"`
	
	// Whether the event was generated from the keypad (default: false).
	IsKeypad	bool	`json:"isKeypad"`
	
	// Whether the event was a system key event (default: false).
	IsSystemKey	bool	`json:"isSystemKey"`
	
	// Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default:
	// 0).
	Location	int	`json:"location"`
	
}

// DispatchKeyEvent returns
type DispatchKeyEventReturns struct {
	
}

// InsertText parameters
type InsertTextParams struct {
	
	// The text to insert.
	Text	string	`json:"text"`
	
}

// InsertText returns
type InsertTextReturns struct {
	
}

// DispatchMouseEvent parameters
type DispatchMouseEventParams struct {
	
	// Type of the mouse event.
	Type	string	`json:"type"`
	
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X	float64	`json:"x"`
	
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to
	// the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y	float64	`json:"y"`
	
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8
	// (default: 0).
	Modifiers	int	`json:"modifiers"`
	
	// Time at which the event occurred.
	Timestamp	TimeSinceEpoch	`json:"timestamp"`
	
	// Mouse button (default: "none").
	Button	string	`json:"button"`
	
	// A number indicating which buttons are pressed on the mouse when a mouse event is triggered.
	// Left=1, Right=2, Middle=4, Back=8, Forward=16, None=0.
	Buttons	int	`json:"buttons"`
	
	// Number of times the mouse button was clicked (default: 0).
	ClickCount	int	`json:"clickCount"`
	
	// X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaX	float64	`json:"deltaX"`
	
	// Y delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY	float64	`json:"deltaY"`
	
}

// DispatchMouseEvent returns
type DispatchMouseEventReturns struct {
	
}

// DispatchTouchEvent parameters
type DispatchTouchEventParams struct {
	
	// Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while
	// TouchStart and TouchMove must contains at least one.
	Type	string	`json:"type"`
	
	// Active touch points on the touch device. One event per any changed point (compared to
	// previous touch event in a sequence) is generated, emulating pressing/moving/releasing points
	// one by one.
	TouchPoints	[]TouchPoint	`json:"touchPoints"`
	
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8
	// (default: 0).
	Modifiers	int	`json:"modifiers"`
	
	// Time at which the event occurred.
	Timestamp	TimeSinceEpoch	`json:"timestamp"`
	
}

// DispatchTouchEvent returns
type DispatchTouchEventReturns struct {
	
}

// EmulateTouchFromMouseEvent parameters
type EmulateTouchFromMouseEventParams struct {
	
	// Type of the mouse event.
	Type	string	`json:"type"`
	
	// X coordinate of the mouse pointer in DIP.
	X	int	`json:"x"`
	
	// Y coordinate of the mouse pointer in DIP.
	Y	int	`json:"y"`
	
	// Mouse button.
	Button	string	`json:"button"`
	
	// Time at which the event occurred (default: current time).
	Timestamp	TimeSinceEpoch	`json:"timestamp"`
	
	// X delta in DIP for mouse wheel event (default: 0).
	DeltaX	float64	`json:"deltaX"`
	
	// Y delta in DIP for mouse wheel event (default: 0).
	DeltaY	float64	`json:"deltaY"`
	
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8
	// (default: 0).
	Modifiers	int	`json:"modifiers"`
	
	// Number of times the mouse button was clicked (default: 0).
	ClickCount	int	`json:"clickCount"`
	
}

// EmulateTouchFromMouseEvent returns
type EmulateTouchFromMouseEventReturns struct {
	
}

// SetIgnoreInputEvents parameters
type SetIgnoreInputEventsParams struct {
	
	// Ignores input events processing when set to true.
	Ignore	bool	`json:"ignore"`
	
}

// SetIgnoreInputEvents returns
type SetIgnoreInputEventsReturns struct {
	
}

// SynthesizePinchGesture parameters
type SynthesizePinchGestureParams struct {
	
	// X coordinate of the start of the gesture in CSS pixels.
	X	float64	`json:"x"`
	
	// Y coordinate of the start of the gesture in CSS pixels.
	Y	float64	`json:"y"`
	
	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor	float64	`json:"scaleFactor"`
	
	// Relative pointer speed in pixels per second (default: 800).
	RelativeSpeed	int	`json:"relativeSpeed"`
	
	// Which type of input events to be generated (default: 'default', which queries the platform
	// for the preferred input type).
	GestureSourceType	GestureSourceType	`json:"gestureSourceType"`
	
}

// SynthesizePinchGesture returns
type SynthesizePinchGestureReturns struct {
	
}

// SynthesizeScrollGesture parameters
type SynthesizeScrollGestureParams struct {
	
	// X coordinate of the start of the gesture in CSS pixels.
	X	float64	`json:"x"`
	
	// Y coordinate of the start of the gesture in CSS pixels.
	Y	float64	`json:"y"`
	
	// The distance to scroll along the X axis (positive to scroll left).
	XDistance	float64	`json:"xDistance"`
	
	// The distance to scroll along the Y axis (positive to scroll up).
	YDistance	float64	`json:"yDistance"`
	
	// The number of additional pixels to scroll back along the X axis, in addition to the given
	// distance.
	XOverscroll	float64	`json:"xOverscroll"`
	
	// The number of additional pixels to scroll back along the Y axis, in addition to the given
	// distance.
	YOverscroll	float64	`json:"yOverscroll"`
	
	// Prevent fling (default: true).
	PreventFling	bool	`json:"preventFling"`
	
	// Swipe speed in pixels per second (default: 800).
	Speed	int	`json:"speed"`
	
	// Which type of input events to be generated (default: 'default', which queries the platform
	// for the preferred input type).
	GestureSourceType	GestureSourceType	`json:"gestureSourceType"`
	
	// The number of times to repeat the gesture (default: 0).
	RepeatCount	int	`json:"repeatCount"`
	
	// The number of milliseconds delay between each repeat. (default: 250).
	RepeatDelayMs	int	`json:"repeatDelayMs"`
	
	// The name of the interaction markers to generate, if not empty (default: "").
	InteractionMarkerName	string	`json:"interactionMarkerName"`
	
}

// SynthesizeScrollGesture returns
type SynthesizeScrollGestureReturns struct {
	
}

// SynthesizeTapGesture parameters
type SynthesizeTapGestureParams struct {
	
	// X coordinate of the start of the gesture in CSS pixels.
	X	float64	`json:"x"`
	
	// Y coordinate of the start of the gesture in CSS pixels.
	Y	float64	`json:"y"`
	
	// Duration between touchdown and touchup events in ms (default: 50).
	Duration	int	`json:"duration"`
	
	// Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	TapCount	int	`json:"tapCount"`
	
	// Which type of input events to be generated (default: 'default', which queries the platform
	// for the preferred input type).
	GestureSourceType	GestureSourceType	`json:"gestureSourceType"`
	
}

// SynthesizeTapGesture returns
type SynthesizeTapGestureReturns struct {
	
}

