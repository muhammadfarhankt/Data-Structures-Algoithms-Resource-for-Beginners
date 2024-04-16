<h2><a href="https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences">Check if All Characters Have Equal Number of Occurrences</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>Given a string <code>s</code>, return <code>true</code><em> if </em><code>s</code><em> is a <strong>good</strong> string, or </em><code>false</code><em> otherwise</em>.</p>

<p>A string <code>s</code> is <strong>good</strong> if <strong>all</strong> the characters that appear in <code>s</code> have the <strong>same</strong> number of occurrences (i.e., the same frequency).</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;abacbc&quot;
<strong>Output:</strong> true
<strong>Explanation:</strong> The characters that appear in s are &#39;a&#39;, &#39;b&#39;, and &#39;c&#39;. All characters occur 2 times in s.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> s = &quot;aaabb&quot;
<strong>Output:</strong> false
<strong>Explanation:</strong> The characters that appear in s are &#39;a&#39; and &#39;b&#39;.
&#39;a&#39; occurs 3 times while &#39;b&#39; occurs 2 times, which is not the same number of times.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> consists of lowercase English letters.</li>
</ul>
