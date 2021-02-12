# ASCII Packets

Draw ASCII packet sequence diagrams!

Create text file containing the contents of the diagram. The format is:
```
node1-node2: SYN
```
node 1 sends a message to node 2, annotated as "SYN".

```
node1=node2: data
```
node 1 and node 2 have a bidrectional connection, annotated as 'data'.

Example files can be found in test_input.txt

## Example:
```
// --- test_input.txt ---
client-server: SYN
server-client: SYN/ACK
client-server: ACK
server=client: data
// ----------------------

./ascii_packets -path test_input.txt -node1 client -node2 server 
*---------------------------------*
|        |-----SYN------>|        |
| client |<---SYN/ACK----| server |
|        |-----ACK------>|        |
|        |<----data----->|        |
*---------------------------------*

```