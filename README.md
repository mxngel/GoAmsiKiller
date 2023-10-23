# GoAmsiKiller
Simple port of [Amsi-Killer](https://github.com/ZeroMemoryEx/Amsi-Killer) (AMSI lifetime bypass via patched byte in AmsiOpenSession) made in Golang to avoid signature-based detections }:)

![image](https://raw.githubusercontent.com/mxngel/GoAmsiKiller/main/PoC.png)

# Note
The patch is applied to byte 0x74 at offset 3 from the base address of AmsiOpenSession. Apparently and according to the tests performed, it can be optimized by avoiding the process memory read and the pattern search algorithm.
