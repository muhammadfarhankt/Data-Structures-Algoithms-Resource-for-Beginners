<h2><a href="https://leetcode.com/problems/largest-even-number">Largest Even Number</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given a string <code>s</code> consisting only of the characters <code>&#39;1&#39;</code> and <code>&#39;2&#39;</code>.</p>

<p>You may delete any number of characters from <code>s</code> without changing the order of the remaining characters.</p>

<p>Return the <strong>largest possible resultant string</strong> that represents an <strong>even</strong> integer. If there is no such string, return the empty string <code>&quot;&quot;</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;1112&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;1112&quot;</span></p>

<p><strong>Explanation:</strong></p>

<p>The string already represents the largest possible even number, so no deletions are needed.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;221&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;22&quot;</span></p>

<p><strong>Explanation:</strong></p>

<p>Deleting <code>&#39;1&#39;</code> results in the largest possible even number which is equal to 22.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;1&quot;</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;&quot;</span></p>

<p><strong>Explanation:</strong></p>

<p>There is no way to get an even number.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> consists only of the characters <code>&#39;1&#39;</code> and <code>&#39;2&#39;</code>.</li>
</ul>
