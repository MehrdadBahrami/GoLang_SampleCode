package main

import (
	"Mehrdad/ConstantsVar"
	forsyntax "Mehrdad/ForSyntax"
	"Mehrdad/HelloWorld"
	"Mehrdad/Values"
	"Mehrdad/Variables"
	"fmt"
)

var p = fmt.Println

func main() {

	p("--------------------------First-----------------------------------------")
	p("It my main Project")

	p("--------------------------SayHello--------------------------------------")
	hi := HelloWorld.SayHello("Mehrdad", "Bahrami")
	p(hi)

	p("--------------------------GetSampleValue--------------------------------")
	valueModules := Values.GetSampleValue()
	p(valueModules)

	p("--------------------------DefineVariables-------------------------------")
	Samplevariables := Variables.DefineVariables("FirstOne", 1, 2, true, 3, "LastOne")
	p(Samplevariables)

	p("--------------------------DefineConstantsVariable-----------------------")
	SampleConstants := ConstantsVar.DefineConstantsVariable()
	p(SampleConstants)

	p("--------------------------ForLoop---------------------------------------")
	ForLoopSyntax := forsyntax.ForLoop()
	p(ForLoopSyntax)

	p("--------------------------Last------------------------------------------")
}
