package drivers

type Board string

const (
	ScriptStickV1 = Board("V1")
)

func BoardName(b Board) string {
	switch b {
	case ScriptStickV1:
		return "scriptstick v1"
	default:
		return "unknown"
	}
}