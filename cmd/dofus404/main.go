package main

import (
 "flag"
 "fmt"
 "log"
 "net/http"
 "os"
 "github.com/uncensored404/404sec-dofus/internal/builds"
 "github.com/uncensored404/404sec-dofus/internal/damage"
 "github.com/uncensored404/404sec-dofus/internal/server"
 "github.com/uncensored404/404sec-dofus/internal/xp"
)
func usage(){fmt.Println("404Sec Dofus Suite\nCommands: builds list | damage --base N --power N --percent N --flat N --resist N | xp --current N --target N | api --addr 127.0.0.1:8787")}
func main(){if len(os.Args)<2{usage();return};switch os.Args[1]{
case "help","--help":usage()
case "builds":if len(os.Args)<3||os.Args[2]!="list"{usage();return};all,e:=builds.Load("data/builds.json");if e!=nil{log.Fatal(e)};for _,b:=range all{fmt.Printf("%-12s %-28s level=%d power=%d vitality=%d\\n",b.ID,b.Name,b.Level,b.Stats.Power,b.Stats.Vitality)}
case "damage":fs:=flag.NewFlagSet("damage",flag.ExitOnError);base:=fs.Float64("base",100,"Base damage");power:=fs.Float64("power",0,"Power stat");percent:=fs.Float64("percent",0,"Percent damage bonus");flat:=fs.Float64("flat",0,"Flat bonus");resist:=fs.Float64("resist",0,"Resistance percent");_ = fs.Parse(os.Args[2:]);r,e:=damage.Calculate(damage.Input{Base:*base,Power:*power,PercentDamage:*percent,FlatBonus:*flat,ResistancePercent:*resist});if e!=nil{log.Fatal(e)};fmt.Printf("Estimated damage: %.2f (generic formula; verify game mechanics)\\n",r.Final)
case "xp":fs:=flag.NewFlagSet("xp",flag.ExitOnError);current:=fs.Int64("current",0,"Current XP");target:=fs.Int64("target",0,"Target XP");_ = fs.Parse(os.Args[2:]);r,e:=xp.Plan(*current,*target);if e!=nil{log.Fatal(e)};fmt.Printf("Remaining XP: %d\\n",r.Remaining)
case "api":fs:=flag.NewFlagSet("api",flag.ExitOnError);addr:=fs.String("addr","127.0.0.1:8787","Listen address");_ = fs.Parse(os.Args[2:]);log.Printf("Dashboard: http://%s",*addr);log.Fatal(http.ListenAndServe(*addr,server.New("data/builds.json").Handler()))
default:usage()}}
