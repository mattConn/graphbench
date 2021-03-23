# Architecture

- entities:
	- graph,node,edge,cycle,star
- `list <entity>`

- Properties:
	- connected, bipartite, cycle, tree
	- `is <property>`
- operations
	- add <entity>

# Tentative Interactive Mode Commands

- n
	- list all nodes
- n a 
	- add node a
- n a b ... 
	- add nodes a,b,...
- N 
	- like n but for removal
- e, E 
	- like n,N, but for edges (node pairs)
- c 
	- list all cycles
- c n
	- list all n cycles
- c a b ... 
	- add cycle a -> b -> ... -> a
- s n
	- list all n stars
- s a b c ...
	- add star a -> b, a -> c, a-> ...