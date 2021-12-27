What happens if you remove the go-command from the Seek call in the main function?
Answer: If the go-command is removed, the last name in the list will not have a reciver in a uneven list

What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?
Answer: When using new(sync.WaitGroup) a pointer is created to wg but not when using var. At the same time, if we not use "*" in 
        the function, wg is then a copy of the wg in main (aka var wg sync.Waitgroup) which then don't correspond correctly. Therefor, the command wg.Done() is not heard by the main function. 

What happens if you remove the buffer on the channel match? 
Answer: Now the sender can't wait for someone to receive item on the channel which then implies that the last name in an uneven list can't get on to the channel. This will not happen when the count is even because then they all split between senders and recievers.

What happens if you remove the default-case from the case-statement in the main function?
Answer: The select-command will wait for match to send a name but that will not happen if the list is even which creates a deadlock

