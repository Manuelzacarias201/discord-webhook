package application
//prueba03
import "fmt"

func GenerateMessageToDiscord(base, titulo, repoFullName, user, urlPullRequest string) string {

	return fmt.Sprintf("Nuevo pull request a la rama %s en el repositorio %s\nTítulo: %s\nAutor: %s\nDetalles: %s", base, repoFullName, titulo, user, urlPullRequest)

}
