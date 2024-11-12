package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceEntityTable(layout string, entity model.EntityTable) string {
	var entityOut string
	entityOut += "[ENTITIES:Table]\n"
	entityTableSectionLayout := getEntitySection(layout, `(?s)\[ENTITIES:Table\](.*?)\[`)
	entityTableSectionLayout = strings.ReplaceAll(entityTableSectionLayout, "{ENTITY NAME}", entity.Name)
	entityOut += entityTableSectionLayout
	entityOut += "\n"
	return entityOut
}

func replaceEntityMember(layout string, entity model.EntityMember) string {
	var entityOut string
	entityOut += "[ENTITIES:Member]\n"
	entityMemberSectionLayout := getEntitySection(layout, `(?s)\[ENTITIES:Member\](.*?)\[`)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY NAME}", entity.Name)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY DATA Name}", entity.Data.Address)
	entityMemberSectionLayout = strings.ReplaceAll(entityMemberSectionLayout, "{ENTITY DATA Address}", entity.Data.Address)
	entityOut += entityMemberSectionLayout
	entityOut += "\n"
	return entityOut
}

func getEntitySection(layout string, regexStr string) (entitySectionLayout string) {
	// get entity table section
	reEntityTable := regexp.MustCompile(regexStr)
	matchesEntity := reEntityTable.FindStringSubmatch(layout)
	if len(matchesEntity) > 1 {
		entitySectionLayout = matchesEntity[1]
	}
	return
}
