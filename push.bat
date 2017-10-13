@echo off
del cmd\preprocessor\*.exe /q  2>nul
del *.log /q  2>nul
rem del temp\*.* /q 2>nul
rem rmdir temp 2>nul
git add *
git commit * -m"%1"
git push

