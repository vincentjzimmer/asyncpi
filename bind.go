package asyncpi

// Binding.
// This file contains functions for name binding.

import (
	"log"
)

// Bind takes a parsed process p and returned a process with valid binding.
func Bind(p Process) Process {
	return bind(p, []Name{})
}

func bind(p Process, boundNames []Name) Process {
	switch proc := p.(type) {
	case *NilProcess:
		return proc
	case *Repeat:
		return bind(proc.Proc, boundNames)
	case *Par:
		for i := range proc.Procs {
			bind(proc.Procs[i], boundNames)
		}
		return proc
	case *Recv:
		names := make([]Name, len(boundNames))
		for i := range boundNames {
			names[i] = boundNames[i]
		}
		names = append(names, proc.Vars...)
		for _, v := range proc.Vars {
			for j := 0; j < len(names)-len(proc.Vars); j++ {
				if v.Name() == names[j].Name() {
					log.Println("Warning: rebinding name", v.Name(), "in recv")
					names = append(names[:j], names[j+1:]...)
				}
			}
		}
		for i, bn := range names {
			if proc.Chan.Name() == bn.Name() { // Found bound Chan
				proc.Chan = names[i]
			}
		}
		proc.Cont = bind(proc.Cont, names)
		return proc
	case *Send:
		count := 0
		for i, bn := range boundNames {
			for j, v := range proc.Vals {
				if bn.Name() == v.Name() { // Found bound name.
					proc.Vals[j] = boundNames[i]
					count++
				}
			}
		}
		for i, bn := range boundNames {
			if bn.Name() == proc.Chan.Name() { // Found bound Chan.
				proc.Chan = boundNames[i]
				count++
			}
		}
		if count < len(proc.Vals)+1 {
			log.Println("Warning:", len(proc.Vals)+1-count, "names are left unbound")
		}
		return proc
	case *Restrict:
		names := append(boundNames, proc.Name)
		for i := 0; i < len(names)-1; i++ {
			if names[i].Name() == proc.Name.Name() {
				log.Println("Warning: rebinding name", proc.Name.Name(), "in restrict")
				names = append(names[:i], names[i+1:]...)
			}
		}
		proc.Proc = bind(proc.Proc, names)
		return proc
	default:
		log.Fatalln("Bind: Unknown process type", proc)
	}
	return proc
}
