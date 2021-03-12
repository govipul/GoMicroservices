# CURL installation in Powershell
For installation i have been using Scoop package manager for windows.
```bash
scoop install curl
```
also to updae to latest cURL i have used following command 
```bash
scoop update curl
```
# cURL issue with url
In the series Nic is using zsh terminal so he was able to curl command directly but i have been using the powershell so even though i have been using the cURL as a command in the backend it call the Invoke-WebRequest which
popups many issues like following 

```bash
curl --data 'Test' http://localhost:9090
Invoke-WebRequest : A positional parameter cannot be found that accepts argument 'Test'.
At line:1 char:1
+ curl --data 'Test' http://localhost:9090
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```
# Fix

Run following command and in that case it will start using cURL.
```bash
Remove-item alias:curl
```

