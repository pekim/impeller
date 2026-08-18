package impeller

func GetVersionVariant() uint32 {
	// #define IMPELLER_VERSION_GET_VARIANT(version) ((uint32_t)(version) >> 29U)
	return GetVersion() >> 29
}

func GetVersionMajor() uint32 {
	// #define IMPELLER_VERSION_GET_MAJOR(version) \
	// 	(((uint32_t)(version) >> 22U) & 0x7FU)
	return (GetVersion() >> 22) & 0x7f
}

func GetVersionMinor() uint32 {
	// #define IMPELLER_VERSION_GET_MINOR(version) \
	// 	(((uint32_t)(version) >> 12U) & 0x3FFU)
	return (GetVersion() >> 12) & 0x3ff
}

func GetVersionPatch() uint32 {
	// #define IMPELLER_VERSION_GET_PATCH(version) ((uint32_t)(version) & 0xFFFU)
	return GetVersion() & 0xfff
}
