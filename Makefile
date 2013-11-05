todo:
	@grep -n TODO *.go || true
	@grep -n println *.go || true

newset:
	@# make newset NEWSET=BitSet > bitset_test.go
	@sed -e 's|HashSet|$(NEWSET)|g' hashset_test.go