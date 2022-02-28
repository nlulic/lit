# Lit

Have you ever wondered how Git works under the hood? I did - so i created `lit`.

Lit implements some of git's basic functionalities in less than 1000 lines using only the Go standard library. 

It currently supports basic versioning of a directory using commits, creating and switching to branches as well as displaying the commit history of a specific branch.

## Try it out

The binary can be built using `go build`. Running the command `lit init` will initialize the repository in the current directory.

## Commands 
The following commands are currently supported. 

### init

`lit init` will initialize the repository in the current directory. 

```
$ lit init

Initialized empty Lit repository in /home/nlulic/.lit
```


### commit

`lit commit -m <msg>` will create a commit of the current working tree. The commit message is optional and the author and committer will always be a default user called _anonymous_. 

Note: The output will print all files that are in a directory where a change was detected - not just single files that changed.

```
$ lit commit -m "initial commit"

[master (commit 1f5f518003d52205d7af3ae6bfe71104a0a217bd)]
2 files changed:
created/updated .litignore
created/updated README.md
```

### log

`lit log` will display the commit history of the current branch.

```
$ lit log

commit b59f7f50dae9c6f526a0f642e1f20e512441f55f (HEAD -> master)
Author: anonymous <anonymous@anonymous>
Date:   Mon, 28 Feb 2022 21:43:29 CET

        updated README.md

commit 5e0edaaee5e844b987deb0d478b78b5f48d84839 (HEAD -> master)
Author: anonymous <anonymous@anonymous>
Date:   Mon, 28 Feb 2022 21:43:17 CET

        initial commit
```

### branch

`lit branch <branch>` will create a new branch. 

```
$ lit branch develop
```

`lit branch` will output all branches.

```
$ lit branch

develop 
* master
```

### checkout 

`lit checkout <branch>` will checkout a branch if it exists. 

```
$ lit checkout develop

Switched to branch 'develop'
```

## .litignore

Just like in Git, specific files and directories can be ignored. Simply add a `.litignore` file to the root directory.

```
# directory
.git/

# binary
*.exe
```