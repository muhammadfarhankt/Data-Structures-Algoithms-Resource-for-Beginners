<h2><a href="https://leetcode.com/problems/find-if-digit-game-can-be-won">Find if Digit Game Can Be Won</a></h2> <img src='https://img.shields.io/badge/Difficulty-Easy-brightgreen' alt='Difficulty: Easy' /><hr><p>You are given an array of <strong>positive</strong> integers <code>nums</code>.</p>

<p>Alice and Bob are playing a game. In the game, Alice can choose <strong>either</strong> all single-digit numbers or all double-digit numbers from <code>nums</code>, and the rest of the numbers are given to Bob. Alice wins if the sum of her numbers is <strong>strictly greater</strong> than the sum of Bob&#39;s numbers.</p>

<p>Return <code>true</code> if Alice can win this game, otherwise, return <code>false</code>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">nums = [1,2,3,4,10]</span></p>

<p><strong>Output:</strong> <span class="example-io">false</span></p>

<p><strong>Explanation:</strong></p>

<p>Alice cannot win by choosing either single-digit or double-digit numbers.</p>
</div>

<p><strong class="example">Example 2:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">nums = [1,2,3,4,5,14]</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>

<p><strong>Explanation:</strong></p>

<p>Alice can win by choosing single-digit numbers which have a sum equal to 15.</p>
</div>

<p><strong class="example">Example 3:</strong></p>

<div class="example-block">
<p><strong>Input:</strong> <span class="example-io">nums = [5,5,5,25]</span></p>

<p><strong>Output:</strong> <span class="example-io">true</span></p>

<p><strong>Explanation:</strong></p>

<p>Alice can win by choosing double-digit numbers which have a sum equal to 25.</p>
</div>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 99</code></li>
</ul>
