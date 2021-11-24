IMPORTANT! WORK-IN-PROGRESS! IMPORTANT!

A terminal prompt, written in Go, that is configurable using only a JSON file.

The config file /config/goprompt.json will give you the following terminal prompt:

![goprompt](/assets/goprompt.png)

For now the config needs to be manually placed in ~/.config/softteam/goprompt/config.json. You also need to edit your ~/.bashrc and make it look like this:

```
    PS1="\$(/path/to/goprompt)"	
```

The config can be edited to fit your needs. The following tags are supported in the config right now:

**main-prefix**

* A prefix to the entire prompt

**main-suffix**

* A suffix to the entire prompt.

**type**

* **text** : Prints a text message.
* **pwd** : Prints out the current working directory.
* **user-name** : Prints out the user name of the person logged in.
* **computer-name** : Prints out the computer name.
* **datetime** : Prints out the current date & time formatted using the Golang date format string in the *Format* tag.
* **git** : Prints out the status of a Git repository.
* **go** : Prints out the Go version in the go.mod file (if it exists).
* **free-space** : Prints out the amount of free space available on the users home drive. Can be formatted using SI or IEC.
* **divider** : Creates a black divider between two sections.

**fg-color**

* The foreground color.

**bg-color**

* The background color.

**styles**

* **none** : Use regular font. An empty string can also be used.
* **bold** : Use bold font.
* **faint** : Use an faint text.
* **italic** : Use italic font.
* **underline** : Use underline font.

**template**

* The text to print. Add $data$ where the actual data should be added, like user name or git status. For a text item, the template is the text that will be printed, i e $data$ will not be replaced in this section.

**format**

* Format is used for the **datetime** section, and the **free-space** section.
    * datetime : Use a Golang date format string (example 2006-01-02 15:04).
    * free-space : Can be SI (1k = 1000) or IEC (1k = 1024).

**remove-if-empty**

* Used for the **git** and **go** section types. If you are not inside a Git repository, or a directory containing a go.mod file, this section is ignored. Remove-if-empty does currently have a bug (issue #1) where the correct color is not being used.

**sep** 

* The separator to the next section.

**sep-fg-color** 

* The separator foreground color.

**sep-bg-color**

* The separator foreground color.

**sep-styles**

* **none** : Use regular font. An empty string can also be used.
* **bold** : Use bold font.
* **faint** : Use an faint text.
* **italic** : Use italic font.
* **underline** : Use underline font.

**sep-prefix**

* A text that will prefix the separator text.

**sep-suffix**

* A text that will suffix the separator text.
