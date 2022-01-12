package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"os"
	"fmt"
	"github.com/mattn/anko/vm"
	"github.com/mattn/anko/env"
	"reflect"
	"math"
	"net/http"
	"io"
)

func ChoiceInitWindow(width int32, height int32, title string, args ...int32) {
	if len(args) != 0 { InitWindowWithFPS(width, height, title, args[0])
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

func webget(url string) (string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
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
	_ = e.Define("type", reflect.TypeOf)

	_ = e.Define("cos", math.Cos)
	_ = e.Define("acos", math.Acos)
	_ = e.Define("sin", math.Sin)
	_ = e.Define("asin", math.Asin)

	_ = e.Define("ceil", math.Ceil)
	
	_ = e.Define("max", math.Max)
	_ = e.Define("min", math.Min)

	_ = e.Define("webget", webget)

	_ = e.Define("quit", func() {os.Exit(0)})
	
	
	
	//### <- This means that it is not yet implemented in go-raylib
	// MODULE : CORE \\


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
	//###_ = e.Define("GetWindowHandle", rl.GetWindowHandle) 
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
	_ = e.Define("BeginMode2D", rl.BeginMode2D)
	_ = e.Define("EndMode2D", rl.EndMode2D)
	_ = e.Define("BeginMode3D", rl.BeginMode3D)
	_ = e.Define("EndMode3D", rl.EndMode3D)
	_ = e.Define("BeginTextureMode", rl.BeginTextureMode)
	_ = e.Define("EndTextureMode", rl.EndTextureMode)
	_ = e.Define("BeginShaderMode", rl.BeginShaderMode)
	_ = e.Define("EndShaderMode", rl.EndShaderMode)
	_ = e.Define("BeginBlendMode", rl.BeginBlendMode)
	_ = e.Define("EndBlendMode", rl.EndBlendMode)
	_ = e.Define("BeginScissorMode", rl.BeginScissorMode)
	_ = e.Define("EndScissorMode", rl.EndScissorMode)
	//###_ = e.Define("BeginVrStereoMode", rl.BeginVrStereoMode)
	//###_ = e.Define("EndVrStereoMode", rl.EndVrStereoMode)


   	 // VR stereo config functions for VR simulator

	//###_ = e.Define("LoadVrStereoConfig", rl.LoadVrStereoConfig)
	//###_ = e.Define("UnloadVrStereoConfig", rl.UnloadVrStereoConfig)


	// Shader management functions
	
	_ = e.Define("LoadShader", rl.LoadShader)
	_ = e.Define("LoadShaderFromMemory", rl.LoadShaderFromMemory)
	_ = e.Define("GetShaderLocation", rl.GetShaderLocation)
	_ = e.Define("GetShaderLocationAttrib", rl.GetShaderLocationAttrib)
	_ = e.Define("SetShaderValue", rl.SetShaderValue)
	_ = e.Define("SetShaderValueV", rl.SetShaderValueV)
	_ = e.Define("SetShaderValueMatrix", rl.SetShaderValueMatrix)
	_ = e.Define("SetShaderValueTexture", rl.SetShaderValueTexture)


   	 // Screen-space-related functions

	_ = e.Define("GetMouseRay", rl.GetMouseRay)
	_ = e.Define("GetCameraMatrix", rl.GetCameraMatrix)
	_ = e.Define("GetCameraMatrix2D", rl.GetCameraMatrix2D)
	_ = e.Define("GetWorldToScreen", rl.GetWorldToScreen)
	//###_ = e.Define("GetWorldToScreenEx", rl.GetWorldToScreenEx)
	_ = e.Define("GetWorldToScreen2D", rl.GetWorldToScreen2D)
	_ = e.Define("GetScreenToWorld2D", rl.GetScreenToWorld2D)


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
	_ = e.Define("TraceLog", rl.TraceLog)
	//###_ = e.Define("SetTraceLogLevel", rl.SetTraceLogLevel)
	//###_ = e.Define("MemAlloc", rl.MemAlloc)
	//###_ = e.Define("MemRealloc", rl.MemRealloc)
	//###_ = e.Define("MemFree", rl.MemFree)


	// Set custom callbacks

	_ = e.Define("SetTraceLogCallback", rl.SetTraceLogCallback)
	//###_ = e.Define("SetLoadFileDataCallback", rl.SetLoadFileDataCallback)
	//###_ = e.Define("SetSaveFileDataCallback", rl.SetSaveFileDataCallback)
	//###_ = e.Define("SetLoadFileTextCallback", rl.SetLoadFileTextCallback)
	//###_ = e.Define("SetSaveFileTextCallback", rl.SetSaveFileTextCallback)


	// Files management functions

	//###_ = e.Define("LoadFileData", rl.LoadFileData)
	//###_ = e.Define("UnloadFileData", rl.UnloadFileData)
	//###_ = e.Define("SaveFileData", rl.SaveFileData)
	//###_ = e.Define("LoadFileText", rl.LoadFileText)
	//###_ = e.Define("UnloadFileText", rl.UnloadFileText)
	//###_ = e.Define("SaveFileText", rl.SaveFileText)
	//###_ = e.Define("FileExists", rl.FileExists)
	//###_ = e.Define("DirectoryExists", rl.DirectoryExists)
	//###_ = e.Define("IsFileExtension", rl.IsFileExtension)
	//###_ = e.Define("GetFileExtension", rl.GetFileExtension)
	//###_ = e.Define("GetFileName", rl.GetFileName)
	//###_ = e.Define("GetFileNameWithoutExt", rl.GetFileNameWithoutExt)
	//###_ = e.Define("GetDirectoryPath", rl.GetDirectoryPath)
	//###_ = e.Define("GetPrevDirectoryPath", rl.GetPrevDirectoryPath)
	//###_ = e.Define("GetWorkingDirectory", rl.GetWorkingDirectory)
	//###_ = e.Define("GetDirectoryFiles", rl.GetDirectoryFiles)
	//###_ = e.Define("ClearDirectoryFiles", rl.ClearDirectoryFiles)
	//###_ = e.Define("ChangeDirectory", rl.ChangeDirectory)
	_ = e.Define("IsFileDropped", rl.IsFileDropped)
	_ = e.Define("GetDroppedFiles", rl.GetDroppedFiles)
	_ = e.Define("ClearDroppedFiles", rl.ClearDroppedFiles)
	//###_ = e.Define("GetFileModTime", rl.GetFileModTime)


	// Compression/Encoding functionality

	//###_ = e.Define("CompressData", rl.CompressData)
	//###_ = e.Define("DecompressData", rl.DecompressData)
	//###_ = e.Define("EncodeDataBase64", rl.EncodeDataBase64)
	//###_ = e.Define("DecodeDataBase64", rl.DecodeDataBase64)


	// Persistent storage management

	_ = e.Define("SaveStorageValue", rl.SaveStorageValue)
	_ = e.Define("LoadStorageValue", rl.LoadStorageValue)


	// Misc.

	//###_ = e.Define("OpenURL", rl.OpenURL)


	// Input-related functions: keyboard

	_ = e.Define("IsKeyPressed", rl.IsKeyPressed)
	_ = e.Define("IsKeyDown", rl.IsKeyDown)
	_ = e.Define("IsKeyReleased", rl.IsKeyReleased)
	_ = e.Define("IsKeyUp", rl.IsKeyUp)
	_ = e.Define("SetExitKey", rl.SetExitKey)
	_ = e.Define("GetKeyPressed", rl.GetKeyPressed)
	_ = e.Define("GetCharPressed", rl.GetCharPressed)

	    
	// Input-related functions: gamepads

	_ = e.Define("IsGamepadAvailable", rl.IsGamepadAvailable)
	_ = e.Define("GetGamepadName", rl.GetGamepadName)
	_ = e.Define("IsGamepadButtonPressed", rl.IsGamepadButtonPressed)
	_ = e.Define("IsGamepadButtonDown", rl.IsGamepadButtonDown)
	_ = e.Define("IsGamepadButtonReleased", rl.IsGamepadButtonReleased)
	_ = e.Define("IsGamepadButtonUp", rl.IsGamepadButtonUp)
	_ = e.Define("GetGamepadButtonPressed", rl.GetGamepadButtonPressed)
	_ = e.Define("GetGamepadAxisCount", rl.GetGamepadAxisCount)
	_ = e.Define("GetGamepadAxisMovement", rl.GetGamepadAxisMovement)
	//###_ = e.Define("SetGamepadMappings", rl.SetGamepadMappings)

	
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
	
	_ = e.Define("GetTouchX", rl.GetTouchX)
	_ = e.Define("GetTouchY", rl.GetTouchY)
	_ = e.Define("GetTouchPosition", rl.GetTouchPosition)
	_ = e.Define("GetTouchPointId", rl.GetTouchPointId)
	_ = e.Define("GetTouchPointCount", rl.GetTouchPointCount)


	// Gestures and Touch Handling Functions (Module: rgestures)

	_ = e.Define("SetGesturesEnabled", rl.SetGesturesEnabled)
	_ = e.Define("IsGestureDetected", rl.IsGestureDetected)
	_ = e.Define("GetGestureDetected", rl.GetGestureDetected)
	_ = e.Define("GetGestureHoldDuration", rl.GetGestureHoldDuration)
	_ = e.Define("GetGestureDragVector", rl.GetGestureDragVector)
	_ = e.Define("GetGestureDragAngle", rl.GetGestureDragAngle)
	_ = e.Define("GetGesturePinchVector", rl.GetGesturePinchVector)
	_ = e.Define("GetGesturePinchAngle", rl.GetGesturePinchAngle)

	
    	// Camera System Functions (Module: rcamera)

	_ = e.Define("SetCameraMode", rl.SetCameraMode)
	_ = e.Define("SetCameraMode", rl.SetCameraMode)

	_ = e.Define("SetCameraPanControl", rl.SetCameraPanControl)
	_ = e.Define("SetCameraAltControl", rl.SetCameraAltControl)
	_ = e.Define("SetCameraSmoothZoomControl", rl.SetCameraSmoothZoomControl)
	_ = e.Define("SetCameraMoveControls", rl.SetCameraMoveControls)



	// MODULE : SHAPES \\


	// Set texture and rectangle to be used on shapes drawing
	_ = e.Define("SetShapesTexture", rl.SetShapesTexture)


	// Basic shapes drawing functions

	_ = e.Define("DrawPixel", rl.DrawPixel)
	_ = e.Define("DrawPixelV", rl.DrawPixelV)
	_ = e.Define("DrawLine", rl.DrawLine)
	_ = e.Define("DrawLineV", rl.DrawLineV)
	_ = e.Define("DrawLineEx", rl.DrawLineEx)
	_ = e.Define("DrawLineBezierQuad", rl.DrawLineBezierQuad)
	_ = e.Define("DrawLineBezierCubic", rl.DrawLineBezierCubic)
	_ = e.Define("DrawLineStrip", rl.DrawLineStrip)
	_ = e.Define("DrawCircle", rl.DrawCircle)
	_ = e.Define("DrawCircleSector", rl.DrawCircleSector)
	_ = e.Define("DrawCircleSectorLines", rl.DrawCircleSectorLines)
	_ = e.Define("DrawCircleGradient", rl.DrawCircleGradient)
	_ = e.Define("DrawCircleV", rl.DrawCircleV)
	_ = e.Define("DrawCircleLines", rl.DrawCircleLines)
	_ = e.Define("DrawEllipse", rl.DrawEllipse)
	_ = e.Define("DrawEllipseLines", rl.DrawEllipseLines)
	_ = e.Define("DrawRing", rl.DrawRing)
	_ = e.Define("DrawRingLines", rl.DrawRingLines)
	_ = e.Define("DrawRectangle", rl.DrawRectangle)
	_ = e.Define("DrawRectangleV", rl.DrawRectangleV)
	_ = e.Define("DrawRectangleRec", rl.DrawRectangleRec)
	_ = e.Define("DrawRectanglePro", rl.DrawRectanglePro)
	_ = e.Define("DrawRectangleGradientV", rl.DrawRectangleGradientV)
	_ = e.Define("DrawRectangleGradientH", rl.DrawRectangleGradientH)
	_ = e.Define("DrawRectangleGradientEx", rl.DrawRectangleGradientEx)
	_ = e.Define("DrawRectangleLines", rl.DrawRectangleLines)
	_ = e.Define("DrawRectangleLinesEx", rl.DrawRectangleLinesEx)
	_ = e.Define("DrawRectangleRounded", rl.DrawRectangleRounded)
	_ = e.Define("DrawRectangleRoundedLines", rl.DrawRectangleRoundedLines)
	_ = e.Define("DrawTriangle", rl.DrawTriangle)
	_ = e.Define("DrawTriangleLines", rl.DrawTriangleLines)
	_ = e.Define("DrawTriangleFan", rl.DrawTriangleFan)
	_ = e.Define("DrawTriangleStrip", rl.DrawTriangleStrip)
	_ = e.Define("DrawPoly", rl.DrawPoly)
	_ = e.Define("DrawPolyLines", rl.DrawPolyLines)
	_ = e.Define("DrawPolyLinesEx", rl.DrawPolyLinesEx)


    	// Basic shapes collision detection functions

	_ = e.Define("CheckCollisionRecs", rl.CheckCollisionRecs)
	_ = e.Define("CheckCollisionCircles", rl.CheckCollisionCircles)
	_ = e.Define("CheckCollisionCircleRec", rl.CheckCollisionCircleRec)
	_ = e.Define("CheckCollisionPointRec", rl.CheckCollisionPointRec)
	_ = e.Define("CheckCollisionPointCircle", rl.CheckCollisionPointCircle)
	_ = e.Define("CheckCollisionPointTriangle", rl.CheckCollisionPointTriangle)
	_ = e.Define("CheckCollisionLines", rl.CheckCollisionLines)
	_ = e.Define("CheckCollisionPointLine", rl.CheckCollisionPointLine)
	_ = e.Define("GetCollisionRec", rl.GetCollisionRec)



	// MODULE: TEXTURES \\

	
	// Image loading functions

	_ = e.Define("LoadImage", rl.LoadImage)
	_ = e.Define("LoadImageRaw", rl.LoadImageRaw)
	_ = e.Define("LoadImageAnim", rl.LoadImageAnim)
	_ = e.Define("LoadImageFromMemory", rl.LoadImageFromMemory)
	_ = e.Define("LoadImageFromTexture", rl.LoadImageFromTexture)
	//###_ = e.Define("LoadImageFromScreen", rl.LoadImageFromScreen)
	_ = e.Define("UnloadImage", rl.UnloadImage)
	_ = e.Define("ExportImage", rl.ExportImage)
	//###_ = e.Define("ExportImageAsCode", rl.ExportImageAsCode)


	// Image generation functions

	_ = e.Define("GenImageColor", rl.GenImageColor)
	_ = e.Define("GenImageGradientV", rl.GenImageGradientV)
	_ = e.Define("GenImageGradientH", rl.GenImageGradientH)
	_ = e.Define("GenImageGradientRadial", rl.GenImageGradientRadial)
	_ = e.Define("GenImageChecked", rl.GenImageChecked)
	_ = e.Define("GenImageWhiteNoise", rl.GenImageWhiteNoise)
	_ = e.Define("GenImageCellular", rl.GenImageCellular)


	// Image manipulation functions

	_ = e.Define("ImageCopy", rl.ImageCopy)
	//###_ = e.Define("ImageFromImage", rl.ImageFromImage)
	_ = e.Define("ImageText", rl.ImageText)
	_ = e.Define("ImageTextEx", rl.ImageTextEx)
	_ = e.Define("ImageFormat", rl.ImageFormat)
	_ = e.Define("ImageToPOT", rl.ImageToPOT)
	_ = e.Define("ImageCrop", rl.ImageCrop)
	_ = e.Define("ImageAlphaCrop", rl.ImageAlphaCrop)
	_ = e.Define("ImageAlphaClear", rl.ImageAlphaClear)
	_ = e.Define("ImageAlphaMask", rl.ImageAlphaMask)
	_ = e.Define("ImageAlphaPremultiply", rl.ImageAlphaPremultiply)
	_ = e.Define("ImageResize", rl.ImageResize)
	_ = e.Define("ImageResizeNN", rl.ImageResizeNN)
	_ = e.Define("ImageResizeCanvas", rl.ImageResizeCanvas)
	_ = e.Define("ImageMipmaps", rl.ImageMipmaps)
	_ = e.Define("ImageDither", rl.ImageDither)
	_ = e.Define("ImageFlipVertical", rl.ImageFlipVertical)
	_ = e.Define("ImageFlipHorizontal", rl.ImageFlipHorizontal)
	_ = e.Define("ImageRotateCW", rl.ImageRotateCW)
	_ = e.Define("ImageRotateCCW", rl.ImageRotateCCW)
	_ = e.Define("ImageColorTint", rl.ImageColorTint)
	_ = e.Define("ImageColorInvert", rl.ImageColorInvert)
	_ = e.Define("ImageColorGrayscale", rl.ImageColorGrayscale)
	_ = e.Define("ImageColorContrast", rl.ImageColorContrast)
	_ = e.Define("ImageColorBrightness", rl.ImageColorBrightness)
	_ = e.Define("ImageColorReplace", rl.ImageColorReplace)
	_ = e.Define("LoadImageColors", rl.LoadImageColors)
	//###_ = e.Define("LoadImagePalette", rl.LoadImagePalette)
	//###_ = e.Define("UnloadImageColors", rl.UnloadImageColors)
	//###_ = e.Define("UnloadImagePalette", rl.UnloadImagePalette)
	//###_ = e.Define("GetImageAlphaBorder", rl.GetImageAlphaBorder)
	_ = e.Define("GetImageColor", rl.GetImageColor)


	// Image drawing functions

	_ = e.Define("ImageClearBackground", rl.ImageClearBackground)
	_ = e.Define("ImageDrawPixel", rl.ImageDrawPixel)
	_ = e.Define("ImageDrawPixelV", rl.ImageDrawPixelV)
	//###_ = e.Define("ImageDrawLine", rl.ImageDrawLine)
	//###_ = e.Define("ImageDrawLineV", rl.ImageDrawLineV)
	_ = e.Define("ImageDrawCircle", rl.ImageDrawCircle)
	_ = e.Define("ImageDrawCircleV", rl.ImageDrawCircleV)
	_ = e.Define("ImageDrawRectangle", rl.ImageDrawRectangle)
	_ = e.Define("ImageDrawRectangleV", rl.ImageDrawRectangleV)
	_ = e.Define("ImageDrawRectangleRec", rl.ImageDrawRectangleRec)
	_ = e.Define("ImageDrawRectangleLines", rl.ImageDrawRectangleLines)
	_ = e.Define("ImageDraw", rl.ImageDraw)
	_ = e.Define("ImageDrawText", rl.ImageDrawText)
	_ = e.Define("ImageDrawTextEx", rl.ImageDrawTextEx)


    	// Texture loading functions

	_ = e.Define("LoadTexture", rl.LoadTexture)
	_ = e.Define("LoadTextureFromImage", rl.LoadTextureFromImage)
	//###_ = e.Define("LoadTextureCubemap", rl.LoadTextureCubemap)
	_ = e.Define("LoadRenderTexture", rl.LoadRenderTexture)
	_ = e.Define("UnloadTexture", rl.UnloadTexture)
	_ = e.Define("UnloadRenderTexture", rl.UnloadRenderTexture)
	_ = e.Define("UpdateTexture", rl.UpdateTexture)
	_ = e.Define("UpdateTextureRec", rl.UpdateTextureRec)


    	// Texture configuration functions

	_ = e.Define("GenTextureMipmaps", rl.GenTextureMipmaps)
	_ = e.Define("SetTextureFilter", rl.SetTextureFilter)
	_ = e.Define("SetTextureWrap", rl.SetTextureWrap)


    	// Texture drawing functions

	_ = e.Define("DrawTexture", rl.DrawTexture)
	_ = e.Define("DrawTextureV", rl.DrawTextureV)
	_ = e.Define("DrawTextureEx", rl.DrawTextureEx)
	_ = e.Define("DrawTextureRec", rl.DrawTextureRec)
	//###_ = e.Define("DrawTextureQuad", rl.DrawTextureQuad)
	//###_ = e.Define("DrawTextureTiled", rl.DrawTextureTiled)
	_ = e.Define("DrawTexturePro", rl.DrawTexturePro)
	//###_ = e.Define("DrawTextureNPatch", rl.DrawTextureNPatch)
	//###_ = e.Define("DrawTexturePoly", rl.DrawTexturePoly)


    	// Color/pixel related functions


	_ = e.Define("Fade", rl.Fade)
	_ = e.Define("ColorToInt", rl.ColorToInt)
	_ = e.Define("ColorNormalize", rl.ColorNormalize)
	_ = e.Define("ColorFromNormalized", rl.ColorFromNormalized)
	_ = e.Define("ColorToHSV", rl.ColorToHSV)
	_ = e.Define("ColorFromHSV", rl.ColorFromHSV)
	_ = e.Define("ColorAlpha", rl.ColorAlpha)
	_ = e.Define("ColorAlphaBlend", rl.ColorAlphaBlend)
	_ = e.Define("GetColor", rl.GetColor)
	//###_ = e.Define("GetPixelColor", rl.GetPixelColor)
	//###_ = e.Define("SetPixelColor", rl.SetPixelColor)
	_ = e.Define("GetPixelDataSize", rl.GetPixelDataSize)



	// MODULE: TEXT \\


	// Font loading/unloading functions

	_ = e.Define("GetFontDefault", rl.GetFontDefault)
	_ = e.Define("LoadFont", rl.LoadFont)
	_ = e.Define("LoadFontEx", rl.LoadFontEx)
	_ = e.Define("LoadFontFromImage", rl.LoadFontFromImage)
	_ = e.Define("LoadFontFromMemory", rl.LoadFontFromMemory)
	_ = e.Define("LoadFontData", rl.LoadFontData)
	//###_ = e.Define("GenImageFontAtlas", rl.GenImageFontAtlas)
	//###_ = e.Define("UnloadFontData", rl.UnloadFontData)
	_ = e.Define("UnloadFont", rl.UnloadFont)


	// Text drawing functions

	_ = e.Define("DrawFPS", rl.DrawFPS)
	_ = e.Define("DrawText", rl.DrawText)
	_ = e.Define("DrawTextEx", rl.DrawTextEx)
	//###_ = e.Define("DrawTextPro", rl.DrawTextPro)
	//###_ = e.Define("DrawTextCodepoint", rl.DrawTextCodepoint)


    	// Text misc. functions

	_ = e.Define("MeasureText", rl.MeasureText)
	_ = e.Define("MeasureTextEx", rl.MeasureTextEx)
	_ = e.Define("GetGlyphIndex", rl.GetGlyphIndex)
	_ = e.Define("GetGlyphInfo", rl.GetGlyphInfo)
	_ = e.Define("GetGlyphAtlasRec", rl.GetGlyphAtlasRec)


    	// Text codepoints management functions (unicode characters)

	//###_ = e.Define("LoadCodepoints", rl.LoadCodepoints)
	//###_ = e.Define("UnloadCodepoints", rl.UnloadCodepoints)
	//###_ = e.Define("GetCodepointCount", rl.GetCodepointCount)
	//###_ = e.Define("GetCodepoint", rl.GetCodepoint)
	//###_ = e.Define("CodepointToUTF8", rl.CodepointToUTF8)
	//###_ = e.Define("TextCodepointsToUTF8", rl.TextCodepointsToUTF8)


    	// Text strings management functions (no utf8 strings, only byte chars)         

	//###_ = e.Define("TextCopy", rl.TextCopy)
	//###_ = e.Define("TextIsEqual", rl.TextIsEqual)
	//###_ = e.Define("TextLength", rl.TextLength)
	//###_ = e.Define("TextFormat", rl.TextFormat)
	//###_ = e.Define("TextSubtext", rl.TextSubtext)
	//###_ = e.Define("TextReplace", rl.TextReplace)
	//###_ = e.Define("TextInsert", rl.TextInsert)
	//###_ = e.Define("TextJoin", rl.TextJoin)
	//###_ = e.Define("TextSplit", rl.TextSplit)
	//###_ = e.Define("TextAppend", rl.TextAppend)
	//###_ = e.Define("TextFindIndex", rl.TextFindIndex)
	//###_ = e.Define("TextToUpper", rl.TextToUpper)
	//###_ = e.Define("TextToLower", rl.TextToLower)
	//###_ = e.Define("TextToPascal", rl.TextToPascal)
	//###_ = e.Define("TextToInteger", rl.TextToInteger)



	// MODULE: TEXT \\


    	// Basic geometric 3D shapes drawing functions

	_ = e.Define("DrawLine3D", rl.DrawLine3D)
	_ = e.Define("DrawPoint3D", rl.DrawPoint3D)
	_ = e.Define("DrawCircle3D", rl.DrawCircle3D)
	//###_ = e.Define("DrawTriangle3D", rl.DrawTriangle3D)
	//###_ = e.Define("DrawTriangleStrip3D", rl.DrawTriangleStrip3D)
	_ = e.Define("DrawCube", rl.DrawCube)
	_ = e.Define("DrawCubeV", rl.DrawCubeV)
	_ = e.Define("DrawCubeWires", rl.DrawCubeWires)
	_ = e.Define("DrawCubeWiresV", rl.DrawCubeWiresV)
	_ = e.Define("DrawCubeTexture", rl.DrawCubeTexture)
	//###_ = e.Define("DrawCubeTextureRec", rl.DrawCubeTextureRec)
	_ = e.Define("DrawSphere", rl.DrawSphere)
	_ = e.Define("DrawSphereEx", rl.DrawSphereEx)
	_ = e.Define("DrawSphereWires", rl.DrawSphereWires)
	_ = e.Define("DrawCylinder", rl.DrawCylinder)
	//###_ = e.Define("DrawCylinderEx", rl.DrawCylinderEx)
	//###_ = e.Define("DrawCylinderWires", rl.DrawCylinderWires)
	//###_ = e.Define("DrawCylinderWiresEx", rl.DrawCylinderWiresEx)
	_ = e.Define("DrawPlane", rl.DrawPlane)
	_ = e.Define("DrawRay", rl.DrawRay)
	_ = e.Define("DrawGrid", rl.DrawGrid)


    	// Model loading/unloading functions

	_ = e.Define("LoadModel", rl.LoadModel)
	_ = e.Define("LoadModelFromMesh", rl.LoadModelFromMesh)
	_ = e.Define("UnloadModel", rl.UnloadModel)
	//###_ = e.Define("UnloadModelKeepMeshes", rl.UnloadModelKeepMeshes)
	_ = e.Define("GetModelBoundingBox", rl.GetModelBoundingBox)


	// Model drawing functions

	_ = e.Define("DrawModel", rl.DrawModel)
	_ = e.Define("DrawModelEx", rl.DrawModelEx)
	_ = e.Define("DrawModelWires", rl.DrawModelWires)
	_ = e.Define("DrawModelWiresEx", rl.DrawModelWiresEx)
	_ = e.Define("DrawBoundingBox", rl.DrawBoundingBox)
	_ = e.Define("DrawBillboard", rl.DrawBillboard)
	_ = e.Define("DrawBillboardRec", rl.DrawBillboardRec)
	//###_ = e.Define("DrawBillboardPro", rl.DrawBillboardPro)


    	// Mesh management functions

	//###_ = e.Define("UploadMesh", rl.UploadMesh)
	//###_ = e.Define("UpdateMeshBuffer", rl.UpdateMeshBuffer)
	_ = e.Define("UnloadMesh", rl.UnloadMesh)
	_ = e.Define("DrawMesh", rl.DrawMesh)
	_ = e.Define("DrawMeshInstanced", rl.DrawMeshInstanced)
	_ = e.Define("ExportMesh", rl.ExportMesh)
	_ = e.Define("GetMeshBoundingBox", rl.GetMeshBoundingBox)
	//###_ = e.Define("GenMeshTangents", rl.GenMeshTangents)
	//###_ = e.Define("GenMeshBinormals", rl.GenMeshBinormals)


	// Mesh generation functions

	_ = e.Define("GenMeshPoly", rl.GenMeshPoly)
	_ = e.Define("GenMeshPlane", rl.GenMeshPlane)
	_ = e.Define("GenMeshCube", rl.GenMeshCube)
	_ = e.Define("GenMeshSphere", rl.GenMeshSphere)
	_ = e.Define("GenMeshHemiSphere", rl.GenMeshHemiSphere)
	_ = e.Define("GenMeshCylinder", rl.GenMeshCylinder)
	_ = e.Define("GenMeshCone", rl.GenMeshCone)
	_ = e.Define("GenMeshTorus", rl.GenMeshTorus)
	_ = e.Define("GenMeshKnot", rl.GenMeshKnot)
	_ = e.Define("GenMeshHeightmap", rl.GenMeshHeightmap)
	_ = e.Define("GenMeshCubicmap", rl.GenMeshCubicmap)


	// Material loading/unloading functions

	_ = e.Define("LoadMaterials", rl.LoadMaterials)
	_ = e.Define("LoadMaterialDefault", rl.LoadMaterialDefault)
	_ = e.Define("UnloadMaterial", rl.UnloadMaterial)
	_ = e.Define("SetMaterialTexture", rl.SetMaterialTexture)
	_ = e.Define("SetModelMeshMaterial", rl.SetModelMeshMaterial)


    	// Model animations loading/unloading functions

	//###_ = e.Define("LoadModelAnimations", rl.LoadModelAnimations)
	//###_ = e.Define("UpdateModelAnimation", rl.UpdateModelAnimation)
	//###_ = e.Define("UnloadModelAnimation", rl.UnloadModelAnimation)
	//###_ = e.Define("UnloadModelAnimations", rl.UnloadModelAnimations)
	//###_ = e.Define("IsModelAnimationValid", rl.IsModelAnimationValid)


   	// Collision detection functions

	_ = e.Define("CheckCollisionSpheres", rl.CheckCollisionSpheres)
	_ = e.Define("CheckCollisionBoxes", rl.CheckCollisionBoxes)
	_ = e.Define("CheckCollisionBoxSphere", rl.CheckCollisionBoxSphere)
	_ = e.Define("GetRayCollisionSphere", rl.GetRayCollisionSphere)
	_ = e.Define("GetRayCollisionBox", rl.GetRayCollisionBox)
	_ = e.Define("GetRayCollisionModel", rl.GetRayCollisionModel)
	_ = e.Define("GetRayCollisionMesh", rl.GetRayCollisionMesh)
	_ = e.Define("GetRayCollisionTriangle", rl.GetRayCollisionTriangle)
	_ = e.Define("GetRayCollisionQuad", rl.GetRayCollisionQuad)



	// MODULE: AUDIO \\


    	// Audio device management functions

	_ = e.Define("InitAudioDevice", rl.InitAudioDevice)
	_ = e.Define("CloseAudioDevice", rl.CloseAudioDevice)
	_ = e.Define("IsAudioDeviceReady", rl.IsAudioDeviceReady)
	_ = e.Define("SetMasterVolume", rl.SetMasterVolume)


    	// Wave/Sound loading/unloading functions

	_ = e.Define("LoadWave", rl.LoadWave)
	_ = e.Define("LoadWaveFromMemory", rl.LoadWaveFromMemory)
	_ = e.Define("LoadSound", rl.LoadSound)
	_ = e.Define("LoadSoundFromWave", rl.LoadSoundFromWave)
	_ = e.Define("UpdateSound", rl.UpdateSound)
	_ = e.Define("UnloadWave", rl.UnloadWave)
	_ = e.Define("UnloadSound", rl.UnloadSound)
	_ = e.Define("ExportWave", rl.ExportWave)
	//###_ = e.Define("ExportWaveAsCode", rl.ExportWaveAsCode)


    	// Wave/Sound management functions

	_ = e.Define("PlaySound", rl.PlaySound)
	_ = e.Define("StopSound", rl.StopSound)
	_ = e.Define("PauseSound", rl.PauseSound)
	_ = e.Define("ResumeSound", rl.ResumeSound)
	_ = e.Define("PlaySoundMulti", rl.PlaySoundMulti)
	_ = e.Define("StopSoundMulti", rl.StopSoundMulti)
	_ = e.Define("GetSoundsPlaying", rl.GetSoundsPlaying)
	_ = e.Define("IsSoundPlaying", rl.IsSoundPlaying)
	_ = e.Define("SetSoundVolume", rl.SetSoundVolume)
	_ = e.Define("SetSoundPitch", rl.SetSoundPitch)
	_ = e.Define("WaveFormat", rl.WaveFormat)
	_ = e.Define("WaveCopy", rl.WaveCopy)
	_ = e.Define("WaveCrop", rl.WaveCrop)
	_ = e.Define("LoadWaveSamples", rl.LoadWaveSamples)
	_ = e.Define("UnloadWaveSamples", rl.UnloadWaveSamples)


    	// Music management functions

	_ = e.Define("LoadMusicStream", rl.LoadMusicStream)
	_ = e.Define("LoadMusicStreamFromMemory", rl.LoadMusicStreamFromMemory)
	_ = e.Define("UnloadMusicStream", rl.UnloadMusicStream)
	_ = e.Define("PlayMusicStream", rl.PlayMusicStream)
	_ = e.Define("IsMusicStreamPlaying", rl.IsMusicStreamPlaying)
	_ = e.Define("UpdateMusicStream", rl.UpdateMusicStream)
	_ = e.Define("StopMusicStream", rl.StopMusicStream)
	_ = e.Define("PauseMusicStream", rl.PauseMusicStream)
	_ = e.Define("ResumeMusicStream", rl.ResumeMusicStream)
	_ = e.Define("SeekMusicStream", rl.SeekMusicStream)
	_ = e.Define("SetMusicVolume", rl.SetMusicVolume)
	_ = e.Define("SetMusicPitch", rl.SetMusicPitch)
	_ = e.Define("GetMusicTimeLength", rl.GetMusicTimeLength)
	_ = e.Define("GetMusicTimePlayed", rl.GetMusicTimePlayed)


    	// AudioStream management functions

	//###_ = e.Define("InitAudioStream", rl.InitAudioStream)
	//###_ = e.Define("CloseAudioStream", rl.CloseAudioStream)
	_ = e.Define("IsAudioStreamProcessed", rl.IsAudioStreamProcessed)
	_ = e.Define("PlayAudioStream", rl.PlayAudioStream)
	_ = e.Define("PauseAudioStream", rl.PauseAudioStream)
	_ = e.Define("ResumeAudioStream", rl.ResumeAudioStream)
	_ = e.Define("IsAudioStreamPlaying", rl.IsAudioStreamPlaying)
	_ = e.Define("StopAudioStream", rl.StopAudioStream)
	_ = e.Define("SetAudioStreamVolume", rl.SetAudioStreamVolume)
	_ = e.Define("SetAudioStreamPitch", rl.SetAudioStreamPitch)
	_ = e.Define("SetAudioStreamBufferSizeDefault", rl.SetAudioStreamBufferSizeDefault)



	// NewSTRUCT funcs \\

	_ = e.Define("NewVector2", rl.NewVector2)
	_ = e.Define("NewVector3", rl.NewVector3)
	_ = e.Define("NewVector4", rl.NewVector4)
	_ = e.Define("NewColor", rl.NewColor)
	_ = e.Define("NewRectangle", rl.NewRectangle)
	_ = e.Define("NewCamera2D", rl.NewCamera2D)
	_ = e.Define("NewCamera3D", rl.NewCamera3D)
	_ = e.Define("NewBoundingBox", rl.NewBoundingBox)
	_ = e.Define("NewRay", rl.NewRay)
	_ = e.Define("NewRayCollision", rl.NewRayCollision)

	
	for key, val := range colors { _ = e.Define(key, val) }

	_, err = vm.Execute(e, nil, string(script))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
