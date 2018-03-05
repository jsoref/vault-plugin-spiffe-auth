#!/bin/bash

nohup spire-server run -config /opt/spire/conf/server/server.conf > server.out 2>&1 &
sleep 10
spire-agent run -config /opt/spire/conf/agent/agent.conf -joinToken $(spire-server token generate -spiffeID spiffe://example.org/host | sed 's/Token: //') 
