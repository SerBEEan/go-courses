package converter

/* Структура приватных полей конвертера */
type converter struct {
	symbols []string
	values  []int
}

/* Интерфейс конвертера. Его публичные методы методы */
type Converter interface {
	IntToRoman(num int) string
}

/* Метод перевода числа из десятичной в римскую систему счисления */
func (c *converter) IntToRoman(num int) string {
	var result string

	for i := 0; i < len(c.values); i++ {
		currentValue := c.values[i]
		currentSymbol := c.symbols[i]

		for num >= currentValue {
			result += currentSymbol
			num -= currentValue
		}
	}
	return result
}

/* Конструктор */
func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
