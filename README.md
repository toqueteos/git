# git

[![GoDoc](https://godoc.org/github.com/toqueteos/git?status.svg)](http://godoc.org/github.com/toqueteos/git)

git implementation in Go.

This is a toy project and it is abandoned, but I'm not deleting it.

Why so? Because of many reasons:

- This repository got me a job, not kidding!
- It helped me learning all the amazing git is made of
- This is a failure project of mine and I'm ashamed of it, it taught me lots of things!

## When...?

*... will this be ready?*

**Never.**

But don't worry there's plenty of git libraries in Go, here's a couple of suggestions:

- https://github.com/kourge/ggit superb documentation, super friendly and amazing guy!
- https://github.com/src-d/go-git almost no documentation but gets the job done most of the time

## Why?

I wanted to learn how git works internally not just the commands to use it.

Also, I want `go get` to be independent of external commands so anyone can `go get` anything without installing git, MsysGit, TortoiseHg, Bazaar or any SCM out there supported by this tool.

By implementing git in Go `go get` can include additional features like fetching a repo at certain tag or commit (just an example).

Go authors still have to approve this, which probably will be rejected just because there's no Mercurial and Bazaar Go implementations yet (if there are I don't know a thing about them).

## Reference

Some useful links that helped me understand some of the deepest concepts of git:

- https://git-scm.com/book/en/v2/Git-Internals-Packfiles
- http://stefan.saasen.me/articles/git-clone-in-haskell-from-the-bottom-up/
- http://git.661346.n2.nabble.com/Exact-format-of-tree-objets-td7588984.html
- http://git.rsbx.net/Documents/Git_Data_Formats.txt
- https://www.youtube.com/watch?v=MYP56QJpDr4

Feel free to email me!
