git
===

git implementation in Go.

Why?
====

I wanted to learn how git works internally not just the commands to use it.

Also, I want `go get` to be independent of external commands so anyone can `go get` anything without installing git, MsysGit, TortoiseHg, Bazaar or any SCM out there supported by this tool.

By implementing git in Go `go get` can include additional features like fetching a repo at certain tag or commit (just an example).

Go authors still have to approve this, which probably will be rejected just because there's no Mercurial and Bazaar Go implementations yet (if there are I don't know a thing about them).

When...?
========================

*... will this be ready?*

I don't know, it's a toy project right now.
