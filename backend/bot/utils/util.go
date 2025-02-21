package utils

var fields = make(map[string]string)
var positions = make(map[string]string)

func FindField(key string) (value string) {
	fields["dev"] = "Разработка"
	fields["sec"] = "Безопасность"
	fields["devops"] = "DevOps"
	fields["science"] = "Научная деятельность"
	fields["org"] = "Руководитель"

	return fields[key]
}

func FindPosition(key string) (value string) {
	positions["front"] = "Фронтенд"
	positions["back"] = "Бекенд"
	positions["ml"] = "Машинное обучение"
	positions["design"] = "Дизайн(UI/UX)"
	positions["full"] = "Фуллстак"
	positions["qa"] = "QA"
	positions["game-dev"] = "ГеймДев"
	positions["mobile"] = "МобайлДев"
	positions["web"] = "Web"
	positions["crypto"] = "Crypto"
	positions["pwn"] = "PWN"
	positions["forensic"] = "FORENSIC"
	positions["admin"] = "Администрирование/сети"
	positions["osint"] = "OSINT"
	positions["joy"] = "JOY"
	positions["stegano"] = "Стеганография"

	return positions[key]
}
