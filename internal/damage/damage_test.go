package damage
import "testing"
func TestCalculate(t *testing.T){r,e:=Calculate(Input{Base:100,Power:250,PercentDamage:20,FlatBonus:12,ResistancePercent:15});if e!=nil{t.Fatal(e)};if r.Final!=327.25{t.Fatalf("got %v",r.Final)}}
func TestInvalidResistance(t *testing.T){if _,e:=Calculate(Input{Base:1,ResistancePercent:101});e==nil{t.Fatal("expected error")}}
