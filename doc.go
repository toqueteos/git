// Package git implements git SCM in Go
package git

/*
Taken from:
http://git-scm.com/book/en/Git-Internals-Plumbing-and-Porcelain

Directories:
- **branches**. Not used by newer Git versions.
- **hooks**. Contains your client or server-side hook scripts (Chapter 7).
- **info**. keeps a global exclude file for ignored patterns that you don't want to track in a .gitignore file.
- **objects**. Stores all the content for your database.
- **refs**. Stores pointers into commit objects in that data (branches).

Files:
- **config**. Contains your project-specific configuration options.
- **description**. Only used by the GitWeb program.
- **HEAD**. Points to the branch you currently have checked out.
- **index**. Where Git stores your staging area information.
*/
