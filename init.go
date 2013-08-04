package git

import (
    "os"
)

type mkMonad struct {
    err error
}

func (m *mkMonad) Mk(name string, perm os.FileMode) {
    if m.err != nil {
        m.err = os.Mkdir(name, 0666)
    }
}

func Init() {
    var m mkMonad
    var list = []string{
        ".git"
        ".git/info"
        ".git/hooks",
        ".git/objects",
        ".git/objects/info",
        ".git/objects/pack",
        ".git/refs",
        ".git/refs/heads",
        ".git/refs/tags",
    }

    for _, path := range list {
        m.MkDir(path, 0666)
    }
}
