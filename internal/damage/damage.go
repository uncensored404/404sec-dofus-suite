package damage
import "errors"
type Input struct{Base float64 `json:"base"`;Power float64 `json:"power"`;PercentDamage float64 `json:"percent_damage"`;FlatBonus float64 `json:"flat_bonus"`;ResistancePercent float64 `json:"resistance_percent"`}
type Result struct{Base float64 `json:"base"`;AfterStats float64 `json:"after_stats"`;BeforeResistance float64 `json:"before_resistance"`;Final float64 `json:"final"`}
func Calculate(in Input)(Result,error){if in.Base<0||in.Power<0||in.PercentDamage < -100||in.FlatBonus<0{return Result{},errors.New("invalid negative damage/stat value")};if in.ResistancePercent<0||in.ResistancePercent>100{return Result{},errors.New("resistance must be between 0 and 100")};after:=in.Base*(1+(in.Power+in.PercentDamage)/100);before:=after+in.FlatBonus;if before<0{before=0};return Result{in.Base,after,before,before*(1-in.ResistancePercent/100)},nil}
