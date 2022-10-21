package helpers

func ParseOid(oid string) string {
	if GetRef(oid) != "" {
		return GetRef(oid)
	}
	return oid
}