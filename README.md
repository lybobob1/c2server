# c2server
First golang project, creating a c2 server

This server receives callbacks through the /callback endpoint (GET for for requesting/receiving tasking, POST for receiving the results of tasking). This server is based around the same design as Silentbreak Security's Throwback

Implant info and tasking are stored in a mysql database

*Caveats: the C2 that this server supports is intentionally insecure and offers literally no obfuscation or encryption
