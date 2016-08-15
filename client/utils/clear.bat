@RD /S /Q "%USERPROFILE%\AppData\Roaming\AsocialFriend"
reg delete "HKCU\Software\AsocialFriend" /f
reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /f /v asocialfriend