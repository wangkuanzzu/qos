# Config Parameters Settings

> config.toml文件记录的是全节点启动时候的节点相关参数设置.

## main base config options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|proxy_app|TCP or UNIX socket address of the ABCI application,or the name of an ABCI application compiled in with the Tendermint binary.<br>abci应用程序的TCP或UNIX套接字地址，或使用tendermint二进制文件编译的abci应用程序的名称。|tcp://127.0.0.1:26658|tcp://127.0.0.1:26658|--|
|moniker|A custom human readable name for this node.<br>节点自定义的可读名称|--|--|--|
|fast_sync（bool）|If this node is many blocks behind the tip of the chain, FastSync allows them to catchup quickly by downloading blocks in parallel and verifying their commits. <br> 如果这个节点落后链的最新高度有许多块，那么fastsync允许它们通过并行下载块并验证它们的提交来快速捕获。|true|true|--|
|db_backend|Database backend: leveldb or memdb or cleveldb<br>数据库后端选择：leveldb or memdb or cleveldb|leveldb|leveldb|--|
|db_dir|Database directory<br>数据库目录|data|data|--|
|log_level|Output level for logging, including package level options<br>日志记录的输出级别，包括包级别选项|main:info,state:info,`*`:error|main:info,state:info,`*`:error|--|
|log_format|Output format: 'plain' (colored text) or 'json'<br>输出格式：带颜色文本 or json格式|plain|plain|--|

## additional base config options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|genesis_file|Path to the JSON file containing the initial validator set and other meta data<br>包含初始验证人集合和其他元数据的json文件路径|config/genesis.json|config/genesis.json|--|
|priv_validator_key_file|Path to the JSON file containing the private key to use as a validator in the consensus protocol<br>包含在共识协议中作为验证人的私钥信息的json文件路径|config/priv_validator_key.json|config/priv_validator_key.json|--|
|priv_validator_state_file|Path to the JSON file containing the last sign state of a validator<br>包含验证人最后一个签名状态的JSON文件的路径|data/priv_validator_state.json|data/priv_validator_state.json|--|
|priv_validator_laddr|TCP or UNIX socket address for Tendermint to listen on for connections from an external PrivValidator process<br>tendermint要侦听外部privvalidator进程的连接的TCP或UNIX套接字地址|||--|
|node_key_file|Path to the JSON file containing the private key to use for node authentication in the p2p protocol<br>包含p2p协议中用于节点身份验证的私钥的JSON文件的路径|config/node_key.json|config/node_key.json|--|
|abci|Mechanism to connect to the ABCI application: socket or grpc<br>连接ABCI应用程序的机制：socket or grpc|socket|socket|--|
|prof_laddr|TCP or UNIX socket address for the profiling server to listen on<br>要监听的资料服务器的TCP或UNIX套接字地址|localhost:6060|localhost:6060|--|
|filter_peers（bool）|If true, query the ABCI app on connecting to a new peer. so the app can decide if we should keep the connection or not.<br>如果为true，在连接到新的peer时候查询abci应用程序，应用程序来决定我们是否应该保持连接|false|false|--|

## rpc server configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|laddr|TCP or UNIX socket address for the RPC server to listen on<br>要侦听的RPC服务器的TCP或UNIX套接字地址|tcp://0.0.0.0:26657|tcp://0.0.0.0:26657|--|
|cors_allowed_origins|A list of origins a cross-domain request can be executed from Default value '[]' disables cors support Use '["*"]' to allow any origin<br>可以从中执行跨域请求的源列表，默认值“[]”禁用CORS支持，使用“[”*“]”允许任何来源|[]|[]|--|
|cors_allowed_methods|A list of methods the client is allowed to use with cross-domain requests<br>允许客户端与跨域请求一起使用的方法列表|["HEAD", "GET", "POST", ]|["HEAD", "GET", "POST", ]|--|
|cors_allowed_headers|A list of non simple headers the client is allowed to use with cross-domain requests<br>允许客户端与跨域请求一起使用的非简单头列表|["Origin", "Accept", "Content-Type", "X-Requested-With", "X-Server-Time", ]|["Origin", "Accept", "Content-Type", "X-Requested-With", "X-Server-Time", ]|--|
|grpc_laddr|TCP or UNIX socket address for the gRPC server to listen on. NOTE: This server only supports /broadcast_tx_commit<br>要侦听的GRPC服务器的TCP或UNIX套接字地址，此服务器只支持/broadcast_tx_commit|||--|
|grpc_max_open_connections|Maximum number of simultaneous connections.If you want to accept a larger number than the default, make sure you increase your OS limits.<br>同时连接的最大数目。如果需要接受比默认值更大的数，那就需要增加操作系统的限制。|900|900|Should be < {ulimit -Sn} - {MaxNumInboundPeers} - {MaxNumOutboundPeers} - {N of wal, db and other open files}. 1024 - 40 - 10 - 50 = 924 = ~900|
|unsafe(bool)|Activate unsafe RPC commands like /dial_seeds and /unsafe_flush_mempool<br>激活不安全的RPC命令，如/dial_seeds和/unsafe_flush_mempool|false|false|--|
|max_open_connections|Maximum number of simultaneous connections (including WebSocket). Does not include gRPC connections. See grpc_max_open_connections<br>同时连接的最大数目（包括WebSocket）。不包括GRPC连接。请参阅GRPC_max_open_connections|900|900|--|
|max_subscription_clients(int)|Maximum number of unique clientIDs that can /subscribe. If you're using /broadcast_tx_commit, set to the estimated maximum number of broadcast_tx_commit calls per block.<br>可以/subscribe的唯一clientIDs的最大数目.如果你正在使用/broadcast_tx_commit，需要为每一块预估最大调用/broadcast_tx_commit的次数。|100|100|--|
|max_subscriptions_per_client(int)|Maximum number of unique queries a given client can /subscribe to. If you're using GRPC (or Local RPC client) and /broadcast_tx_commit, set to the estimated # maximum number of broadcast_tx_commit calls per block.<br>给定客户端可以/subscribe to的唯一查询的最大数目。如果使用GRPC（或本地RPC客户端）和/broadcast_tx_commit，请设置为： 估计每块的最大broadcast_tx_commit 调用次数。|5|5|--|
|timeout_broadcast_tx_commit|How long to wait for a tx to be committed during /broadcast_tx_commit. WARNING: Using a value larger than 10s will result in increasing the global HTTP write timeout, which applies to all connections and endpoints.<br>在调用/broadcast_tx_commit期间等待tx提交的时间。如果使用大于10s的值，会导致全局的http写入超时增加，这会适用于所有的连接和终结点。|10s|10s|--|
|tls_cert_file|The name of a file containing certificate that is used to create the HTTPS server. If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.<br>包含用于创建HTTPS服务器的证书的文件名。如果证书由证书颁发机构签名，那么证书文件应该是服务器证书、任何中间文件和CA证书的串联。|||both tls_cert_file and tls_key_file must be present for Tendermint to create HTTPS server. Otherwise, HTTP server is run.<br>必须同时存在tls_cert_file和tls_key_file，Tendermint才能创建HTTPS服务器。否则，将运行HTTP服务器。|
|tls_key_file|The name of a file containing matching private key that is used to create the HTTPS server.<br>包含用于创建HTTPS服务器的匹配私钥的文件名。|||--|

## peer to peer configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|laddr|Address to listen for incoming connections<br>监听向我们的连接请求的地址|tcp://0.0.0.0:26656|tcp://0.0.0.0:26656|--|
|external_address|Address to advertise to peers for them to dial. If empty, will use the same port as the laddr, and will introspect on the listener or use UPnP to figure out the address.<br>向peers公布的地址，供他们拨号连接。如果为空，将使用与laddr相同的端口，并对监听器进行内省或使用upnp计算出地址。|||--|
|seeds|Comma separated list of seed nodes to connect to<br>要连接到的种子节点的逗号分隔列表||指定地址|--|
|persistent_peers|Comma separated list of nodes to keep persistent connections to<br>要一直保持连接的peer的逗号分隔列表|||--|
|upnp(bool)|UPNP port forwarding<br>UPNP端口转发|false|false|--|
|addr_book_file|Path to address book<br>地址表文件路径|config/addrbook.json|config/addrbook.json|--|
|addr_book_strict(bool)|Set true for strict address routability rules Set false for private or local networks<br>为严格的地址可路由性规则设置为真，为专用网络或本地网络设置为假|true|true|--|
|max_num_inbound_peers(int)|Maximum number of inbound peers<br>允许其他节点连接我们节点的最大数目|40|?|--|
|max_num_outbound_peers(int)|Maximum number of outbound peers to connect to, excluding persistent peers<br>我们节点向外节点连接的最大数目，不包含始终连接的peer数量|10|?|--|
|flush_throttle_timeout|Time to wait before flushing messages out on the connection<br>在清除连接上的消息之前等待的时间|100ms|100ms|--|
|max_packet_msg_payload_size(int)|Maximum size of a message packet payload, in bytes<br>消息包负载的最大大小（字节）|1024|?|--|
|send_rate(int)|Rate at which packets can be sent, in bytes/second<br>发送数据包的速率（字节/秒）|5120000|5120000|--|
|recv_rate(int)|Rate at which packets can be received, in bytes/second<br>接受数据包的速率（字节/秒）|5120000|5120000|--|
|pex(bool)|Set true to enable the peer-exchange reactor<br>设置为true以启用peer-exchange反应器|true|true|--|
|seed_mode(bool)|Seed mode, in which node constantly crawls the network and looks for peers. If another node asks it for addresses, it responds and disconnects. Does not work if the peer-exchange reactor is disabled.<br>种子模式，在这种模式下，节点不断地爬行网络并寻找对等点。如果另一个节点向它请求地址，它会响应并断开连接。如果对等交换反应器被禁用，则不工作。|true|false|--|
|private_peer_ids|Comma separated list of peer IDs to keep private (will not be gossiped to other peers)<br>以逗号分隔的对等ID列表以保持私有（不会与其他对等方交换信息）|||--|
|allow_duplicate_ip(bool)|Toggle to disable guard against peers connecting from the same ip.<br>切换到禁用：用来防护来自同一IP的peer连接|false|false|--|
|handshake_timeout|Peer connection configuration.<br>peer连接的握手超时时间|20s|20s|--|
|dial_timeout|Peer connection configuration.<br>peer拨号连接超时时间|3s|3s|--|

## mempool configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|recheck|对mempool的交易更新时进行recheck|true|true|--|
|broadcast|节点是否可以将收到的交易进行全网广播|true|true|--|
|wal_dir|日志路径|||--|
|size|Maximum number of transactions in the mempool<br>mempool中可以存储的最大交易数量|5000|5000|--|
|max_txs_bytes(int)|Limit the total size of all txs in the mempool.<br>mempool中可以存储的最大交易大小|1073741824b=1G|1073741824b=1G|--|
|cache_size|Size of the cache (used to filter transactions we saw earlier) in transactions<br>mempool缓存中可存放的最大交易数量|10000|10000|--|

## consensus configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|wal_file|wal路径|data/cs.wal/wal|data/cs.wal/wal|--|
|timeout_propose|提议阶段超时初始时长|3s|3s|--|
|timeout_propose_delta|提议阶段超时后续增加时长，round*delta|500ms|500ms|--|
|timeout_prevote|预投票阶段超时初始时长|1s|1s|--|
|timeout_prevote_delta|预投票超时后续增加时长，round*delta|500ms|500ms|--|
|timeout_precommit|预提交超时初始时长|1s|1s|--|
|timeout_precommit_delta|预提交超时后续增加时长，round*delta|500ms|500ms|--|
|timeout_commit|提交超时初始时长|5s|5s|--|
|skip_timeout_commit(bool)|Make progress as soon as we have all the precommits (as if TimeoutCommit = 0)<br>设置为true：当我们拥有了所有precommits，进入新一轮共识|false|false|--|
|create_empty_blocks|EmptyBlocks mode and possible interval between empty blocks<br>是否允许打空块|true|true|--|
|create_empty_blocks_interval|打空块的时间间隔|0s|0s|--|
|peer_gossip_sleep_duration|Reactor sleep duration parameters<br>peer之间的数据交换休眠间隔|100ms|100ms|--|
|peer_query_maj23_sleep_duration|Reactor sleep duration parameters<br>peer收到大于2/3的vote时候休眠间隔|2s|2s|--|

## transactions indexer configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|indexer|1) "null" <br> 2) "kv" (default) - the simplest possible indexer, backed by key-value storage (defaults to levelDB; see DBBackend).<br>最简单的索引器，由键值存储支持（默认为leveldb；请参阅dbbackend）|kv|kv|--|
|index_tags|It's recommended to index only a subset of tags due to possible memory bloat. This is, of course, depends on the indexer's DB and the volume of transactions.<br>指定tags用于创建kv索引|||--|
|index_all_tags(bool)|When set to true, tells indexer to index all tags (predefined tags: "tx.hash", "tx.height" and all tags from DeliverTx responses).<br>将DeliverTx中所有的tags用于创建kv索引|true|true|--|

## instrumentation configuration options

|参数名称（无特殊说明参数类型均为string）|参数释义(en & cn)|自己节点|全节点部署推荐|备注|
|--|--|--|--|--|
|prometheus|When true, Prometheus metrics are served under /metrics on PrometheusListenAddr.Check out the documentation for the list of available metrics.<br>设置为true，开启prometheus服务，prometheus metrics在prometheusListenAddr上的/metrics下提供。|false|？|--|
|prometheus_listen_addr|Address to listen for Prometheus collector(s) connections<br>监听Prometheus收集器连接的地址|:26660|:26660|--|
|max_open_connections（int）|Maximum number of simultaneous connections. If you want to accept a larger number than the default, make sure you increase your OS limits. 0 - unlimited.<br>同时连接的最大数目|3|3|--|
|namespace|Instrumentation namespace<br>命名空间|tendermint|tendermint|--|
