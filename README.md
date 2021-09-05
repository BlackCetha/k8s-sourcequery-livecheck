# k8s-sourcequery-livecheck

This is a very simple program that checks wether it can reach a server using valve's Server Query/A2S-protocol.

It will perform a full info request including challenge negotiation and validation.
The return code tells you wether it was successful (0) or not (1). Any error messages will be printed to stdout.

This program is intended to be used as a [Kubernetes liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/),
but other applications are certainly possible.

Configuration is kept to the minimum: The program accepts a single, optional parameter for the target address or DNS name.
The default is querying localhost on port 27015.

The timeout is set to three seconds and is not configurable.
