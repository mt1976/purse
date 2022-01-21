package dao

import dm "github.com/mt1976/purse/datamodel"

func template_ExtraField_Extra_ExtraStore(recItem dm.Template) string {
	return "Extra"
}

func template_ExtraField2_Extra_ExtraStore(recItem dm.Template) string {
	return "Extra"
}

func template_FIELD2_OverrideStore(recItem dm.Template) string {
	return dm.Template_FIELD2 + " Override"
}

func template_FIELD2_OverrideFetch(recItem dm.Template) string {
	return dm.Template_FIELD2 + " Override"
}

func template_ExtraField_Extra(recItem dm.Template) string {
	return template_ExtraField_Extra_ExtraStore(recItem)
}

func template_ExtraField2_Extra(recItem dm.Template) string {
	return template_ExtraField2_Extra_ExtraStore(recItem)
}
