package xp
import "errors"
type Result struct{Current int64 `json:"current"`;Target int64 `json:"target"`;Remaining int64 `json:"remaining"`}
func Plan(current,target int64)(Result,error){if current<0||target<0{return Result{},errors.New("XP must be non-negative")};if target<current{return Result{},errors.New("target must be >= current XP")};return Result{current,target,target-current},nil}
