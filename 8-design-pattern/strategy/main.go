package main

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	pp := NewPasswordProtector("Carlos", "Email", md5)
	pp.Hash()

	pp.SetHashAlgorithm(sha)
	pp.Hash()
}
