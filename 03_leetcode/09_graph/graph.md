<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [graph](#graph)
  - [二叉树 DFS 与网格图 DFS 的区别](#%E4%BA%8C%E5%8F%89%E6%A0%91-dfs-%E4%B8%8E%E7%BD%91%E6%A0%BC%E5%9B%BE-dfs-%E7%9A%84%E5%8C%BA%E5%88%AB)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# graph



## 二叉树 DFS 与网格图 DFS 的区别

|      | 二叉树 |  网格图 |
|:----:| :--: |:-:|
| 递归入口 | 根节点 | 网格图的某个格子 |
| 递归方向 | 左儿子和右儿子 | 一般为左右上下的相邻格子 |
| 递归边界 | 空节点（或者叶节点） | 出界、遇到障碍或者已访问 |
