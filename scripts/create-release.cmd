@echo off

rem This script create a release w/ the exe's as part of the release
rem Assumed is 
rem - a Github repo 
rem - the presence of the Github CLI in the PATH
rem - Linux tools from the Git client (cp, sed)

if [%1]==[] goto usage

choice /C YN /M "create release %1? "
if ERRORLEVEL 2 goto end

echo GO!
rem 1. overwrite logic/version.txt first with correct version
cp version-template.txt logic/version.txt
sed -i "s/${version}/%1/" logic/version.txt

rem 2. commit the file
git add logic/version.txt
git commit -m "bumped version to %1"

rem 3. create, commit an annotated tag
git tag -a %1 -m "release: %1"

rem 4. push commit and tag to origin
git push 
git push origin %1

rem 5. create the release based on the tag
gh release create %1 --generate-notes

goto end

:usage
@echo Usage: %0 ^<version^>
@echo ^<version^> must have v prefix 
exit /B 1

:end
exit /B 0