package builds
import("encoding/json";"fmt";"os")
type Stats struct{Power int `json:"power"`;Vitality int `json:"vitality"`;Agility int `json:"agility"`;Intelligence int `json:"intelligence"`;Chance int `json:"chance"`;Strength int `json:"strength"`;Critical int `json:"critical"`}
type Build struct{ID string `json:"id"`;Name string `json:"name"`;Class string `json:"class"`;Level int `json:"level"`;Server string `json:"server"`;Stats Stats `json:"stats"`;Gear []string `json:"gear"`;Notes string `json:"notes"`}
func Load(path string)([]Build,error){b,e:=os.ReadFile(path);if e!=nil{return nil,fmt.Errorf("read builds: %w",e)};var out []Build;if e=json.Unmarshal(b,&out);e!=nil{return nil,fmt.Errorf("parse builds: %w",e)};return out,nil}
