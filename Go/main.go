package main
 
import (
"fmt"
"github.com/shirou/gopsutil/cpu"
"time"
)
 
func main() {
   //Percent calculates the percentage of cpu used either per CPU or combined.
   percent, _ := cpu.Percent(time.Second,true)
   fmt.Printf("  User: %.2f\n",percent[cpu.CPUser])
   fmt.Printf("  Nice: %.2f\n",percent[cpu.CPNice])
   fmt.Printf("   Sys: %.2f\n",percent[cpu.CPSys])
   fmt.Printf("  Intr: %.2f\n",percent[cpu.CPIntr])
   fmt.Printf("  Idle: %.2f\n",percent[cpu.CPIdle])
   fmt.Printf("States: %.2f\n",percent[cpu.CPUStates])
}