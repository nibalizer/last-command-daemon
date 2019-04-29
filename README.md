# Last-Command-Daemon

Warning: This is extremely beta. It will probably set your computer on fire. 

I find that my twitch streams involve a lot of live terminal usage. Probably more than I use an editor. In order to make it easier for the audience to follow along and to enhance engagement, I display the last command my terminal ran as a browser source.



## Quickstart

1. Download the go binary from the [releases page](https://github.com/nibalizer/last-command-daemon/releases)

2. Run the binary:


```shell
chmod +x  commandprinter-linux 
./commandprinter-linux 
```


3. Use the zsh shell. Source `last_command.zsh`


```shell
source last_command.zsh
Loaded last-command hook. preexec function overloaded
```

4. Create an OBS browser source. The daemon is listining on `localhost:8080` you can use either `/latestCommand` or `/latestCommand.html` as sources. 



## Details

This works by using a zsh hook called `preexec`. For that reason it doesn't work with bash. I've not tested it, but [this repository](https://github.com/rcaloras/bash-preexec) claims to add similar functionality to bash.

The `lastcommand` binary simply holds data and moves it around, it would be easy to implement this in any number of ways, including just overwriting a text file and using a text source in OBS.


### Links:

* https://stackoverflow.com/questions/12580675/zsh-preexec-command-modification
* http://zshwiki.org/home/
* https://github.com/rcaloras/bash-preexec
* http://zsh.sourceforge.net/Doc/Release/Functions.html ( search preexec)

```
preexec

    Executed just after a command has been read and is about to be executed.
If the history mechanism is active (regardless of whether the line was discarded
from the history buffer), the string that the user typed is passed as the first
argument, otherwise it is an empty string. The actual command that will be
executed (including expanded aliases) is passed in two different forms: the second
argument is a single-line, size-limited version of the command (with things like
function bodies elided); the third argument contains the full text that is being
executed. 
```

### preexec functions


```shell
preexec () {
    echo "Latest Command: $2"
}
```

```shell
preexec () {
    cmd=$(echo $2 | base64)
    echo "Latest Command (base64): $cmd"
    /usr/bin/curl -d "${cmd}" localhost:8080/setCommand 2>1 >/dev/null
}
```
