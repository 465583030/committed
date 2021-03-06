package db

// // Topic is a replicated state machine that accepts a partitioned type of data
// type Topic struct {
// 	Name    string
// 	Nodes   []*node
// 	network *raftNetwork
// 	wal     *wal.WAL
// }

// func newTopic(name string, nodeCount int) *Topic {
// 	peers := make([]raft.Peer, 0)
// 	ids := make([]uint64, 0)
// 	for i := 1; i <= nodeCount; i++ {
// 		id := uint64(i)
// 		peers = append(peers, raft.Peer{ID: id, Context: nil})
// 		ids = append(ids, id)
// 	}

// 	nt := newRaftNetwork(ids...)

// 	nodes := make([]*node, 0)

// 	for _, id := range ids {
// 		n := startNode(id, peers, nt.nodeNetwork(id))
// 		nodes = append(nodes, n)
// 	}

// 	waitLeader(nodes)

// 	return &Topic{Name: name, Nodes: nodes, network: nt}
// }

// func (t *Topic) stop() {
// 	for i := 0; i < len(t.Nodes); i++ {
// 		t.Nodes[i].Stop()
// 	}
// }

// func (t *Topic) up() bool {
// 	v := false

// 	for i := 0; i < len(t.Nodes); i++ {
// 		if t.Nodes[i].Node.Status().ID > 0 {
// 			v = true
// 		}
// 	}

// 	return v
// }

// // Append a proposal to the topic
// func (t *Topic) Append(ctx context.Context, proposal string) {
// 	n := t.Nodes[0]
// 	n.Propose(ctx, []byte(proposal))
// }

// func (t *Topic) size(ctx context.Context) uint64 {
// 	storage := t.Nodes[0].storage
// 	first, _ := storage.FirstIndex()
// 	last, _ := storage.LastIndex()

// 	entries, error := storage.Entries(first, last+1, uint64(1024*1024))
// 	if error != nil {
// 		fmt.Println("[topic] Error retrieving entries from storage")
// 	}

// 	count := uint64(0)
// 	for _, e := range entries {
// 		if e.Type == raftpb.EntryNormal && len(e.Data) != 0 {
// 			count++
// 		}
// 	}

// 	return count
// }

// // ReadIndex from the topic
// func (t *Topic) ReadIndex(ctx context.Context, index uint64) string {
// 	storage := t.Nodes[0].storage
// 	first, _ := storage.FirstIndex()
// 	last, _ := storage.LastIndex()

// 	entries, error := storage.Entries(first, last+1, uint64(1024*1024))
// 	if error != nil {
// 		fmt.Println("[topic] Error retrieving entries from storage")
// 	}

// 	count := uint64(0)
// 	for _, e := range entries {
// 		if e.Type == raftpb.EntryNormal && len(e.Data) != 0 {
// 			if count == index {
// 				return string(e.Data[:])
// 			}
// 			count++
// 		}
// 	}

// 	// TODO This should be an error
// 	fmt.Println("[topic] Could not find index")
// 	return ""
// }

// // Sync the contents of the topic into a Syncable
// func (t *Topic) Sync(ctx context.Context, s syncable.Syncable) {
// 	size := t.size(ctx)

// 	for i := uint64(0); i < size; i++ {
// 		s.Sync(ctx, []byte(t.ReadIndex(ctx, uint64(i))))
// 	}

// 	for _, n := range t.Nodes {
// 		syncNode(ctx, s, n)
// 	}
// }

// func syncNode(ctx context.Context, s syncable.Syncable, n *node) {
// 	subc := n.syncp.Sub("StoredData")
// 	go func() {
// 		for {
// 			select {
// 			case e := <-subc:
// 				s.Sync(ctx, e.(raftpb.Entry).Data)
// 			default:
// 				time.Sleep(time.Millisecond * 1)
// 			}
// 		}
// 	}()
// }
