<h2><a href="https://leetcode.com/problems/check-if-two-chessboard-squares-have-the-same-color">Check if Two Chessboard Squares Have the Same Color</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given two strings, <code>coordinate1</code> and <code>coordinate2</code>, representing the coordinates of a square on an <code>8 x 8</code> chessboard.</p>

<p>Below is the chessboard for reference.</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/17/screenshot-2021-02-20-at-22159-pm.png" style="width: 400px; height: 396px;" /></p>

<p>Return <code>true</code> if these two squares have the same color and <code>false</code> otherwise.</p>

<p>The coordinate will always represent a valid chessboard square. The coordinate will always have the letter first (indicating its column), and the number second (indicating its row).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">coordinate1 = &quot;a1&quot;, coordinate2 = &quot;c3&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>

<p><strong>Explanation:</strong></p>

<p>Both squares are black.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">coordinate1 = &quot;a1&quot;, coordinate2 = &quot;h3&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">false</span></p>

<p><strong>Explanation:</strong></p>

<p>Square <code>&quot;a1&quot;</code> is black and <code>&quot;h3&quot;</code> is white.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>coordinate1.length == coordinate2.length == 2</code></li>
	<li><code>&#39;a&#39; &lt;= coordinate1[0], coordinate2[0] &lt;= &#39;h&#39;</code></li>
	<li><code>&#39;1&#39; &lt;= coordinate1[1], coordinate2[1] &lt;= &#39;8&#39;</code></li>
</ul>
