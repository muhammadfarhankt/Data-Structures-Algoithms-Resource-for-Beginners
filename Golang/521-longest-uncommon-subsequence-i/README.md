<h2><a href="https://leetcode.com/problems/longest-uncommon-subsequence-i">Longest Uncommon Subsequence I</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>Given two strings <code>a</code> and <code>b</code>, return <em>the length of the <strong>longest uncommon subsequence</strong> between </em><code>a</code> <em>and</em> <code>b</code>. <em>If no such uncommon subsequence exists, return</em> <code>-1</code><em>.</em></p>

<p>An <strong>uncommon subsequence</strong> between two strings is a string that is a <strong><span data-keyword="subsequence-string">subsequence</span> of exactly one of them</strong>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> a = &quot;aba&quot;, b = &quot;cdc&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong> One longest uncommon subsequence is &quot;aba&quot; because &quot;aba&quot; is a subsequence of &quot;aba&quot; but not &quot;cdc&quot;.
Note that &quot;cdc&quot; is also a longest uncommon subsequence.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> a = &quot;aaa&quot;, b = &quot;bbb&quot;
<strong>Output:</strong> 3
<strong>Explanation:</strong>&nbsp;The longest uncommon subsequences are &quot;aaa&quot; and &quot;bbb&quot;.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> a = &quot;aaa&quot;, b = &quot;aaa&quot;
<strong>Output:</strong> -1
<strong>Explanation:</strong>&nbsp;Every subsequence of string a is also a subsequence of string b. Similarly, every subsequence of string b is also a subsequence of string a. So the answer would be <code>-1</code>.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= a.length, b.length &lt;= 100</code></li>
	<li><code>a</code> and <code>b</code> consist of lower-case English letters.</li>
</ul>
