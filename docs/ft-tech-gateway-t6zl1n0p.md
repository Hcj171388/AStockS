# FT Tech Gateway 接口文档备份

来源：https://market.ft.tech/gateway/doc/p/t6zl1n0p

保存时间：2026-08-14

> 说明：接口文档页面为动态页面，当前环境只能保存入口地址和本项目使用约定。后续如需离线完整文档，可将浏览器中导出的接口说明追加到本文件。

## 本项目资金流字段约定

资金净流入根据超大单、大单、中单、小单净流入合成：

```text
net_inflow = super_net_inflow + large_net_inflow + medium_net_inflow + small_net_inflow
```

资金分层评分使用：

- F 资金强度：资金净流入 / 流通市值
- E 资金效率：涨幅 / F
- T 资金持续性：连续资金净流入天数
- C 资金一致性：资金净流入 / 总成交额
- S 资金变化趋势：资金流入占比斜率

综合评分：

```text
Score = 0.35 * fund_score + 0.25 * price_score + 0.2 * chip_score + 0.2 * market_score
```
