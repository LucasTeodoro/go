import(
	"soma"
)

func TestSoma() {
	result := soma(5, 5)
	if result != 10 {
		t.Error("Valor diferente de 10")
	}
}