/*
 * 程序名：Graph.cpp，此程序是图章节各种算法的伪代码。
 * 作者：C语言技术网(www.freecplus.net) 日期：20200302
*/
//////////////////////////////////////////////////////////////////
// 图的广度优先搜索伪代码开始。

bool visited[MAXVERNUM];   // 标记顶点是否已访问的数组。

// 图的广度优先搜索
void BFST(Graph G,int v)  
{
  // 初始化访问标记数组为false。
  for ( i=0; i<MAXVERNUM; i++ ) visited[i]=false;
  
  InitQueue(Q);  // 初始化辅助队列。

  BFS(G,v);      // 从顶点v开始对图进行广度优先搜索。

  // 对其它的连通分量进行广度优先搜索。
  for (i=1;i<G.vexnum;i++) 
    if ((visited[i]==false) BFS(G,i); 
}

// 图的广度优先搜索主函数。
void BFS(Graph G,int v)  // 从顶点v开始，遍历图G。
{
  visit(v);         // 访问顶点v。
  visited[v]=true;  // 标记顶点v为已访问。
  EnQueue(Q,v);     // 顶点v入队。

  while (IsEmpty(Q)==false) { // 如果队列不为空。
    DeQueue(Q,u);        // 队头顶点出队，保存在u中。
    // for循环获取顶点u全部的边。
    for ( w=FirstAdjVex(G,u); w>=0; w=NextAdjVex(G,u,w) ) {
      if (visited[w]==false) {
        visit(w);        // 访问顶点w。
        visited[w]=true; // 标记顶点w为已访问。
        EnQueue(Q,w);    // 顶点w入队。
      } // if
    } // for
  } // while
}

// 图的广度优先搜索伪代码结束。
//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////
// 图的深度优先搜索伪代码开始。

bool visited[MAXVERNUM];   // 标记顶点是否已访问的数组。

// 图的深度优先搜索
void DFST(Graph G,int v)  
{
  // 初始化访问标记数组为false。
  for ( i=0; i<MAXVERNUM; i++ ) visited[i]=false;
  
  DFS(G,v);   // 从顶点v开始对图进行深度优先搜索。

  // 对其它的连通分量进行深度优先搜索。
  for (i=1;i<G.vexnum;i++) 
    if ((visited[i]==false) DFS(G,i); 
}

// 图的深度优先搜索主函数。
void DFS(Graph G,int v)  // 从顶点v开始。
{
  visit(v);         // 访问顶点v。
  visited[v]=true;  // 标记顶点v为已访问。

  // for循环获取顶点v全部的边。
  for ( w=FirstAdjVex(G,v); w>=0; w=NextAdjVex(G,v,w) ) 
    if (visited[w]==false) 
      DFS(G,w);     // 对顶点w递归调用，往深处走。
}

// 图的深度优先搜索伪代码结束。
//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////
// 采用图的广度优先搜索求无权图的最短路径伪代码开始。

bool visited[MAXVERNUM];   // 标记顶点是否已访问的数组。
bool dist[MAXVERNUM];      // 顶点的最短路径，-1-无路径。
bool path[MAXVERNUM];      // 顶点的前驱顶点，-1-无前驱。

void BFSMinDist(Graph G,int v)  // 从顶点v到其它顶点的最短路径。
{
  // 初始化访问标记、最短路径和前驱顶点数组。
  for ( i=0; i<MAXVERNUM; i++ ) {
    visited[i]=false; dist[i]=-1; path[i]=-1;
  }

  dist[v]=0;        // 顶点v的路径为0。
  visited[v]=true;  // 标记顶点v为已访问。
  EnQueue(Q,v);     // 顶点v入队。

  while (IsEmpty(Q)==false) { // 如果队列不为空。
    DeQueue(Q,u);             // 队头顶点出队，保存在u中。
    // for循环获取顶点u全部的边。
    for ( w=FirstAdjVex(G,u); w>=0; w=NextAdjVex(G,u,w) ) {
      if (visited[w]==false) {
        dist[w]=dist[u]+1;  // 顶点w的路径长度加1。
        path[w]=u;          // 顶点w的前驱顶点为v。
        visited[w]=true;    // 标记顶点w为已访问。
        EnQueue(Q,w);       // 顶点w入队。
      } // if
    } // for
  } // while
}

// 采用图的广度优先搜索求无权图的最短路径伪代码结束。
//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////
// 拓扑排序伪代码开始。

// 拓扑排序主函数，若图G无环，排序结果保存在result数组中，否则返回失败。
bool TopoSort(Graph G)  
{
  int indegree[MAXVERNUM];   // 存放全部顶点的入度。
  int result[MAXVERNUM];     // 存放拓扑排序的结果。
  int count=0;               // 已排序顶点的计数器。

  // 求图G中全部顶点的入度，保存在数组indegree中。
  FindInDegree(G,indegree);  

  InitStack(S);     // 初始化栈，用于保存入度为0的顶点。

  for ( i=0; i<G.vexnum; i++ ) // 全部入度为0的顶点入栈。
    if (indegree[i]==0) Push(S,i);

  while (IsEmpty(S)==false) {  // 如果栈不为空。
    Pop(S,u);                  // 出栈一个元素，保存在u中。
    result[count++]=u;         // 输出顶点u，计数器count加1。
    // for循环获取顶点u全部的出边。
    for ( w=FirstAdjVex(G,u); w>=0; w=NextAdjVex(G,u,w) ) {
      indegree[w]--;   // 将出边指向的顶点的入度减1。
      if (indegree[w]==0) Push(S,w);  // 入度为0的顶点入栈。
    } 
  } 

  if (count<G.vexnum) return false; // 图中有环，排序失败。
  else return true;                 // 排序成功。
}

// 拓扑排序伪代码开始。
//////////////////////////////////////////////////////////////////
