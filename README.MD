A terminal prompt, written in Go, that is configurable using only a JSON file.

The config file /config/goprompt.json will give you the following terminal prompt:

![goprompt](/assets/goprompt.png)

For now the config needs to be manually placed in ~/.config/softteam/goprompt/config.json. You also need to edit your ~/.bashrc and make it look like this:

```
    PS1="\$(/path/to/goprompt)"	
```

The config can be edited to fir your needs. The following sections are supported right now:

* **text** : Prints a text message
* **pwd** : Prints out the current working directory
* **user-name** : Prints out the user name of the person logged in
* **computer-name** : Prints out the computer name
* **datetime** : Prints out the current date & time formatted using the Golang date format string in the *Format* tag.
* **git** : Prints out the status of a Git repository 
* **go** : Prints out the Go version in the go.mod file (if it exists)
* **free-space** : Prints out the amount of free space available on the users home drive. Can be formatted using SI or IEC.
