# Lit

Have you ever wondered how git works under the hood? I did - so i created `lit`.

Lit implements some of git's basic functionalities in less than 1000 lines using only the Go standard library. 

It currently supports basic versioning of a directory using commits, creating and switching to branches as well as displaying the commit history of a specific branch.

## Try it out

The binary can be built using `go build`. Running the command `lit init` will initialize the repository in the current directory.

## Commands 
The following commands are currently supported. 

### init

`lit init` will initialize the repository in the current directory. 

### commit

`lit commit -m <msg>` will create a commit of the current working tree. The commit message is optional and the author and committer will always be a default user called _anonymous_. 

### branch

`lit branch` will output all branches.

`lit branch <branch>` will create a new branch. 

### checkout 

`lit checkout <branch>` will checkout a branch if it exists. 
