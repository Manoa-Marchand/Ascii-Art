# Ascii-Art

## *Authors*

* Hugo JOYET
* Pierre-Louis BERTIN
* Manoa MARCHAND

## *Installing*

### *Cloning an Existing Repository* : 
If you want to get a copy of this repository  — the command you need is `git clone`. If you’re familiar with other VCS systems such as Subversion, you’ll notice that the command is "clone" and not "checkout". Instead of getting just a working copy, Git receives a full copy of nearly all data that the server has. Every version of every file for the history of the project is pulled down by default when you run git clone.

You clone a repository with `git clone <ssh link>` or `git clone <url>`. you can do so like this for the current repository :

    $ git clone git@git.ytrack.learn.ynov.com:HJOYET/ascii-art.git
    
## *Examples*

```
~/ascii-art$ go build
~/ascii-art$ ./ascii-art "hello"
```
 ```
  _              _   _
 | |            | | | |
 | |__     ___  | | | |   ___
 |  _ \   / _ \ | | | |  / _ \
 | | | | |  __/ | | | | | (_) |
 |_| |_|  \___| |_| |_|  \___/
 ```
 
 **~/ascii-art$ ./ascii-art "1Hello 2There"**
```
      _    _           _    _                       _______   _
  _  | |  | |         | |  | |               ____  |__   __| | |
 / | | |__| |   ___   | |  | |    ___       |___ \    | |    | |__      ___    _ __     ___
 | | |  __  |  / _ \  | |  | |   / _ \        __) |   | |    |  _ \    / _ \  | '__|   / _ \
 | | | |  | | |  __/  | |  | |  | (_) |      / __/    | |    | | | |  |  __/  | |     |  __/
 |_| |_|  |_|  \___|  |_|  |_|   \___/      |_____|   |_|    |_| |_|   \___|  |_|      \___|
```

**~/ascii-art$ ./ascii-art "1Hello 2There" "Layout"**

```
ERREUR: Veuillez rentrer un argument
```