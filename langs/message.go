package langs

import (
	"fmt"
)

func GenerateValidationMessage(Field, Rule string) (message string) {
	switch Rule {
	case "required":
		return fmt.Sprintf("Le champs %s est %s", Field, Rule)
	default:
		return fmt.Sprintf("Le champs %s n'est pas valide", Field)
	}
}
