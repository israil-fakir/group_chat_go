<h1>group_chat_go</h1>

<p>
A simple <b>multi-user chat application</b> built with Go (Golang).  
This project demonstrates Go's concurrency features (<code>goroutines</code>, <code>channels</code>, <code>mutex</code>) 
by creating a real-time <b>terminal-based chatroom</b> over TCP.
</p>

<hr>

<h2>Features</h2>
<ul>
  <li>Multi-client support (chatroom style)</li>
  <li>Username registration</li>
  <li>Real-time messaging</li>
  <li>Broadcast system to all connected users</li>
  <li>Join/leave notifications</li>
  <li>Ignores empty messages for clean output</li>
</ul>

<hr>

<h2>Project Structure</h2>
<pre>
group_chat_go/
│── server.go   # Chat server
│── client.go   # Chat client
</pre>

<hr>

<h2>Getting Started</h2>

<h3>1. Clone the repo</h3>
<pre><code>
git clone https://github.com/your-username/group_chat_go.git
cd group_chat_go
</code></pre>

<h3>2. Run the server</h3>
<pre><code>
go run server.go
</code></pre>

<h3>3. Run a client</h3>
<pre><code>
go run client.go
</code></pre>

<p>You can start <b>multiple clients</b> in different terminals, all connecting to the same server.</p>

<hr>

<h2>📸 Example (Terminal Screenshot)</h2>
<pre>
# Terminal 1 (server)
Chat server started on :8080
New client joined: Alice
New client joined: Bob

# Terminal 2 (Alice)
Enter your username: Alice
Bob joined the chat
Bob: Hello Alice!
Alice: Hi Bob!

# Terminal 3 (Bob)
Enter your username: Bob
Alice: Hi Bob!
Bob: Hello Alice!
</pre>

<hr>

 
 
<!-- <h2>📌 Next Steps (Future Improvements)</h2>
<ul>
  <li>Add <code>/quit</code> command to exit gracefully</li>
  <li>List active users</li>
  <li>Save chat history to file</li>
  <li>WebSocket support for a web-based version</li>
</ul>

<hr> -->

<h2>Built With</h2>
<ul>
  <li><a href="https://golang.org/">Go</a> 1.23+</li>
</ul>

