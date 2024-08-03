<h2><a href="https://leetcode.com/problems/find-the-number-of-winning-players">Find the Number of Winning Players</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given an integer <code>n</code> representing the number of players in a game and a 2D array <code>pick</code> where <code>pick[i] = [x<sub>i</sub>, y<sub>i</sub>]</code> represents that the player <code>x<sub>i</sub></code> picked a ball of color <code>y<sub>i</sub></code>.</p>

<p>Player <code>i</code> <strong>wins</strong> the game if they pick <strong>strictly more</strong> than <code>i</code> balls of the <strong>same</strong> color. In other words,</p>

<ul>
	<li>Player 0 wins if they pick any ball.</li>
	<li>Player 1 wins if they pick at least two balls of the <em>same</em> color.</li>
	<li>...</li>
	<li>Player <code>i</code> wins if they pick at least<code>i + 1</code> balls of the <em>same</em> color.</li>
</ul>

<p>Return the number of players who <strong>win</strong> the game.</p>

<p><strong>Note</strong> that <em>multiple</em> players can win the game.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 4, pick = [[0,0],[1,0],[1,0],[2,1],[2,1],[2,0]]</span></p>

<p><strong>Output:</strong> <span class="example-io">2</span></p>

<p><strong>Explanation:</strong></p>

<p>Player 0 and player 1 win the game, while players 2 and 3 do not win.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 5, pick = [[1,1],[1,2],[1,3],[1,4]]</span></p>

<p><strong>Output:</strong> <span class="example-io">0</span></p>

<p><strong>Explanation:</strong></p>

<p>No player wins the game.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">n = 5, pick = [[1,1],[2,4],[2,4],[2,4]]</span></p>

<p><strong>Output:</strong> <span class="example-io">1</span></p>

<p><strong>Explanation:</strong></p>

<p>Player 2 wins the game by picking 3 balls with color 4.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= pick.length &lt;= 100</code></li>
	<li><code>pick[i].length == 2</code></li>
	<li><code>0 &lt;= x<sub>i</sub> &lt;= n - 1 </code></li>
	<li><code>0 &lt;= y<sub>i</sub> &lt;= 10</code></li>
</ul>
