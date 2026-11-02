package xp
import "testing"
func TestPlan(t *testing.T){r,e:=Plan(120,250);if e!=nil{t.Fatal(e)};if r.Remaining!=130{t.Fatalf("got %d",r.Remaining)}}
