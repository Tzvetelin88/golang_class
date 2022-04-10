
# Go Routine is a lightweight thread managed by the Go runtime. Executing code line be line that can be used to handle blocking code.
### Adding <go> keyword will add a new thread go routine.
### One CPU is running only one routine!
### When we have multple CPUs - we use go Scheduler which will spread go routine to available CPUs
### Concurency !== parallelism
### Concurency - our program have the availability to run different tasks, multiple routines.
### Parallelism - our program have the availability to run different tasks at the same exact time, at the same nano second!
### We always have one main routine when we launched the program.
### Child routines created by the keyword <go>

### Channel - sharing data. To have communication between all routines we can connect the channel.
### Main Routine will check when all child routines finished and then main will finish too, no sooner!
### Channels are tight and has to be from same type like 'string'
