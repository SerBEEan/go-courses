package converter

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(num int) string
}

func (c *converter) IntToRoman(num int) string {
	number := num
	result := ""

	for i := 0; i < len(c.values); i++ {
		currentValue := c.values[i]
		currentSymbol := c.symbols[i]

		for number >= currentValue {
			result += currentSymbol
			number -= currentValue
		}
	}

	return result
}

func NewConverter(values []int, symbols []string) Converter {
	converter := &converter{
		values:  values,
		symbols: symbols,
	}

	return converter
}
