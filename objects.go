package git

type ObjectType byte

const (
	Blob   ObjectType = iota // holds a block of user data (eg the contents of a file, but not filename or other metadata)
	Tree                     // holds a map of (filename => checksum) - and therefore implicitly (filename => object)
	Commit                   // holds (commit-message, patch-checksum, list of parent-checksums, root-dir-object-checksum)
	Tag                      // holds (tagname, commit-object-checksum)
)

func (o ObjectType) String() {
	switch o {
	case Blob:
		return "blob"
	case Tree:
		return "tree"
	case Commit:
		return "commit"
	case Tag:
		return "tag"
	}
}
