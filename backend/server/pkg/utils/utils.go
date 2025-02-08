package utils

func RemoveDuplicates(input []string) []string {
	uniqueMap := make(map[string]struct{}) // Используем map для отслеживания уникальных значений
	var result []string                    // Результирующий массив

	for _, value := range input {
		if _, exists := uniqueMap[value]; !exists {
			uniqueMap[value] = struct{}{}  // Добавляем значение в map
			result = append(result, value) // Добавляем значение в результат
		}
	}

	return result
}
