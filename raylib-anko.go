package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"os"
	"fmt"
	"github.com/mattn/anko/vm"
	"github.com/mattn/anko/env"
	//"reflect"
)

func ChoiceInitWindow(width int32, height int32, title string, args ...interface{}) {
	if len(args) == 0 { InitWindowWithFPS(width, height, title, args[0].(int32))
	} else { InitWindowWithoutFPS(width, height, title) }
	if !rl.IsWindowReady() {
		fmt.Println("ERR - Failed to initialise window")
		os.Exit(0)
	}
}
func InitWindowWithFPS(width int32, height int32, title string, fps int32) {
	rl.InitWindow(width, height, title) ; rl.SetTargetFPS(fps)
}
func InitWindowWithoutFPS(width int32, height int32, title string) {
	rl.InitWindow(width, height, title)
}

func main() {
	script_path := ""

	if len(os.Args) >= 2 { script_path = os.Args[1] } else { fmt.Println("ERR - No script!") ; os.Exit(0) }

	script, err := os.ReadFile(script_path)
	if err != nil { fmt.Println("ERR - Script at '" + script_path + "' not found!") ; os.Exit(0)}

	colors := map[string]rl.Color {
		"LIGHTGRAY": rl.LightGray,
		"GRAY": rl.Gray,
		"DARKGRAY": rl.DarkGray,
		"YELLOW": rl.Yellow,
		"GOLD": rl.Gold,
		"ORANGE": rl.Orange,
		"PINK": rl.Pink,
		"RED": rl.Red,
		"MAROON": rl.Maroon,
		"GREEN": rl.Green,
		"LIME": rl.Lime,
		"DARKGREEN": rl.DarkGreen,
		"SKYBLUE": rl.SkyBlue,
		"BLUE": rl.Blue,
		"DARKBLUE": rl.DarkBlue,
		"PURPLE": rl.Purple,
		"VIOLET": rl.Violet,
		"DARKPURPLE": rl.DarkPurple,
		"BEIGE": rl.Beige,
		"BROWN": rl.Brown,
		"DARKBROWN": rl.DarkBrown,
		"WHITE": rl.White,
		"BLACK": rl.Black,
		"BLANK": rl.Blank,
		"MAGENTA": rl.Magenta,
		"RAYWHITE": rl.RayWhite,
	}

	e := env.NewEnv()

	
	// Non-Raylib functions
	_ = e.Define("print", fmt.Println)


	// MODULE : CORE \\ | //## = ORANGE | //#### = RED


	// Window-related functions

	_ = e.Define("InitWindow", ChoiceInitWindow)
	_ = e.Define("WindowShouldClose", rl.WindowShouldClose)
	_ = e.Define("CloseWindow", rl.CloseWindow)
	_ = e.Define("IsWindowReady", rl.IsWindowReady)
	_ = e.Define("IsWindowFullscreen", rl.IsWindowFullscreen)
	_ = e.Define("IsWindowHidden", rl.IsWindowHidden)
	_ = e.Define("IsWindowMinimized", rl.IsWindowMinimized)
	_ = e.Define("IsWindowMaximized", rl.IsWindowMaximized)
	_ = e.Define("IsWindowFocused", rl.IsWindowFocused)
	_ = e.Define("IsWindowResized", rl.IsWindowResized)
	_ = e.Define("IsWindowState", rl.IsWindowState)
	_ = e.Define("SetWindowState", rl.SetWindowState)
	_ = e.Define("ClearWindowState", rl.ClearWindowState)
	_ = e.Define("ToggleFullscreen", rl.ToggleFullscreen)
	_ = e.Define("MaximizeWindow", rl.MaximizeWindow)
	_ = e.Define("MinimizeWindow", rl.MinimizeWindow)
	_ = e.Define("RestoreWindow", rl.RestoreWindow)
	_ = e.Define("SetWindowIcon", rl.SetWindowIcon)
	_ = e.Define("SetWindowTitle", rl.SetWindowTitle)
	_ = e.Define("SetWindowPosition", rl.SetWindowPosition)
	_ = e.Define("SetWindowMonitor", rl.SetWindowMonitor)
	_ = e.Define("SetWindowMinSize", rl.SetWindowMinSize)
	_ = e.Define("SetWindowSize", rl.SetWindowSize)
	//####_ = e.Define("GetWindowHandle", rl.GetWindowHandle) -- NOT IN GO-raylib
	_ = e.Define("GetScreenWidth", rl.GetScreenWidth)
	_ = e.Define("GetScreenHeight", rl.GetScreenHeight)
	_ = e.Define("GetMonitorCount", rl.GetMonitorCount)
	_ = e.Define("GetCurrentMonitor", rl.GetCurrentMonitor)
	_ = e.Define("GetMonitorPosition", rl.GetMonitorPosition)
	_ = e.Define("GetMonitorWidth", rl.GetMonitorWidth)
	_ = e.Define("GetMonitorHeight", rl.GetMonitorHeight)
	_ = e.Define("GetMonitorPhysicalWidth", rl.GetMonitorPhysicalWidth)
	_ = e.Define("GetMonitorPhysicalHeight", rl.GetMonitorPhysicalHeight)
	_ = e.Define("GetMonitorRefreshRate", rl.GetMonitorRefreshRate)
	_ = e.Define("GetWindowPosition", rl.GetWindowPosition)
	_ = e.Define("GetWindowScaleDPI", rl.GetWindowScaleDPI)
	_ = e.Define("GetMonitorName", rl.GetMonitorName)
	_ = e.Define("SetClipboardText", rl.SetClipboardText)
	_ = e.Define("GetClipboardText", rl.GetClipboardText)


	// Cursor-related functions

	_ = e.Define("ShowCursor", rl.ShowCursor)
	_ = e.Define("HideCursor", rl.HideCursor)
	_ = e.Define("IsCursorHidden", rl.IsCursorHidden)
	_ = e.Define("EnableCursor", rl.EnableCursor)
	_ = e.Define("DisableCursor", rl.DisableCursor)
	_ = e.Define("IsCursorOnScreen", rl.IsCursorOnScreen)


    // Drawing-related functions

	_ = e.Define("ClearBackground", rl.ClearBackground)
	_ = e.Define("BeginDrawing", rl.BeginDrawing)
	_ = e.Define("EndDrawing", rl.EndDrawing)
	//####_ = e.Define("BeginMode2D", rl.BeginMode2D)
	//####_ = e.Define("EndMode2D", rl.EndMode2D)
	//####_ = e.Define("BeginMode3D", rl.BeginMode3D)
	//####_ = e.Define("EndMode3D", rl.EndMode3D)
	//####_ = e.Define("BeginTextureMode", rl.BeginTextureMode)
	//####_ = e.Define("EndTextureMode", rl.EndTextureMode)
	//####_ = e.Define("BeginShaderMode", rl.BeginShaderMode)
	//####_ = e.Define("EndShaderMode", rl.EndShaderMode)
	//####_ = e.Define("BeginBlendMode", rl.BeginBlendMode)
	//####_ = e.Define("EndBlendMode", rl.EndBlendMode)
	//####_ = e.Define("BeginScissorMode", rl.BeginScissorMode)
	//####_ = e.Define("EndScissorMode", rl.EndScissorMode)
	//####_ = e.Define("BeginVrStereoMode", rl.BeginVrStereoMode)
	//####_ = e.Define("EndVrStereoMode", rl.EndVrStereoMode)


    // VR stereo config functions for VR simulator

	//####_ = e.Define("LoadVrStereoConfig", rl.LoadVrStereoConfig)
	//####_ = e.Define("UnloadVrStereoConfig", rl.UnloadVrStereoConfig)


	// Shader management functions
	
	//####_ = e.Define("LoadShader", rl.LoadShader)
	//####_ = e.Define("LoadShaderFromMemory", rl.LoadShaderFromMemory)
	//####_ = e.Define("GetShaderLocation", rl.GetShaderLocation)
	//####_ = e.Define("GetShaderLocationAttrib", rl.GetShaderLocationAttrib)
	//####_ = e.Define("SetShaderValue", rl.SetShaderValue)
	//####_ = e.Define("SetShaderValueV", rl.SetShaderValueV)
	//####_ = e.Define("SetShaderValueMatrix", rl.SetShaderValueMatrix)
	//####_ = e.Define("SetShaderValueTexture", rl.SetShaderValueTexture)


    // Screen-space-related functions

	//####_ = e.Define("GetMouseRay", rl.GetMouseRay)
	//####_ = e.Define("GetCameraMatrix", rl.GetCameraMatrix)
	//####_ = e.Define("GetCameraMatrix2D", rl.GetCameraMatrix2D)
	//####_ = e.Define("GetWorldToScreen", rl.GetWorldToScreen)
	//####_ = e.Define("GetWorldToScreenEx", rl.GetWorldToScreenEx)
	//####_ = e.Define("GetWorldToScreen2D", rl.GetWorldToScreen2D)
	//####_ = e.Define("GetScreenToWorld2D", rl.GetScreenToWorld2D)


	// Timing-related functions

	_ = e.Define("SetTargetFPS", rl.SetTargetFPS)
	_ = e.Define("GetFPS", rl.GetFPS)
	_ = e.Define("GetFrameTime", rl.GetFrameTime)
	_ = e.Define("GetTime", rl.GetTime)


    // Misc. functions

	_ = e.Define("GetRandomValue", rl.GetRandomValue)
	_ = e.Define("SetRandomSeed", rl.SetRandomSeed)
	_ = e.Define("TakeScreenshot", rl.TakeScreenshot)
	_ = e.Define("SetConfigFlags", rl.SetConfigFlags)
	//####_ = e.Define("TraceLog", rl.TraceLog)
	//####_ = e.Define("SetTraceLogLevel", rl.SetTraceLogLevel)
	//####_ = e.Define("MemAlloc", rl.MemAlloc)
	//####_ = e.Define("MemRealloc", rl.MemRealloc)
	//####_ = e.Define("MemFree", rl.MemFree)


	// Set custom callbacks

	//####_ = e.Define("SetTraceLogCallback", rl.SetTraceLogCallback)
	//####_ = e.Define("SetLoadFileDataCallback", rl.SetLoadFileDataCallback)
	//####_ = e.Define("SetSaveFileDataCallback", rl.SetSaveFileDataCallback)
	//####_ = e.Define("SetLoadFileTextCallback", rl.SetLoadFileTextCallback)
	//####_ = e.Define("SetSaveFileTextCallback", rl.SetSaveFileTextCallback)


	// Files management functions

	//####_ = e.Define("LoadFileData", rl.LoadFileData)
	//####_ = e.Define("UnloadFileData", rl.UnloadFileData)
	//####_ = e.Define("SaveFileData", rl.SaveFileData)
	//####_ = e.Define("LoadFileText", rl.LoadFileText)
	//####_ = e.Define("UnloadFileText", rl.UnloadFileText)
	//####_ = e.Define("SaveFileText", rl.SaveFileText)
	//####_ = e.Define("FileExists", rl.FileExists)
	//####_ = e.Define("DirectoryExists", rl.DirectoryExists)
	//####_ = e.Define("IsFileExtension", rl.IsFileExtension)
	//####_ = e.Define("GetFileExtension", rl.GetFileExtension)
	//####_ = e.Define("GetFileName", rl.GetFileName)
	//####_ = e.Define("GetFileNameWithoutExt", rl.GetFileNameWithoutExt)
	//####_ = e.Define("GetDirectoryPath", rl.GetDirectoryPath)
	//####_ = e.Define("GetPrevDirectoryPath", rl.GetPrevDirectoryPath)
	//####_ = e.Define("GetWorkingDirectory", rl.GetWorkingDirectory)
	//####_ = e.Define("GetDirectoryFiles", rl.GetDirectoryFiles)
	//####_ = e.Define("ClearDirectoryFiles", rl.ClearDirectoryFiles)
	//####_ = e.Define("ChangeDirectory", rl.ChangeDirectory)
	//####_ = e.Define("IsFileDropped", rl.IsFileDropped)
	//####_ = e.Define("GetDroppedFiles", rl.GetDroppedFiles)
	//####_ = e.Define("ClearDroppedFiles", rl.ClearDroppedFiles)
	//####_ = e.Define("GetFileModTime", rl.GetFileModTime)


	// Compression/Encoding functionality

	//####_ = e.Define("CompressData", rl.CompressData)
	//####_ = e.Define("DecompressData", rl.DecompressData)
	//####_ = e.Define("EncodeDataBase64", rl.EncodeDataBase64)
	//####_ = e.Define("DecodeDataBase64", rl.DecodeDataBase64)


	// Persistent storage management

	//####_ = e.Define("SaveStorageValue", rl.SaveStorageValue)
	//####_ = e.Define("LoadStorageValue", rl.LoadStorageValue)


	// Misc.

	//####_ = e.Define("OpenURL", rl.OpenURL)


	// Input-related functions: keyboard

	_ = e.Define("IsKeyPressed", rl.IsKeyPressed)
	_ = e.Define("IsKeyDown", rl.IsKeyDown)
	_ = e.Define("IsKeyReleased", rl.IsKeyReleased)
	_ = e.Define("IsKeyUp", rl.IsKeyUp)
	_ = e.Define("SetExitKey", rl.SetExitKey)
	_ = e.Define("GetKeyPressed", rl.GetKeyPressed)
	_ = e.Define("GetCharPressed", rl.GetCharPressed)

	    
	// Input-related functions: gamepads

	//####_ = e.Define("IsGamepadAvailable", rl.IsGamepadAvailable)
	//####_ = e.Define("GetGamepadName", rl.GetGamepadName)
	//####_ = e.Define("IsGamepadButtonPressed", rl.IsGamepadButtonPressed)
	//####_ = e.Define("IsGamepadButtonDown", rl.IsGamepadButtonDown)
	//####_ = e.Define("IsGamepadButtonReleased", rl.IsGamepadButtonReleased)
	//####_ = e.Define("IsGamepadButtonUp", rl.IsGamepadButtonUp)
	//####_ = e.Define("GetGamepadButtonPressed", rl.GetGamepadButtonPressed)
	//####_ = e.Define("GetGamepadAxisCount", rl.GetGamepadAxisCount)
	//####_ = e.Define("GetGamepadAxisMovement", rl.GetGamepadAxisMovement)
	//####_ = e.Define("SetGamepadMappings", rl.SetGamepadMappings)

	
    // Input-related functions: mouse

	_ = e.Define("IsMouseButtonPressed", rl.IsMouseButtonPressed)
	_ = e.Define("IsMouseButtonDown", rl.IsMouseButtonDown)
	_ = e.Define("IsMouseButtonReleased", rl.IsMouseButtonReleased)
	_ = e.Define("IsMouseButtonUp", rl.IsMouseButtonUp)
	_ = e.Define("GetMouseX", rl.GetMouseX)
	_ = e.Define("GetMouseY", rl.GetMouseY)
	_ = e.Define("GetMousePosition", rl.GetMousePosition)
	_ = e.Define("GetMouseDelta", rl.GetMouseDelta)
	_ = e.Define("SetMousePosition", rl.SetMousePosition)
	_ = e.Define("SetMouseOffset", rl.SetMouseOffset)
	_ = e.Define("SetMouseScale", rl.SetMouseScale)
	_ = e.Define("GetMouseWheelMove", rl.GetMouseWheelMove)
	_ = e.Define("SetMouseCursor", rl.SetMouseCursor)

	
    // Input-related functions: touch
	
	//####_ = e.Define("GetTouchX", rl.GetTouchX)
	//####_ = e.Define("GetTouchY", rl.GetTouchY)
	//####_ = e.Define("GetTouchPosition", rl.GetTouchPosition)
	//####_ = e.Define("GetTouchPointId", rl.GetTouchPointId)
	//####_ = e.Define("GetTouchPointCount", rl.GetTouchPointCount)


	// Gestures and Touch Handling Functions (Module: rgestures)

	//####_ = e.Define("SetGesturesEnabled", rl.SetGesturesEnabled)
	//####_ = e.Define("IsGestureDetected", rl.IsGestureDetected)
	//####_ = e.Define("GetGestureDetected", rl.GetGestureDetected)
	//####_ = e.Define("GetGestureHoldDuration", rl.GetGestureHoldDuration)
	//####_ = e.Define("GetGestureDragVector", rl.GetGestureDragVector)
	//####_ = e.Define("GetGestureDragAngle", rl.GetGestureDragAngle)
	//####_ = e.Define("GetGesturePinchVector", rl.GetGesturePinchVector)
	//####_ = e.Define("GetGesturePinchAngle", rl.GetGesturePinchAngle)

	
    // Camera System Functions (Module: rcamera)

	//####_ = e.Define("SetCameraMode", rl.SetCameraMode)
	//####_ = e.Define("SetCameraMode", rl.SetCameraMode)

	//####_ = e.Define("SetCameraPanControl", rl.SetCameraPanControl)
	//####_ = e.Define("SetCameraAltControl", rl.SetCameraAltControl)
	//####_ = e.Define("SetCameraSmoothZoomControl", rl.SetCameraSmoothZoomControl)
	//####_ = e.Define("SetCameraMoveControls", rl.SetCameraMoveControls)








	/*

	//####_ = e.Define("aaaaa", rl.aaaaa)

	_ = e.Define("aaaaa", rl.aaaaa)
	_ = e.Define("aaaaa", rl.aaaaa)
	_ = e.Define("aaaaa", rl.aaaaa)
	_ = e.Define("aaaaa", rl.aaaaa)
	
	*/
	

	

	for key, val := range colors { _ = e.Define(key, val) }

	_, _ = vm.Execute(e, nil, string(script))
}