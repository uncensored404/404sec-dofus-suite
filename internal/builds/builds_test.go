package builds
import "testing"
func TestLoad(t *testing.T){b,e:=Load("../../data/builds.json");if e!=nil{t.Fatal(e)};if len(b)!=2{t.Fatalf("got %d builds",len(b))}}
