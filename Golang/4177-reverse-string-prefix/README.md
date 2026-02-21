<h2><a href="https://leetcode.com/problems/reverse-string-prefix">Reverse String Prefix</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given a string <code>s</code> and an integer <code>k</code>.</p>

<p>Reverse the first <code>k</code> characters of <code>s</code> and return the resulting string.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;abcd&quot;, k = 2</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;bacd&quot;</span></p>

<p><strong>Explanation:</strong>​​​​​​​</p>

<p>The first <code>k = 2</code> characters <code>&quot;ab&quot;</code> are reversed to <code>&quot;ba&quot;</code>. The final resulting string is <code>&quot;bacd&quot;</code>.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;xyz&quot;, k = 3</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;zyx&quot;</span></p>

<p><strong>Explanation:</strong></p>

<p>The first <code>k = 3</code> characters <code>&quot;xyz&quot;</code> are reversed to <code>&quot;zyx&quot;</code>. The final resulting string is <code>&quot;zyx&quot;</code>.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">s = &quot;hey&quot;, k = 1</span></p>

<p><strong>Output:</strong> <span class="example-io">&quot;hey&quot;</span></p>

<p><strong>Explanation:</strong></p>

<p>The first <code>k = 1</code> character <code>&quot;h&quot;</code> remains unchanged on reversal. The final resulting string is <code>&quot;hey&quot;</code>.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> consists of lowercase English letters.</li>
	<li><code>1 &lt;= k &lt;= s.length</code></li>
</ul>
